#!/bin/sh
cat /load-users.sql | mysql -h $MYSQL_HOST -u $MYSQL_USER --password=$MYSQL_PASSWORD --database=$MYSQL_DATABASE
