package Music

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql" //aliased so we don't depend on driver's functions
)

var db *sqlx.DB

type MusicDataOperations struct{
	Test int
}

func init(){
	db = sqlx.MustConnect("mysql", "root:admin@tcp(localhost:3306)/Spa")
	
	err := db.Ping()

	if err != nil {
		fmt.Println("No se pudo conectar.")
	}
}

func NewMusicDataOperations() *MusicDataOperations{
	return new (MusicDataOperations);
}

func InsertAlbum(title string, price float32, rating int) {
	insertQuery := "INSERT INTO Albums (title, price, rating) VALUES (?, ?, ?)"

	tx := db.MustBegin()
	tx.MustExec(insertQuery, title, price, rating)
	tx.Commit();
}

func SelectAll() ( []Album){
	selectQuery := "SELECT * FROM Albums"

	rows, err := db.Queryx(selectQuery)
	
	if err != nil {
		return nil
	}

	albums := make([]Album, 0, 2)

	for rows.Next() {
		var album Album
		err = rows.StructScan(&album)
		albums = append(albums, album)
	}

	return albums
}

func SelectAlbum(albumId int) (album Album){
	selectQuery := "SELECT title, price, rating FROM Albums WHERE album_id = ?"

	err := db.QueryRowx(selectQuery, albumId).StructScan(&album)
	album.AlbumId = albumId
	
	if err != nil {
		fmt.Println(err)
	}

	return
}

func UpdateAlbum(albumId int, title string, price float32, rating int) bool {
	updateQuery := "UPDATE Albums SET title = ?, price = ?, rating = ? WHERE album_id = ?"
	
	tx := db.MustBegin()
	rows := tx.MustExec(updateQuery, title, price, rating, albumId)
	tx.Commit();

	rowsAffected, _ := rows.RowsAffected();

	if rowsAffected == 0 {
		return false
	}

	return true
}

func DeleteAlbum(albumId int) bool{
	deleteQuery := "DELETE FROM Albums WHERE album_id = ?"

	tx := db.MustBegin()
	rows := db.MustExec(deleteQuery, albumId)
	tx.Commit();

	rowsAffected, _ := rows.RowsAffected();

	if rowsAffected == 0 {
		return false
	}

	return true
}

