CREATE TABLE users(
    id TEXT PRIMARY KEY  not null,
    firstName VARCHAR(30) not null,
    lastName VARCHAR(30)not null,
    email CHAR(50) not null,
    password TEXT not null,
    createdAt TIMESTAMP,
    updatedAt   TIMESTAMP
);


CREATE TABLE users(
    id TEXT PRIMARY KEY  not null,
    firstName VARCHAR(30) not null,
    lastName VARCHAR(30)not null,
    email CHAR(50) not null,
    password TEXT not null,
    createdAt TIMESTAMP,
    updatedAt   TIMESTAMP
);





`INSERT INTO users(id,firstName,lastName,email,password,createdAt,updatedAt) VALUES($1,$2,$3,$4,$5,$6,$7)`,