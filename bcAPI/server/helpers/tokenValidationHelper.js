const yamler = require('yamler')
const jwt    = require('jsonwebtoken')
const token  = yamler(`${__dirname}/../../config/config.yml`).parameters.token


function generateToken() {

	let expiration = undefined
	return jwt.sign({ created: new Date() }, token.secret_key, expiration)

}

function verifyToken(userToken) {
    return new Promise((resolve, reject) => {
        jwt.verify(userToken, token.secret_key, function(err, decoded) {
            if (err) {
                reject({ message:'TOKEN_INVALID' })
            } else {
                resolve(userToken)
            }
        })
    })
}

module.exports = {
    verifyToken: verifyToken,
    generateToken: generateToken
}
