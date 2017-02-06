/*
 * Generate HTML files from Markdown files embedded into `handlebars` templates.
 * Also generate an index HTML using a template and data from `git`.
 */

const fs = require('fs')
const path = require('path')
const child_process = require('child_process')
const marked = require('marked')
const highlight = require('highlight.js')
const Handlebars = require('handlebars')

// Synchronous highlighting with highlight.js
marked.setOptions({
  highlight: function (code) {
    return highlight.highlightAuto(code).value;
  }
})

const glikiRenderer = new marked.Renderer()
glikiRenderer.link = (href, title, text) => {
  // NOTE: We break sanitization because we know we won't do it
  // Replace .md extension for local links only with .html
  href = /^(?!https?:\/\/).+\.md(#.*)?$/.test(href)
    ? href.replace('.md', '.html')
    : href
  const titleAttr = title ? `title="${title}"` : ''
  const out = `<a href="${href}" ${titleAttr}>${text}</a>`
  return out
}

function gliki(options) {
  
  GLIKI_DIRNAME = '.gliki'
  WORKING_DIR = path.resolve('./') // TODO: pass in
  GLIKI_DIR = path.join(WORKING_DIR, GLIKI_DIRNAME) // FIXME
  BUILD_DIR = path.join(GLIKI_DIR, 'build')
  
  const markedOptions = {
    renderer: glikiRenderer,
    gfm:      true,
    tables:   true,
    breaks:   false,
  }
  
  const files = readGlikiFiles()
  const filesWithoutREADME = files.filter(file => file.filename != 'README.md') // FIXME: this is janky

  // create build directory
  child_process.execSync(`mkdir -p ${BUILD_DIR}`)

  return compileHandlebarsMap(path.join(GLIKI_DIR, 'hbs'))
    .then(templateMap => {
      return Promise.all(files.map(file => readFile(file.filename)
                                              .then(md => Object.assign({ files: filesWithoutREADME }, file, { 
                                                markdown: marked(md, markedOptions) 
                                              }))
                                              .then(context => {
                                                // look for individual templates or use default one
                                                const template = templateMap[file.filekey]
                                                  ? templateMap[file.filekey]
                                                  : templateMap['__default__']
                                                return template(context)
                                              })
                                              .then(html => {
                                                const outfile = file.filename === 'README.md'
                                                  ? 'index.html'
                                                  : file.filename.replace('.md', '.html')
                                                return writeFile(path.join(BUILD_DIR, outfile), html)
                                              })))
    })
    .catch(console.error)
}

//
// Helpers
//

/**
 * Return a hashmap (object) of compiled Handlebars templates.
 *
 */
function compileHandlebarsMap(dirpath) {
  return readdir(dirpath)
    .then(filenames => Promise.all(filenames.map(filename => 
      readFile(path.join(dirpath, filename))
        .then(hbs => ({
          key:      path.basename(filename).replace('.hbs', ''),
          template: Handlebars.compile(hbs)
        }))))
      .then(tmpls => tmpls.reduce((prev, curr) => {
        prev[curr.key] = curr.template
        return prev
      }, {})))
}

function readGlikiFiles() {
  // generate file candidates
  // assumes all files are flat in working dir
  // See https://git-scm.com/docs/pretty-formats
  const files = child_process.execSync(`
      git ls-files *.md \
        | while read filename; do
            modified_by=$(git log -1 --format="%ad,%an,%ae,%H" --date=short -- $filename)
            created_by=$(git log --diff-filter=A --format="%ad,%an,%ae,%H" --date=short -- $filename)
            echo "$modified_by,$created_by,$filename"
          done \
        | sort -r
      `, { encoding: 'utf-8' })
    .trim()
    .split('\n')
    .map(line => {
      return ((fields) => {
        // Something like github's schema -- https://developer.github.com/v3/git/commits/
        return {
          modified_by: {
            sha: fields[3],
            author: {
              date:  fields[0],
              name:  fields[1],
              email: fields[2],
            }
          },
          created_by: {
            sha: fields[7],
            author: {
              date:  fields[4],
              name:  fields[5],
              email: fields[6],
            }
          },
          filename: fields[8],
          filekey:  fields[8].replace('.md', '')
        }
      })(line.split(','))
    })
  
  return files
}

/**
 * Promise-version of fs.readdir (kinda)
 * @param String file
 */
function readdir(path) {
  return new Promise((resolve, reject) => {
    fs.readdir(path, (err, files) => {
      if (err) reject(err)
      resolve(files)
    })
  })
}

/**
 * Promise-version of fs.readFile (kinda)
 * @param String file
 */
function readFile(file) {
  return new Promise((resolve, reject) => {
    fs.readFile(file, 'utf8', (err, data) => {
      if (err) reject(err)
      resolve(data)
    })
  })
}

/**
 * Promise-version of fs.writeFile (kinda)
 * @param String file
 * @param String data
 */
function writeFile(file, data) {
  return new Promise((resolve, reject) => {
    fs.writeFile(file, data, (err) => {
      if (err) reject(err)
      resolve()
    })
  })
}

module.exports = gliki