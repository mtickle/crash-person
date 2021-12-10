-- Table: public.crash

-- DROP TABLE IF EXISTS public.crash;

CREATE TABLE IF NOT EXISTS public.crash
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1000 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    crash_key text COLLATE pg_catalog."default",
    city text COLLATE pg_catalog."default",
    drivers_license_class text COLLATE pg_catalog."default",
    drivers_license_restrictions text COLLATE pg_catalog."default",
    drivers_license_state text COLLATE pg_catalog."default" NOT NULL,
    commercial_drivers_license text COLLATE pg_catalog."default" NOT NULL,
    age text COLLATE pg_catalog."default",
    vehicle_seizure text COLLATE pg_catalog."default",
    alcohol_suspected text COLLATE pg_catalog."default",
    alcohol_testtext COLLATE pg_catalog."default",
    alcohol_result_type text COLLATE pg_catalog."default",
    airbag_switch text COLLATE pg_catalog."default",
    airbag_deployed text COLLATE pg_catalog."default",
    ejection text COLLATE pg_catalog."default",
    gender text COLLATE pg_catalog."default",
    race text COLLATE pg_catalog."default",
    injury text COLLATE pg_catalog."default",
    protection text COLLATE pg_catalog."default",
    trapped COLLATE pg_catalog."default",
    person_type text COLLATE pg_catalog."default",
    vision_obstruction text COLLATE pg_catalog."default",
    contributing_circumstance_1 text COLLATE pg_catalog."default",
    contributing_circumstance_2 text COLLATE pg_catalog."default",
    contributing_circumstance_3 text COLLATE pg_catalog."default",
    vehicletype text COLLATE pg_catalog."default",
    date_of_crash date,
    CONSTRAINT crash_key UNIQUE (crash_key)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.crash
    OWNER to pi;

GRANT ALL ON TABLE public.crash TO pi;