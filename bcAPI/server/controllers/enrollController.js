const enrollHelper = apprequire('helpers/enrollHelper.js')

var path = require('path')

const options = {
    wallet_path: path.join(__dirname, './creds'),
    network_url: 'grpc://localhost:7051',
    ca_url: 'http://localhost:7054',
    channel_id: 'mychannel',
    chaincode_id: 'mycc',
    user_id: 'admin',
    user_secret: 'adminpw',
    org: 'Org1MSP'
}

class Enroll {
    constructor() {}

    registerNewUser(req, res) {
        let body = req.body

        enrollHelper.registerNewUser(options, body)
            .then((result) => {
                res.status(200).send(result)
            })
            .catch(err => {
                res.status(400).send(err)
            })
    }
}

module.exports = Enroll