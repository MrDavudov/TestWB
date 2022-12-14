CREATE TABLE datatemp (
    id      serial          not null primary key,
    city    varchar(255)    not null,
    temp    varchar(30)     not null,
    dt      date          
);

CREATE UNIQUE INDEX city_dt_idx ON datatemp (city, dt);