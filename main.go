package main

import(
	"probe-challenge/database"
	"probe-challenge/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
	// Initialize the database connection
	if err := database.Connect(); err != nil {
		panic(err)
	}
	
	// Create a new gin router
	router := gin.Default()
	routes.SetupProbeRoutes(router)

	// Start the server
	router.Run(":8080") // Listen and serve on port 8080
}