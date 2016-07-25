package Music

import (
	"strconv"
	"github.com/kataras/iris"
	"strings"
)

type AlbumAPI struct {
	*iris.Context
}

//GET /:parameter id
func (albumHandler AlbumAPI) GetBy(id string){
	albumId, _ := strconv.Atoi(id)
	albumHandler.Write("Get from /Albums/%s", id)
	album := SelectAlbum(albumId)
	println(string(album.Title))
	albumHandler.JSON(iris.StatusOK, album)
}

//PUT /Albums
func (albumHandler AlbumAPI) PutBy(title string) {
	comando := strings.Split(title, "&")
	//price := albumHandler.FormValueString(comando[1])
	//rating := albumHandler.FormValueString(comando[2])
	
	fPrice, _ := strconv.ParseFloat(comando[1], 32)
	iRating, _ := strconv.Atoi(comando[2])

	InsertAlbum(comando[0], float32(fPrice), iRating)

	println(string(title))
	println("Put from /Albums"+comando[1])
}


//POST /Albums/:params
func (albumHandler AlbumAPI) PostBy(id string) {
	albumId, _ := strconv.Atoi(id)
	title := albumHandler.FormValue("title")
	price := albumHandler.FormValueString("price")
	rating := albumHandler.FormValueString("rating")
		
	fPrice, _ := strconv.ParseFloat(price, 32)
	iRating, _ := strconv.Atoi(rating)

	UpdateAlbum(albumId, string(title), float32(fPrice), iRating)
	
	println(string(title))
	println("Post from /Albums/" + id)
}

//DELETE /Albums/:paramId
func (albumHandler AlbumAPI) DeleteBy(id string){
	albumId, _ := strconv.Atoi(id)
	
	DeleteAlbum(albumId)

	println("Delete from /" + id)
}