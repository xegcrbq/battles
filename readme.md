для работы требуется файл 
.env
в корневой папке проекта

Для создания бд необходимо запустить db/init.sql

```
DB_URL_SQLX="url до базы данных"
DB_DRIVER_NAME_SQLX=postgres
logging=debug
jwtKey="ключ для шифрования cookie"
```

Рабочие ссылки:

1)http://127.0.0.1:8080/api/login?publicaddress=address

2)http://127.0.0.1:8080/api/portfolio/