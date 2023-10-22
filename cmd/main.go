package main

import "backend/internal/database"

func main() {
	// database.CreateDatabase()

	// space_crafts := parse.GrabAPI("https://api.spacexdata.com/v3/payloads")
	// pretty.Println(string(space_crafts))
	// build_space_crafts := types.CreateDonki(space_crafts)

	// pretty.Println(build_space_crafts)

	// database.InsertIntoTable("spacex_payloads", build_space_crafts)
	database.QueryTable("satellites")

}
