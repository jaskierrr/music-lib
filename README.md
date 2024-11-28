# Стек
Создал основу API при помощи openAPI Swagger (go-swagger)
БД Postgres
Миграции golang-migrate
Сборка docker-compose
Старался придерживаться Clean Architecture

# Сборка
make build в корневой папке

Применил многоэтапную сборку для golang image

# API
Узнать все эндпоинты можно при помощи https://editor.swagger.io/ скопировав туда swagger.yaml целиком

# URL для внешнего API
В файле .env есть параметр **EXTERNALAPIURL** для установки части строки URL для вызова внешнего API (нужно вставить все что идет до `/info?group=Muse&song=Supermassive%20Black%20Hole`)

# Уровень логирования
В файле .env есть параметр **LOGLEVEL** для установки уровня логирования
