const tokenValidator = apprequire('helpers/tokenValidationHelper')
const ignoredRoutes = [
    '/invoke',
    '/query'
]

module.exports = (req, res, next) => {
    if(_isIgnored(req.url)) { return next() }

    let token = req.get('X-Authtoken')

    if(!token) return res.status(401).json({ message: "TOKEN_NOT_PRESENT"})

    tokenValidator.verifyToken(token)
    .then(token => {
        next()
    })
    .catch(err => res.status(401).json(err))
}

function _isIgnored(url) {
	let ignored = ignoredRoutes.map(route => { if(url.match(route)) { return true } })

	if (ignored.indexOf(true) > -1 ) { return true }
	else { return false }
}
