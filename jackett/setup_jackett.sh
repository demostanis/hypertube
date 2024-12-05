#!/bin/sh

echo hello from setup_jackett.sh

login_url=http://localhost:9117/UI/Login'?cookiesChecked=1'
cookies=$(curl "$login_url" -sD- | \
	grep Set-Cookie | awk -F'[:;]' '{print$2}')

curl() {
	command curl -sH 'Cookie: '"$cookies" "$@"
}

enabled_indexers='[
	"cpasbien",
	"torrent9",
	"oxtorrent",
	"zetorrents"
]'

indexers=$(curl http://localhost:9117/api/v2.0/indexers | \
	jq -j '.[] | . as $f |
		select('"$enabled_indexers"' |
		any(. == $f.id)) |
		.id, " ", .site_link, "\n"')

mkdir -p /config/Jackett/Indexers/
echo "$indexers" | while read i; do
	name=${i%% *}
	url=${i##* }

	cat >/config/Jackett/Indexers/$name.json <<-EOF
	[
		{
			"id": "sitelink",
			"type": "inputstring",
			"name": "Site Link",
			"value": "$url"
		}
	]
EOF
done
chmod 755 -R /config/Jackett/Indexers/

# restart jackett so it reads /config
s6-svc -r /run/service/svc-jackett

sleep infinity
