package weather

type OpenMeteoStruct struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	GenerationtimeMs float64 `json:"generationtime_ms"`
	UtcOffsetSeconds int `json:"utc_offset_seconds"`
	Timezone string `json:"timezone"`
	TimezoneAbbreviation string `json:"timezone_abbreviation"`
	Elevation int `json:"elevation"`
	DailyUnits DailyUnits `json:"daily_units"`
	Daily Daily `json:"daily"`
}
type DailyUnits struct {
	Time string `json:"time"`
	Temperature2MMax string `json:"temperature_2m_max"`
	Temperature2MMin string `json:"temperature_2m_min"`
}
type Daily struct {
	Time []string `json:"time"`
	Temperature2MMax []float64 `json:"temperature_2m_max"`
	Temperature2MMin []float64 `json:"temperature_2m_min"`
}

