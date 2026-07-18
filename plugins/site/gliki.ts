/**
 * gliki — a git-based bliki (blog + wiki) static site generator.
 *
 * Reads Markdown (and passes through other assets) from a source tree, renders
 * each file through a Handlebars layout, and derives ordering + metadata from
 * git history. A TypeScript revival of the original 2016 gliki weekend hack,
 * upgraded for nested directories, YAML front matter, and a second projection
 * (an agent skill registry — not emitted in this engine-swap step).
 *
 * Every unit is exported so the pipeline can be reused and unit-tested; running
 * the file directly (see {@link main}) builds the site and optionally serves it.
 *
 * Run:  node --experimental-strip-types plugins/site/gliki.ts [--serve]
 *
 * @module gliki
 */

import { execFileSync } from "node:child_process";
import fs from "node:fs";
import http from "node:http";
import path from "node:path";
import { pathToFileURL } from "node:url";
import matter from "gray-matter";
import Handlebars from "handlebars";
import hljs from "highlight.js";
import { Marked } from "marked";
import { markedHighlight } from "marked-highlight";

// --- config (inline for the engine-swap step; will move to gliki.config.ts) ---

/** Repository root; all other paths are resolved relative to it. */
export const ROOT = process.cwd();
/** Source content tree that gliki reads from. */
export const SRC = path.join(ROOT, "docs");
/** Build output directory (git-ignored; uploaded as the Pages artifact). */
export const OUT = path.join(ROOT, "public");
/** Directory holding the Handlebars layout templates. */
export const TPL_DIR = path.join(ROOT, "plugins/site/templates");
/** Directory of gliki's own static assets (CSS) copied verbatim into OUT. */
export const STATIC_DIR = path.join(ROOT, "plugins/site/static");

/** Site-wide values exposed to templates and used to build source/edit links. */
export const SITE = {
  title: "HGPA",
  url: "https://hgpa.tv",
  repo: "hparra/hparra.github.io",
  branch: "main",
  srcPrefix: "docs", // where content lives in-repo, for edit/source links
};

/** Non-`.md` files under {@link SRC} that must never be published (Jekyll leftovers). */
export const SKIP_STATIC = new Set([
  "Gemfile",
  "Gemfile.lock",
  "Makefile",
  "_config.yml",
  "README.md",
]);

/**
 * Shared Markdown renderer: syntax highlighting via highlight.js plus a
 * `walkTokens` pass that rewrites local `.md` links to their `.html` targets.
 */
export const marked = new Marked(
  markedHighlight({
    highlight(code, lang) {
      if (lang && hljs.getLanguage(lang)) {
        return hljs.highlight(code, { language: lang }).value;
      }
      return hljs.highlightAuto(code).value;
    },
  }),
  {
    gfm: true,
    breaks: false,
    walkTokens(token) {
      if (token.type !== "link") return;
      token.href = rewriteLocalLink(token.href);
    },
  },
);

/**
 * Rewrite a link href for static output: local `foo.md` becomes `foo.html`, and
 * `README.md`/`index.md` collapse to their containing directory. External,
 * anchor (`#…`), `mailto:`, and `tel:` links are returned unchanged.
 *
 * @param href - Raw link target as written in the Markdown source.
 * @returns The href rewritten for the generated site.
 */
export function rewriteLocalLink(href: string): string {
  // leave external, anchor, and mailto links alone
  if (/^(https?:\/\/|mailto:|#|tel:)/i.test(href)) return href;
  if (!/\.md(#|\?|$)/i.test(href)) return href;
  // README.md / index.md are directory indexes; everything else .md -> .html
  return href
    .replace(/(^|\/)(README|index)\.md(#|\?|$)/i, "$1$3")
    .replace(/\.md(#|\?|$)/i, ".html$1");
}

/**
 * Recursively collect absolute paths of every file under `dir`, skipping any
 * entry whose name starts with a dot (`.git`, `.jekyll`, dotfiles).
 *
 * @param dir - Directory to walk.
 * @returns Absolute paths of all non-hidden files found, depth-first.
 */
export function walk(dir: string): string[] {
  const out: string[] = [];
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    if (entry.name.startsWith(".")) continue; // skip .jekyll, .git, dotfiles
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) out.push(...walk(full));
    else out.push(full);
  }
  return out;
}

/**
 * Run `git` with the given arguments and return trimmed stdout, swallowing any
 * failure (missing repo, untracked path, …) as an empty string.
 *
 * @param args - Argument vector passed to the `git` executable.
 * @returns Trimmed stdout, or `""` if the command errors.
 */
export function git(args: string[]): string {
  try {
    return execFileSync("git", args, { encoding: "utf8" }).trim();
  } catch {
    return "";
  }
}

/**
 * Resolve a file's created/modified dates and latest content-commit sha from
 * git history.
 *
 * Renames are followed (`--follow`) so a file traces back to its true birth,
 * and pure rename/move commits are skipped (`--diff-filter=AM` keeps Adds and
 * content Modifies, drops Renames) so a bulk "move everything" reorg does not
 * stamp every file with the same date. The newest surviving commit is the
 * modified date; the original Add is the created date. Untracked/uncommitted
 * files fall back to filesystem mtime.
 *
 * @param srcRel - Repo-relative path to the file.
 * @returns `modified` and `created` dates (`YYYY-MM-DD`) and the modified sha
 *   (empty when the file is not yet committed).
 */
export function gitDates(srcRel: string): {
  modified: string;
  created: string;
  sha: string;
} {
  const log = git([
    "log",
    "--follow",
    "--diff-filter=AM",
    "--format=%ad|%H",
    "--date=short",
    "--",
    srcRel,
  ]);
  if (log) {
    const lines = log.split("\n");
    const [modified, sha] = lines[0].split("|");
    const created = lines[lines.length - 1].split("|")[0];
    return { modified, created, sha };
  }
  // untracked / uncommitted: fall back to filesystem mtime
  const mtime = fs
    .statSync(path.join(ROOT, srcRel))
    .mtime.toISOString()
    .slice(0, 10);
  return { modified: mtime, created: mtime, sha: "" };
}

/**
 * Extract the text of the first Markdown H1 (`# Title`) in the given source.
 *
 * @param md - Markdown body (front matter already stripped).
 * @returns The heading text, or `null` if there is no H1.
 */
export function firstH1(md: string): string | null {
  const m = md.match(/^#\s+(.+?)\s*$/m);
  return m ? m[1] : null;
}

/**
 * Turn a slug such as `binary-search` into title case (`Binary Search`) for use
 * as a last-resort page title.
 *
 * @param slug - Hyphen/underscore-delimited slug.
 * @returns The title-cased string.
 */
export function titleCase(slug: string): string {
  return slug.replace(/[-_]/g, " ").replace(/\b\w/g, (c) => c.toUpperCase());
}

/** A single content file resolved into everything the templates and index need. */
export interface Leaf {
  /** Repo-relative source path, e.g. `docs/wiki/bash.md`. */
  srcRel: string;
  /** Path relative to {@link SRC}, e.g. `wiki/bash.md`. */
  rel: string;
  /** Top-level directory (`wiki`, `notes`, …), or `""` at the root. */
  section: string;
  /** Basename without `.md` — matches Jekyll's `page.name` display. */
  name: string;
  /** Whether this file is a directory index (`README.md`/`index.md`). */
  isIndex: boolean;
  /** Parsed YAML front matter (empty object when absent). */
  meta: Record<string, unknown>;
  /** Markdown body with front matter removed. */
  content: string;
  /** Resolved page title (front matter → first H1 → title-cased name). */
  title: string;
  /** Output path relative to {@link OUT}, e.g. `wiki/bash.html`. */
  outRel: string;
  /** Absolute site URL, e.g. `/wiki/bash.html`. */
  url: string;
  /** {@link url} without the `.html` suffix (the current site's link style). */
  href: string;
  /** Created/modified dates and modified sha from {@link gitDates}. */
  git: { modified: string; created: string; sha: string };
}

/**
 * Read a Markdown file and resolve it into a {@link Leaf}: parse front matter,
 * derive the title, compute the output path/URL (honoring an explicit
 * `permalink`), and attach git dates.
 *
 * @param file - Absolute path to a `.md` file under {@link SRC}.
 * @returns The fully-resolved leaf.
 */
export function toLeaf(file: string): Leaf {
  const rel = path.relative(SRC, file).split(path.sep).join("/");
  const srcRel = path.relative(ROOT, file).split(path.sep).join("/");
  const raw = fs.readFileSync(file, "utf8");
  const { data: meta, content } = matter(raw);
  const base = path.basename(rel);
  const isIndex = base === "README.md" || base === "index.md";
  const dir = path.dirname(rel) === "." ? "" : path.dirname(rel);
  const section = rel.includes("/") ? rel.split("/")[0] : "";
  const name = base.replace(/\.md$/, "");

  let outRel: string;
  let url: string;
  if (typeof meta.permalink === "string" && meta.permalink) {
    const p = meta.permalink.replace(/^\//, "");
    if (p.endsWith("/")) {
      outRel = `${p}index.html`;
      url = `/${p}`;
    } else if (/\.[a-z0-9]+$/i.test(p)) {
      outRel = p;
      url = `/${p}`;
    } else {
      outRel = `${p}.html`;
      url = `/${p}.html`;
    }
  } else if (isIndex) {
    outRel = dir ? `${dir}/index.html` : "index.html";
    url = dir ? `/${dir}/` : "/";
  } else {
    outRel = rel.replace(/\.md$/, ".html");
    url = `/${outRel}`;
  }

  const title =
    (typeof meta.title === "string" && meta.title) ||
    firstH1(content) ||
    titleCase(name);

  return {
    srcRel,
    rel,
    section,
    name,
    isIndex,
    meta,
    content,
    title,
    outRel,
    url,
    href: url.replace(/\.html$/, ""),
    git: gitDates(srcRel),
  };
}

/**
 * Build the whole site into {@link OUT}: render every Markdown leaf through the
 * default Handlebars layout (index files get a prior Handlebars pass so their
 * ported section TOCs resolve), copy static passthrough assets and gliki's own
 * CSS, and write a `.nojekyll` marker so Pages does not re-process the artifact.
 */
export function build(): void {
  fs.rmSync(OUT, { recursive: true, force: true });
  fs.mkdirSync(OUT, { recursive: true });

  const files = walk(SRC);
  const leaves = files.filter((f) => f.endsWith(".md")).map(toLeaf);

  // page list exposed to index templates, git-recency sorted (newest first)
  const byRecency = (a: Leaf, b: Leaf) =>
    b.git.modified.localeCompare(a.git.modified);
  const pages = [...leaves].sort(byRecency);

  // Handlebars helper: `{{#each (section "wiki")}}` -> that section's pages,
  // excluding directory indexes, newest first.
  Handlebars.registerHelper("section", (name: string) =>
    pages.filter((p) => p.section === name && !p.isIndex),
  );

  const layout = Handlebars.compile(
    fs.readFileSync(path.join(TPL_DIR, "__default__.hbs"), "utf8"),
  );

  for (const leaf of leaves) {
    // drop now-meaningless Jekyll {% raw %} guards (they only ever wrapped
    // literal Liquid examples; without Jekyll the inner text stands on its own)
    const content = leaf.content.replace(/\{%\s*(end)?raw\s*%\}\s*\n?/g, "");
    // index files may contain Handlebars (the ported section TOCs); render that
    // pass first, then markdown. Regular content is passed straight through.
    const md = leaf.isIndex
      ? Handlebars.compile(content)({ site: SITE, pages, page: leaf })
      : content;
    const body = marked.parse(md) as string;

    const html = layout({
      site: SITE,
      title: leaf.title,
      body,
      meta: leaf.meta,
      filename: leaf.rel,
      git: leaf.git,
      sourceUrl: `https://github.com/${SITE.repo}/blob/${SITE.branch}/${leaf.srcRel}`,
      editUrl: `https://github.com/${SITE.repo}/edit/${SITE.branch}/${leaf.srcRel}`,
    });

    const dest = path.join(OUT, leaf.outRel);
    fs.mkdirSync(path.dirname(dest), { recursive: true });
    fs.writeFileSync(dest, html);
  }

  // static passthrough (media, CNAME, ...) from the content tree
  for (const file of files) {
    if (file.endsWith(".md")) continue;
    const rel = path.relative(SRC, file).split(path.sep).join("/");
    if (SKIP_STATIC.has(path.basename(rel))) continue;
    const dest = path.join(OUT, rel);
    fs.mkdirSync(path.dirname(dest), { recursive: true });
    fs.copyFileSync(file, dest);
  }

  // gliki's own static assets (css)
  if (fs.existsSync(STATIC_DIR)) {
    for (const file of walk(STATIC_DIR)) {
      const rel = path.relative(STATIC_DIR, file).split(path.sep).join("/");
      const dest = path.join(OUT, rel);
      fs.mkdirSync(path.dirname(dest), { recursive: true });
      fs.copyFileSync(file, dest);
    }
  }

  // tell GitHub Pages not to re-run Jekyll on our artifact
  fs.writeFileSync(path.join(OUT, ".nojekyll"), "");

  console.log(
    `gliki: built ${leaves.length} pages -> ${path.relative(ROOT, OUT)}/`,
  );
}

/**
 * Resolve a request pathname to an absolute path guaranteed to stay under
 * `base`, or `null` when the input is unsafe. Guards the dev server against
 * directory traversal (`/../package.json`) and malformed percent-encoding
 * (which would otherwise throw out of `decodeURIComponent`).
 *
 * @param base - Directory the result must remain inside (e.g. {@link OUT}).
 * @param urlPath - Raw URL pathname (query string already stripped).
 * @returns The safe absolute path, or `null` if malformed or escaping `base`.
 */
export function safeResolve(base: string, urlPath: string): string | null {
  let decoded: string;
  try {
    decoded = decodeURIComponent(urlPath);
  } catch {
    return null; // malformed percent-encoding
  }
  const resolved = path.join(base, decoded);
  const rel = path.relative(base, resolved);
  if (rel.startsWith("..") || path.isAbsolute(rel)) return null;
  return resolved;
}

/**
 * Start a minimal static file server over {@link OUT}, resolving `/` and
 * extensionless paths to their `.html`/`index.html` files the way GitHub Pages
 * does. Intended as a local `make serve` replacement, not for production.
 *
 * @param port - TCP port to listen on (default `4000`).
 */
export function serve(port = 4000): void {
  const types: Record<string, string> = {
    ".html": "text/html; charset=utf-8",
    ".css": "text/css; charset=utf-8",
    ".js": "text/javascript; charset=utf-8",
    ".gif": "image/gif",
    ".png": "image/png",
    ".jpg": "image/jpeg",
    ".svg": "image/svg+xml",
  };
  http
    .createServer((req, res) => {
      const urlPath = (req.url || "/").split("?")[0];
      const resolved = safeResolve(OUT, urlPath);
      if (resolved === null) {
        res.writeHead(400, { "content-type": "text/plain" });
        res.end("400 Bad Request");
        return;
      }
      let file = resolved;
      if (urlPath.endsWith("/")) file = path.join(file, "index.html");
      else if (!fs.existsSync(file) && fs.existsSync(`${file}.html`))
        file = `${file}.html`;
      if (fs.existsSync(file) && fs.statSync(file).isDirectory())
        file = path.join(file, "index.html");
      if (!fs.existsSync(file)) {
        res.writeHead(404, { "content-type": "text/plain" });
        res.end("404");
        return;
      }
      res.writeHead(200, {
        "content-type": types[path.extname(file)] || "application/octet-stream",
      });
      res.end(fs.readFileSync(file));
    })
    .listen(port, () =>
      console.log(`gliki: serving ${OUT} at http://127.0.0.1:${port}`),
    );
}

/**
 * CLI entry point: build the site, then start the dev server when `--serve` is
 * passed. Runs only when the file is executed directly, not when imported.
 */
export function main(): void {
  build();
  if (process.argv.includes("--serve")) serve();
}

if (import.meta.url === pathToFileURL(process.argv[1] ?? "").href) {
  main();
}
