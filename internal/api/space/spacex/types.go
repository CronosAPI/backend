package spacex

const KEY_PAYLOADS = "spacex:payload:"
const KEY_ROCKETS = "spacex:rocket:"

const API_SPACEX = "https://api.spacexdata.com/v3"
const API_SPACEX_PAYLOADS = API_SPACEX + "/payloads"
const API_SPACEX_ROCKETS = API_SPACEX + "/rockets"

/*** Payloads /payloads ***/
type OrbitParams struct {
	ReferenceSystem string  `json:"reference_system" redis:"reference_system"`
	Regime          string  `json:"regime" redis:"regime"`
	Longitude       float64 `json:"longitude" redis:"longitude"`
	SemiMajorAxisKM float64 `json:"semi_major_axis_km" redis:"semi_major_axis_km"`
	Eccentricity    float64 `json:"eccentricity" redis:"eccentricity"`
	PeriapsisKM     float64 `json:"periapsis_km" redis:"periapsis_km"`
	ApoapsisKM      float64 `json:"apoapsis_km" redis:"apoapsis_km"`
	InclinationDeg  float64 `json:"inclination_deg" redis:"inclination_deg"`
	PeriodMin       float64 `json:"period_min" redis:"period_min"`
	LifespanYears   float64 `json:"lifespan_years" redis:"lifespan_years"`
}

type Payload struct {
	PayloadID      string      `json:"payload_id" redis:"payload_id"`
	NoradID        []int       `json:"norad_id" redis:"norad_id"`
	Reused         bool        `json:"reused" redis:"reused"`
	Customers      []string    `json:"customers" redis:"customers"`
	Nationality    string      `json:"nationality" redis:"nationality"`
	Manufacturer   string      `json:"manufacturer" redis:"manufacturer"`
	PayloadType    string      `json:"payload_type" redis:"payload_type"`
	PayloadMassKG  float64     `json:"payload_mass_kg" redis:"payload_mass_kg"`
	PayloadMassLbs float64     `json:"payload_mass_lbs" redis:"payload_mass_lbs"`
	Orbit          string      `json:"orbit" redis:"orbit"`
	OrbitParams    OrbitParams `json:"orbit_params"`
}

/*** Rockets /rockets ***/
type Height struct {
	Meters float64 `json:"meters"`
	Feet   float64 `json:"feet"`
}

type Diameter struct {
	Meters float64 `json:"meters"`
	Feet   float64 `json:"feet"`
}

type Thrust struct {
	KN  float64 `json:"kN"`
	Lbf float64 `json:"lbf"`
}

type ThrustInfo struct {
	SeaLevel Thrust `json:"thrust_sea_level"`
	Vacuum   Thrust `json:"thrust_vacuum"`
}

type CompositeFairing struct {
	Height   Height   `json:"height"`
	Diameter Diameter `json:"diameter"`
}

type FirstStage struct {
	Reusable       bool       `json:"reusable"`
	Engines        int        `json:"engines"`
	FuelAmountTons float64    `json:"fuel_amount_tons"`
	BurnTimeSec    int        `json:"burn_time_sec"`
	ThrustSeaLevel ThrustInfo `json:"thrust_sea_level"`
	ThrustVacuum   ThrustInfo `json:"thrust_vacuum"`
}

type PayloadsSecondStage struct {
	Option1          string           `json:"option_1"`
	CompositeFairing CompositeFairing `json:"composite_fairing"`
}
type SecondStage struct {
	Engines        int                 `json:"engines"`
	FuelAmountTons float64             `json:"fuel_amount_tons"`
	BurnTimeSec    int                 `json:"burn_time_sec"`
	Thrust         Thrust              `json:"thrust"`
	Payloads       PayloadsSecondStage `json:"payloads"`
}

type Engines struct {
	Number         int        `json:"number"`
	Type           string     `json:"type"`
	Version        string     `json:"version"`
	Layout         string     `json:"layout"`
	EngineLossMax  int        `json:"engine_loss_max"`
	Propellant1    string     `json:"propellant_1"`
	Propellant2    string     `json:"propellant_2"`
	ThrustSeaLevel ThrustInfo `json:"thrust_sea_level"`
	ThrustVacuum   ThrustInfo `json:"thrust_vacuum"`
	ThrustToWeight float64    `json:"thrust_to_weight"`
}

type LandingLegs struct {
	Number   int    `json:"number"`
	Material string `json:"material"`
}

type PayloadWeights struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Kg   float64 `json:"kg"`
	Lb   float64 `json:"lb"`
}

type RocketMass struct {
	Kg float64 `json:"kg"`
	Lb float64 `json:"lb"`
}

type Rocket struct {
	ID             int              `json:"id"`
	Active         bool             `json:"active"`
	Stages         int              `json:"stages"`
	Boosters       int              `json:"boosters"`
	CostPerLaunch  int              `json:"cost_per_launch"`
	SuccessRatePct int              `json:"success_rate_pct"`
	FirstFlight    string           `json:"first_flight"`
	Country        string           `json:"country"`
	Company        string           `json:"company"`
	Height         Height           `json:"height"`
	Diameter       Diameter         `json:"diameter"`
	Mass           RocketMass       `json:"mass"`
	PayloadWeights []PayloadWeights `json:"payload_weights"`
	FirstStage     FirstStage       `json:"first_stage"`
	SecondStage    SecondStage      `json:"second_stage"`
	Engines        Engines          `json:"engines"`
	LandingLegs    LandingLegs      `json:"landing_legs"`
	Wikipedia      string           `json:"wikipedia"`
	Description    string           `json:"description"`
	RocketID       string           `json:"rocket_id"`
	RocketName     string           `json:"rocket_name"`
	RocketType     string           `json:"rocket_type"`
}
