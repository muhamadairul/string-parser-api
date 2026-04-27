DROP TABLE IF EXISTS "public"."parsed_results";
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS parsed_results_id_seq;

-- Table Definition
CREATE TABLE "public"."parsed_results" (
    "id" int8 NOT NULL DEFAULT nextval('parsed_results_id_seq'::regclass),
    "name" bpchar(30) NOT NULL,
    "age" bpchar(3) NOT NULL,
    "city" bpchar(20) NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);


-- Indices
CREATE INDEX idx_parsed_results_deleted_at ON public.parsed_results USING btree (deleted_at);

