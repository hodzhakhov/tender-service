--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Debian 16.4-1.pgdg120+1)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: organization_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);


ALTER TYPE public.organization_type OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: bid; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bid (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    description text,
    status character varying(20) DEFAULT 'CREATED'::character varying,
    tender_id integer,
    organization_id uuid,
    creator_username character varying(50) NOT NULL,
    version integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bid OWNER TO postgres;

--
-- Name: bid_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bid_id_seq OWNER TO postgres;

--
-- Name: bid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bid_id_seq OWNED BY public.bid.id;


--
-- Name: bid_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bid_versions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    bid_id integer,
    name character varying(50) NOT NULL,
    description text,
    status character varying(20) DEFAULT 'CREATED'::character varying,
    tender_id integer,
    organization_id uuid,
    creator_username character varying(50) NOT NULL,
    version integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bid_versions OWNER TO postgres;

--
-- Name: bid_votes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bid_votes (
    id integer NOT NULL,
    bid_id integer,
    username character varying(50) NOT NULL,
    decision boolean,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bid_votes OWNER TO postgres;

--
-- Name: bid_votes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bid_votes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bid_votes_id_seq OWNER TO postgres;

--
-- Name: bid_votes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bid_votes_id_seq OWNED BY public.bid_votes.id;


--
-- Name: employee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employee (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying(50) NOT NULL,
    first_name character varying(50),
    last_name character varying(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.employee OWNER TO postgres;

--
-- Name: organization; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organization (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(100) NOT NULL,
    description text,
    type public.organization_type,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.organization OWNER TO postgres;

--
-- Name: organization_responsible; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organization_responsible (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    organization_id uuid,
    user_id uuid
);


ALTER TABLE public.organization_responsible OWNER TO postgres;

--
-- Name: review; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.review (
    id integer NOT NULL,
    bid_id integer,
    username character varying(50) NOT NULL,
    comment text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.review OWNER TO postgres;

--
-- Name: review_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.review_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.review_id_seq OWNER TO postgres;

--
-- Name: review_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.review_id_seq OWNED BY public.review.id;


--
-- Name: tender; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tender (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    description text,
    service_type character varying(50) NOT NULL,
    status character varying(20) DEFAULT 'CREATED'::character varying,
    organization_id uuid,
    creator_username character varying(50) NOT NULL,
    version integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.tender OWNER TO postgres;

--
-- Name: tender_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tender_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tender_id_seq OWNER TO postgres;

--
-- Name: tender_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tender_id_seq OWNED BY public.tender.id;


--
-- Name: tender_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tender_versions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    tender_id integer,
    name character varying(50) NOT NULL,
    description text,
    service_type character varying(50) NOT NULL,
    status character varying(20) DEFAULT 'CREATED'::character varying,
    organization_id uuid,
    creator_username character varying(50) NOT NULL,
    version integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.tender_versions OWNER TO postgres;

--
-- Name: bid id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid ALTER COLUMN id SET DEFAULT nextval('public.bid_id_seq'::regclass);


--
-- Name: bid_votes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_votes ALTER COLUMN id SET DEFAULT nextval('public.bid_votes_id_seq'::regclass);


--
-- Name: review id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review ALTER COLUMN id SET DEFAULT nextval('public.review_id_seq'::regclass);


--
-- Name: tender id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tender ALTER COLUMN id SET DEFAULT nextval('public.tender_id_seq'::regclass);


--
-- Data for Name: bid; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bid (id, name, description, status, tender_id, organization_id, creator_username, version, created_at, updated_at) FROM stdin;
2	big bid	test	CREATED	6	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	1	2024-09-13 14:17:40.491007	2024-09-13 14:17:40.491007
1	big bid	test	PUBLISHED	5	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	4	2024-09-13 13:46:43.743844	2024-09-13 16:53:24.667763
\.


--
-- Data for Name: bid_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bid_versions (id, bid_id, name, description, status, tender_id, organization_id, creator_username, version, created_at, updated_at) FROM stdin;
06adb66d-1245-45a9-961e-1f51d5dc6643	1	big bid	test	CREATED	5	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	1	2024-09-13 13:46:43.743844	2024-09-13 13:46:43.743844
23f4dd54-be08-40ff-ac9b-cae3bc59ce47	1	biiig biiid	test	CREATED	5	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	2	2024-09-13 13:46:43.743844	2024-09-13 16:52:26.410243
d4b2bcfa-68e4-4694-8721-d2ec421204d5	1	biiig biiid aaa	wow biid	CREATED	5	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	3	2024-09-13 13:46:43.743844	2024-09-13 16:52:44.472496
f2f18f0c-c3bc-49db-9e52-4863f76b30f3	1	big bid	test	CREATED	5	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	4	2024-09-13 13:46:43.743844	2024-09-13 16:53:24.667763
0558b492-66d8-4f18-9716-43ef321ee897	2	big bid	test	CREATED	6	6273247c-fbdb-4828-93cb-f94edd71da6f	mgkeda	1	2024-09-13 14:17:40.491007	2024-09-13 14:17:40.491007
\.


--
-- Data for Name: bid_votes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bid_votes (id, bid_id, username, decision, created_at) FROM stdin;
1	1	c4stom	t	2024-09-13 15:11:35.938009
2	1	m2ma	t	2024-09-13 15:15:40.248948
3	1	ghost	t	2024-09-13 15:15:50.245405
\.


--
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employee (id, username, first_name, last_name, created_at, updated_at) FROM stdin;
4724d825-7b38-4eb6-9f9d-056b9996a4c5	mgkeda	Gleb	Matveev	2024-09-12 19:21:08.648837	2024-09-12 19:21:08.648837
a737b4a0-c884-4e23-893e-04583e4b3d44	c4stom	Emil	Khodzhakhov	2024-09-12 19:21:08.648837	2024-09-12 19:21:08.648837
3a9968db-8869-445f-b5a2-11acfc264f25	m2ma	Maks	Markov	2024-09-12 19:21:08.648837	2024-09-12 19:21:08.648837
3e790cad-50e0-494d-b73f-61823daf0ff3	welty	Ramil	Amirov	2024-09-12 19:21:08.648837	2024-09-12 19:21:08.648837
728c1fa9-aa15-4c6b-a56b-328864c4b58c	ghost	Dima	Onshin	2024-09-12 19:21:08.648837	2024-09-12 19:21:08.648837
\.


--
-- Data for Name: organization; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organization (id, name, description, type, created_at, updated_at) FROM stdin;
6273247c-fbdb-4828-93cb-f94edd71da6f	Yandex	It giant	LLC	2024-09-11 20:21:52.753282	2024-09-11 20:21:52.753282
d184c167-6bb3-4563-9f01-cbb49ead740d	T-Bank	Aka tinkoff	JSC	2024-09-11 20:22:13.870051	2024-09-11 20:22:13.870051
afeb01d4-7942-4554-aa10-7cbbc4c78706	MTC	Egg	IE	2024-09-11 20:22:29.690561	2024-09-11 20:22:29.690561
\.


--
-- Data for Name: organization_responsible; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organization_responsible (id, organization_id, user_id) FROM stdin;
1b75bfce-efb1-4c49-9c49-6543a745f7e0	6273247c-fbdb-4828-93cb-f94edd71da6f	4724d825-7b38-4eb6-9f9d-056b9996a4c5
780f8e6d-f408-4bd4-b640-a14e14b0ee13	d184c167-6bb3-4563-9f01-cbb49ead740d	a737b4a0-c884-4e23-893e-04583e4b3d44
d70cbd41-16ed-4445-ab2e-a6a180f6c17e	afeb01d4-7942-4554-aa10-7cbbc4c78706	3e790cad-50e0-494d-b73f-61823daf0ff3
fd66963d-56c6-448c-90da-10abd113f76b	d184c167-6bb3-4563-9f01-cbb49ead740d	3a9968db-8869-445f-b5a2-11acfc264f25
2d3f167c-df51-4cba-97c9-a67dbbfd08f4	d184c167-6bb3-4563-9f01-cbb49ead740d	728c1fa9-aa15-4c6b-a56b-328864c4b58c
\.


--
-- Data for Name: review; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.review (id, bid_id, username, comment, created_at) FROM stdin;
1	1	c4stom	super good	2024-09-13 13:55:08.063228
\.


--
-- Data for Name: tender; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tender (id, name, description, service_type, status, organization_id, creator_username, version, created_at, updated_at) FROM stdin;
6	Tender2	Very big tender	Construction	CREATED	afeb01d4-7942-4554-aa10-7cbbc4c78706	welty	1	2024-09-13 14:17:14.227012	2024-09-13 14:17:14.227012
5	Tender1	Very big tender	Construction	CANCELED	d184c167-6bb3-4563-9f01-cbb49ead740d	c4stom	1	2024-09-13 13:42:04.985938	2024-09-13 13:42:04.985938
\.


--
-- Data for Name: tender_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tender_versions (id, tender_id, name, description, service_type, status, organization_id, creator_username, version, created_at, updated_at) FROM stdin;
76574011-dc20-426b-9a7f-9d0aa293e1d4	5	Tender1	Very big tender	Construction	CREATED	d184c167-6bb3-4563-9f01-cbb49ead740d	c4stom	1	2024-09-13 13:42:04.985938	2024-09-13 13:42:04.985938
2d6be192-f09c-4cbb-8542-fe1650ae9dfb	6	Tender2	Very big tender	Construction	CREATED	afeb01d4-7942-4554-aa10-7cbbc4c78706	welty	1	2024-09-13 14:17:14.227012	2024-09-13 14:17:14.227012
\.


--
-- Name: bid_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bid_id_seq', 2, true);


--
-- Name: bid_votes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bid_votes_id_seq', 3, true);


--
-- Name: review_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.review_id_seq', 1, true);


--
-- Name: tender_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tender_id_seq', 6, true);


--
-- Name: bid bid_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid
    ADD CONSTRAINT bid_pkey PRIMARY KEY (id);


--
-- Name: bid_versions bid_versions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_versions
    ADD CONSTRAINT bid_versions_pkey PRIMARY KEY (id);


--
-- Name: bid_votes bid_votes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_votes
    ADD CONSTRAINT bid_votes_pkey PRIMARY KEY (id);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


--
-- Name: employee employee_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_username_key UNIQUE (username);


--
-- Name: organization organization_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- Name: organization_responsible organization_responsible_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_responsible
    ADD CONSTRAINT organization_responsible_pkey PRIMARY KEY (id);


--
-- Name: review review_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT review_pkey PRIMARY KEY (id);


--
-- Name: tender tender_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tender
    ADD CONSTRAINT tender_pkey PRIMARY KEY (id);


--
-- Name: tender_versions tender_versions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tender_versions
    ADD CONSTRAINT tender_versions_pkey PRIMARY KEY (id);


--
-- Name: bid bid_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid
    ADD CONSTRAINT bid_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: bid bid_tender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid
    ADD CONSTRAINT bid_tender_id_fkey FOREIGN KEY (tender_id) REFERENCES public.tender(id) ON DELETE CASCADE;


--
-- Name: bid_versions bid_versions_bid_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_versions
    ADD CONSTRAINT bid_versions_bid_id_fkey FOREIGN KEY (bid_id) REFERENCES public.bid(id) ON DELETE CASCADE;


--
-- Name: bid_versions bid_versions_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_versions
    ADD CONSTRAINT bid_versions_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: bid_versions bid_versions_tender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_versions
    ADD CONSTRAINT bid_versions_tender_id_fkey FOREIGN KEY (tender_id) REFERENCES public.tender(id) ON DELETE CASCADE;


--
-- Name: bid_votes bid_votes_bid_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bid_votes
    ADD CONSTRAINT bid_votes_bid_id_fkey FOREIGN KEY (bid_id) REFERENCES public.bid(id) ON DELETE CASCADE;


--
-- Name: organization_responsible organization_responsible_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_responsible
    ADD CONSTRAINT organization_responsible_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: organization_responsible organization_responsible_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_responsible
    ADD CONSTRAINT organization_responsible_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.employee(id) ON DELETE CASCADE;


--
-- Name: review review_bid_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT review_bid_id_fkey FOREIGN KEY (bid_id) REFERENCES public.bid(id) ON DELETE CASCADE;


--
-- Name: tender tender_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tender
    ADD CONSTRAINT tender_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: tender_versions tender_versions_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tender_versions
    ADD CONSTRAINT tender_versions_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: tender_versions tender_versions_tender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tender_versions
    ADD CONSTRAINT tender_versions_tender_id_fkey FOREIGN KEY (tender_id) REFERENCES public.tender(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

