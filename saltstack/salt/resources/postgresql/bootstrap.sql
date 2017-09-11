CREATE USER plugdj WITH LOGIN PASSWORD 'plugdj';
CREATE DATABASE plugdj OWNER plugdj;
\connect plugdj;
CREATE TABLE laters (
    "id" serial,
    "name" text,
    PRIMARY KEY ("id")
);
ALTER TABLE laters OWNER TO plugdj;
