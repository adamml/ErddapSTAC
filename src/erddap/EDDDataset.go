package erddap

import (
	"ErddapSTAC/src/stac"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// EDDDataset describes an Erddap dataset, either
type EDDDataset struct {
	// The name of the Erddap dataset
	Name string
	// The name of the Erddap server hosting the dataset
	HostName string
	// The uri (normally the url) of the Erddap dataset
	Uri string
	// A more setailed textual description of the Erddap dataset
	Description string
	// The type of the Erddap dataset - either EDD_GRID or EDD_TABLE
	EddDatastType string
	License       string
	Id            string
	InfoUrl       string

	// The maximum latitude (northerly extent) of the dataset
	BoundingBoxMaxLat float32
	// The minium latitude (southerly extent) of the dataset
	BoundingBoxMinLat float32
	// The maximum longitude (easternmost extent) of the dataset
	BoundingBoxMaxLon float32
	// The minimum longitude (westernmost extent) of the dataset
	BoundingBoxMinLon float32

	TimeMin time.Time
	TimeMax time.Time

	Keywords []string
}

func NewEDDDataset(url string,
	id string, title string,
	description string, eddDatastType string,
	hostName string, infoUrl string,
	startTime time.Time, endTime time.Time) EDDDataset {
	r, e := http.Get(url)
	if e == nil {
		c, e := ioutil.ReadAll(r.Body)
		if e == nil {
			var jsTable EDDJsonTable
			e := json.Unmarshal([]byte(string(c)), &jsTable)
			bboxMaxLat := float64(90)
			bboxMinLat := float64(-90)
			bboxMaxLon := float64(180)
			bboxMinLon := float64(-180)

			var keywords []string
			var license string

			if e == nil {
				for i := 0; i < len(jsTable.Table.Rows); i++ {
					if jsTable.Table.Rows[i][1] == "NC_GLOBAL" {
						if jsTable.Table.Rows[i][0] == "attribute" {
							if jsTable.Table.Rows[i][2] == "geospatial_lat_max" {
								bboxMaxLat, e = strconv.ParseFloat(jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lat_min" {
								bboxMinLat, e = strconv.ParseFloat(jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_max" {
								bboxMaxLon, e = strconv.ParseFloat(jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_min" {
								bboxMinLon, e = strconv.ParseFloat(jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_min" {
								bboxMinLon, e = strconv.ParseFloat(jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_min" {
								bboxMinLon, e = strconv.ParseFloat(jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "keywords" {
								keywords = strings.Split(jsTable.Table.Rows[i][4], ",")
								for ii := 0; ii < len(keywords); ii++ {
									keywords[ii] = strings.Trim(keywords[ii], " ")
								}
							} else if jsTable.Table.Rows[i][2] == "license" {
								license = jsTable.Table.Rows[i][4]
							}
						}
					}
				}
				edd := EDDDataset{
					Uri:               url,
					BoundingBoxMaxLat: float32(bboxMaxLat),
					BoundingBoxMinLat: float32(bboxMinLat),
					BoundingBoxMaxLon: float32(bboxMaxLon),
					BoundingBoxMinLon: float32(bboxMinLon),
					Keywords:          keywords,
					License:           license,
					Name:              title,
					Description:       description,
					Id:                id,
					EddDatastType:     eddDatastType,
					HostName:          hostName,
					InfoUrl:           infoUrl,
					TimeMin:           startTime,
					TimeMax:           endTime,
				}
				return edd
			} else {
				log.Fatal(e)
			}
		} else {
			log.Fatal(e)
		}
	} else {
		log.Fatal(e)
	}
	var edd EDDDataset
	return edd
}

func EDDDatasetToSTACCollection(dataset EDDDataset) stac.Collection {
	var provider stac.Provider
	provider.Name = dataset.HostName
	provider.Url = dataset.InfoUrl

	sc := stac.NewCollection()
	sc.Keywords = dataset.Keywords
	sc.License = dataset.License
	sc.CollectionExtent.Spatial.Bbox[0][0] = dataset.BoundingBoxMinLon
	sc.CollectionExtent.Spatial.Bbox[0][1] = dataset.BoundingBoxMinLat
	sc.CollectionExtent.Spatial.Bbox[0][2] = dataset.BoundingBoxMaxLon
	sc.CollectionExtent.Spatial.Bbox[0][3] = dataset.BoundingBoxMaxLat
	// TODO: handle null times
	sc.CollectionExtent.Temporal.Interval[0][0] = dataset.TimeMin
	sc.CollectionExtent.Temporal.Interval[0][1] = dataset.TimeMax
	sc.Title = dataset.Name
	sc.Description = dataset.Description
	sc.Id = dataset.Id //TODO: need to add host id
	sc.Providers = append(sc.Providers, provider)
	return sc
}
