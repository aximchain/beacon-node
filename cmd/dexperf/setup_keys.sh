#!/usr/bin/env bash

for i in {0..1000}
do
 ./add_key.exp node0_user$i taxccli /home/test/.axcli/
done
