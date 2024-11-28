#!/bin/sh

env_file=.env

if ! grep -q POSTGRES_PASSWORD= "$env_file"; then
	echo POSTGRES_PASSWORD="$(tr -dc '[:print:]' < /dev/urandom | \
		head -c 32)" >> "$env_file"
fi
