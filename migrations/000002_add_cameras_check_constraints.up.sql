--alter database mycameraapp owner to mycameraapp;
ALTER TABLE cameras
    ADD CONSTRAINT дmodel_check CHECK (model IS NOT NULL);

ALTER TABLE cameras
    ADD CONSTRAINT resolution_check CHECK (resolution IS NOT NULL);

ALTER TABLE cameras
    ADD CONSTRAINT weight_check CHECK (weight >= 0);

ALTER TABLE cameras
    ADD CONSTRAINT zoom_check CHECK (zoom >= 0);

--migrate -path=./migrations -database=$MYCAMERAAPP_DB_DSN up
-- migrate -path=./migrations -database=$EXAMPLE_DSN down
--go run ./cmd/api
-- curl localhost:4000/v1/cameras/1
--source $HOME/.profile
--BODY='{"Name":"Film Camera","Model":"Nikon FM2","Resolution":"35mm", "Weight":650.0"}'
--curl -X PUT -d "$BODY" localhost:4000/v1/cameras