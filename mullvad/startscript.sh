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

# Create symlink to avoid libgmodule-2.0.so.0 not found message
ln -s /usr/lib/x86_64-linux-gnu/libgm2.so.0 /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0

exec /usr/bin/xterm &
exec '/opt/Mullvad VPN/mullvad-gui' &
# exec '/opt/Mullvad VPN/mullvad-gui' %U
# exec '/opt/Mullvad VPN/mullvad-vpn'

# Work to infinity
echo "Here the infinity begins ..."
tail -f /dev/null
