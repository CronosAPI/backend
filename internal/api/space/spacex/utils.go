package spacex

import (
	"backend/internal/parse"
	"context"
	"encoding/json"
	"fmt"

	"backend/internal/utils"

	"github.com/go-redis/redis/v8"
)

func (i Payload) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
func InsertData_SPACEX_Payload(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var payloads []Payload
	if err := json.Unmarshal(data, &payloads); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, payload := range payloads {
		key := utils.BuildKey(KEY_PAYLOADS, payload.PayloadID)

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "payload_id", payload.PayloadID)
		jsonNorad, _ := json.Marshal(payload.NoradID)
		pipe.HSet(ctx, key, "norad_id", jsonNorad)

		pipe.HSet(ctx, key, "reused", payload.Reused)

		jsonCustomers, _ := json.Marshal(payload.Customers)
		pipe.HSet(ctx, key, "customers", jsonCustomers)
		pipe.HSet(ctx, key, "nationality", payload.Nationality)
		pipe.HSet(ctx, key, "manufacturer", payload.Manufacturer)
		pipe.HSet(ctx, key, "payload_type", payload.PayloadType)
		pipe.HSet(ctx, key, "payload_mass_kg", payload.PayloadMassLbs)
		pipe.HSet(ctx, key, "payload_mass_lbs", payload.PayloadMassLbs)
		pipe.HSet(ctx, key, "orbit", payload.Orbit)

		jsonOrbitParams := utils.MarshalObject(payload.OrbitParams)
		pipe.HSet(ctx, key, "orbit_params", jsonOrbitParams)

		_, errPipe := pipe.Exec(ctx)

		if errPipe != nil {
			fmt.Println("Error storing data in Redis:", errPipe)
			return errPipe
		}
	}

	return nil
}

func GetAllValues_SPACEX_Payload(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_PAYLOADS+"*"))

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {
		GetValue_SPACEX_Payload(Client, key, print)
	}
	if print {
		fmt.Println("}")
	}

	return nil
}

func GetValue_SPACEX_Payload(Client *redis.Client, key string, print bool) error {
	var payload Payload
	ctx := context.Background()

	hashResult := Client.HGetAll(ctx, key)
	result, _ := hashResult.Result()
	err1 := hashResult.Scan(&payload)

	if err1 != nil {
		fmt.Println(err1)
		return err1
	}

	if print {
		fmt.Printf("%s: {", payload.PayloadID)
		fmt.Printf("\n\tNoradID: %v,\n", result["norad_id"])
		fmt.Printf("\tReused: %t,\n", payload.Reused)
		fmt.Printf("\tCustomers: %v,\n", result["customers"])
		fmt.Printf("\tNationality: %s,\n", utils.Nully(payload.Nationality))
		fmt.Printf("\tManufacturer: %s,\n", utils.Nully(payload.Manufacturer))
		fmt.Printf("\tPayloadType: %s,\n", utils.Nully(payload.PayloadType))
		fmt.Printf("\tPayloadMassKG: %f,\n", payload.PayloadMassKG)
		fmt.Printf("\tPayloadMassLbs: %f,\n", payload.PayloadMassLbs)
		fmt.Printf("\tOrbit: %s,\n", payload.Orbit)

		utils.UnmarshalData(result["orbit_params"], &payload.OrbitParams)

		fmt.Printf("\tOrbitParams: {")
		fmt.Printf("\n\t\tReferenceSystem: %s,\n\t", utils.Nully(payload.OrbitParams.ReferenceSystem))
		fmt.Printf("\tRegime: %s,\n\t", utils.Nully(payload.OrbitParams.Regime))
		fmt.Printf("\tLogitude: %f\n\t", payload.OrbitParams.Longitude)
		fmt.Printf("\tSemiMajorAxisKM: %f\n\t", payload.OrbitParams.SemiMajorAxisKM)
		fmt.Printf("\tEccentricity: %f\n\t", payload.OrbitParams.Eccentricity)
		fmt.Printf("\tPeriapsisKM: %f\n\t", payload.OrbitParams.PeriapsisKM)
		fmt.Printf("\tApoaapsisKM: %f\n\t", payload.OrbitParams.ApoapsisKM)
		fmt.Printf("\tInclinationDeg: %f\n\t", payload.OrbitParams.InclinationDeg)
		fmt.Printf("\tPeriodMin: %f\n\t", payload.OrbitParams.PeriodMin)
		fmt.Printf("\tLifespanYears: %f\n\t", payload.OrbitParams.LifespanYears)
		fmt.Printf("},\n")

		fmt.Printf("},\n")
	}
	return nil
}

func InsertData_SPACEX_Rocket(Client *redis.Client, data []byte) error {
	ctx := context.Background()
	var rockets []Rocket
	if err := json.Unmarshal(data, &rockets); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}
	for _, rocket := range rockets {
		key := utils.BuildKey(KEY_ROCKETS, fmt.Sprint(rocket.ID))
		pipe := Client.TxPipeline()
		pipe.HSet(ctx, key, "id", rocket.ID)
		pipe.HSet(ctx, key, "active", rocket.Active)
		pipe.HSet(ctx, key, "stages", rocket.Stages)
		pipe.HSet(ctx, key, "boosters", rocket.Boosters)
		pipe.HSet(ctx, key, "cost_per_launch", rocket.CostPerLaunch)
		pipe.HSet(ctx, key, "success_rate_pct", rocket.SuccessRatePct)
		pipe.HSet(ctx, key, "first_flight", rocket.FirstFlight)
		pipe.HSet(ctx, key, "country", rocket.Country)
		pipe.HSet(ctx, key, "company", rocket.Company)
		objHeight := utils.MarshalObject(rocket.Height)
		pipe.HSet(ctx, key, "height", objHeight)
		objDiameter := utils.MarshalObject(rocket.Diameter)
		pipe.HSet(ctx, key, "diameter", objDiameter)
		objMass := utils.MarshalObject(rocket.Mass)
		pipe.HSet(ctx, key, "mass", objMass)
		jsonPayloadWeights, _ := json.Marshal(rocket.PayloadWeights)
		pipe.HSet(ctx, key, "payload_weights", jsonPayloadWeights)
		objFirstStage := utils.MarshalObject(rocket.FirstStage)
		pipe.HSet(ctx, key, "first_stage", objFirstStage)
		objSecondStage := utils.MarshalObject(rocket.SecondStage)
		pipe.HSet(ctx, key, "second_stage", objSecondStage)
		objEngines := utils.MarshalObject(rocket.Engines)
		pipe.HSet(ctx, key, "engines", objEngines)
		objLandingLegs := utils.MarshalObject(rocket.LandingLegs)
		pipe.HSet(ctx, key, "landing_legs", objLandingLegs)
		pipe.HSet(ctx, key, "wikipedia", rocket.Wikipedia)
		pipe.HSet(ctx, key, "description", rocket.Description)
		pipe.HSet(ctx, key, "rocket_id", rocket.RocketID)
		pipe.HSet(ctx, key, "rocket_name", rocket.RocketName)
		pipe.HSet(ctx, key, "rocket_type", rocket.RocketType)
		_, errPipe := pipe.Exec(ctx)
		if errPipe != nil {
			fmt.Println("Error storing data in Redis:", errPipe)
			return errPipe
		}
	}
	return nil
}

func GetAllValues_SPACEX_Rocket(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_ROCKETS+"*"))

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {
		GetValue_SPACEX_Rocket(Client, key, print)
	}
	if print {
		fmt.Println("}")
	}
	return nil
}

func GetValue_SPACEX_Rocket(Client *redis.Client, key string, print bool) error {
	var rocket Rocket
	ctx := context.Background()

	if err := Client.HGetAll(ctx, key).Scan(&rocket); err != nil {
		panic(err)
	}

	// result, err := Client.HGetAll(ctx, key).Result()
	// if err != nil {
	// 	fmt.Println("Error retrieving data from Redis:", err)
	// 	return err
	// }

	if print {
		fmt.Printf("%d: {", rocket.ID)

		fmt.Printf("\tActive: %t,\n", rocket.Active)
		fmt.Printf("},\n")
	}
	return nil
}

func RetrieveAllAndStore_SPACEX(Client *redis.Client, retrieve bool) {
	fmt.Println("Requesting data from SpaceX...")

	fmt.Println("Requesting data for Payloads from SpaceX...")
	// response1 := parse.GrabAPI(API_SPACEX_PAYLOADS, nil)
	fmt.Println("Finished Getting data for Payloads")

	// fmt.Println("Requesting data for Rockets from SpaceX...")
	response2 := parse.GrabAPI(API_SPACEX_ROCKETS, nil)
	// fmt.Println("Finished Getting data for Rockets")

	fmt.Println("Storing SpaceX data in Redis...")
	// InsertData_SPACEX_Payload(Client, response1)
	fmt.Println("Finished Storing Payloads data")

	InsertData_SPACEX_Rocket(Client, response2)
	// fmt.Println("Finished Storing Rockets data")

	fmt.Println("SpaceX Data Store finished")

	if retrieve {
		fmt.Println("Retrieving SpaceX data from Redis...")

		// GetAllValues_SPACEX_Payload(Client, true)
		// GetAllValues_SPACEX_Rocket(Client, true)
		fmt.Println("Finished retrieving SpaceX data!")
	}
}
