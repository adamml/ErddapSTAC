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

func (dataset EDDDataset) ToSTACCollection(baseURL string) stac.Collection {
	var provider stac.Provider
	provider.Name = dataset.HostName
	provider.Url = dataset.InfoUrl

	links := []stac.Link{}
	links = append(links, stac.Link{Href: "./" + strings.ToLower(dataset.HostName) + "_catalog.json", Rel: stac.STAC_LINK_RELTYPE_PARENT, Type: stac.STAC_CATALOG_MIME_TYPE})
	links = append(links, stac.Link{Href: "./" + strings.ToLower(dataset.HostName) + "_" + strings.ToLower(dataset.Id) + "_item.json", Rel: stac.STAC_LINK_RELTYPE_ITEM, Type: stac.STAC_ITEM_MIME_TYPE})
	links = append(links, stac.Link{Href: baseURL + strings.ToLower(dataset.HostName) + "_" + strings.ToLower(dataset.Id) + "_collection.json", Rel: stac.STAC_LINK_RELTYPE_SELF, Type: stac.STAC_COLLECTION_MIME_TYPE})

	sc := stac.NewCollection()
	sc.Keywords = dataset.Keywords
	sc.License = dataset.License

	if dataset.EddDatastType == EDD_TABLE {
		sc.StacExtensions = append(sc.StacExtensions, string(stac.STAC_EXTENSIONS_TABLE))
	}

	sc.CollectionExtent.Spatial.Bbox[0][0] = dataset.BoundingBoxMinLon
	sc.CollectionExtent.Spatial.Bbox[0][1] = dataset.BoundingBoxMinLat
	sc.CollectionExtent.Spatial.Bbox[0][2] = dataset.BoundingBoxMaxLon
	sc.CollectionExtent.Spatial.Bbox[0][3] = dataset.BoundingBoxMaxLat
	// TODO: handle null times
	sc.CollectionExtent.Temporal.Interval[0][0] = dataset.TimeMin
	sc.CollectionExtent.Temporal.Interval[0][1] = dataset.TimeMax
	sc.Title = dataset.Name
	sc.Description = dataset.Description
	sc.Id = strings.ToLower(dataset.HostName) + "_" + strings.ToLower(dataset.Id) + "_collection" //TODO: need to add host id
	sc.Providers = append(sc.Providers, provider)
	sc.Links = links
	sc.Keywords = dataset.Keywords
	return sc
}

func (dataset EDDDataset) DatasetUriToMetadataUri() string {
	uri := dataset.Uri
	uri = strings.Split(uri, "erddap/")[0]
	return uri + "erddap/metadata/iso19115/xml/" + dataset.Id + "_iso19115.xml"
}

func (dataset EDDDataset) ToSTACItem(baseURL string) stac.Item {
	item := stac.NewItem()

	item.Id = strings.ToLower(dataset.HostName) + "_ " + strings.ToLower(dataset.Id) + "_item"

	item.Bbox = append(item.Bbox, dataset.BoundingBoxMinLon)
	item.Bbox = append(item.Bbox, dataset.BoundingBoxMinLat)
	item.Bbox = append(item.Bbox, dataset.BoundingBoxMaxLon)
	item.Bbox = append(item.Bbox, dataset.BoundingBoxMaxLat)

	item.Properties.Title = dataset.Name
	item.Properties.Description = dataset.Description

	item.Properties.StartTime = dataset.TimeMin.String()
	item.Properties.EndTime = dataset.TimeMax.String()

	md := stac.Asset{
		Href:  dataset.DatasetUriToMetadataUri(),
		Type:  stac.STAC_ASSET_MEDIA_TYPE_XML,
		Title: "ISO19115 metadata",
		Roles: []string{},
	}
	md.Roles = append(md.Roles, stac.STAC_ASSET_ROLE_METADATA)
	item.Assets["metadata"] = md

	d := stac.Asset{
		Href:  dataset.Uri,
		Type:  stac.STAC_ASSET_MEDIA_TYPE_JSON,
		Title: "JSON Data",
		Roles: []string{},
	}
	d.Roles = append(d.Roles, stac.STAC_ASSET_ROLE_DATA)
	item.Assets["json_data"] = d

	d2 := stac.Asset{
		Href:  dataset.Uri,
		Type:  stac.STAC_ASSET_MEDIA_TYPE_NETCDF,
		Title: "NetCDF Data",
		Roles: []string{},
	}
	d2.Roles = append(d2.Roles, stac.STAC_ASSET_ROLE_DATA)
	item.Assets["netcdf_data"] = d2

	item.Properties.Keywords = dataset.Keywords

	item.Geometry.Coordinates = append(item.Geometry.Coordinates, []float32{})

	return item
}
