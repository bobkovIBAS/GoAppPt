# Проект: Клиент-Серверное Приложение для Вычисления Числа Фибоначчи
## Описание
Это клиент-серверное приложение, в котором пользователь задает число n, а сервер вычисляет соответствующее число Фибоначчи. Взаимодействие между клиентом и сервером осуществляется через HTTP API. Сервер использует очередь сообщений RabbitMQ для обработки запросов.

## Содержание
 1. [Требования]
 2. [Установка Go]
   - На Windows
   - На Linux
 3. [Запуск проекта]
   - Клонирование репозитория
   - Сборка и запуск с помощью Docker
     
## Требования
 - Docker и Docker Compose установлены на запускаемой системе.
 - Go 1.23.х
   
## Установка Go
  ### Установка Go на Windows
  Скачать установщик:
  Перейдите на официальный сайт Go и скачайте MSI-установщик для Windows.
  
  После установки:
  Откройте командную строку (cmd) или PowerShell и введите:
  `go version`
  Должны увидеть скаченную версию Go

  ### Установка Go на Linux
  - Скачать архив Go:   `wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz`
  - Распаковать архив в /usr/local:
  ```
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz
  ```
  - Настроить переменные окружения: Добавьте следующие строки в ваш ~/.bashrc:
  ```
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
  ```
  - Примените изменения:
    `source ~/.profile`
  - Проверка установки:
    `go version`
  Вы должны увидеть установленную версию Go.

### Запуск проекта
1. Клонирование репозитория
Склонируйте репозиторий проекта и перейди в клонированную директорию:
```
  git clone https://github.com/bobkovIBAS/GoAppPt
  cd GoAppPt
```
2. Сборка и запуск с помощью Docker
- Шаг 1: Сборка Docker-образов
Выполните команду для сборки образов:
```sudo docker-compose build```
- Шаг 2: Запуск контейнеров
```docker-compose up```
