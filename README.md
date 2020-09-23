# trip-calc-service
**Установка приложения**

git clone https://github.com/molchanovvg/trip-calc-service.git

go mod vendor 

**Конфигурация приложения**

Для локального запуска приложения можно отредактировать файл .env .
Содержит следующие ключи:
- REDIS_URL = адрес redis 
- SERVICE_PORT = порт на котором нужно запустить приложение
- CALC_ROUTE_URL = url для запроса на вычисления маршрута

Для запуска приложения в докере, можно отредактировать файл Dockerfile, ключи ENV такие же.

**Запуск приложения локально**

Для локального запуска необходим установленный и запущенный https://redis.io/

go run main.go 

или 

go build -o main . && ./main

**Запуск приложения docker**

docker-compose up

**Пример работы**

Для создания запроса на вычисление маршрута:


POST: http://127.0.0.1:8080/trip/calc/request (8080 порт из конфига по умолчанию в ENV)

params:
latitudeFrom
longitudeFrom
latitudeTo
longitudeTo

В ответ будет token.

Для получения обработанной информации о маршруте:

http://127.0.0.1:8080/trip/calc/result?token={token}

Сервер в ответ вернет Distance и Time. Если информация еще не обработана, сервер вернет 425 HTTP ответ.

Для Graceful Shutdown можно выполнить команду
kill SIGTERM {PID} где PID - id запущенного процесса build приложения.
