package geo_ip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type GeoJs struct {
	AreaCode         string `json:"area_code"`
	OrganizationName string `json:"organization_name"`
	CountryCode      string `json:"country_code"`
	CountryCode3     string `json:"country_code3"`
	ContinentCode    string `json:"continent_code"`
	IP               string `json:"ip"`
	Region           string `json:"region"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	Accuracy         int    `json:"accuracy"`
	Timezone         string `json:"timezone"`
	City             string `json:"city"`
	Organization     string `json:"organization"`
	Asn              int    `json:"asn"`
	Country          string `json:"country"`
}

func (geoJs GeoJs) Get() (GeoIP, error) {
	var httpClient = &http.Client{Timeout: 4 * time.Second}
	r, err := httpClient.Get("https://get.geojs.io/v1/ip/geo.json")
	if err != nil {
		return GeoIP{}, fmt.Errorf("retrieving IP from GeoJS: %v", err)
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return GeoIP{}, fmt.Errorf("retrieving IP from GeoJS: %v", r.Status)
	}
	if err := json.NewDecoder(r.Body).Decode(&geoJs); err != nil {
		return GeoIP{}, fmt.Errorf("decoding IP from GeoJS: %v", err)
	}
	lat, err := strconv.ParseFloat(geoJs.Latitude, 64)
	if err != nil {
		return GeoIP{}, fmt.Errorf("parsing latitude from Geo: %v", err)
	}
	lon, err := strconv.ParseFloat(geoJs.Longitude, 64)
	if err != nil {
		return GeoIP{}, fmt.Errorf("parsing longitude from Geo: %v", err)
	}

	geoIP := GeoIP{
		Country:     geoJs.Country,
		Region:      geoJs.Region,
		City:        geoJs.City,
		CountryCode: geoJs.CountryCode,
		IP:          geoJs.IP,
		Timezone:    geoJs.Timezone,
		Latitude:    lat,
		Longitude:   lon,
	}
	return geoIP, nil
}
