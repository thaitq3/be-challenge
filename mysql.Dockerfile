FROM mysql
COPY database-script/*.sql /docker-entrypoint-initdb.d/