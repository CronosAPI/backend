package isro

const KEY_SPACECRAFT = "isro:spacecrafts:"
const KEY_LAUNCHER = "isro:launchers:"
const KEY_SATELLITE = "isro:satellites:"
const KEY_CENTER = "isro:centers:"

const API_ISRO = "https://isro.vercel.app/api"

const API_ISRO_SPACECRAFTS = API_ISRO + "/spacecrafts"
const API_ISRO_LAUNCHERS = API_ISRO + "/launchers"
const API_ISRO_SATELLITES = API_ISRO + "/customer_satellites"
const API_ISRO_CENTERS = API_ISRO + "/centres"

/* Space Craft ISRO Struct */
type SpaceCraft struct {
	ID   int    `json:"id" redis:"id"`
	Name string `json:"name" redis:"name"`
}

type SpaceCrafts struct {
	SpaceCraftArray []SpaceCraft `json:"spacecrafts"`
}

/* ISRO launchers */
type Launcher struct {
	ID string `json:"id" redis:"id"`
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

/* Customer Sattelite */
type Satellite struct {
	ID         string `json:"id"`
	Country    string `json:"country"`
	LaunchDate string `json:"launch_date"`
	Mass       string `json:"mass"`
	Launcher   string `json:"launcher"`
}

/* Centre */
type Centre struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Place string `json:"Place"`
	State string `json:"State"`
}
