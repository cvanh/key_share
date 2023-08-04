--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8 (Ubuntu 14.8-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.8 (Ubuntu 14.8-0ubuntu0.22.04.1)

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
-- Name: user; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA "user";


ALTER SCHEMA "user" OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: key; Type: TABLE; Schema: user; Owner: root
--

CREATE TABLE "user".key (
    user_id uuid,
    fingerprint character varying
);


ALTER TABLE "user".key OWNER TO root;

--
-- Name: user; Type: TABLE; Schema: user; Owner: root
--

CREATE TABLE "user"."user" (
    user_id character varying,
    username character varying
);


ALTER TABLE "user"."user" OWNER TO root;

--
-- Data for Name: key; Type: TABLE DATA; Schema: user; Owner: root
--

COPY "user".key (user_id, fingerprint) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: user; Owner: root
--

COPY "user"."user" (user_id, username) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

