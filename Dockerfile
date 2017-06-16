FROM postgres:alpine

ADD query.sql /queries/query.sql
ADD rsql /usr/local/bin/

WORKDIR /queries
