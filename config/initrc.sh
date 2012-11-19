#!/bin/bash

DATA_PATH=/var/www-data/roaming-initiative-http/
VENV_PATH=./develop/bin/activate

WSGI_PIDFILE=/var/run/ri-http.pid
DAEMON="uwsgi -w application.main -s localhost:5555"

case "$1" in
  start)
    echo "Starting server"

    cd $DATA_PATH

    # Activate the virtual environment
    . $VENV_PATH

    # Run start-stop-daemon, the $DAEMON variable contains the path to the
    # application to run
    start-stop-daemon --start --pidfile $WSGI_PIDFILE \
        --user www-data --group www-data \
        --chuid www-data \
        --exec "$DAEMON"
    ;;
  stop)
    echo "Stopping WSGI Application"

    # Start-stop daemon can also stop the application by sending sig 15
    # (configurable) to the process id contained in the run/pid file
    start-stop-daemon --stop --pidfile $WSGI_PIDFILE --verbose
    ;;
  restart)
    /etc/init.d/ri-http stop
    /etc/init.d/ri-http start
    ;;
  *)
    # Refuse to do other stuff
    echo "Usage: /etc/init.d/ri-http {start|stop|restart}"
    exit 1
    ;;
esac

exit 0
