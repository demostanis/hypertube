#!/bin/sh

env_file=.env

pwgen() {
	pw=$(tr -dc '[:alnum:]' < /dev/urandom | head -c 32)

	if ! grep -q "$1"_DB_PASSWORD= "$env_file"; then
		echo "$1"_DB_PASSWORD="$pw" >> "$env_file"
	fi
}

pwgen HYPERTUBE
pwgen KEYCLOAK
pwgen ROOT
