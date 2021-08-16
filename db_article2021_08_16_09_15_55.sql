--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 13.3

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
-- Name: add_article(text, text, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.add_article(author_param text, title_param text, body_param text) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO article (author, title, body)
    VALUES (author_param,title_param,body_param);
    RETURN TRUE;
END;
$$;


ALTER FUNCTION public.add_article(author_param text, title_param text, body_param text) OWNER TO postgres;

--
-- Name: get_article(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_article() RETURNS TABLE(id integer, author text, title text, body text, created text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT a.id,
                        a.author,
                        a.title,
                        a.body,
                        a.created::text
                 FROM article a
                 ORDER BY a.created DESC;
END;
$$;


ALTER FUNCTION public.get_article() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: article; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.article (
    id integer NOT NULL,
    author text,
    title text,
    body text,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.article OWNER TO postgres;

--
-- Name: article_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.article_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.article_id_seq OWNER TO postgres;

--
-- Name: article_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.article_id_seq OWNED BY public.article.id;


--
-- Name: article id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article ALTER COLUMN id SET DEFAULT nextval('public.article_id_seq'::regclass);


--
-- Data for Name: article; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.article (id, author, title, body, created) FROM stdin;
1	Calvin	Artikel	Artikel baru terbit	2021-08-15 11:03:36.782342
2	Calvin	Artikel 2	Artikel terbit 	2021-08-15 23:51:23.800703
7	Calvin	Artikel 3	Artikel terbit  Bulan 3	2021-08-16 00:06:54.79416
8	Fadhil	Artikel	Artikel baru terbit	2021-08-16 07:01:24.75282
9	Fadhil	Artikel	Artikel baru terbit	2021-08-16 07:02:14.32113
10	Fadhil	Artikel	Artikel baru terbit	2021-08-16 07:04:41.646085
11	Fadhil	Artikel	Artikel baru terbit	2021-08-16 07:05:11.067806
\.


--
-- Name: article_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.article_id_seq', 11, true);


--
-- Name: article article_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

