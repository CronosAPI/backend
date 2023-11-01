package nasa

import (
	"backend/internal/parse"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func InsertData_NASA_GeomagneticStorm(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var geoStorms []GeoStorm
	if err := json.Unmarshal(data, &geoStorms); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, geoStorm := range geoStorms {
		key := utils.BuildKey(KEY_GEOMAGNETIC_STORM, fmt.Sprint(geoStorm.GSTID))

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "gstID", geoStorm.GSTID)
		pipe.HSet(ctx, key, "startTime", geoStorm.StartTime)
		allKpIndexJSON, err := json.Marshal(geoStorm.AllKpIndex)
		if err != nil {
			fmt.Println("Error marshaling AllKpIndex to JSON:", err)
			return err
		}
		pipe.HSet(ctx, key, "allKpIndex", allKpIndexJSON)
		linkedEventsJSON, err := json.Marshal(geoStorm.LinkedEvents)
		if err != nil {
			fmt.Println("Error marshaling LinkedEvents to JSON:", err)
			return err
		}
		pipe.HSet(ctx, key, "linkedEvents", linkedEventsJSON)

		pipe.HSet(ctx, key, "link", geoStorm.Link)
		_, errPipe := pipe.Exec(ctx)

		if errPipe != nil {
			fmt.Println("Error storing data in Redis:", errPipe)
			return errPipe
		}

	}

	return nil
}

func GetAllValues_NASA_GeomagneticStorm(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_GEOMAGNETIC_STORM+"*"))

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {
		GetValue_NASA_GeomagneticStorm(Client, key, print)
	}
	if print {
		fmt.Println("}")
	}

	return nil
}

func GetValue_NASA_GeomagneticStorm(Client *redis.Client, key string, print bool) error {
	var geoStorm GeoStorm
	ctx := context.Background()

	if err := Client.HGetAll(ctx, key).Scan(&geoStorm); err != nil {
		panic(err)
	}

	result, err := Client.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println("Error retrieving data from Redis:", err)
		return err
	}
	if print {
		fmt.Printf("%s: {", geoStorm.GSTID)
		fmt.Printf("\n\tStartTime: %s,\n", geoStorm.StartTime)
		fmt.Printf("\tAllKpIndex: {")
		var allKpIndex []__KpIndex

		allKpIndexStr := result["allKpIndex"]
		if err := json.Unmarshal([]byte(allKpIndexStr), &allKpIndex); err != nil {
			fmt.Println("Error unmarshaling allKpIndex:", err)
			return err
		}
		for _, kpIndex := range allKpIndex {
			fmt.Printf("\n\t\tObservedTime: %s,", kpIndex.ObservedTime)
			fmt.Printf("KpIndex: %f,", kpIndex.KpIndex)
			fmt.Printf("Source: %s\n\t", kpIndex.Source)
		}
		fmt.Printf("},\n")

		var linkedEvents []__LinkedEvent

		linkedEventsStr := result["linkedEvents"]
		if err := json.Unmarshal([]byte(linkedEventsStr), &linkedEvents); err != nil {
			fmt.Println("Error unmarshaling linkedEvents:", err)
			return err
		}
		fmt.Printf("\tLinkedEvents: {")
		for _, linkedEvent := range geoStorm.LinkedEvents {
			fmt.Printf("\n\t\tActivityID: %s\n\t", linkedEvent.ActivityID)
		}
		fmt.Printf("},\n")
		fmt.Printf("\tLink: %s,\n", geoStorm.Link)
		fmt.Printf("},\n")
	}
	return nil
}

func InsertData_NASA_Solarflare(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var solarflares []Solarflare
	if err := json.Unmarshal(data, &solarflares); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, solarflare := range solarflares {

		key := utils.BuildKey(KEY_SOLARFLARE, fmt.Sprint(solarflare.FlrID))

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "flrID", solarflare.FlrID)
		pipe.HSet(ctx, key, "beginTime", solarflare.BeginTime)
		pipe.HSet(ctx, key, "peakTime", solarflare.PeakTime)
		pipe.HSet(ctx, key, "endTime", solarflare.EndTime)
		pipe.HSet(ctx, key, "classType", solarflare.ClassType)
		pipe.HSet(ctx, key, "sourceLocation", solarflare.SourceLocation)
		pipe.HSet(ctx, key, "activeRegionNum", solarflare.ActiveRegionNum)
		pipe.HSet(ctx, key, "link", solarflare.Link)

		instrumentsJson, err := json.Marshal(solarflare.Instruments)
		if err != nil {
			fmt.Println("Error marshaling instruments to JSON:", err)
			return err
		}

		pipe.HSet(ctx, key, "instruments", instrumentsJson)

		linkedEventsJSON, err := json.Marshal(solarflare.LinkedEvents)
		if err != nil {
			fmt.Println("Error marshaling LinkedEvents to JSON:", err)
			return err
		}
		pipe.HSet(ctx, key, "linkedEvents", linkedEventsJSON)

		_, errPipe := pipe.Exec(ctx)

		if errPipe != nil {
			fmt.Println("Error storing data in Redis:", errPipe)
			return errPipe
		}

	}

	return nil
}

func GetAllValues_NASA_Solarflare(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_SOLARFLARE+"*"))

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {
		GetValue_NASA_Solarflare(Client, key, print)
	}
	if print {
		fmt.Println("}")
	}
	return nil
}

func GetValue_NASA_Solarflare(Client *redis.Client, key string, print bool) error {
	var solarflare Solarflare
	ctx := context.Background()

	if err := Client.HGetAll(ctx, key).Scan(&solarflare); err != nil {
		panic(err)
	}

	result, err := Client.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println("Error retrieving data from Redis:", err)
		return err
	}
	if print {
		fmt.Printf("%s: {", solarflare.FlrID)

		var instruments []__Instrument

		instrumentsStr := result["instruments"]
		if err := json.Unmarshal([]byte(instrumentsStr), &instruments); err != nil {
			fmt.Println("Error unmarshaling instruments:", err)
			return err
		}
		fmt.Printf("\n\tInstruments: {")
		for _, instrument := range instruments {
			fmt.Printf("\n\t\tDisplayName: %s\n\t", instrument.DisplayName)
		}
		fmt.Printf("},\n")

		fmt.Printf("\tBeginTime: %s,\n", solarflare.BeginTime)
		fmt.Printf("\tPeakTime: %s,\n", solarflare.PeakTime)
		fmt.Printf("\tEndTime: %s,\n", solarflare.EndTime)
		fmt.Printf("\tClassType: %s,\n", solarflare.ClassType)
		fmt.Printf("\tSourceLocation: %s,\n", solarflare.SourceLocation)
		fmt.Printf("\tActiveRegionNum: %d,\n", solarflare.ActiveRegionNum)

		var linkedEvents []__LinkedEvent

		linkedEventsStr := result["linkedEvents"]
		if err := json.Unmarshal([]byte(linkedEventsStr), &linkedEvents); err != nil {
			fmt.Println("Error unmarshaling linkedEvents:", err)
			return err
		}
		fmt.Printf("\tLinkedEvents: {")
		for _, linkedEvent := range linkedEvents {
			fmt.Printf("\n\t\tActivityID: %s\n\t", linkedEvent.ActivityID)
		}
		fmt.Printf("},\n")
		fmt.Printf("\tLink: %s,\n", solarflare.Link)
		fmt.Printf("},\n")
	}
	return nil
}

func RetrieveAllAndStore_NASA(Client *redis.Client, retrieve bool) {
	fmt.Println("Requesting data from NASA...")

	params := map[string]string{
		"startDate": "1990-01-01",
		"endDate":   "2023-10-30",
		"api_key":   os.Getenv("API_KEY_NASA"),
	}

	fmt.Println("Requesting data for Geomagnetic Storm from NASA...")
	response1 := parse.GrabAPI(API_NASA_GEOMAGNETIC_STORM, params)
	fmt.Println("Finished Getting data for Geomagnetic Storm")

	fmt.Println("Requesting data for Solarflare from NASA...")
	response2 := parse.GrabAPI(API_NASA_SOLARFLARE, params)
	fmt.Println("Finished Getting data for Solarflare")

	fmt.Println("Storing NASA data in Redis...")
	InsertData_NASA_GeomagneticStorm(Client, response1)
	fmt.Println("Finished Storing Geomagnetic Storm data")

	InsertData_NASA_Solarflare(Client, response2)
	fmt.Println("Finished Storing Solarflare data")

	fmt.Println("NASA Data Store finished")

	if retrieve {
		fmt.Println("Retrieving NASA data from Redis...")

		GetAllValues_NASA_GeomagneticStorm(Client, true)
		GetAllValues_NASA_Solarflare(Client, true)
		fmt.Println("Finished retrieving NASA data!")
	}
}
