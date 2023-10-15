CREATE TABLE IF NOT EXISTS cameras (
    Id bigserial PRIMARY KEY,
    Created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    Name text NOT NULL,
    Model text NOT NULL,
    Resolution text NOT NULL,
    Weight real NOT NULL,
    Zoom real NOT NULL,
    Version integer NOT NULL DEFAULT 1
    );