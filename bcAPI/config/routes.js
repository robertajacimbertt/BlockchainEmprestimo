const routes          = require('ike-router')(`${__dirname}/../server/controllers/`)
const interceptor     = apprequire('helpers/authInterceptor')

/*
* Interceptor
 */
routes.mountMiddleware(interceptor)

/*
* Generic Routes
 */
routes.post('/invoke', 'invokeController#invoke')
routes.post('/query', 'queryController#query')


module.exports = routes.draw();
