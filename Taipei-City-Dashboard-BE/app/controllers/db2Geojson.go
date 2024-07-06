package controllers

import (
	"fmt"
	"TaipeiCityDashboardBE/app/models"
)

// Geojson is a struct that contains the data of a geojson file
type Geojson struct {
	Type     string       `json:"type"`
	Features []GeoFeature `json:"features"`
}

// GeoJson for dot data
type GeoFeature struct {
	Type       string          `json:"type"`
	Geometry   GeoGeometry     `json:"geometry"`
	Properties GeoProperties   `json:"properties"`
	
}

geojson := Geojson{
	Type: "FeatureCollection",
	Features: []GeoFeature,
}

func GetGeojson() string {
	// Get the geojson data from the database
	// todo finish this function
	var noResourceLocation = []NoResourceLocation
	table := models.DBManager.Table("report")
	if err := table.Find(&geofrature).Error; err != nil {
		fmt.Println("Error retrieving messages from database: %v", err)
		return
	}



	// Return the geojson data



	return Geojson{}
}
