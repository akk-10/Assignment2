CREATE INDEX IF NOT EXISTS cameras_name_idx ON cameras USING GIN (to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS cameras_model_idx ON cameras USING GIN (model);