CREATE TABLE weather (
    place           varchar(80),
    temp_lo         int,           -- low temperature
    temp_hi         int,           -- high temperature
    prcp            real,          -- precipitation
    date            date
);

CREATE TABLE places (
    name            varchar(80),
    location        point
);

INSERT INTO weather VALUES ('Snøhetta', -16.7, -4.8, 0, '2019-03-16');
INSERT INTO places VALUES ('Snøhetta', '(62.3198, 9.2677)');
INSERT INTO weather VALUES('Glittertind', -11.2, -1.4, 0, '2019-01-01');
INSERT INTO places VALUES('Glittertind', '(61.6512, 8.5575)');

