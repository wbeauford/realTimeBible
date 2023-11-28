package main

func initializeRoutes() {
	router.GET("/", renderIndex)
	router.POST("/getBooks", renderBooks)
}
