import assert from "node:assert/strict";
import { describe, it } from "node:test";
import {
  build,
  firstH1,
  rewriteLocalLink,
  serve,
  titleCase,
  walk,
} from "./gliki.ts";

describe("rewriteLocalLink", () => {
  it("rewrites local .md links to .html", () => {
    assert.equal(rewriteLocalLink("editors.md"), "editors.html");
    assert.equal(rewriteLocalLink("../notes/x.md"), "../notes/x.html");
  });

  it("preserves anchors and query strings", () => {
    assert.equal(rewriteLocalLink("editors.md#usage"), "editors.html#usage");
    assert.equal(rewriteLocalLink("x.md?v=1"), "x.html?v=1");
  });

  it("collapses README.md / index.md to their directory", () => {
    assert.equal(rewriteLocalLink("/README.md"), "/");
    assert.equal(rewriteLocalLink("sub/README.md"), "sub/");
    assert.equal(rewriteLocalLink("sub/index.md"), "sub/");
  });

  it("leaves external, mailto, tel, and anchor links untouched", () => {
    assert.equal(
      rewriteLocalLink("https://example.com/a.md"),
      "https://example.com/a.md",
    );
    assert.equal(rewriteLocalLink("mailto:a@b.com"), "mailto:a@b.com");
    assert.equal(rewriteLocalLink("tel:+15551234"), "tel:+15551234");
    assert.equal(rewriteLocalLink("#section"), "#section");
  });

  it("leaves non-markdown targets untouched", () => {
    assert.equal(rewriteLocalLink("bash"), "bash");
    assert.equal(rewriteLocalLink("media/cat.gif"), "media/cat.gif");
  });
});

describe("firstH1", () => {
  it("returns the first H1's text", () => {
    assert.equal(firstH1("# Hello\n\nbody"), "Hello");
  });

  it("ignores H2+ and finds the first real H1", () => {
    assert.equal(firstH1("## Sub\n\n# Real Title\n"), "Real Title");
  });

  it("trims trailing whitespace", () => {
    assert.equal(firstH1("#   Spaced   "), "Spaced");
  });

  it("returns null when there is no H1", () => {
    assert.equal(firstH1("no heading here\n## only h2"), null);
  });
});

describe("titleCase", () => {
  it("title-cases hyphen and underscore slugs", () => {
    assert.equal(titleCase("binary-search"), "Binary Search");
    assert.equal(titleCase("a_mind_for_numbers"), "A Mind For Numbers");
  });

  it("capitalizes a single word", () => {
    assert.equal(titleCase("bash"), "Bash");
  });
});

describe("module surface", () => {
  it("exports the pipeline functions", () => {
    assert.equal(typeof build, "function");
    assert.equal(typeof serve, "function");
    assert.equal(typeof walk, "function");
  });
});
