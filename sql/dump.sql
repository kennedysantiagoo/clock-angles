--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Debian 15.3-1.pgdg120+1)
-- Dumped by pg_dump version 15.3 (Debian 15.3-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: geometry; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA geometry;


ALTER SCHEMA geometry OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: clock_angles; Type: TABLE; Schema: geometry; Owner: postgres
--

CREATE TABLE geometry.clock_angles (
    id integer NOT NULL,
    hour integer,
    minute integer,
    angle integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE geometry.clock_angles OWNER TO postgres;

--
-- Name: clock_angles_id_seq; Type: SEQUENCE; Schema: geometry; Owner: postgres
--

CREATE SEQUENCE geometry.clock_angles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE geometry.clock_angles_id_seq OWNER TO postgres;

--
-- Name: clock_angles_id_seq; Type: SEQUENCE OWNED BY; Schema: geometry; Owner: postgres
--

ALTER SEQUENCE geometry.clock_angles_id_seq OWNED BY geometry.clock_angles.id;


--
-- Name: clock_angles id; Type: DEFAULT; Schema: geometry; Owner: postgres
--

ALTER TABLE ONLY geometry.clock_angles ALTER COLUMN id SET DEFAULT nextval('geometry.clock_angles_id_seq'::regclass);


--
-- Data for Name: clock_angles; Type: TABLE DATA; Schema: geometry; Owner: postgres
--

COPY geometry.clock_angles (id, hour, minute, angle, created_at) FROM stdin;
1	6	0	180	2023-08-04 10:48:18.857091
2	9	0	90	2023-08-05 15:23:19.797139
3	2	40	160	2023-08-05 15:28:13.177689
4	3	0	90	2023-08-05 15:55:11.622045
5	3	30	75	2023-08-05 16:02:23.852566
6	3	10	35	2023-08-05 21:10:01.180064
\.


--
-- Name: clock_angles_id_seq; Type: SEQUENCE SET; Schema: geometry; Owner: postgres
--

SELECT pg_catalog.setval('geometry.clock_angles_id_seq', 6, true);


--
-- Name: clock_angles clock_angles_pkey; Type: CONSTRAINT; Schema: geometry; Owner: postgres
--

ALTER TABLE ONLY geometry.clock_angles
    ADD CONSTRAINT clock_angles_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

