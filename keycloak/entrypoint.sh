#!/bin/sh

rm -f /opt/keycloak/ready

ok() {
	( exec 3<>/dev/tcp/localhost/8000
	(
	# imagine having your /health/ready brokenfajsdlfjk
		echo GET /admin/master/console/ HTTP/1.1
		echo Host: keycloak.localhost
		echo
	) >&3
	head -n1 <&3 | grep -q 200 )
}

wait_ok() {
	while ! ok 2>/dev/null; do
		sleep 1
	done
}

if [ -e /opt/keycloak/did-initial-configuration ]; then
	echo Starting optimized Keycloak...
	bin/kc.sh "$@" --optimized &
	wait_ok
	touch /opt/keycloak/ready
	wait
fi

bin/kc.sh "$@" &

wait_ok

conf=$(mktemp)
bin/kcadm.sh config credentials \
	--config "$conf" \
	--server http://localhost:8000 \
	--realm master \
	--user "$KC_BOOTSTRAP_ADMIN_USERNAME" \
	--password "$KC_BOOTSTRAP_ADMIN_PASSWORD" 

kcadm() {
	bin/kcadm.sh "$@" --config "$conf"
}

kcadm create realms -s realm=default -s enabled=true
# TODO: find the URL and pass it to the golang app
kcadm create clients -r default \
	-s clientId='crocotube-auth' \
	-s 'redirectUris=["http://localhost:8000"]' \
	-s publicClient=true \
	-s directAccessGrantsEnabled=true 

FORWARD_AUTH_ID=$(kcadm create clients -r master \
	-s clientId='forward-auth' \
	-s 'redirectUris=["http://jackett.localhost:8000/oauth2/callback", "http://grafana.localhost:8000/oauth2/callback", "http://smtp4dev.localhost:8000/oauth2/callback"]' \
	-s publicClient=false \
	-s "secret=$KC_FORWARD_AUTH_SECRET" --id)

ADMIN_ID=$(kcadm get users -r master -q exact=true -q username=$KC_BOOTSTRAP_ADMIN_USERNAME | grep id | sed -e 's/"id" : "//' -e 's/",//' | xargs)

kcadm update "users/$ADMIN_ID" \
	-s emailVerified=true \
	-s email=$KC_ADMIN_EMAIL

if [ -n "$FORWARD_AUTH_ID" ]; then
	kcadm create "clients/$FORWARD_AUTH_ID/protocol-mappers/models" -r master \
		-s name=aud-mapper-forward-auth \
		-s protocol=openid-connect \
		-s protocolMapper=oidc-audience-mapper \
		-s 'config."included.client.audience"=forward-auth' \
		-s 'config."id.token.claim"=true' \
		-s 'config."access.token.claim"=true'
fi

touch /opt/keycloak/ready
touch /opt/keycloak/did-initial-configuration

wait
