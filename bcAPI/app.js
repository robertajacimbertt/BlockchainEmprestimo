/**
 * APP.JS - Main entry file
 * The app starts here. The current NODE_ENV will be booted.
 */

/**
 * Improve the require path. Controllers, models and helpers can
 * be required direclty, without the need of traversing the
 * relative path.
 *
 * This will work, anywhere on the app.
 * const controller = apprequire('controller/samples');
 */
global.apprequire = function(fileName) { return require(__dirname + '/server/' + fileName)};

/** * Boot the right environment based on NODE_ENV. */
require('./config/environment.js');

