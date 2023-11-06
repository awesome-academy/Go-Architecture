DROP DATABASE IF EXISTS animal;
CREATE DATABASE animal;
USE animal;
CREATE TABLE animals(
                        id int NOT NULL AUTO_INCREMENT,
                        name varchar(50),
                        age int,
                        PRIMARY KEY(id));
INSERT INTO animals VALUES(1,'Hippo',10);
INSERT INTO animals VALUES(2,'Ele',20);