--liquibase formatted sql

--changeset rayprastya:1
CREATE TABLE Topic (
    -- cannot use auto_increment since it use pgsql so use serial instead, also not null not needed, and no need to specify the data type
    ID SERIAL PRIMARY KEY,
    Name varchar(255) NOT NULL
);

--changeset rayprastya:2
ALTER TABLE Topic ADD COLUMN Description TEXT;


--changeset rayprastya:3
INSERT INTO Topic (Name, Description) Values ('Test', 'This is a test topic');


--changeset rayprastya:4
ALTER TABLE Topic 
ADD Status varchar(15) DEFAULT 'Active',
ADD Created_At TIMeSTAMP DEFAULT CURRENT_TIMESTAMP,
ADD Updated_At TIMeSTAMP DEFAULT CURRENT_TIMESTAMP,
ADD Deleted_At TIMeSTAMP DEFAULT CURRENT_TIMESTAMP;

--changeset raypras:5
ALTER TABLE Topic RENAME TO Topics
