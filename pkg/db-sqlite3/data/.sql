
-- users definition
CREATE TABLE users (
	id_user INT PRIMARY KEY,
	name TEXT,
	nickname TEXT UNIQUE,
	password TEXT,
	logado int
);

-- products definition
CREATE TABLE products(
   id_product INTEGER PRIMARY KEY,
   name TEXT,
   price DECIMAL,
   quantidade int
);

