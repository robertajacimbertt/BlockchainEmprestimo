#!/bin/bash
#
# Author: Juliana M Destro
#
# Exit on first error
set -e

# Print the usage message
function printHelp () {
  echo ""
  echo "Script to automatically upgrade the version of a local fabric"
  echo ""
  echo "Usage: "
  echo "  chaincode_upgrade.sh"
  echo ""
  echo "Optional: "
  echo "  chaincode_upgrade.sh <increment>"
  echo "  (Increment can be 0.1 or 1)"
}

# Parse commandline args
while getopts "h?" opt; do
  case "$opt" in
    h|\?)
      printHelp
      exit 0
    ;;
  esac
done

CHANNEL_ID=mychannel
CHAINCODE_NAME=mycc
ORDERER=orderer.example.com:7050
CHAINCODE_PATH=github.com/chaincode/chaincode_example02/go/
CORE_PEER_CONTAINER=peer0.org1.example.com

INCREMENT=${1:-1}

# Get version of the last chaincode installed and increment
CCVER=`docker exec -it ${CORE_PEER_CONTAINER} ls -1tr /var/hyperledger/production/chaincodes|grep ${CHAINCODE_NAME}|awk 'END{print}'`
CCVER=${CCVER[@]#${CHAINCODE_NAME}.}
CCVER=`echo -e $CCVER | tr -d '[:space:]'`
CCVER=`echo $CCVER + $INCREMENT | bc`
CHAINCODE_VERSION=$CCVER

setGlobals () {
	if [ $1 -eq 0 -o $1 -eq 1 ] ; then
		CORE_PEER_LOCALMSPID="Org1MSP"
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
		if [ $1 -eq 0 ]; then
			CORE_PEER_ADDRESS=peer0.org1.example.com:7051
		else
			CORE_PEER_ADDRESS=peer1.org1.example.com:7051
		fi
	else
		CORE_PEER_LOCALMSPID="Org2MSP"
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
		if [ $1 -eq 2 ]; then
			CORE_PEER_ADDRESS=peer0.org2.example.com:7051
		else
			CORE_PEER_ADDRESS=peer1.org2.example.com:7051
		fi
		echo $CORE_PEER_LOCALMSPID
	fi
}


echo
echo "======================================================================================"
echo "      ASSETCHAIN - Install and Upgrade chaincode to version ${CHAINCODE_VERSION}"
echo "======================================================================================"
echo
echo
echo "============== Starting cli container..."
if [ ! "$(docker ps -q -f name=cli)" ]; then 
    docker start cli
fi

echo
setGlobals 0
echo "============== Installing version ${CHAINCODE_VERSION} in Peer ${CORE_PEER_ADDRESS}"
docker exec -e "CORE_PEER_ADDRESS=${CORE_PEER_ADDRESS}" -e "CORE_PEER_LOCALMSPID=${CORE_PEER_LOCALMSPID}" -e "CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH}" cli peer chaincode install -n ${CHAINCODE_NAME} -v ${CHAINCODE_VERSION} -p ${CHAINCODE_PATH}

echo
setGlobals 2
echo "============== Installing version ${CHAINCODE_VERSION} in Peer ${CORE_PEER_ADDRESS}"
docker exec -e "CORE_PEER_ADDRESS=${CORE_PEER_ADDRESS}" -e "CORE_PEER_LOCALMSPID=${CORE_PEER_LOCALMSPID}" -e "CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH}" cli peer chaincode install -n ${CHAINCODE_NAME} -v ${CHAINCODE_VERSION} -p ${CHAINCODE_PATH}

echo
setGlobals 3
echo "============== Installing version ${CHAINCODE_VERSION} in Peer ${CORE_PEER_ADDRESS}"
docker exec -e "CORE_PEER_ADDRESS=${CORE_PEER_ADDRESS}" -e "CORE_PEER_LOCALMSPID=${CORE_PEER_LOCALMSPID}" -e "CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH}" cli peer chaincode install -n ${CHAINCODE_NAME} -v ${CHAINCODE_VERSION} -p ${CHAINCODE_PATH}

echo
echo "============== Upgrade chaincode to version ${CHAINCODE_VERSION}"
docker exec -e "CORE_PEER_ADDRESS=${CORE_PEER_ADDRESS}" -e "CORE_PEER_LOCALMSPID=${CORE_PEER_LOCALMSPID}" -e "CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH}" cli peer chaincode upgrade -o ${ORDERER} -C ${CHANNEL_ID} -n ${CHAINCODE_NAME} -v ${CHAINCODE_VERSION} -c '{"Args":["init","c","100","d","200"]}' -P "OR	('Org1MSP.member','Org2MSP.member')" 

echo
echo "============== Stopping cli container..."
if [ "$(docker ps -q -f name=cli)" ]; then 
    docker stop cli
fi


echo
echo "All done!"
echo
echo

exit 0
