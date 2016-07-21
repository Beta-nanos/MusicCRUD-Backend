package Music

import (
	"fmt"
	_ "github.com/jmoiron/sqlx" //check alias
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //aliased so we don't depend on driver's functions
)

type User struct {
	AlbumId int
	Title string
	Price float
	Rating int
}