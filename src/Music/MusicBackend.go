package Music

import (
	"fmt"
	"sqlx"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	AlbumId int
	Title string
	Price float
	Rating int
}