package main

import (
	"backend/internal/database"
	"backend/internal/parse"
	"backend/internal/types"
	"fmt"
)

func main() {
	database.CreateDatabase()

	// space_crafts := parse.GrabAPI("https://isro.vercel.app/api/spacecrafts")
	// build_space_crafts := types.CreateSpaceCraft(space_crafts)
	// fmt.Println(build_space_crafts)

	// launchers := parse.GrabAPI("https://isro.vercel.app/api/launchers")
	// build_launchers := types.CreateLauncher(launchers)
	// fmt.Println(build_launchers)

	// satellite := parse.GrabAPI("https://isro.vercel.app/api/customer_satellites")
	// build_sattelites := types.CreateSattelite(satellite)
	// fmt.Println(build_sattelites)

	centre := parse.GrabAPI("https://isro.vercel.app/api/centres")
	build_centre := types.CreateCentre(centre)
	fmt.Println(build_centre)

	// database.InsertIntoTable("centres", build_centre)
	// database.QueryTable("centres")

	launchers := parse.GrabAPI("https://isro.vercel.app/api/launchers")
	build_launcher := types.CreateLauncher(launchers)
	// fmt.Println(build_launcher)

	database.InsertIntoTable("launchers", build_launcher)
	database.QueryTable("launchers")

}
