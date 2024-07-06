package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// // Geojson is a struct that contains the data of a geojson file
// type Geojson struct {
// 	Type     string       `json:"type"`
// 	Features []GeoFeature `json:"features"`
// }

// // GeoJson for dot data
// type GeoFeature struct {
// 	Type       string          `json:"type"`
// 	Geometry   GeoGeometry     `json:"geometry"`
// 	Properties GeoProperties   `json:"properties"`

// }
type Geojson struct {
	Type	string
	Features []GeoFeature
}
type GeoFeature struct {
	Type	string `json:"type"`
	Geometry GeoGeometries `json:"geometry`
	Properties	GeoProperties `json:"properties"`
}
type GeoGeometries struct {
	Type        string	`json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type GeoProperties struct {
	Content	string	`json:"content" gorm:"column:content"`
	Type	string	`json:"type" gorm:"column:type"`
	Time 	time.Time	`json:"time" gorm:"column:time"`
}
func GetGeojson(c *gin.Context) {
	str := GetGeojsonStr()
	c.JSON(200, str)
}
func GetGeojsonStr() string {
    var reports []NoResourceLocation

    // 执行查询
    rows, err := models.DBDashboard.Table("report").Select("*").Rows()
    if err != nil {
        fmt.Printf("Error executing query: %v\n", err)
        return ""
    }
    defer rows.Close()

    // 逐行扫描结果
    for rows.Next() {
        var report NoResourceLocation
        if err := rows.Scan(&report.message, &report.theType, &report.theTime, &report.lng, &report.lat); err != nil {
            fmt.Printf("Error scanning row: %v\n", err)
            return ""
        }
        reports = append(reports, report)
    }

    // 检查是否有扫描过程中的错误
    if err := rows.Err(); err != nil {
        fmt.Printf("Error during rows iteration: %v\n", err)
        return ""
    }

    // 转换为 GeoJSON 格式
    var features []GeoFeature
    for _, report := range reports {
        feature := GeoFeature{
            Type: "Feature",
            Geometry: GeoGeometries{
                Type: "Point",
                Coordinates: []float64{report.lng, report.lat},
            },
            Properties: GeoProperties{
                Content: report.message,
                Type:    report.theType,
                Time:    report.theTime,
            },
        }
        features = append(features, feature)
    }

    geojson := Geojson{
        Type:     "FeatureCollection",
        Features: features,
    }
	// fmt.Print(geojson)
    geojsonString, err := json.Marshal(geojson)
    if err != nil {
        fmt.Printf("Error marshalling GeoJSON: %v\n", err)
        return ""
    }

    return string(geojsonString)
}