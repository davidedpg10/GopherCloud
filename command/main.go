package main

import (
	"encoding/json"
	"fmt"

	geo_ip "github.com/davidedpg10/GopherCloud/geo-ip"
)

func main() {
	geoJs := geo_ip.GeoJs{}
	r, err := geoJs.Get()
	if err != nil {
		panic(err)
	}
	prettyPrintStruct(r)
}

func prettyPrintStruct(g geo_ip.GeoIP) {
	umJson, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(string(umJson)))
}
