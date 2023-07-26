
-- users definition
CREATE TABLE users (
	id_user INT PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	nickname TEXT UNIQUE,
	password TEXT,
	logado int
);

-- products definition
CREATE TABLE products(
   id_product INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT,
   price DECIMAL,
   count int
);

