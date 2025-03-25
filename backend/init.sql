-- Создание таблицы пользователей
CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Создание таблицы устройств
CREATE TABLE devices
(
    id            SERIAL PRIMARY KEY,
    serial_number VARCHAR(255) NOT NULL,
    model         VARCHAR(255) NOT NULL,
    address       VARCHAR(255) NOT NULL
);

-- Создание таблицы телеметрии
CREATE TABLE telemetry
(
    id          SERIAL PRIMARY KEY,
    device_id   INT REFERENCES devices (id),
    timestamp   TIMESTAMP NOT NULL,
    temperature FLOAT     NOT NULL,
    humidity    FLOAT     NOT NULL
);

-- Добавление тестового пользователя
INSERT INTO users (username, password)
VALUES ('admin', 'admin123');

-- Добавление 3 устройств
INSERT INTO devices (serial_number, model, address)
VALUES ('12345', 'Model A', 'Address 1'),
       ('67890', 'Model B', 'Address 2'),
       ('11223', 'Model C', 'Address 3');

-- Добавление 100 записей телеметрии
DO $$
DECLARE
device_id INT;
    start_time TIMESTAMP := '2023-01-01 00:00:00';
    i INT;
BEGIN
    -- Получаем ID устройств
SELECT id INTO device_id FROM devices WHERE serial_number = '12345';

-- Генерация 100 записей телеметрии для первого устройства
FOR i IN 1..100 LOOP
        INSERT INTO telemetry (device_id, timestamp, temperature, humidity)
        VALUES (
            device_id,
            start_time + (i || ' hours')::INTERVAL,
            20 + RANDOM() * 10,  -- Температура от 20 до 30
            50 + RANDOM() * 20   -- Влажность от 50 до 70
        );
END LOOP;

    -- Получаем ID второго устройства
SELECT id INTO device_id FROM devices WHERE serial_number = '67890';

-- Генерация 100 записей телеметрии для второго устройства
FOR i IN 1..100 LOOP
        INSERT INTO telemetry (device_id, timestamp, temperature, humidity)
        VALUES (
            device_id,
            start_time + (i || ' hours')::INTERVAL,
            20 + RANDOM() * 10,  -- Температура от 20 до 30
            50 + RANDOM() * 20   -- Влажность от 50 до 70
        );
END LOOP;

    -- Получаем ID третьего устройства
SELECT id INTO device_id FROM devices WHERE serial_number = '11223';

-- Генерация 100 записей телеметрии для третьего устройства
FOR i IN 1..100 LOOP
        INSERT INTO telemetry (device_id, timestamp, temperature, humidity)
        VALUES (
            device_id,
            start_time + (i || ' hours')::INTERVAL,
            20 + RANDOM() * 10,  -- Температура от 20 до 30
            50 + RANDOM() * 20   -- Влажность от 50 до 70
        );
END LOOP;
END $$;