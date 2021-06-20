FROM mysql:latest
COPY initCandyStorage.sql /docker-entrypoint-initdb.d