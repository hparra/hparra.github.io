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
  href = /^(?!https?:\/\/).+\.md$/.test(href)
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
  
  // generate file candidates
  // assumes all files are flat in working dir
  const files = child_process.execSync(`
      git ls-files \
        | grep .md \
        | while read filename; do
            modified_on=$(git log -1 --format="%cd" --date=short -- $filename)
            created_on=$(git log --diff-filter=A --format="%cd" --date=short -- $filename)
            echo "$modified_on,$created_on,$filename"
          done \
        | sort -r
      `, { encoding: 'utf-8' })
    .trim()
    .split('\n')
    .map(line => {
      return ((fields) => {
        return {
          modified_on: fields[0],
          created_on:  fields[1],
          filename:    fields[2]
        }
      })(line.split(','))
    })  

  // create build directory
  child_process.execSync(`mkdir -p ${BUILD_DIR}`)

  return readFile(path.join(GLIKI_DIR, 'hbs/markdown.hbs'))
    .then(hbs => Handlebars.compile(hbs))
    .then(template => {
      return Promise.all(files.map(file => readFile(file.filename)
                                              .then(md => Object.assign({ files: files }, file, { 
                                                markdown: marked(md, { 
                                                  renderer: glikiRenderer 
                                                }) 
                                              }))
                                              .then(context => template(context))
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