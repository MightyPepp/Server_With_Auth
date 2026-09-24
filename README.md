# Server_With_Auth

## Автоматизация

Теперь сервис запускается через `docker run` посредством сборки контейнера из образа на основе Dockerfile'а. К папке logs корня проекта примонтировали том для логов, снова поменяли структуру проекта, избавились от скриптов.

## Структура проекта

.
├── cmd
│   └── api
│       ├── configs
│       │   └── config.yml
│       ├── init.go - это убрать
│       └── main.go - запуск приложения
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── db - функции, подключающие сервер к бд
│   ├── handlers - обработчики ручек
│   │   ├── authHandler.go
│   │   └── structures.go
│   ├── logging - настройка логирования
│   └── server - регистрация обработчиков и middleware
├── logs
│   └── log.txt
├── README.md
└── self_signed
    ├── certs
    │   ├── ca.crt
    │   └── server.crt
    └── private
        ├── ca.key
        └── server.key