package Music

import (
	"bytes"
	"strconv"
	"encoding/binary"
	"github.com/kataras/iris"
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
func (albumHandler AlbumAPI) Put() {
	title := albumHandler.FormValue("title")
	price := albumHandler.FormValueString("price")
	rating := albumHandler.FormValueString("rating")
	
	fPrice, _ := strconv.ParseFloat(price, 32)
	iRating, _ := strconv.Atoi(rating)

	InsertAlbum(string(title), float32(fPrice), iRating)

	println(string(title))
	println("Put from /Albums")
}


//Put /Albums/:params
func (albumHandler AlbumAPI) Put(title, price, rating string) {
	fPrice, _ := strconv.ParseFloat(price, 32)
	iRating, _ := strconv.Atoi(rating)
	
	InsertAlbum(title, float32(fPrice), iRating)

	println("Put from /Albums/params")	
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