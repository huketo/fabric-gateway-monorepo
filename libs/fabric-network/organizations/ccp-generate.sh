#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/g" \
        -e "s/\${P0PORT}/$2/g" \
        -e "s/\${P1PORT}/$3/g" \
        -e "s/\${CAPORT}/$4/g" \
        -e "s#\${PEERPEM}#$PP#g" \
        -e "s#\${CAPEM}#$CP#g" \
        organizations/ccp-template.json
}

function json_local_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/g" \
        -e "s/\${P0PORT}/$2/g" \
        -e "s/\${P1PORT}/$3/g" \
        -e "s/\${CAPORT}/$4/g" \
        -e "s#\${PEERPEM}#$PP#g" \
        -e "s#\${CAPEM}#$CP#g" \
        organizations/ccp-local-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/g" \
        -e "s/\${P0PORT}/$2/g" \
        -e "s/\${P1PORT}/$3/g" \
        -e "s/\${CAPORT}/$4/g" \
        -e "s#\${PEERPEM}#$PP#g" \
        -e "s#\${CAPEM}#$CP#g" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

function explorer_ca_json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/g" \
        -e "s/\${P0PORT}/$2/g" \
        -e "s/\${P1PORT}/$3/g" \
        -e "s/\${CAPORT}/$4/g" \
        -e "s#\${PEERPEM}#$PP#g" \
        -e "s#\${CAPEM}#$CP#g" \
        organizations/ccp-explorer-ca-template.json
}

function explorer_json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    local PEERPRIV=$7
    sed -e "s/\${ORG}/$1/g" \
        -e "s/\${P0PORT}/$2/g" \
        -e "s/\${P1PORT}/$3/g" \
        -e "s/\${CAPORT}/$4/g" \
        -e "s#\${PEERPEM}#$PP#g" \
        -e "s#\${CAPEM}#$CP#g" \
        -e "s#\${PEERPRIV}#$PEERPRIV#g" \
        organizations/ccp-explorer-template.json
}

function caliper_yaml_ccp {
    sed -e "s#\${CHANNEL_NAME}#$CHANNEL_NAME#g" \
        -e "s#\${CHAINCODE_NAME}#$CHAINCODE_NAME#g" \
        -e "s#\${O1PRIV}#$O1PRIV#g" \
        -e "s#\${O2PRIV}#$O2PRIV#g" \
        -e "s#\${O3PRIV}#$O3PRIV#g" \
        organizations/caliper-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=1
P0PORT=7051
P1PORT=7151
CAPORT=7054
PEERPEM=organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
CAPEM=organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem
O1PRIV=$(ls organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/*_sk | head -n 1 | sed 's/^organizations\///')

# Copy the certificate and private key to the api config folder
[[ -d ../../apps/api/src/configs/org1/user1 ]] || mkdir -p ../../apps/api/src/configs/org1/user1
cp -r organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore ../../apps/api/src/configs/org1/user1/keystore
cp -r organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts ../../apps/api/src/configs/org1/user1/signcerts
cp -r organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/cacerts ../../apps/api/src/configs/org1/user1/cacerts

echo "$(json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org1.example.com/connection-org1.json
echo "$(yaml_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org1.example.com/connection-org1.yaml
echo "$(explorer_ca_json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > explorer/connection-profile/connection-ca-org1.json
echo "$(explorer_json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM $O1PRIV)" > explorer/connection-profile/connection-org1.json

ORG=2
P0PORT=9051
P1PORT=9151
CAPORT=8054
PEERPEM=organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem
CAPEM=organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem
O2PRIV=$(ls organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore/*_sk | head -n 1 | sed 's/^organizations\///')

# Copy the certificate and private key to the api config folder
[[ -d ../../apps/api/src/configs/org2/user1 ]] || mkdir -p ../../apps/api/src/configs/org2/user1
cp -r organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore ../../apps/api/src/configs/org2/user1/keystore
cp -r organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/signcerts ../../apps/api/src/configs/org2/user1/signcerts
cp -r organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/cacerts ../../apps/api/src/configs/org2/user1/cacerts

echo "$(json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org2.example.com/connection-org2.json
echo "$(yaml_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org2.example.com/connection-org2.yaml
echo "$(explorer_ca_json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > explorer/connection-profile/connection-ca-org2.json
echo "$(explorer_json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM $O2PRIV)" > explorer/connection-profile/connection-org2.json

ORG=3
P0PORT=11051
P1PORT=11151
CAPORT=11054
PEERPEM=organizations/peerOrganizations/org3.example.com/tlsca/tlsca.org3.example.com-cert.pem
CAPEM=organizations/peerOrganizations/org3.example.com/ca/ca.org3.example.com-cert.pem
O3PRIV=$(ls organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp/keystore/*_sk | head -n 1 | sed 's/^organizations\///')

# Copy the certificate and private key to the api config folder
[[ -d ../../apps/api/src/configs/org3/user1 ]] || mkdir -p ../../apps/api/src/configs/org3/user1
cp -r organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp/keystore ../../apps/api/src/configs/org3/user1/keystore
cp -r organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp/signcerts ../../apps/api/src/configs/org3/user1/signcerts
cp -r organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp/cacerts ../../apps/api/src/configs/org3/user1/cacerts

echo "$(json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org3.example.com/connection-org3.json
echo "$(yaml_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org3.example.com/connection-org3.yaml
echo "$(explorer_ca_json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > explorer/connection-profile/connection-ca-org3.json
echo "$(explorer_json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM $O3PRIV)" > explorer/connection-profile/connection-org3.json

echo "$(caliper_yaml_ccp)" > caliper.yaml
