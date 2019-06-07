#!/bin/sh

# Start MySQL connector
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://devdataservicesdebezium-connect.marathon.l4lb.thisdcos.directory:8083/connectors/ -d @register-mysql.json

