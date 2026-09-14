#!/bin/bash

set -e

# Порядок параметров CLI:
# ca.key ca.crt server.key server.crt

cd /etc/ssl/self_signed

CAkey="private/${1:-ca.key}"
CAcrt="certs/${2:-ca.crt}"
SRVkey="private/${3:-server.key}"
SRVcrt="certs/${4:-server.crt}"

if ! [ -e "$CAkey" ] || ! [ -e "$CAcrt" ] || ! [ -e "$SRVkey" ] || ! [ -e "$SRVcrt" ]; then
    rm -f $CAkey $CAcrt $SRVkey $SRVcrt
    echo "Удалили всё, чтобы заново нагенерить!"
    /home/mighty-pepe/Desktop/Server_With_Auth/server/scripts/genKeys.sh
else 
    echo "Все файлы ключей и сертификатов на месте!"
fi