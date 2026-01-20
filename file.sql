CREATE TABLE users (
	user_id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email varchar(100) UNIQUE,
	password varchar(256) NOT NULL,
	token varchar(256)
)

CREATE TABLE payment_methods (
	payment_id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	company VARCHAR(100) NOT NULL
)

CREATE TABLE films (
	film_id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	genre VARCHAR(100) NOT NULL,
	language VARCHAR(100) NOT NULL,
	duration_minute INTEGER  NOT NULL,
	release_date DATE  NOT NULL,
	rating FLOAT  NOT NULL,
	review_count INTEGER  NOT NULL,
	storyline TEXT  NOT NULL,
	status VARCHAR(30)  NOT NULL,
	image_url TEXT  NOT NULL
)

CREATE TABLE cinemas (
	cinema_id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	film_id INT NOT NULL,
	CONSTRAINT filmID
		FOREIGN KEY (film_id)
		REFERENCES films(film_id),
	time TIMESTAMPTZ  NOT NULL,
	capacity INTEGER  NOT NULL,
	available INTEGER  NOT NULL,
	price NUMERIC NOT NULL
)

CREATE TABLE seats (
	seat_id SERIAL PRIMARY KEY,
	cinema_id INT NOT NULL,
	CONSTRAINT cinemaID
		FOREIGN KEY (cinema_id)
		REFERENCES cinemas(cinema_id),
	status BOOL
)

CREATE TABLE bookings (
	booking_id SERIAL PRIMARY KEY,
	cinema_id INT NOT NULL,
	CONSTRAINT cinemaID
		FOREIGN KEY (cinema_id)
		REFERENCES cinemas(cinema_id),
	seat_id INT NOT NULL,
	CONSTRAINT seatID
		FOREIGN KEY (seat_id)
		REFERENCES seats(seat_id),
	user_id INT NOT NULL,
	CONSTRAINT userID
		FOREIGN KEY (user_id)
		REFERENCES users(user_id),
	payment_id INT NOT NULL,
	CONSTRAINT paymentID
		FOREIGN KEY (payment_id)
		REFERENCES payment_methods(payment_id),
	created_at TIMESTAMPTZ NOT NULL
)