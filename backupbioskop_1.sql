--
-- PostgreSQL database dump
--

\restrict cXjbojpqWYnyLjNXtB8kZSq27er61PKgbcvLBW8v0cjc9T0Or4lV9zuEvJJFrme

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    booking_id integer NOT NULL,
    cinema_id integer NOT NULL,
    seat_id integer NOT NULL,
    user_id integer NOT NULL,
    payment_id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    status boolean NOT NULL
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- Name: bookings_booking_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_booking_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_booking_id_seq OWNER TO postgres;

--
-- Name: bookings_booking_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_booking_id_seq OWNED BY public.bookings.booking_id;


--
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    cinema_id integer NOT NULL,
    name character varying(100) NOT NULL,
    film_id integer NOT NULL,
    "time" timestamp with time zone NOT NULL,
    price numeric NOT NULL
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- Name: cinemas_cinema_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_cinema_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cinemas_cinema_id_seq OWNER TO postgres;

--
-- Name: cinemas_cinema_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_cinema_id_seq OWNED BY public.cinemas.cinema_id;


--
-- Name: films; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.films (
    film_id integer NOT NULL,
    name character varying(100) NOT NULL,
    genre character varying(100) NOT NULL,
    language character varying(100) NOT NULL,
    duration_minute integer NOT NULL,
    release_date date NOT NULL,
    rating double precision NOT NULL,
    review_count integer NOT NULL,
    storyline text NOT NULL,
    status character varying(30) NOT NULL,
    image_url text NOT NULL
);


ALTER TABLE public.films OWNER TO postgres;

--
-- Name: films_film_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.films_film_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.films_film_id_seq OWNER TO postgres;

--
-- Name: films_film_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.films_film_id_seq OWNED BY public.films.film_id;


--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    payment_id integer NOT NULL,
    name character varying(100) NOT NULL,
    company character varying(100) NOT NULL
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- Name: payment_methods_payment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_payment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_payment_id_seq OWNER TO postgres;

--
-- Name: payment_methods_payment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_payment_id_seq OWNED BY public.payment_methods.payment_id;


--
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    seat_id integer NOT NULL,
    cinema_id integer NOT NULL,
    status boolean
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- Name: seats_seat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_seat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_seat_id_seq OWNER TO postgres;

--
-- Name: seats_seat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_seat_id_seq OWNED BY public.seats.seat_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(256) NOT NULL,
    token character varying(256)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: bookings booking_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN booking_id SET DEFAULT nextval('public.bookings_booking_id_seq'::regclass);


--
-- Name: cinemas cinema_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN cinema_id SET DEFAULT nextval('public.cinemas_cinema_id_seq'::regclass);


--
-- Name: films film_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.films ALTER COLUMN film_id SET DEFAULT nextval('public.films_film_id_seq'::regclass);


--
-- Name: payment_methods payment_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN payment_id SET DEFAULT nextval('public.payment_methods_payment_id_seq'::regclass);


--
-- Name: seats seat_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN seat_id SET DEFAULT nextval('public.seats_seat_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bookings (booking_id, cinema_id, seat_id, user_id, payment_id, created_at, status) FROM stdin;
1	1	45	1	1	2026-01-21 15:28:47.541693+07	t
3	4	199	1	2	2026-01-22 14:09:18.928616+07	t
2	1	46	2	2	2026-01-21 23:06:43.787396+07	t
\.


--
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cinemas (cinema_id, name, film_id, "time", price) FROM stdin;
1	Cinema XXI Plaza A	1	2026-01-31 10:00:00+07	35000
2	Cinema XXI Plaza A	2	2026-01-31 13:00:00+07	35000
4	Cinema CGV Mall B	4	2026-01-31 11:30:00+07	40000
5	Cinema Cinepolis C	5	2026-01-31 12:00:00+07	30000
3	Cinema CGV Mall B	3	2026-01-31 11:00:00+07	40000
\.


--
-- Data for Name: films; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.films (film_id, name, genre, language, duration_minute, release_date, rating, review_count, storyline, status, image_url) FROM stdin;
1	The Last Horizon	Action, Sci-Fi	English	135	2025-12-10	8.4	1245	Seorang pilot luar angkasa harus menyelamatkan koloni manusia terakhir dari kehancuran.	Now Showing	https://example.com/images/the_last_horizon.jpg
2	Cinta di Ujung Senja	Romance, Drama	Indonesian	110	2025-11-20	7.9	856	Kisah cinta dua insan yang bertemu kembali setelah terpisah selama bertahun-tahun.	Now Showing	https://example.com/images/cinta_di_ujung_senja.jpg
3	Mystery of Aruna	Mystery, Thriller	English	125	2025-10-05	8.1	980	Seorang detektif mengungkap rahasia gelap di sebuah kota kecil yang misterius.	Now Showing	https://example.com/images/mystery_of_aruna.jpg
4	Petualangan Si Kiko	Animation, Family	Indonesian	95	2025-12-01	8.6	1500	Seekor kucing kecil memulai petualangan besar untuk menemukan rumah barunya.	Now Showing	https://example.com/images/petualangan_si_kiko.jpg
5	Shadow Protocol	Action, Thriller	English	140	2025-09-18	7.8	1120	Agen rahasia terjebak dalam konspirasi global yang mengancam stabilitas dunia.	Now Showing	https://example.com/images/shadow_protocol.jpg
\.


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_methods (payment_id, name, company) FROM stdin;
1	OVO	PT Visionet Internasional
2	GoPay	PT GoTo Gojek Tokopedia Tbk
3	ShopeePay	PT Shopee International Indonesia
4	DANA	PT Espay Debit Indonesia Koe
5	Transfer Bank	Perbankan Indonesia
\.


--
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seats (seat_id, cinema_id, status) FROM stdin;
1	1	t
2	1	t
3	1	t
4	1	t
5	1	t
6	1	t
7	1	t
8	1	t
9	1	t
10	1	t
11	1	t
12	1	t
13	1	t
14	1	t
15	1	t
16	1	t
17	1	t
18	1	t
19	1	t
20	1	t
21	1	f
22	1	f
23	1	f
24	1	f
25	1	f
26	1	f
27	1	f
28	1	f
29	1	f
30	1	f
31	1	f
32	1	f
33	1	f
34	1	f
35	1	f
36	1	f
37	1	f
38	1	f
39	1	f
40	1	f
41	1	f
42	1	f
43	1	f
44	1	f
47	1	f
48	1	f
49	1	f
50	1	f
51	2	t
52	2	t
53	2	t
54	2	t
55	2	t
56	2	t
57	2	t
58	2	t
59	2	t
60	2	t
61	2	t
62	2	t
63	2	t
64	2	t
65	2	t
66	2	f
67	2	f
68	2	f
69	2	f
70	2	f
71	2	f
72	2	f
73	2	f
74	2	f
75	2	f
76	2	f
77	2	f
78	2	f
79	2	f
80	2	f
81	2	f
82	2	f
83	2	f
84	2	f
85	2	f
86	2	f
87	2	f
88	2	f
89	2	f
90	2	f
91	2	f
92	2	f
93	2	f
94	2	f
95	2	f
96	2	f
97	2	f
98	2	f
99	2	f
100	2	f
101	3	t
102	3	t
103	3	t
104	3	t
105	3	t
106	3	t
107	3	t
108	3	t
109	3	t
110	3	f
111	3	f
112	3	f
113	3	f
114	3	f
115	3	f
116	3	f
117	3	f
118	3	f
119	3	f
120	3	f
121	3	f
122	3	f
123	3	f
124	3	f
125	3	f
126	3	f
127	3	f
128	3	f
129	3	f
130	3	f
131	3	f
132	3	f
133	3	f
134	3	f
135	3	f
136	3	f
137	3	f
138	3	f
139	3	f
140	3	f
141	3	f
142	3	f
143	3	f
144	3	f
145	3	f
146	3	f
147	3	f
148	3	f
149	3	f
150	3	f
151	4	t
152	4	t
153	4	t
154	4	t
155	4	t
156	4	t
157	4	t
158	4	t
159	4	t
160	4	t
161	4	t
162	4	t
163	4	t
164	4	t
165	4	t
166	4	t
167	4	t
168	4	t
169	4	t
170	4	t
171	4	t
172	4	t
173	4	t
174	4	t
175	4	t
176	4	t
177	4	t
178	4	t
179	4	t
180	4	t
181	4	t
182	4	t
183	4	t
184	4	t
185	4	t
46	1	t
186	4	t
187	4	t
188	4	t
189	4	f
190	4	f
191	4	f
192	4	f
193	4	f
194	4	f
195	4	f
196	4	f
197	4	f
198	4	f
200	4	f
201	5	t
202	5	t
203	5	t
204	5	t
205	5	t
206	5	t
207	5	t
208	5	t
209	5	t
210	5	t
211	5	t
212	5	t
213	5	t
214	5	t
215	5	t
216	5	t
217	5	t
218	5	t
219	5	t
220	5	t
221	5	t
222	5	t
223	5	t
224	5	t
225	5	t
226	5	t
227	5	t
228	5	t
229	5	t
230	5	t
231	5	f
232	5	f
233	5	f
234	5	f
235	5	f
236	5	f
237	5	f
238	5	f
239	5	f
240	5	f
241	5	f
242	5	f
243	5	f
244	5	f
245	5	f
246	5	f
247	5	f
248	5	f
249	5	f
250	5	f
45	1	t
199	4	t
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (user_id, name, email, password, token) FROM stdin;
1	habibi	habibi@gmail.com	12345678	625c0329-b24e-4b04-9d85-b8a72aa63f46
6	indira	indira@gmail.com	12345678	\N
4	senyum sumringah	senyum@gmail.com	12345678	60cee4e3-51e8-4f62-8417-7cf54e4b6563
2	budi cahyo	budi@gmail.com	12345678	4e00c710-d022-4622-83fc-40cb1bff061e
\.


--
-- Name: bookings_booking_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_booking_id_seq', 3, true);


--
-- Name: cinemas_cinema_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_cinema_id_seq', 5, true);


--
-- Name: films_film_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.films_film_id_seq', 10, true);


--
-- Name: payment_methods_payment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_payment_id_seq', 5, true);


--
-- Name: seats_seat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_seat_id_seq', 250, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_user_id_seq', 6, true);


--
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (booking_id);


--
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (cinema_id);


--
-- Name: films films_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.films
    ADD CONSTRAINT films_pkey PRIMARY KEY (film_id);


--
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (payment_id);


--
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (seat_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: bookings cinemaid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT cinemaid FOREIGN KEY (cinema_id) REFERENCES public.cinemas(cinema_id);


--
-- Name: seats cinemaid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT cinemaid FOREIGN KEY (cinema_id) REFERENCES public.cinemas(cinema_id);


--
-- Name: cinemas filmid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT filmid FOREIGN KEY (film_id) REFERENCES public.films(film_id);


--
-- Name: bookings paymentid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT paymentid FOREIGN KEY (payment_id) REFERENCES public.payment_methods(payment_id);


--
-- Name: bookings seatid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT seatid FOREIGN KEY (seat_id) REFERENCES public.seats(seat_id);


--
-- Name: bookings userid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT userid FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- PostgreSQL database dump complete
--

\unrestrict cXjbojpqWYnyLjNXtB8kZSq27er61PKgbcvLBW8v0cjc9T0Or4lV9zuEvJJFrme

