package Music

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql" //aliased so we don't depend on driver's functions
)

var db *sqlx.DB

func init(){
	db, err := sql.Open("mysql", "root:admin@/Spa")

	err = db.Ping()

	if (err != nil){
		fmt.Println("No se pudo conectar.")
	}
}

func InsertAlbum(title string, price float, rating int) {
	insertQuery := "INSERT INTO Albums VALUES (?, ?, ?)"

	db.Exec(insertQuery, title, price, rating)
}

func SelectAlbum(albumId int) (album *Album){
	selectQuery := "SELECT title, price, rating FROM Albums WHERE album_id = ?"
	row := db.QueryRow(selectQuery, albumId)

	err := row.StructScan(&album)

	if (err != nil){
		fmt.Println("No pudo encontrarse.")
	}
}

func UpdateAlbum(albumId int, title string, price float, rating int) bool {
	updateQuery := "UPDATE Albums SET title = ?, price = ?, rating = ? WHERE album_id = ?"
	
	_, err := db.Exec(updateQuery, title, price, rating, albumId)

	if (err != nil){
		return false
	}

	return true
}

func DeleteAlbum(albumId int) {
	deleteQuery := "DELETE FROM Albums WHERE album_id = ?"
	db.Exec(deleteQuery, albumId)
}

