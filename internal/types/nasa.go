package types

import "encoding/json"

type DONKI struct {
	GstId        string   `json:"gstID"`
	StartTime    string   `json:"startTime"`
	Link         string   `json:"link"`
	KpIndex      []string `json:"allKpIndex"`
	LinkedEvents []string `json:"linkedEvents"`
}

// type DONKIS struct {
// 	DonkiArray []DONKI `json:"Donki"`
// }

func CreateNasa(inputJSON []byte) []SOLARFLARE {
	var resp []SOLARFLARE
	err := json.Unmarshal(inputJSON, &resp)

	if err != nil {
		return resp
	}
	return resp
}

type SOLARFLARE struct {
	FlareId         string `json:"flrID"`
	Instruments     string `json:"instruments"`
	BeginTime       string `json:"beginTime"`
	PeakTime        string `json:"peakTime"`
	EndTime         string `json:"endTime"`
	ClassType       string `json:"classType"`
	SourceLocation  string `json:"sourceLocation"`
	ActiveRegionNum string `json:"activeRegionNum"`
	LinkedEvents    string `json:"linkedEvents"`
	Link            string `json:"link"`
}
