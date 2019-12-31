-- Table: pooble.stored_product

-- DROP TABLE pooble.stored_product;

CREATE TABLE pooble.stored_product
(
    id integer NOT NULL,
    storer_id integer,
    product_id integer,
    entry_date date,
    output_date date,
    CONSTRAINT id_stored_product PRIMARY KEY (id),
    CONSTRAINT fk_id_product FOREIGN KEY (product_id)
        REFERENCES pooble.product (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_id_storer FOREIGN KEY (storer_id)
        REFERENCES pooble.storer (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE pooble.stored_product
    OWNER to postgres;