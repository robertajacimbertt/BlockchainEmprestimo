var hfc = require('fabric-client')
var path = require('path')

const enrollHelper = apprequire('helpers/enrollHelper.js')

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

module.exports = function(body) {
    var channel = {}
    var client = null

    return new Promise(function(resolve, reject) {
        console.log("Create a client")
        client = new hfc();
        return enrollHelper.enrollUser(client, options)
        .then((user) => {
            console.log("Check user is enrolled, and set a query URL in the network");
            if (user === undefined || user.isEnrolled() === false) {
                console.error("User not defined, or not enrolled - error");
            }

            channel = client.newChannel(options.channel_id);
            channel.addPeer(client.newPeer(options.network_url));
            return;
        }).then(() => {
            console.log("Make query");
            var transaction_id = client.newTransactionID();
            console.log("Assigning transaction_id: ", transaction_id._transaction_id);

            const request = {
                chaincodeId: options.chaincode_id,
                txId: transaction_id,
                fcn: body.fcn,
                args: body.args
            };
            return channel.queryByChaincode(request);
        }).then((query_responses) => {
            console.log("returned from query");
            if (!query_responses.length) {
                console.log("No payloads were returned from query");
                reject("No payloads were returned from query")
            } else {
                console.log("Query result count = ", query_responses.length)
            }
            if (query_responses[0] instanceof Error) {
                console.error("error from query = ", query_responses[0]);
                reject(query_responses[0].toString())
            } else {
				console.log("Response is ", query_responses[0].toString());
				resolve(query_responses[0].toString())
			}
        }).catch((err) => {
            console.error("Caught Error", err);
            reject(err)
        });
    })
}
