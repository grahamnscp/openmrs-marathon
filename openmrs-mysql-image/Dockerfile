FROM mysql:5.7

RUN chown -R mysql:root /var/lib/mysql/

ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_ROOT_PASSWORD

ENV MYSQL_DATABASE=openmrs
ENV MYSQL_USER=mysqluser
ENV MYSQL_PASSWORD=mysqlpw
ENV MYSQL_ROOT_PASSWORD=root

COPY mysqld.cnf /etc/mysql/conf.d/mysqld.cnf

ADD debezium.sql /etc/mysql/debezium.sql
ADD demo.sql /etc/mysql/demo.sql

RUN sed -i 's/MYSQL_DATABASE/openmrs/g' /etc/mysql/debezium.sql
RUN sed -i 's/MYSQL_DATABASE/openmrs/g' /etc/mysql/demo.sql
RUN cp /etc/mysql/debezium.sql /docker-entrypoint-initdb.d
RUN cp /etc/mysql/demo.sql /docker-entrypoint-initdb.d

EXPOSE 3306
