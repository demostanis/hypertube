#!/bin/sh

env_file=.env

pwgen() {
	pw=$(tr -dc '[:alnum:]' < /dev/urandom | head -c 32)

	if ! grep -q "$1"= "$env_file"; then
		echo "$1"="$pw" >> "$env_file"
	fi
}

pwgen HYPERTUBE_DB_PASSWORD
pwgen KEYCLOAK_DB_PASSWORD
pwgen ROOT_DB_PASSWORD
pwgen KEYCLOAK_FORWARD_AUTH_SECRET
