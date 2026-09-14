#!/bin/bash

cd /etc/ssl/self_signed

openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 365 -key ca.key -out ca.crt \
  -subj "/C=RU/ST=State/L=City/O=MyDevCA/CN=MyDevCA"

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr \
  -subj "/C=RU/ST=State/L=City/O=MyServer/CN=localhost"

openssl x509 -req -days 365 -in server.csr \
  -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt \
  -extfile /home/mighty-pepe/Desktop/Server_With_Auth/configs/san.cnf -extensions v3_req

mv ca.key ./private
mv server.key ./private
mv ca.crt ./certs
mv server.crt ./certs

rm -rf server.csr ca.srl