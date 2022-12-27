#!/bin/bash

### Allow Local Network connections
if [ -n "${LOCAL_NETWORK}" ]; then
    echo "Variable LOCAL_NETWORK defined!"
    echo "Therefore adding ip-range as follows:"
    eval "$(ip r l | grep -v 'tun0\|kernel' | grep default |awk '{print "GW="$3"\nINT="$5}')"
    execute="ip route add $LOCAL_NETWORK via $GW dev $INT"
    echo "$execute"
    eval $execute
    echo "DONE!"
    echo ""
fi

### Start mullvad daemon
exec mullvad-daemon
