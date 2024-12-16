#!/bin/sh

# can't use -i here, we can't create temporary files in /meilisync...
sed '
s/${MEILISEARCH_DB_PASSWORD}/'"$MEILISEARCH_DB_PASSWORD"'/;
' /meilisync/config.yml.template > /tmp/config.yml
mv /tmp/config.yml /meilisync/config.yml
rm -f /tmp.config.yml

exec meilisync start
