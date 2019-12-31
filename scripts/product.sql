-- Table: pooble.product

-- DROP TABLE pooble.product;

CREATE TABLE pooble.product
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    price double precision,
    due_date date,
    brand_name character varying COLLATE pg_catalog."default",
    barcode character varying COLLATE pg_catalog."default",
    is_vegan boolean,
    measurement_value double precision,
    measurement_code character varying COLLATE pg_catalog."default",
    CONSTRAINT product_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE pooble.product
    OWNER to postgres;