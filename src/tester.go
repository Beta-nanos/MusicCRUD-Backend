package main

import (
	"fmt"
	"Music"
	"github.com/kataras/iris"
)

/*type Album struct {
	AlbumId int
	Title string
	Price float32
	Rating int
}*/

func main() {
	//dbOps := Music.NewMusicDataOperations();
	//var db *sqlx.DB
	//db, err := sqlx.Connect("mysql", "user=root password=admin dbname=Spa sslmode=disable")
	//Music.Init();
	/*Music.InsertAlbum("Ruido Blanco", 9.99, 3);
	Music.InsertAlbum("Soda Stereo", 4.99, 4);
	Music.InsertAlbum("Me Veras Volver", 12.99, 5);*/

	//Music.UpdateAlbum(2, "Sueño Stereo", 4.4, 3)
	//album := Music.SelectAlbum(2);
	//fmt.Println(album.Title);

	/*if Music.DeleteAlbum(5) {
		fmt.Println("Soda Stereo borrado.")
	}else {
		fmt.Println("Didn't even exist")
	}*/
	//dbOps.Test = 0;
	iris.API("/Albums", Music.AlbumAPI{})
	iris.Listen(":8080")
	fmt.Println("Server mounted")
}