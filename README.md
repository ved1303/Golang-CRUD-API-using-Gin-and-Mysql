# Golang-CRUD-API-using-Gin-and-Mysql


Mysql Database setup:

use mydb;

create table products(
id int auto_increment PRIMARY KEY,
name varchar(20),
description varchar(100),
price decimal(8,3) 
);
