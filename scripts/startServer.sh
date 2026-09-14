#!/bin/sh

# Первый этап: 
# Генерим ключи и сертификаты если их нет

/home/mighty-pepe/Desktop/Server_With_Auth/scripts/existanceCrtsAndKeys.sh

# Второй этап: 
# Настройка логирования внутри кода сервера, пока что просто перед каждым запуском чистим файл логов

if [ -e /home/mighty-pepe/Desktop/Server_With_Auth/server/logs/log.txt ] && [ -s /home/mighty-pepe/Desktop/Server_With_Auth/server/logs/log.txt ]; then
  > /home/mighty-pepe/Desktop/Server_With_Auth/logs/log.txt
fi

# Третий этап: 
# Идём в директорию main.go и собираем бинарь сервера

cd /home/mighty-pepe/Desktop/Server_With_Auth/internal/cmd/api 

if [ -e srv ]; then
  rm srv
fi

/usr/local/go/bin/go build -o srv main.go 

echo "Запускаем сервер"

nohup ./srv > /home/mighty-pepe/Desktop/Server_With_Auth/logs/log.txt 2>&1 &