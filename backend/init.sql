CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE devices
(
    id            SERIAL PRIMARY KEY,
    serial_number VARCHAR(255) NOT NULL,
    model         VARCHAR(255) NOT NULL,
    address       VARCHAR(255) NOT NULL
);

CREATE TABLE telemetry
(
    id          SERIAL PRIMARY KEY,
    device_id   INT REFERENCES devices (id),
    timestamp   TIMESTAMP NOT NULL,
    temperature FLOAT     NOT NULL,
    humidity    FLOAT     NOT NULL
);

-- Добавление тестовых данных
INSERT INTO users (username, password)
VALUES ('admin', 'admin123');
INSERT INTO devices (serial_number, model, address)
VALUES ('12345', 'Model A', 'Address 1'),
       ('67890', 'Model B', 'Address 2'),
       ('11223', 'Model C', 'Address 3');
INSERT INTO telemetry (device_id, timestamp, temperature, humidity)
VALUES (1, '2023-01-01 12:00:00', 25.5, 60.0),
       (1, '2023-01-01 13:00:00', 26.0, 61.0);