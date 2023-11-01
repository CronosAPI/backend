package isro

import (
	"context"
	"encoding/json"
	"fmt"

	"backend/internal/parse"
	"backend/internal/utils"

	"github.com/go-redis/redis/v8"
)

func InsertData_ISRO_Spacecrafts(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var spacecrafts SpaceCrafts
	if err := json.Unmarshal(data, &spacecrafts); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, spacecraft := range spacecrafts.SpaceCraftArray {
		key := utils.BuildKey(KEY_SPACECRAFT, fmt.Sprint(spacecraft.ID))

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "id", spacecraft.ID)
		pipe.HSet(ctx, key, "name", spacecraft.Name)
		_, err := pipe.Exec(ctx)

		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return err
		}

	}

	return nil
}

func GetAllValues_ISRO_Spacecrafts(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_SPACECRAFT+"*"))

	var spacecraft SpaceCraft

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {

		if err := Client.HGetAll(ctx, key).Scan(&spacecraft); err != nil {
			panic(err)
		}
		if print {
			fmt.Printf("%d : %s,\n", spacecraft.ID, spacecraft.Name)
		}
	}
	if print {
		fmt.Println("}")
	}

	return nil
}

func InsertData_ISRO_Launchers(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var launchers Launchers
	if err := json.Unmarshal(data, &launchers); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, launcher := range launchers.LauncherArray {
		key := utils.BuildKey(KEY_LAUNCHER, fmt.Sprint(launcher.ID))

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "id", launcher.ID)
		_, err := pipe.Exec(ctx)

		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return err
		}

	}

	return nil
}

func GetAllValues_ISRO_Launchers(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_LAUNCHER+"*"))

	var launcher Launcher

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {

		if err := Client.HGetAll(ctx, key).Scan(&launcher); err != nil {
			panic(err)
		}
		if print {
			fmt.Printf("ID: %s\n", launcher.ID)
		}
	}
	if print {
		fmt.Println("}")
	}

	return nil
}
func InsertData_ISRO_Satellites(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var satellites Satellites
	if err := json.Unmarshal(data, &satellites); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, satellite := range satellites.SatelliteArray {
		key := utils.BuildKey(KEY_SATELLITE, fmt.Sprint(satellite.ID))

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "id", satellite.ID)
		pipe.HSet(ctx, key, "country", satellite.Country)
		pipe.HSet(ctx, key, "launch_date", satellite.LaunchDate)
		pipe.HSet(ctx, key, "mass", satellite.Mass)
		pipe.HSet(ctx, key, "launcher", satellite.Launcher)
		_, err := pipe.Exec(ctx)

		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return err
		}

	}

	return nil
}

func GetAllValues_ISRO_Satellites(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_SATELLITE+"*"))

	var satellite Satellite

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {

		if err := Client.HGetAll(ctx, key).Scan(&satellite); err != nil {
			panic(err)
		}
		if print {
			fmt.Printf("%s: {\n", satellite.ID)
			fmt.Printf("\tCountry: %s,\n", satellite.Country)
			fmt.Printf("\tLaunchDate: %s,\n", satellite.LaunchDate)
			fmt.Printf("\tMass: %s,\n", satellite.Mass)
			fmt.Printf("\tLauncher: %s,\n", satellite.Launcher)
			fmt.Printf("\t},\n")
		}
	}
	if print {
		fmt.Println("}")
	}

	return nil
}
func InsertData_ISRO_Centers(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var centers Centers
	if err := json.Unmarshal(data, &centers); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, center := range centers.CenterArray {
		key := utils.BuildKey(KEY_CENTER, fmt.Sprint(center.ID))

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "id", center.ID)
		pipe.HSet(ctx, key, "name", center.Name)
		pipe.HSet(ctx, key, "place", center.Place)
		pipe.HSet(ctx, key, "state", center.State)
		_, err := pipe.Exec(ctx)

		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return err
		}

	}

	return nil
}

func GetAllValues_ISRO_Centers(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_CENTER+"*"))

	var center Center

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {

		if err := Client.HGetAll(ctx, key).Scan(&center); err != nil {
			panic(err)
		}
		if print {
			fmt.Printf("%d: {\n", center.ID)
			fmt.Printf("\tName: %s\n", center.Name)
			fmt.Printf("\tPlace: %s\n", center.Place)
			fmt.Printf("\tState: %s\n", center.State)
			fmt.Printf("\t},\n")
		}
	}
	if print {
		fmt.Println("}")
	}

	return nil
}

func RetrieveAllAndStore_ISRO(Client *redis.Client, retrieve bool) {
	fmt.Println("Requesting data from ISRO...")

	response1 := parse.GrabAPI(API_ISRO_SPACECRAFTS, nil)
	fmt.Println("Finished Getting data from ISRO Spacecrafts")

	response2 := parse.GrabAPI(API_ISRO_LAUNCHERS, nil)
	fmt.Println("Finished Getting data from ISRO Launchers")

	response3 := parse.GrabAPI(API_ISRO_SATELLITES, nil)
	fmt.Println("Finished Getting data from ISRO Satellites")

	response4 := parse.GrabAPI(API_ISRO_CENTERS, nil)
	fmt.Println("Finished Getting data from ISRO Centers")
	fmt.Println("ISRO API Request complete!")

	fmt.Println("Storing ISRO data in redis...")
	InsertData_ISRO_Spacecrafts(Client, response1)
	fmt.Println("Finished Storing Spacecrafts data")

	InsertData_ISRO_Launchers(Client, response2)
	fmt.Println("Finished Storing Launchers data")

	InsertData_ISRO_Satellites(Client, response3)
	fmt.Println("Finished Storing Satellites data")

	InsertData_ISRO_Centers(Client, response4)
	fmt.Println("Finished Storing Centers data")
	fmt.Println("ISRO Data Store finished")

	if retrieve {
		fmt.Println("Retrieving ISRO data from Redis...")

		GetAllValues_ISRO_Spacecrafts(Client, true)
		GetAllValues_ISRO_Launchers(Client, true)
		GetAllValues_ISRO_Satellites(Client, true)
		GetAllValues_ISRO_Centers(Client, true)
		fmt.Println("Finished ISRO!")

	}
}
