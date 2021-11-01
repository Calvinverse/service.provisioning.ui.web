#!/bin/sh

# `/sbin/setuser memcache` runs the given command as the user `memcache`.
# If you omit that part, the command will be run as root.
echo "Starting provisioning service"
/usr/bin/bff server --config $PROVISION_CONFIG_PATH
