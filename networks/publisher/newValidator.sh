#!/usr/bin/env bash

home=$HOME
src="${home}/go/src/github.com/axc-chain/node"
executable="${home}/go/src/github.com/axc-chain/node/build/axcchaind"
cli="${home}/go/src/github.com/axc-chain/node/build/axccli"

## clean history data
#rm -r ${home}/.axcchaind_val2
#
## init a witness node
#${executable} init --name val2 --home ${home}/.axcchaind_val2 > ~/init2.json

# config witness node
cp ${home}/.axcchaind/config/genesis.json ${home}/.axcchaind_val2/config/

sed -i -e "s/26/30/g" ${home}/.axcchaind_val2/config/config.toml
sed -i -e "s/6060/10060/g" ${home}/.axcchaind_val2/config/config.toml

# get validator id
validator_pid=$(ps aux | grep "axcchaind start$" | awk '{print $2}')
validatorStatus=$(${cli} status)
validatorId=$(echo ${validatorStatus} | grep -o "\"id\":\"[a-zA-Z0-9]*\"" | sed "s/\"//g" | sed "s/id://g")
#echo ${validatorId}

# set witness peer to validator and start witness
sed -i -e "s/persistent_peers = \"\"/persistent_peers = \"${validatorId}@127.0.0.1:26656\"/g" ${home}/.axcchaind_val2/config/config.toml
sed -i -e "s/index_all_tags = false/index_all_tags = true/g" ${home}/.axcchaind_val2/config/config.toml
${executable} start --home ${home}/.axcchaind_val2 > ${home}/.axcchaind_val2/log.txt 2>&1 &
validator_pid=$!
echo ${validator_pid}
