CREATE USER events_api WITH PASSWORD 'ooMuUa44xBsiLcbicTX8';
CREATE USER dashboard WITH PASSWORD  'mDZrO1XKudNeTUq0MPuE';

CREATE DATABASE grafana;
CREATE DATABASE events;

GRANT ALL PRIVILEGES ON DATABASE events TO events_api;
GRANT ALL PRIVILEGES ON DATABASE grafana TO dashboard;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO dashboard;

\connect "events";

CREATE TABLE IF NOT EXISTS mm_events (
    id SERIAL PRIMARY KEY,
    nro_mm INT NOT NULL,
    clasif_seg CHAR(10) NOT NULL,
    precedencia CHAR(10) NOT NULL,
    cifrado BOOLEAN NOT NULL,
    destino VARCHAR(50) NOT NULL,
    origen VARCHAR(50) NOT NULL,
    evento VARCHAR(50) NOT NULL,
    gfh TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

GRANT ALL PRIVILEGES ON TABLE mm_events TO events_api;
GRANT ALL ON SEQUENCE mm_events_id_seq to events_api;