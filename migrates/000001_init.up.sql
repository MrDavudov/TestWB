CREATE TABLE datatemp (
    id      serial          not null unique,
    city    varchar(255)    not null,
    temp    varchar(30)     not null,
    dt      datetime        not null
);