-- Table: rblock.customer

-- DROP TABLE rblock.customer;

CREATE TABLE rblock.customer
(
    customer_id bigint NOT NULL DEFAULT nextval('rblock.seq_customer_id'::regclass),
    customer_email_c character varying(162) COLLATE pg_catalog."default" NOT NULL,    --encrypted
    customer_password_h character varying(75) COLLATE pg_catalog."default" NOT NULL,  --hashed
    customer_reg_key character varying(50) COLLATE pg_catalog."default",
    customer_created timestamp with time zone NOT NULL,
    customer_activated timestamp with time zone,  -- if NULL then not activated
    CONSTRAINT customer_id_pk PRIMARY KEY (customer_id),
    CONSTRAINT customer_email_uq UNIQUE (customer_email_c),
    CONSTRAINT customer_reg_key_uq UNIQUE (customer_reg_key)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE rblock.customer
    OWNER to postgres;

-- Index: customer_id_indx

-- DROP INDEX rblock.customer_id_indx;

CREATE UNIQUE INDEX customer_id_indx
    ON rblock.customer USING btree
    (customer_id)
    TABLESPACE pg_default;

-- Index: customer_email_indx

-- DROP INDEX rblock.customer_email_indx;

CREATE UNIQUE INDEX customer_email_indx
  ON rblock.customer USING btree
  (customer_email_c)
TABLESPACE pg_default;