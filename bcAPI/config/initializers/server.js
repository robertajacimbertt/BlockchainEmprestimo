/**
 * Starts the express server
 * @param  {object} config Holds the configuration used to start the server
 * @return {Express}        The server Up & Running
 */

let CronJob = require('cron').CronJob;
var XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;

function bootServer(config) {
    // Setting modules
    const express = require('express')
    const bodyParser = require('body-parser')
    const routes = require('../routes')
    const app = express()
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

    new CronJob('0 * * * * *', function() {

        // Execute code here
        console.log('Hello puppies!');
        const Http = new XMLHttpRequest();
        const url = 'http://localhost:8080/api/query';
        Http.open("POST", url);
        Http.setRequestHeader("Content-Type", "application/json")
        const data = {
            "fcn": "query",
            "args": ["fin1"]
        };
        Http.send(JSON.stringify(data));
        Http.onreadystatechange = (e) => {
            console.log("Chamou -> ", Http.responseText)
        }

    }, null, true, 'America/Sao_Paulo');
}

module.exports = bootServer;