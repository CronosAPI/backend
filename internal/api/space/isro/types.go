package isro

const KEY_SPACECRAFT = "isro:spacecraft:"
const KEY_LAUNCHER = "isro:launcher:"
const KEY_SATELLITE = "isro:satellite:"
const KEY_CENTER = "isro:center:"

const API_ISRO = "https://isro.vercel.app/api"

const API_ISRO_SPACECRAFTS = API_ISRO + "/spacecrafts"
const API_ISRO_LAUNCHERS = API_ISRO + "/launchers"
const API_ISRO_SATELLITES = API_ISRO + "/customer_satellites"
const API_ISRO_CENTERS = API_ISRO + "/centres"

type SpaceCraft struct {
	ID   int    `json:"id" redis:"id"`
	Name string `json:"name" redis:"name"`
}

type SpaceCrafts struct {
	SpaceCraftArray []SpaceCraft `json:"spacecrafts"`
}

type Launcher struct {
	ID string `json:"id" redis:"id"`
}

type Launchers struct {
	LauncherArray []Launcher `json:"launchers"`
}

type Satellite struct {
	ID         string `json:"id" redis:"id"`
	Country    string `json:"country" redis:"country"`
	LaunchDate string `json:"launch_date" redis:"launch_date"`
	Mass       string `json:"mass" redis:"mass"`
	Launcher   string `json:"launcher" redis:"launcher"`
}

type Satellites struct {
	SatelliteArray []Satellite `json:"customer_satellites"`
}

type Center struct {
	ID    int    `json:"id" redis:"id"`
	Name  string `json:"name" redis:"name"`
	Place string `json:"Place" redis:"place"`
	State string `json:"State" redis:"state"`
}

type Centers struct {
	CenterArray []Center `json:"centres"`
}
