
const gulp = require('gulp')
const webserver = require('gulp-webserver')
const less = require('gulp-less')

const path = require('path')
const child_process = require('child_process')
const gliki = require('./.gliki/gliki')

gulp.task('default', ['build', 'watch', 'webserver'])

gulp.task('build', ['media', 'gliki', 'less'])

gulp.task('watch', () => {
  gulp.watch('media/**/*', ['media'])
  gulp.watch(['*.md', '.gliki/hbs/*.hbs'], ['gliki'])
  gulp.watch('.gliki/less/*.less', ['less'])
})

gulp.task('webserver', () => {
  gulp.src('.gliki/build')
    .pipe(webserver({
      port: 4472, // HGPA
      livereload: true
    }))
})

gulp.task('media', () => {
  gulp.src('media/**/*')
    .pipe(gulp.dest('.gliki/build/media/'))
})

gulp.task('gliki', () => {
  return gliki()
})

gulp.task('less', () => {
  gulp.src('.gliki/less/*.less')
    .pipe(less({
      paths: [ 
        path.join(__dirname, 'less', 'includes'),
        path.join(__dirname, 'node_modules/highlight.js/styles'),
      ]
    }))
    .pipe(gulp.dest('.gliki/build/css'))
})
