#!/bin/bash

DATA_PATH=/var/www-data/roaming-initiative-http/

cd $DATA_PATH

uwsgi -w application.main -s localhost:5555
