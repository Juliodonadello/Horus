CREATE USER dashboard WITH PASSWORD 'dashboard';

CREATE DATABASE events;

GRANT ALL PRIVILEGES ON DATABASE events TO dashboard;

\connect "events";

CREATE TABLE IF NOT EXISTS mm_events (
    id INT PRIMARY KEY,
    clasif_seg CHAR(10) NOT NULL,
    precedencia CHAR(10) NOT NULL,
    cifrado BOOLEAN NOT NULL,
    destino VARCHAR(50) NOT NULL,
    origen VARCHAR(50) NOT NULL,
    evento VARCHAR(50) NOT NULL,
    gfh TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);