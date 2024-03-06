#!/bin/bash

#Demo script fie for updating site once deployed on a web server
#using supervisor


git pull
soda migrate

go build -o bookings cmd/web/*.go

sudo supervisorctl stop book
sudo supervisorctl start book


