var gulp = require('gulp');
var less = require('gulp-less');
var cssmin = require('gulp-cssmin');
var rename = require("gulp-rename");
var watch = require('gulp-watch');

gulp.task('less', function () {
    return gulp.src('static/less/app.less')
        .pipe(less())
        .pipe(gulp.dest('static/css'));
});
;

gulp.task('cssmin', ['less'], function () {
    gulp.src('static/css/app.css')
        .pipe(cssmin({
            keepSpecialComments: 0
        }))
        .pipe(rename({suffix: '.min'}))
        .pipe(gulp.dest('static/css'));
});

gulp.task('watch', function() {
    gulp.watch('static/less/**/*.less', ['cssmin']);
});

gulp.task('build', ['less', 'cssmin'], function () {});

gulp.task('default', function() {
    gulp.start('cssmin')
});