--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)

-- Started on 2023-08-28 10:53:17 +05

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 16414)
-- Name: client; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.client (
    id integer NOT NULL,
    login character varying(128),
    password character varying(255),
    email character varying(128)
);


ALTER TABLE public.client OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16413)
-- Name: client_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.client_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.client_id_seq OWNER TO postgres;

--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 209
-- Name: client_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.client_id_seq OWNED BY public.client.id;


--
-- TOC entry 213 (class 1259 OID 16435)
-- Name: client_info; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.client_info (
    id integer,
    first_name character varying(128) NOT NULL,
    last_name character varying(128) NOT NULL,
    reg_date date DEFAULT CURRENT_DATE,
    teacher boolean DEFAULT false,
    language_items json DEFAULT '[]'::json,
    about_me character varying(2500) DEFAULT ''::character varying,
    avatar character varying(500) DEFAULT ''::character varying,
    interface_lang character varying(24) DEFAULT 'ru'::bpchar
);


ALTER TABLE public.client_info OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16451)
-- Name: course; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course (
    id integer NOT NULL,
    user_id integer,
    name character varying(200) NOT NULL,
    language character varying(30) DEFAULT 'english'::bpchar,
    main_img character varying(128) DEFAULT ''::bpchar,
    main_video character varying(128) DEFAULT ''::bpchar,
    description character varying(5000) DEFAULT ''::bpchar,
    hidden boolean DEFAULT false,
    price integer DEFAULT 0,
    created_date date DEFAULT CURRENT_DATE,
    editing_date date DEFAULT CURRENT_DATE
);


ALTER TABLE public.course OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16450)
-- Name: course_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.course_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_id_seq OWNER TO postgres;

--
-- TOC entry 3416 (class 0 OID 0)
-- Dependencies: 214
-- Name: course_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.course_id_seq OWNED BY public.course.id;


--
-- TOC entry 217 (class 1259 OID 16473)
-- Name: lesson; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lesson (
    id integer NOT NULL,
    course_id integer,
    name character varying(255) NOT NULL,
    date_created date DEFAULT CURRENT_DATE,
    deadline integer DEFAULT 0,
    date_editing date DEFAULT CURRENT_DATE,
    exercise json DEFAULT '[]'::json,
    task_description text DEFAULT ''::text,
    lesson_number integer NOT NULL
);


ALTER TABLE public.lesson OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16472)
-- Name: lesson_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lesson_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lesson_id_seq OWNER TO postgres;

--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 216
-- Name: lesson_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lesson_id_seq OWNED BY public.lesson.id;


--
-- TOC entry 212 (class 1259 OID 16423)
-- Name: registration_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.registration_token (
    id integer NOT NULL,
    token character varying(64)
);


ALTER TABLE public.registration_token OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16422)
-- Name: registration_token_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.registration_token_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.registration_token_id_seq OWNER TO postgres;

--
-- TOC entry 3418 (class 0 OID 0)
-- Dependencies: 211
-- Name: registration_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.registration_token_id_seq OWNED BY public.registration_token.id;


--
-- TOC entry 3228 (class 2604 OID 16417)
-- Name: client id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client ALTER COLUMN id SET DEFAULT nextval('public.client_id_seq'::regclass);


--
-- TOC entry 3236 (class 2604 OID 16454)
-- Name: course id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course ALTER COLUMN id SET DEFAULT nextval('public.course_id_seq'::regclass);


--
-- TOC entry 3245 (class 2604 OID 16476)
-- Name: lesson id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lesson ALTER COLUMN id SET DEFAULT nextval('public.lesson_id_seq'::regclass);


--
-- TOC entry 3229 (class 2604 OID 16426)
-- Name: registration_token id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.registration_token ALTER COLUMN id SET DEFAULT nextval('public.registration_token_id_seq'::regclass);


--
-- TOC entry 3402 (class 0 OID 16414)
-- Dependencies: 210
-- Data for Name: client; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.client VALUES (1, 'test', '$2a$05$FbIMcSKpDxQ58K8K/eYRE.nosRwg92/v9SoCRpikz8Hb5QWWOKbTK', 'test');
INSERT INTO public.client VALUES (2, 'test1', '$2a$05$mEjYL785MP2k6UnAyZhtyu.q9FUE7JwBqWlIkBVpmQwSzMdh4ekoC', 'test1');
INSERT INTO public.client VALUES (3, 'test12', '$2a$05$KZH7h27jBhL26NOVbcf.T.1BeP49Pr0dOW6uAhSPGyNHMZIoandV2', 'test12');


--
-- TOC entry 3405 (class 0 OID 16435)
-- Dependencies: 213
-- Data for Name: client_info; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.client_info VALUES (3, 'Dmitry', 'Ishkov', '2023-08-08', true, '[{"lang":"en","lvl":"C1"}]', 'fdghdf dfghdf ghdfg dfghdfghdfgh dfghdfgh', '/teachers/teacher_3/avatar/avatar_150120-grafika-graficeskij_dizajn-vyshka-svet-nebesnoe_yavlenie-2560x1440.jpg', 'ru');
INSERT INTO public.client_info VALUES (2, 'Gim', 'Berg', '2023-08-08', false, '[]', '', '', 'ru');


--
-- TOC entry 3407 (class 0 OID 16451)
-- Dependencies: 215
-- Data for Name: course; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.course VALUES (1, 3, 'fdgdfgdfg', 'english', '/teachers/teacher_3/course_1/main/mainImg_f59467ed05a91c23b344b43684e60b2c.jpeg', '/teachers/teacher_3/course_1/main/mainVideo_The Neighbourhood - Sweater Weather (speedboys remix).mp4', 'dfgdfgsdfgsd', false, 32423, '2023-08-28', '2023-08-28');


--
-- TOC entry 3409 (class 0 OID 16473)
-- Dependencies: 217
-- Data for Name: lesson; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.lesson VALUES (2, 1, '2222', '2023-08-28', 0, '2023-08-28', '[{"name":"2222","type":"image","taskDescription":"2222","exerciseNumber":1,"path":"/teachers/teacher_3/course_1/lessons/image/lesson_2_exercise_1_image_cyber_punk.jpg","task":{"type":"","correctAnswer":null,"expandBrackets":{"text":"","answerItems":null}}}]', '', 1);


--
-- TOC entry 3404 (class 0 OID 16423)
-- Dependencies: 212
-- Data for Name: registration_token; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.registration_token VALUES (1, '123123');


--
-- TOC entry 3419 (class 0 OID 0)
-- Dependencies: 209
-- Name: client_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.client_id_seq', 3, true);


--
-- TOC entry 3420 (class 0 OID 0)
-- Dependencies: 214
-- Name: course_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.course_id_seq', 1, true);


--
-- TOC entry 3421 (class 0 OID 0)
-- Dependencies: 216
-- Name: lesson_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lesson_id_seq', 2, true);


--
-- TOC entry 3422 (class 0 OID 0)
-- Dependencies: 211
-- Name: registration_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.registration_token_id_seq', 1, true);


--
-- TOC entry 3252 (class 2606 OID 16421)
-- Name: client client_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_pkey PRIMARY KEY (id);


--
-- TOC entry 3256 (class 2606 OID 16466)
-- Name: course course_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course
    ADD CONSTRAINT course_pkey PRIMARY KEY (id);


--
-- TOC entry 3258 (class 2606 OID 16485)
-- Name: lesson lesson_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lesson
    ADD CONSTRAINT lesson_pkey PRIMARY KEY (id);


--
-- TOC entry 3254 (class 2606 OID 16428)
-- Name: registration_token registration_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.registration_token
    ADD CONSTRAINT registration_token_pkey PRIMARY KEY (id);


--
-- TOC entry 3259 (class 2606 OID 16443)
-- Name: client_info client_info_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client_info
    ADD CONSTRAINT client_info_id_fkey FOREIGN KEY (id) REFERENCES public.client(id);


--
-- TOC entry 3260 (class 2606 OID 16467)
-- Name: course course_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course
    ADD CONSTRAINT course_client_id_fkey FOREIGN KEY (user_id) REFERENCES public.client(id);


--
-- TOC entry 3261 (class 2606 OID 16486)
-- Name: lesson lesson_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lesson
    ADD CONSTRAINT lesson_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.course(id);


-- Completed on 2023-08-28 10:53:17 +05

--
-- PostgreSQL database dump complete
--

