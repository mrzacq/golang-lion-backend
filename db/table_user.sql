CREATE TABLE users (
	id serial PRIMARY KEY,
	full_name VARCHAR ( 255 ) NOT NULL,
	username VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 255 ) NOT NULL,
	gender VARCHAR ( 255 ) NOT NULL,
	created_on TIMESTAMP NOT null DEFAULT CURRENT_TIMESTAMP
);