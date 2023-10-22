package types

import "encoding/json"

type ROCKETS struct {
	Payload_Id   string `json:"payload_id"`
	Manufacturer string `json:"manufacturer"`
	Payload_Mass int    `json:"payload_mass_kg"`
	Orbit        string `json:"orbit"`
}

// type DONKIS struct {
// 	DonkiArray []DONKI `json:"Donki"`
// }

func CreateDonki(inputJSON []byte) []ROCKETS {
	var resp []ROCKETS
	err := json.Unmarshal(inputJSON, &resp)

	if err != nil {
		return resp
	}
	return resp
}

// type SOLARFLARE struct {
// 	FlareId         string `json:"flrID"`
// 	Instruments     string `json:"instruments"`
// 	BeginTime       string `json:"beginTime"`
// 	PeakTime        string `json:"peakTime"`
// 	EndTime         string `json:"endTime"`
// 	ClassType       string `json:"classType"`
// 	SourceLocation  string `json:"sourceLocation"`
// 	ActiveRegionNum string `json:"activeRegionNum"`
// 	LinkedEvents    string `json:"linkedEvents"`
// 	Link            string `json:"link"`
// }
