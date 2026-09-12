#!/bin/sh

set -e

# Запускать скрипт с sudo и -E флагом!

# Первый этап: 
# Генерим ключи и сертификаты если их нет

/home/mighty-pepe/Desktop/Server_With_Auth/BasicAuthTLSHashedPassword/serverTLSHashedPassword/scripts/existanceCrtsAndKeys.sh

# Второй этап: 
# Настройка логирования: пусть будет (пока что) большой log.txt с логами между запусками.
# Пусть при каждом запуске по скрипту создаётся временный файл для логов, а при закрытии сервера этот файл копирует себя в log.txt и удаляется.
# Потом можно сделать не log.txt а архивацию и чистку.

# Третий этап: 
# Идём в директорию main.go и собираем бинарь сервера

cd /home/mighty-pepe/Desktop/Server_With_Auth/BasicAuthTLSHashedPassword/serverTLSHashedPassword/cmd/api 

if [ -e srv ]; then
  rm srv
fi

/usr/local/go/bin/go build -o srv main.go 

echo "Запускаем сервер"

nohup ./srv > /dev/null 2>&1 &

# Четвёртый этап:
# При graceful shutdown (и не только) надо:
# - Удалить файлы сертификатов и всякой бурмалды (потом это поменяем);
# - Логи из временного файла скопировать в log.txt;

# Ещё (ПОТОМ!) можно закидывать сертификаты куда-нибудь в /, и при запуске сервера не генерить их если уже есть;
# Читать сертификаты с клиента тоже из /;
