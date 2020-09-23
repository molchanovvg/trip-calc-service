# trip-calc-service
**Установка приложения**

**Конфигурация приложения**

Для локального запуска приложения можно отредактировать файл .env .
Содержит следующие ключи:
- REDIS_URL = адрес redis 
- SERVICE_PORT = порт на котором нужно запустить приложение
- CALC_ROUTE_URL = url для запроса на вычисления маршрута

Для запуска приложения в докере, можно отредактировать файл Dockerfile, ключи ENV такие же.

**Запуск приложения локально**

**Запуск приложения docker**

**Пример работы**


 docker build -t trip-calc-service .
 docker run -d -p 8080:8080 trip-calc-service
 
 docker run -d -p 6379:6379 --name redis
    

start docker-compose up


stop docker compose stop


ps -ef | grep redis

Для Graceful Shutdown можно выполнить команду
kill SIGTERM {PID}
