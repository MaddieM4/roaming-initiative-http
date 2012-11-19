#!/bin/bash

DATA_PATH=/var/www-data/roaming-initiative-http/
VENV_PATH=./develop/bin/activate

cd $DATA_PATH

# Activate the virtual environment
. $VENV_PATH

/usr/local/bin/uwsgi -w application.main -s localhost:5555
