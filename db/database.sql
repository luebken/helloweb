DROP TABLE IF EXISTS USER;

CREATE TABLE USER(
    email TEXT NOT NULL UNIQUE,
    name TEXT
);

INSERT INTO USER(email, name) values ('test@test.de', 'Test Test');
