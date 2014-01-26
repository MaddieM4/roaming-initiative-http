#!/bin/bash

DATA_PATH=/var/www-data/roaming-initiative-http/
LOG_PATH=/var/log/ri-http

cd $DATA_PATH

./roaming-initiative-http --fcgi > $LOG_PATH
