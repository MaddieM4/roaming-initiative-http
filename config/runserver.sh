#!/bin/bash

DATA_PATH=/var/www-data/roaming-initiative-http/

cd $DATA_PATH

/usr/local/bin/uwsgi -w application.main -s localhost:5555
