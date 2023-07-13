package main

import (
	"fmt"
	"strings"
)

const (
	SELECT = "SELECT"
	FROM   = "FROM"
	WHERE  = "WHERE"
	DELETE = "DELETE"
	UPDATE = "UPDADE"
	TABLE  = "TABLE"
	SET    = "SET"
)

type user struct {
	ID       int    `db:"id"`
	Nickname string `db:"nickname"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

var fieldsList = []string{
	"id", "nickname", "password", "name",
}

func getCompare(field string) string {
	result := "%s %s %s %s %s %s=?"
	fields := strings.Join([]string{"id", "password", "nickname"}, ",")
	return fmt.Sprintf(result, SELECT, fields, FROM, "users", WHERE, field)
}

/*
	required
	- nome da entidade
	- lista de nome dos campos da entidade
	- estrutura representativa da entidade
*/

func main() {
	fmt.Println(getCompare("id"), 1)
	// fmt.Println(SELECT, strings.Join([]string{"id_product", "name", "price", "quantidade"}, ","), FROM, "products")
	// fmt.Println(SELECT, strings.Join([]string{"id", "password", "nickname"}, ","), FROM, "users")
	// fmt.Println(SELECT, strings.Join([]string{"id", "password", "nickname"}, ","), FROM, "users", WHERE, "id=?")
}
