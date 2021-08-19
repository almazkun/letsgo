package main

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album[
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 59.99}
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99}
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 29.99}
]

func getAlbums(c *gin.Context) {
	c.indentedJSON(http,StatusOK, albums)
}