package types

import (
	"encoding/json"
)

/* Space Craft ISRO Struct */
type SpaceCraft struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SpaceCrafts struct {
	SpaceCraftArray []SpaceCraft `json:"spacecrafts"`
}

/* ISRO launchers */
type Launcher struct {
	ID string `json:"id"`
}

type Launchers struct {
	LauncherArray []Launcher `json:"launchers"`
}

type Satellites struct {
	CustomerSatellites []Satellite `json:"customer_satellites"`
}

type Centres struct {
	CustomerSatellites []Centre `json:"centres"`
}

// loop and send to redis
func CreateSpaceCraft(inputJSON []byte) SpaceCrafts {
	var resp SpaceCrafts
	err := json.Unmarshal(inputJSON, &resp)

	if err != nil {
		return resp
	}
	return resp
}

func CreateLauncher(inputJSON []byte) Launchers {
	var resp Launchers

	err := json.Unmarshal(inputJSON, &resp)

	if err != nil {
		return resp
	}

	return resp
}

/* Customer Sattelite */
type Satellite struct {
	ID         string `json:"id"`
	Country    string `json:"country"`
	LaunchDate string `json:"launch_date"`
	Mass       string `json:"mass"`
	Launcher   string `json:"launcher"`
}

func CreateSattelite(inputJSON []byte) Satellites {
	var resp Satellites
	err := json.Unmarshal(inputJSON, &resp)

	if err != nil {
		return resp
	}

	return resp
}

/* Centre */
type Centre struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Place string `json:"Place"`
	State string `json:"State"`
}

func CreateCentre(inputJSON []byte) Centres {
	var resp Centres
	err := json.Unmarshal(inputJSON, &resp)

	if err != nil {
		return resp
	}

	return resp
}
