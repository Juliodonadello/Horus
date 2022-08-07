CREATE USER api_events WITH PASSWORD 'ooMuUa44xBsiLcbicTX8';
CREATE USER dashboard WITH PASSWORD  'mDZrO1XKudNeTUq0MPuE';
CREATE USER api_conf WITH PASSWORD '724KR24H7MXFBczNrhgD';

CREATE DATABASE grafana;
CREATE DATABASE events;
CREATE DATABASE api_horus;

GRANT ALL PRIVILEGES ON DATABASE events TO api_events;
GRANT ALL PRIVILEGES ON DATABASE grafana TO dashboard;
GRANT ALL PRIVILEGES ON DATABASE api_horus to api_conf;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO dashboard;

\connect "events";

CREATE TABLE IF NOT EXISTS mm_events (
    id SERIAL PRIMARY KEY,
    nro_mm INT NOT NULL,
    clasif_seg CHAR(15) NOT NULL,
    precedencia CHAR(15) NOT NULL,
    cifrado BOOLEAN NOT NULL,
    destino VARCHAR(50) NOT NULL,
    origen VARCHAR(50) NOT NULL,
    evento VARCHAR(50) NOT NULL,
    gfh TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
GRANT ALL PRIVILEGES ON TABLE mm_events TO api_events;
GRANT ALL ON SEQUENCE mm_events_id_seq to api_events;
GRANT SELECT ON TABLE mm_events TO dashboard;

CREATE TABLE IF NOT EXISTS sensor_events (
    id SERIAL PRIMARY KEY,
    facilidad VARCHAR(50) NOT NULL,
    evento VARCHAR(50) NOT NULL,
    valor BOOLEAN NOT NULL,
    gfh TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE sensor_events
    ADD CONSTRAINT sensor_events_uq
    UNIQUE (facilidad, evento);

GRANT ALL PRIVILEGES ON TABLE sensor_events TO api_events;
GRANT ALL ON SEQUENCE sensor_events_id_seq to api_events;
GRANT SELECT ON TABLE sensor_events TO dashboard;

CREATE TABLE IF NOT EXISTS gps_events(
    id SERIAL PRIMARY KEY,
    facilidad VARCHAR(50) NOT NULL,
    latitude FLOAT,
    longitude FLOAT,
    altitude FLOAT,
    gfh TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

GRANT ALL PRIVILEGES ON TABLE gps_events TO api_events;
GRANT ALL ON SEQUENCE gps_events_id_seq TO api_events;
GRANT SELECT ON TABLE gps_events TO dashboard;

\connect "api_horus";
CREATE TABLE IF NOT EXISTS device_tokens (
    id SERIAL PRIMARY KEY,
    token CHAR(32) NOT NULL,
    device VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);
GRANT ALL PRIVILEGES ON TABLE device_tokens TO api_conf;
GRANT ALL ON SEQUENCE device_tokens_id_seq to api_conf;
