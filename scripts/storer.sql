-- Table: pooble.storer

-- DROP TABLE pooble.storer;

CREATE TABLE pooble.storer
(
    id numeric NOT NULL,
    name character varying COLLATE pg_catalog."default",
    CONSTRAINT id_storer PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE pooble.storer
    OWNER to postgres;