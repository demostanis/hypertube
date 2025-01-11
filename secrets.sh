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
pwgen FORWARD_AUTH_COOKIE_SECRET
pwgen MEILI_MASTER_KEY
pwgen MEILISEARCH_DB_PASSWORD

if ! grep -q TMDB_API_KEY "$env_file"; then
	echo missing TMDB_API_KEY >&2
	echo please generate one at https://www.themoviedb.org/settings/api >&2
	exit 1
fi
