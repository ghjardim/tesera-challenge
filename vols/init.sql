CREATE DATABASE customerdb;

CREATE TABLE customer
(id INTEGER PRIMARY KEY,
firstname TEXT,
lastname TEXT,
email TEXT NOT NULL UNIQUE,
password TEXT);
