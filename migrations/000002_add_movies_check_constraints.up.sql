alter database mycameraapp owner to mycameraapp;
ALTER TABLE cameras
    ADD CONSTRAINT model_check CHECK (model IS NOT NULL);

ALTER TABLE cameras
    ADD CONSTRAINT resolution_check CHECK (resolution IS NOT NULL);

ALTER TABLE cameras
    ADD CONSTRAINT weight_check CHECK (weight >= 0);

ALTER TABLE cameras
    ADD CONSTRAINT zoom_check CHECK (zoom >= 0);
