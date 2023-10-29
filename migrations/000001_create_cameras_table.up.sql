CREATE TABLE IF NOT EXISTS cameras(
    Id bigserial PRIMARY KEY,
    Created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    Name text NOT NULL,
    Model text NOT NULL,
    Resolution text NOT NULL,
    Weight real NOT NULL,
    Zoom real NOT NULL,
    Version integer NOT NULL DEFAULT 1
    );

--alter database mycameraapp owner to mycameraapp;
--env:MYCAMERAAPP_DB_DSN = "postgres://mycameraapp:mycamera@localhost/mycameraapp?sslmode=disable"
--migrate -path=./migrations -database=$MYCAMERAAPP_DB_DSN up