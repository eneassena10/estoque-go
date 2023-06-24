package sqlite3_repository

const (
	QUERY_SELECT_ALL_PRODUCT   = "SELECT id_product, name, price, quantidade FROM products;"
	QUERY_SELECT_BY_ID_PRODUCT = "SELECT id_product, name, price, quantidade FROM products WHERE id_product=?;"
	QUERY_UPDATE_COUNT_PRODUCT = "UPDATE products SET quantidade=? WHERE id_product=?;"
	QUERY_DELETE_PRODUCT       = "DELETE FROM products WHERE id_product=?;"
	QUERY_CREATE_PRODUCT       = "INSERT INTO products (name, price, quantidade) VALUES(?, ?, ?);"
)
