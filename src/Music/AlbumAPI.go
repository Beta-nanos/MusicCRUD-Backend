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

func byteConverterToFloat(data []byte) float32{
	var number float32
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &number)

	if err != nil {
		return 0
	}

	return number
}

func byteConverterToInt(data []byte) int{
	var number int
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &number)

	if err != nil {
		return 0
	}

	return number
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
	price := albumHandler.FormValue("price")
	rating := albumHandler.FormValue("rating")
	
	InsertAlbum(string(title), byteConverterToFloat(price), byteConverterToInt(rating))

	println(string(title))
	println("Put from /Albums")
}

//POST /Albums/:params
func (albumHandler AlbumAPI) PostBy(id string) {
	albumId, _ := strconv.Atoi(id)
	title := albumHandler.FormValue("title")
	price := albumHandler.FormValue("price")
	rating := albumHandler.FormValue("rating")
	
	UpdateAlbum(albumId, string(title), byteConverterToFloat(price), byteConverterToInt(rating))
	
	println(string(title))
	println("Post from /Albums/" + id)
}

//DELETE /Albums/:paramId
func (albumHandler AlbumAPI) DeleteBy(id string){
	albumId, _ := strconv.Atoi(id)
	
	DeleteAlbum(albumId)

	println("Delete from /" + id)
}