FROM mariadb:10.5
ENV MARIADB_ROOT_PASSWORD="abc123"
ADD ./seed.sql /docker-entrypoint-initdb.d/seed.sql
ADD ./my.cnf /etc/mysql/my.cnf
EXPOSE 3306
