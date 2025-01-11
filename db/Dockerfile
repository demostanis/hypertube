FROM postgres:17.2-bookworm

RUN apt-get update && \
	apt-get install postgresql-17-cron && \
	apt-get install postgresql-17-wal2json

RUN echo "shared_preload_libraries = 'pg_cron,wal2json'" >> /usr/share/postgresql/postgresql.conf.sample
RUN echo "wal_level = logical" >> /usr/share/postgresql/postgresql.conf.sample
