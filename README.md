# Bilar Åland

https://bil.edonalds.com/

Create custom RSS feeds for search terms on cars listed for car sale listings on Åland.

## Requirements
- Go 1.22.2
- [templ 0.2.524](https://templ.guide/)
- docker

## Running locally
Start the test database

```
docker build . -t bilkoll-mariadb 
docker run -d --name bilkoll-mariadb -p 3306:3306 bilkoll-mariadb
```

Then run with env variables:

```
cd src

MYSQL_HOST=localhost \
MYSQL_PWD=abc123 \
templ generate --watch --cmd 'go run .'
 ```
