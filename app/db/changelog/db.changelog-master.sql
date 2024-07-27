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
    lng double precision,
);

CREATE TABLE Wishlists (
    id SERIAL PRIMARY KEY,
    merchant_id int REFERENCES Merchants(id),
    name varchar(255) NOT NULL,
    recipes varchar(255),
    steps varchar(255)
);

--changeset crewsakan:2
ALTER TABLE Wishlists
ALTER COLUMN recipes TYPE TEXT,
ALTER COLUMN steps TYPE TEXT;

--changeset crewsakan:3
Insert into Merchants (name, lat, lng) values ('Warung mak ijah', 12.3, 3.21);

--changeset crewsakan:4
ALTER TABLE Orders
ADD COLUMN order_datetime TIMESTAMP DEFAULT NOW();

--changeset crewsakan:5
INSERT INTO Merchants (name, lat, lng) VALUES
    ('Merchant 1', 12.34, 56.78),
    ('Merchant 2', 13.45, 57.89),
    ('Merchant 3', 14.56, 58.90),
    ('Merchant 4', 15.67, 59.01),
    ('Merchant 5', 16.78, 60.12),
    ('Merchant 6', 17.89, 61.23),
    ('Merchant 7', 18.90, 62.34),
    ('Merchant 8', 19.01, 63.45),
    ('Merchant 9', 20.12, 64.56),
    ('Merchant 10', 21.23, 65.67);

--changeset crewsakan:6
-- Insert dummy data into the Menus table

INSERT INTO Menus (merchant_id, name, description, price) VALUES
    (1, 'Beef Burger', 'A juicy beef burger with lettuce, tomato, and cheese.', 8.99),
    (2, 'Chicken Sandwich', 'A grilled chicken sandwich with mayonnaise and pickles.', 7.49),
    (3, 'Veggie Pizza', 'A delicious veggie pizza with bell peppers, onions, and olives.', 9.99),
    (4, 'Pasta Carbonara', 'Creamy pasta carbonara with bacon and parmesan cheese.', 11.49),
    (5, 'Caesar Salad', 'A classic Caesar salad with croutons and Caesar dressing.', 6.99),
    (6, 'Sushi Platter', 'An assorted sushi platter with salmon, tuna, and avocado rolls.', 14.99),
    (7, 'Taco Combo', 'Three soft tacos with your choice of beef, chicken, or veggie.', 10.49),
    (8, 'Steak Dinner', 'A tender steak dinner with mashed potatoes and green beans.', 19.99),
    (9, 'Fish and Chips', 'Crispy fish and chips served with tartar sauce.', 12.49),
    (10, 'Pad Thai', 'Traditional Pad Thai with shrimp, peanuts, and bean sprouts.', 13.99);
