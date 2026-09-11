#!/bin/sh

set -e

# Первый этап: 
# Генерим ключи и сертификаты

openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 365 -key ca.key -out ca.crt \
  -subj "/C=RU/ST=State/L=City/O=MyDevCA/CN=MyDevCA"

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr \
  -subj "/C=RU/ST=State/L=City/O=MyServer/CN=localhost"

openssl x509 -req -days 365 -in server.csr \
  -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt \
  -extfile san.cnf -extensions v3_req
rm -rf server.csr ca.srl

# Второй этап: 
# Настройка логирования: пусть будет (пока что) большой log.txt с логами между запусками.
# Пусть при каждом запуске по скрипту создаётся временный файл для логов, а при закрытии сервера этот файл копирует себя в log.txt и удаляется.
# Потом можно сделать не log.txt а архивацию и чистку.

# Третий этап: 
# Идём в директорию main.go и собираем бинарь сервера

cd serverTLSHashedPassword && go build -o srv main.go && cd ..

echo "Запускаем сервер"

cd serverTLSHashedPassword && nohup ./srv > ../log/log.txt 2>&1 &

# Четвёртый этап:
# При graceful shutdown (и не только) надо:
# - Удалить файлы сертификатов и всякой бурмалды (потом это поменяем);
# - Логи из временного файла скопировать в log.txt;

# Ещё (ПОТОМ!) можно закидывать сертификаты куда-нибудь в /, и при запуске сервера не генерить их если уже есть;
# Читать сертификаты с клиента тоже из /;
