/**
 * Responsible to configure the server initialization for development
 */
const env = process.env.NODE_ENV || 'development'
const yamler    = require('yamler')
const appConfig = yamler('config/config.yml').environments[env].application

require('./initializers/server.js')(appConfig) //bootServer
