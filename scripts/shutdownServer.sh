#!/bin/bash

declare -r srvPID="$(pgrep srv)"
kill $srvPID
rm /home/mighty-pepe/Desktop/Server_With_Auth/server/internal/cmd/api/srv
echo "Сервер остановлен"