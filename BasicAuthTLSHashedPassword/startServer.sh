#!/bin/sh

#!/bin/sh
set -e

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

cd serverTLSHashedPassword && go build -o srv main.go && cd ..

# mkdir log && touch log/log.txt

echo "Запускаем сервер"
cd serverTLSHashedPassword && nohup ./srv > ../log/log.txt 2>&1 &

# Почитать про синтаксис bash и сделать следующее:
# - Создание директорий если нет;
# - Проброс ошибок;
# - Реакцию на kill PID сервера (в сервере написать graceful shutdown) чтоб удалялись временные файлы;
# Ещё можно закидывать сертификаты куда-нибудь в /, и при запуске сервера не генерить их если уже есть;
# Читать сертификаты с клиента тоже из /;

# rm -rf ca.crt ca.key ca.srl server.csr server.key server.crt