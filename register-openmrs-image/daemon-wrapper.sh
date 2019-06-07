#!/bin/bash

trap '{ echo "Received Interupt. Shutting down, bye." ; exit 1; }' INT KILL TERM HUP

while sleep 60; do
  DATETIME=`date`
  echo $DATETIME: Alpine image running..
done
