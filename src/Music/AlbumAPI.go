package Music

import (
	"strconv"
	"github.com/kataras/iris"
)

type AlbumAPI struct {
	*iris.Context
}

//GET /:parameter id
func (albumHandler AlbumAPI) GetBy(id string){
	albumId, _ := strconv.Atoi(id)
	albumHandler.Write("Get from /Albums/%s", albumId)
	/**
	TO-DO:
		Call to SelectAlbum (MusicDataOps)		
	**/
}

//PUT /Albums
func (albumHandler AlbumAPI) Put() {
	title := albumHandler.FormValue("title")
	price := albumHandler.FormValue("price")
	rating := albumHandler.FormValue("rating")
	/**
	TO-DO:
		Call to InsertAlbum (MusicDataOps)
	**/
	println(string(name))
	println("Put from /Albums")
}

//Post /Albums/:params
func (albumHandler AlbumAPI) PostBy(id string) {
	albumId, _ := strconv.Atoi(id)
	title := albumHandler.FormValue("title")
	price := albumHandler.FormValue("price")
	rating := albumHandler.FormValue("rating")

	/**
	TO-DO:
		Call to UpdateAlbum (MusicDataOps)
	**/
	println(string(title))
	println("Post from /Albums/" + id)
}

//DELETE /Albums/:paramId
func (albumHandler AlbumAPI) DeleteBy(id string){
	albumId, _ := strconv.Atoi(id)
	/**
	TO-DO:
		Call to DeleteAlbum (MusicDataOps)
	**/
	println("Delete from /" + id)
}