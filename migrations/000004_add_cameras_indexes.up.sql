CREATE INDEX IF NOT EXISTS cameras_title_idx ON cameras USING GIN (to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS cameras_genres_idx ON cameras USING GIN (genres);