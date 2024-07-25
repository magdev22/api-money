-- db/migrations/create_Db.sql
    -- +goose Up

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(20),
    usersurname VARCHAR(20),
    bill INT
);

   
