#!/bin/bash

declare -r srvPID="$(pgrep srv)"
kill $srvPID
echo "Сервер остановлен"