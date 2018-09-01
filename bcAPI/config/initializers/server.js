/**
 * Starts the express server
 * @param  {object} config Holds the configuration used to start the server
 * @return {Express}        The server Up & Running
 */
function bootServer(config) {
	// Setting modules
	const express    = require('express')
	const bodyParser = require('body-parser')
	const routes     = require('../routes')
	const app        = express()
	var logStdout = process.stdout;

	//Injecting the middlewares
	app.use(bodyParser.urlencoded({ extended: true }))
	app.use(bodyParser.json())

	//Serving the front-end
	app.use(express.static('client'))

	//Serving the back-end
	app.use('/api', routes)

	//starts the server
	app.listen(config.port, () => { console.log(`Express server is up on port ${config.port}!`) })
}

module.exports = bootServer
