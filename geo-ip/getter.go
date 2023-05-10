package geo_ip

type GeoIP struct {
	Country     string
	City        string
	CountryCode string
	IP          string
	Latitude    float64
	Longitude   float64
	Timezone    string
}

type IGeoIP interface {
	Get() (*GeoIP, error)
}
