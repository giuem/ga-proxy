const gulp = require("gulp");
const buble = require("gulp-buble");
const uglify = require("gulp-uglify");
const rename = require("gulp-rename");

function build() {
  return gulp
    .src("./js/src/**/*.js")
    .pipe(buble())
    .pipe(gulp.dest("./js/dist"))
    .pipe(
      uglify({
        mangle: true,
        ie8: true
      })
    )
    .pipe(rename({ suffix: ".min" }))
    .pipe(gulp.dest("./js/dist"));
}

exports.build = build;
