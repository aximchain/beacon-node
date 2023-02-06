#!/usr/bin/env bash

kill -9 $(ps aux | grep "axcchaind start" | grep -v "grep" | awk '{print $2}')
kill -9 $(ps aux | grep "ordergen" | grep -v "grep" | awk '{print $2}')