var hfc = require('fabric-client')
var path = require('path')
var util = require('util')

const enrollHelper = apprequire('helpers/enrollHelper.js')

const options = {
    wallet_path: path.join(__dirname, './creds'),
    channel_id: 'mychannel',
    chaincode_id: 'mycc',
    peer_url: 'grpc://localhost:7051',
    event_url: 'grpc://localhost:7053',
    orderer_url: 'grpc://localhost:7050',
    ca_url: 'http://localhost:7054',
    user_id: 'admin',
    user_secret: 'adminpw',
    org: 'Org1MSP'
}

module.exports = function(body) {
    var channel = {};
    var client = null;
    var targets = [];
    var tx_id = null;
    var chaincodeResponse = {
        message: ""
    };

    return new Promise(function(resolve, reject) {
        console.log("Create a client")
        client = new hfc()
        return enrollHelper.enrollUser(client, options)
            .then((user) => {
                console.log("Check user is enrolled, and set a query URL in the network");
                if (user === undefined || user.isEnrolled() === false) {
                    console.error("User not defined, or not enrolled - error");
                }
                channel = client.newChannel(options.channel_id);
                var peerObj = client.newPeer(options.peer_url);
                channel.addPeer(peerObj);
                channel.addOrderer(client.newOrderer(options.orderer_url));
                targets.push(peerObj);
                return;
            })
            .then(() => {
                tx_id = client.newTransactionID();
                console.log("Assigning transaction_id: ", tx_id._transaction_id);
                // createCar - requires 5 args, ex: args: ['CAR11', 'Honda', 'Accord', 'Black', 'Tom'],
                // changeCarOwner - requires 2 args , ex: args: ['CAR10', 'Barry'],
                // send proposal to endorser

                var request = {
                    targets: targets,
                    chaincodeId: options.chaincode_id,
                    fcn: body.fcn,
                    args: body.args,
                    chainId: options.channel_id,
                    txId: tx_id
                };
                return channel.sendTransactionProposal(request);
            })
            .then((results) => {
                var proposalResponses = results[0];
                var proposal = results[1];
                var header = results[2];
                let isProposalGood = false;
                if (proposalResponses && proposalResponses[0].response &&
                    proposalResponses[0].response.status === 200) {
                    isProposalGood = true;
                    console.log('transaction proposal was good');
                } else {
                    console.error('transaction proposal was bad');
                }
                if (isProposalGood) {
                    chaincodeResponse.message = "" + proposalResponses[0].response.payload
                    console.log(util.format(
                        'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s", metadata - "%s", endorsement signature: %s',
                        proposalResponses[0].response.status, proposalResponses[0].response.message,
                        proposalResponses[0].response.payload, proposalResponses[0].endorsement.signature));
                    var request = {
                        proposalResponses: proposalResponses,
                        proposal: proposal,
                        header: header
                    };
                    // set the transaction listener and set a timeout of 30sec
                    // if the transaction did not get committed within the timeout period,
                    // fail the test
                    var transactionID = tx_id.getTransactionID();
                    var eventPromises = [];
                    let eh = client.newEventHub();
                    eh.setPeerAddr(options.event_url);
                    eh.connect();

                    let txPromise = new Promise((resolve, reject) => {
                        let handle = setTimeout(() => {
                            eh.disconnect();
                            reject();
                        }, 30000);

                        eh.registerTxEvent(transactionID, (tx, code) => {
                            clearTimeout(handle);
                            eh.unregisterTxEvent(transactionID);
                            eh.disconnect();

                            if (code !== 'VALID') {
                                console.error(
                                    'The transaction was invalid, code = ' + code);
                                reject();
                            } else {
                                console.log(
                                    'The transaction has been committed on peer ' +
                                    eh._ep._endpoint.addr);
                                resolve();
                            }
                        });
                    });
                    eventPromises.push(txPromise);
                    var sendPromise = channel.sendTransaction(request);
                    return Promise.all([sendPromise].concat(eventPromises)).then((results) => {
                            console.log(' event promise all complete and testing complete');
                            return results[0]; // the first returned value is from the 'sendPromise' which is from the 'sendTransaction()' call
                        })
                        .catch((err) => {
                            console.error('Failed to send transaction and get notifications within the timeout period.');
                        });
                } else {
                    console.error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
                    var rsp = Object.getOwnPropertyNames(proposalResponses[0]);
                    console.log(proposalResponses[0]["message"]);
                    reject(proposalResponses[0]["message"]);
                }
            })
            .catch(err => {
                console.error('Failed to send proposal due to error: ' + err.stack ? err.stack : err);
                reject('Failed to send proposal due to error: ' + err.stack ? err.stack : err)
            })
            .then((response) => {
                if (response && response.status && response.status === 'SUCCESS') {
                    console.log('Successfully sent transaction to the orderer.');

                    if (chaincodeResponse.message == "") {
                        resolve(tx_id.getTransactionID());
                    } else {
                        resolve(chaincodeResponse);
                    }

                } else if (response && response.status) {
                    console.error('Failed to order the transaction. Error code: ' + response.status);
                    reject('Failed to order the transaction. Error code: ' + response.status);
                } else {
                    console.error('Failed to order the transaction.');
                    reject('Failed to order the transaction.');
                }
            })
            .catch(err => {
                console.error('Failed to send transaction due to error: ' + err.stack ? err.stack : err);
                reject('Failed to send transaction due to error: ' + err.stack ? err.stack : err)
            });
    })
}