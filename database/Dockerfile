# bd mysql
FROM mysql:8.0

# copa de init de bd
COPY init.sql /docker-entrypoint-initdb.d/

# usuario 'mysql' para evitar problemas de permisos.
RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

# puerto
EXPOSE 3306
