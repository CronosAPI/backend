package nasa

const KEY_GEOMAGNETIC_STORM = "nasa:geomagnetic_storm:"
const KEY_SOLARFLARE = "nasa:solarflare:"

const API_NASA = "https://api.nasa.gov"
const API_NASA_GEOMAGNETIC_STORM = API_NASA + "/DONKI/GST"
const API_NASA_SOLARFLARE = API_NASA + "/DONKI/FLR"

/****** DONKI ******/

/*** Geomagnetic Storm /GST ***/
type __KpIndex struct {
	ObservedTime string  `json:"observedTime" redis:"observedTime"`
	KpIndex      float64 `json:"kpIndex" redis:"kpIndex"`
	Source       string  `json:"source" redis:"source"`
}

type __LinkedEvent struct {
	ActivityID string `json:"activityID" redis:"activityID"`
}

type GeoStorm struct {
	GSTID        string          `json:"gstID" redis:"gstID"`
	StartTime    string          `json:"startTime" redis:"startTime"`
	AllKpIndex   []__KpIndex     `json:"allKpIndex" redis:"allKpIndex"`
	LinkedEvents []__LinkedEvent `json:"linkedEvents" redis:"linkedEvents"`
	Link         string          `json:"link" redis:"link"`
}

/*** Solarflare /FLR ***/
type __Instrument struct {
	DisplayName string `json:"displayName" redis:"displayName"`
}
type Solarflare struct {
	FlrID           string          `json:"flrID" redis:"flrID"`
	Instruments     []__Instrument  `json:"instruments" redis:"instruments"`
	BeginTime       string          `json:"beginTime" redis:"beginTime"`
	PeakTime        string          `json:"peakTime" redis:"peakTime"`
	EndTime         string          `json:"endTime" redis:"endTime"`
	ClassType       string          `json:"classType" redis:"classType"`
	SourceLocation  string          `json:"sourceLocation" redis:"sourceLocation"`
	ActiveRegionNum int             `json:"activeRegionNum" redis:"activeRegionNum"`
	LinkedEvents    []__LinkedEvent `json:"linkedEvents" redis:"linkedEvents"`
	Link            string          `json:"link" redis:"link"`
}
