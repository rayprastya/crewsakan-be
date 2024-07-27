--liquibase formatted sql

--changeset crewsakan:1
CREATE TABLE Topic (
    -- cannot use auto_increment since it use pgsql so use serial instead, also not null not needed, and no need to specify the data type
    ID SERIAL PRIMARY KEY,
    Name varchar(255) NOT NULL
);

--changeset crewsakan:2
ALTER TABLE Topic ADD COLUMN Description TEXT;


--changeset crewsakan:3
INSERT INTO Topic (Name, Description) Values ('Test', 'This is a test topic');


--changeset crewsakan:4
ALTER TABLE Topic 
ADD Status varchar(15) DEFAULT 'Active',
ADD Created_At TIMeSTAMP DEFAULT CURRENT_TIMESTAMP,
ADD Updated_At TIMeSTAMP DEFAULT CURRENT_TIMESTAMP,
ADD Deleted_At TIMeSTAMP DEFAULT CURRENT_TIMESTAMP;

--changeset crewsakan:5
ALTER TABLE Topic RENAME TO Topics


--changeset crewsakan:6
CREATE TABLE Menu (
    -- cannot use auto_increment since it use pgsql so use serial instead, also not null not needed, and no need to specify the data type
    ID SERIAL PRIMARY KEY,
    Name varchar(255) NOT NULL
)

--changeset crewsakan:7
ALTER TABLE Menu RENAME TO Menus

--changeset crewsakan:8
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL
)

--changeset crewsakan:9
ALTER TABLE Menus
ADD user_id int REFERENCES Users(id)