CREATE SEQUENCE rblock.seq_customer_id
    INCREMENT 1
    START 1
    MINVALUE 0
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE rblock.seq_customer_id
    OWNER TO postgres;