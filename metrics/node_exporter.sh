#!/bin/sh

ARGS="--path.rootfs=/host --collector.processes"

if [ -e /host/var/run/dbus/system_bus_socket ]; then
	ARGS+="--collector.systemd"
fi

exec /bin/node_exporter $ARGS
