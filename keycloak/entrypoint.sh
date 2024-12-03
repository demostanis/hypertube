#!/bin/sh

bin/kc.sh "$@" &

ok() {
	( exec 3<>/dev/tcp/localhost/8080
	(
	# imagine having your /health/ready brokenfajsdlfjk
		echo GET /admin/master/console/ HTTP/1.1
		echo Host: keycloak.localhost
		echo
	) >&3
	head -n1 <&3 | grep -q 200 )
}

while ! ok 2>/dev/null; do
	sleep 1
done

conf=$(mktemp)
bin/kcadm.sh config credentials \
	--config "$conf" \
	--server http://localhost:8080 \
	--realm master \
	--user "$KC_BOOTSTRAP_ADMIN_USERNAME" \
	--password "$KC_BOOTSTRAP_ADMIN_PASSWORD" 

kcadm() {
	bin/kcadm.sh "$@" --config "$conf"
}

kcadm create realms -s realm=default -s enabled=true
# TODO: find the URL and pass it to the golang app
kcadm create clients -r default \
	-s clientId='hypertube-auth' \
	-s 'redirectUris=["http://localhost:8000"]' \
	-s publicClient=true \
	-s directAccessGrantsEnabled=true 

wait
