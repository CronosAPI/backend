package main

import (
	"backend/internal/parse"
	"backend/internal/types"
	"fmt"
)

func main() {
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

}
