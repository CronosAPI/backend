package main

import (
	"backend/internal/api/space/isro"
	"backend/internal/database"
	"backend/internal/parse"
)

func main() {
	var Client = database.InitializeConnection()

	response1 := parse.GrabAPI(isro.API_ISRO_SATELLITES)
	response2 := parse.GrabAPI(isro.API_ISRO_CENTERS)

	isro.InsertData_ISRO_Satellites(Client, response1)
	isro.GetAllValues_ISRO_Satellites(Client, true)

	isro.InsertData_ISRO_Centers(Client, response2)
	isro.GetAllValues_ISRO_Centers(Client, true)

}
