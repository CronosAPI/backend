package main

import (
	"backend/internal/api/space/spacex"
	"backend/internal/database"
	"backend/internal/utils"
)

func main() {
	utils.LoadEnvironmentVariables()
	var Client = database.InitializeConnection()

	// isro.RetrieveAllAndStore_ISRO(Client, true)
	// nasa.RetrieveAllAndStore_NASA(Client, true)
	spacex.RetrieveAllAndStore_SPACEX(Client, true)
}
