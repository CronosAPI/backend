package main

import (
	"backend/internal/api/space/isro"
	"backend/internal/database"
	"backend/internal/parse"
)

func main() {
	var Client = database.InitializeConnection()

	response1 := parse.GrabAPI(isro.API_ISRO_SPACECRAFTS)
	response2 := parse.GrabAPI(isro.API_ISRO_LAUNCHERS)

	isro.InsertData_ISRO_Spacecrafts(Client, response1)
	isro.GetAllValues_ISRO_Spacecrafts(Client, true)

	isro.InsertData_ISRO_Launchers(Client, response2)
	isro.GetAllValues_ISRO_Launchers(Client, true)

}
