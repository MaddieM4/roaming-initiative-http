#!/bin/bash

# nginx config
cp config/nginx-site /etc/nginx/sites-available/ri-http
ln -s /etc/nginx/sites-available/ri-http /etc/nginx/sites-enabled/ri-http

/etc/init.d/nginx restart

# golang server
cp config/initrc.sh /etc/init.d/ri-http
chmod +x /etc/init.d/ri-http

/etc/init.d/ri-http restart
