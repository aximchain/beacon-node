#!/bin/bash

# ./show.sh -l ADA_ --from alice

chain_id=$CHAIN_ID

while true ; do
    case "$1" in
        -l|--list-pair )
            pair=$2
            shift 2
        ;;
		*)
            break
        ;;
    esac
done;

./axccli dex show -l $pair
