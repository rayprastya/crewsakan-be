--liquibase formatted sql

--changeset crewsakan:1
CREATE TABLE Merchants (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    lat double precision,
    lng double precision
);

CREATE TABLE Menus (
    id SERIAL PRIMARY KEY,
    merchant_id int REFERENCES Merchants(id),
    name varchar(255) NOT NULL,
    description varchar(255),
    price double precision
);

CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    menu_id int REFERENCES Menus(id),
    user_id varchar(255),
    optional varchar(255),
    lat double precision,
    lng double precision
);

CREATE TABLE Wishlists (
    id SERIAL PRIMARY KEY,
    merchant_id int REFERENCES Merchants(id),
    name varchar(255) NOT NULL,
    recipes varchar(255),
    steps varchar(255)
);