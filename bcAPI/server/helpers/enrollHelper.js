var hfc = require('fabric-client')
var CaService = require('fabric-ca-client')
var User = require('fabric-client/lib/User.js');
var path = require('path')

/**
 * Gets User Enrollment info.
 * @param  {fabric-client} client  The fabric client that will execute the routines
 * @param  {Object}        options JSON with user information
 * @return {User}                  The enrolled user
 */
function enrollUser(client, options) {
    return new Promise( (resolve, reject) => {
        console.log("Set up wallet location")
        return hfc.newDefaultKeyValueStore({path: options.wallet_path})
        .then(wallet => {
            client.setStateStore(wallet)
            return client.getUserContext(options.user_id, true)
        })
        .then(user => {
            if(user && user.isEnrolled) {
                resolve(user)
            }
            else {
                let tlsOptions = {
                trustedRoots: "",
                verify: false
                } 
               
                let ca_client = new CaService(options.ca_url)
                user = ca_client.enroll({
                    enrollmentID:      options.user_id,
                    enrollmentSecret:  options.user_secret
                })
                .then(enrollment => {
                    member = new User(options.user_id, client)
                    return member.setEnrollment(enrollment.key, enrollment.certificate, options.org)
                })
                .then( () => {
                    return client.setUserContext(member);
                })
                .then(() => {
                    return member
                })
                .catch(err => {
                    console.error("Error during enrollment. " + err.stack ? err.stack : err)
                    reject("Error during enrollment. " + err.stack ? err.stack : err)
                })
                resolve(user)
            }
        })
        .catch(err => {
            console.error("Error getting user context. " + err.stack ? err.stack : err)
            reject("Error getting user context. " + err.stack ? err.stack : err)
        })
    })
}

/**
 * Gets the Admin enrollment file
 * @param  {[type]} client [description]
 * @return {[type]}        [description]
 */
function getAdmin (client) {

    let options = {
        wallet_path: path.join(__dirname, './creds'),
        user_id: 'admin',
        user_secret: 'adminpw',
        org: 'Org1MSP',
        ca_url: 'http://localhost:7054'
    }

    return enrollUser(client, options)

}


module.exports = {
    enrollUser: enrollUser
    
}
