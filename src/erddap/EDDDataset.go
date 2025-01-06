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

	"github.com/gofrs/uuid/v5"
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
	// The license under which the Erddap dataset is made available
	License string
	// A lexical identifier for the dataset
	Id string
	// An HTTP URL to more information about the dataset
	InfoUrl string
	// The maximum latitude (northerly extent) of the dataset
	BoundingBoxMaxLat float32
	// The minium latitude (southerly extent) of the dataset
	BoundingBoxMinLat float32
	// The maximum longitude (easternmost extent) of the dataset
	BoundingBoxMaxLon float32
	// The minimum longitude (westernmost extent) of the dataset
	BoundingBoxMinLon float32
	// The minimum timetamp of the dataset
	TimeMin time.Time
	// The maximum timestamp of the dataset
	TimeMax time.Time
	// The timestamp of the dataset's creation
	CreatedTime time.Time
	// The timestamp of the last update of the dataset
	ModifiedTime time.Time
	// A link to more detailed metadata describing the dataset
	MetadataLink string
	// A slice of strings listing keywords associated with the dataset
	Keywords []string
	// A slice of stac.TableColumns describing the variables in the dataset
	Variables []stac.TableColumn
	// A list of the variable names used to subset the dataset
	SubsetVariables []string
	// Collection GUID
	CollectionGUID uuid.UUID
	// Catalog GUID
	CatalogGUID uuid.UUID
	// Item GUIDs (to be properly implemented later)
	// TODO: Add extra GUIDs...
	ItemGUIDs []uuid.UUID
}

// NewEDDDataset builds a new EDDDataset struct from the following inputs:
//   - url: An HTTP URL to the dataset info.json as a string
//   - id: A lexical identifier for the dataset as a string
//   - descrition: A more detailed lexlical description of the dataset as a string
//   - eddDatasetType: A string identifying the type of Erddap dataset, either EDD_GRID or EDD_TABLE
//   - hostName
//   - infoUrl
//   - startTime
//   - endTime
func NewEDDDataset(url string,
	id string, title string,
	description string, eddDatastType string,
	hostGUID uuid.UUID,
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

			var variables []stac.TableColumn

			var mdLink string

			var createdTime time.Time
			var modifiedTime time.Time

			var subsetVariables []string

			var itemGUIDs []uuid.UUID

			if e == nil {
				var err error
				for i := 0; i < len(jsTable.Table.Rows); i++ {
					if jsTable.Table.Rows[i][1] == "NC_GLOBAL" {
						if jsTable.Table.Rows[i][0] == "attribute" {
							if jsTable.Table.Rows[i][2] == "geospatial_lat_max" {
								bboxMaxLat, e = strconv.ParseFloat(
									jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lat_min" {
								bboxMinLat, e = strconv.ParseFloat(
									jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_max" {
								bboxMaxLon, e = strconv.ParseFloat(
									jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_min" {
								bboxMinLon, e = strconv.ParseFloat(
									jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_min" {
								bboxMinLon, e = strconv.ParseFloat(
									jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "geospatial_lon_min" {
								bboxMinLon, e = strconv.ParseFloat(
									jsTable.Table.Rows[i][4], 32)
							} else if jsTable.Table.Rows[i][2] == "keywords" {
								keywords = strings.Split(jsTable.Table.Rows[i][4], ",")
								for ii := 0; ii < len(keywords); ii++ {
									keywords[ii] = strings.Trim(keywords[ii], " ")
								}
							} else if jsTable.Table.Rows[i][2] == "license" {
								license = jsTable.Table.Rows[i][4]
							} else if jsTable.Table.Rows[i][2] == "date_created" {
								err = nil
								createdTime, err = time.Parse(EDD_PROVENANCE_TIME_LAYOUT, jsTable.Table.Rows[i][4])
								if err != nil {
									createdTime, err = time.Parse("2006-01-02", jsTable.Table.Rows[i][4][0:10])
								}
							} else if jsTable.Table.Rows[i][2] == "date_modified" {
								err = nil
								modifiedTime, err = time.Parse(EDD_PROVENANCE_TIME_LAYOUT, jsTable.Table.Rows[i][4])
								if err != nil {
									modifiedTime, err = time.Parse("2006-01-02", jsTable.Table.Rows[i][4][0:10])
								}
							} else if jsTable.Table.Rows[i][2] == "metadata_link" {
								mdLink = jsTable.Table.Rows[i][4]
							} else if jsTable.Table.Rows[i][2] == "subsetVariables" {
								subsetVariables = strings.Split(jsTable.Table.Rows[i][4], ",")
							} else if jsTable.Table.Rows[i][2] == "time_coverage_start" {
								if startTime.Year() == 1 {
									if startTime.YearDay() == 1 {
										startTime, _ = time.Parse(EDD_TIME_LAYOUT, jsTable.Table.Rows[i][4])
									}
								}
							} else if jsTable.Table.Rows[i][2] == "time_coverage_end" {
								if endTime.Year() == 1 {
									if endTime.YearDay() == 1 {
										endTime, _ = time.Parse(EDD_TIME_LAYOUT, jsTable.Table.Rows[i][4])
									}
								}
							}
						}
					} else if jsTable.Table.Rows[i][0] == "variable" {
						variables = append(variables,
							stac.TableColumn{Name: jsTable.Table.Rows[i][1],
								Type:     jsTable.Table.Rows[i][3],
								DataType: jsTable.Table.Rows[i][3]})
					} else if jsTable.Table.Rows[i][0] == "attribute" {
						if jsTable.Table.Rows[i][2] == "units" {
							for j := 0; j < len(variables); j++ {
								if variables[j].Name == jsTable.Table.Rows[i][1] {
									variables[j].Unit = jsTable.Table.Rows[i][4]
								}
							}
						} else if jsTable.Table.Rows[i][2] == "long_name" {
							for j := 0; j < len(variables); j++ {
								if variables[j].Name == jsTable.Table.Rows[i][1] {
									variables[j].Description = jsTable.Table.Rows[i][4]
								}
							}
						} else if jsTable.Table.Rows[i][2] == "_FillValue" {
							for j := 0; j < len(variables); j++ {
								if variables[j].Name == jsTable.Table.Rows[i][1] {
									if jsTable.Table.Rows[i][3] == "double" {
										variables[j].NoData, _ = strconv.ParseFloat(jsTable.Table.Rows[i][4], 64)
									} else if jsTable.Table.Rows[i][3] == "int" {
										variables[j].NoData, _ = strconv.ParseInt(jsTable.Table.Rows[i][4], 0, 64)
									} else {
										variables[j].NoData = jsTable.Table.Rows[i][4]
									}
								}
							}
						} else if jsTable.Table.Rows[i][2] == "actual_range" {
							for j := 0; j < len(variables); j++ {
								if variables[j].Name == jsTable.Table.Rows[i][1] {
									sSplit := strings.SplitN(jsTable.Table.Rows[i][4], ",", 2)
									mn, _ := strconv.ParseFloat(sSplit[0], 64)
									mx, _ := strconv.ParseFloat(strings.Trim(sSplit[1], " "), 64)
									variables[j].Statistics.Minimum = mn
									variables[j].Statistics.Maximum = mx
								}
							}
						}
					}
				}
				iUid, _ := uuid.NewV4()
				cUid, _ := uuid.NewV4()
				itemGUIDs = append(itemGUIDs, iUid)

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
					Variables:         variables,
					MetadataLink:      mdLink,
					CreatedTime:       createdTime,
					ModifiedTime:      modifiedTime,
					SubsetVariables:   subsetVariables,
					CollectionGUID:    cUid,
					ItemGUIDs:         itemGUIDs,
					CatalogGUID:       hostGUID,
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

// toSTACCollection(baseURL) exports an EDDDataset as a stac.Collection. As STAC Collections
// require an absolute URI, baseURL provides the absolute URI reference to the parent
// folder of the JSON file that the stac.Collection will be serialised to.
func (dataset EDDDataset) ToSTACCollection(baseURL string) stac.Collection {
	var provider stac.Provider
	provider.Name = dataset.HostName
	provider.Url = dataset.InfoUrl

	links := []stac.Link{}
	links = append(links, stac.Link{Href: "./" +
		dataset.CatalogGUID.String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_PARENT,
		Type: stac.STAC_CATALOG_MIME_TYPE})
	//TODO: Refactor the items...
	links = append(links, stac.Link{Href: "./" +
		dataset.ItemGUIDs[0].String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_ITEM,
		Type: stac.STAC_ITEM_MIME_TYPE})
	links = append(links, stac.Link{Href: baseURL +
		dataset.CollectionGUID.String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_SELF,
		Type: stac.STAC_COLLECTION_MIME_TYPE})

	sc := stac.NewCollection()
	sc.Keywords = dataset.Keywords
	sc.License = dataset.License

	if dataset.EddDatastType == EDD_TABLE {
		sc.StacExtensions = append(sc.StacExtensions,
			string(stac.STAC_EXTENSIONS_TABLE))
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
	sc.Id = dataset.CollectionGUID.String()
	sc.Providers = append(sc.Providers, provider)
	sc.Links = links
	sc.Keywords = dataset.Keywords
	return sc
}

// GetMetadataUri returns an HTTP URL pointing to the ISO19115 XML metadata record for the Erddap dataset as a string
func (dataset EDDDataset) GetMetadataUri() string {
	uri := dataset.Uri
	uri = strings.Split(uri, "erddap/")[0]
	return uri + "erddap/metadata/iso19115/xml/" + dataset.Id + "_iso19115.xml"
}

// GetDataUri(format) returns, as a string, the HTTP URL to download the full
// dataset from an Erddap server in the specified format. The erddap package
// provides some helper consts for format, e.g. ERDDAP_FORMAT_JSON and
// ERDDAP_FORMAT_NETCDF.
func (dataset EDDDataset) GetDataUri(format string) string {
	uri := dataset.Uri
	uri = strings.Split(uri, "erddap/")[0]
	if dataset.EddDatastType == EDD_TABLE {
		return uri + "erddap/tabledap/" + dataset.Id + format
	} else {
		return uri + "erddap/griddap/" + dataset.Id + format
	}

}

// GetTableRowCount returns the number of rows in an Erddap Tabledap dataset.
// If the dataset is an Erddap Griddap then -1 is returned.
func (dataset EDDDataset) GetTableRowCount() int64 {
	if dataset.EddDatastType == EDD_GRID {
		return -1
	} else {
		gotAVariable := false
		var varForGetTableLength string
		var subsetVariablesCheck bool

		if len(dataset.Variables) > 0 {
			for i := 0; i < len(dataset.Variables); i++ {
				subsetVariablesCheck = false
				if len(dataset.SubsetVariables) > 0 {
					for j := 0; j < len(dataset.SubsetVariables); j++ {
						if dataset.Variables[i].Name == dataset.SubsetVariables[j] {
							subsetVariablesCheck = true
						}
					}
				}
				if !subsetVariablesCheck {
					varForGetTableLength = dataset.Variables[i].Name
					i = len(dataset.Variables)
					gotAVariable = true
				}
			}
		}

		if gotAVariable {
			uri := dataset.Uri
			uri = strings.Split(uri, "erddap/")[0]
			uri = uri + "erddap/tabledap/" + dataset.Id + ".json?" + varForGetTableLength
			r, e := http.Get(uri)
			if e == nil {
				c, e := ioutil.ReadAll(r.Body)

				if e == nil {
					var jsTable EDDJsonTable
					e := json.Unmarshal([]byte(string(c)), &jsTable)
					if e == nil {
						return int64(len(jsTable.Table.Rows))
					}
				}
			}
			return -1
		} else {
			return -1
		}
	}
}

// GetSummaryMapUri returns the HTTP URL to a summary geographic map of the
// Erddap dataset as a string, uesful for creating thumbnail images.
func (dataset EDDDataset) GetSummaryMapUri() string {
	uri := dataset.Uri
	uri = strings.Split(uri, "erddap/")[0]
	if dataset.EddDatastType == EDD_TABLE {
		return uri + "erddap/tabledap/" + dataset.Id + ".png?longitude,latitude&.draw=markers&.marker=3%7C5&.color=0xFF9900&.colorBar=%7C%7C%7C%7C%7C&.bgColor=0xffccccff"
	} else {
		return uri + "erddap/griddap/" + dataset.Id + ".png?longitude,latitude&.draw=markers&.marker=3%7C5&.color=0xFF9900&.colorBar=%7C%7C%7C%7C%7C&.bgColor=0xffccccff"
	}
}

// ToSTACItem(baseURL) exports a stac.Item based on the entire EDDDataset.
// As a STAC Item requires an absolute URI, baseURL provides the absolute URI reference to the parent
// folder of the JSON file that the stac.Collection will be serialised to.
func (dataset EDDDataset) ToSTACItem(baseURL string) stac.Item {
	item := stac.NewItem()

	item.Id = dataset.ItemGUIDs[0].String()

	item.Bbox = append(item.Bbox, dataset.BoundingBoxMinLon)
	item.Bbox = append(item.Bbox, dataset.BoundingBoxMinLat)
	item.Bbox = append(item.Bbox, dataset.BoundingBoxMaxLon)
	item.Bbox = append(item.Bbox, dataset.BoundingBoxMaxLat)

	item.Properties.Title = dataset.Name
	item.Properties.Description = dataset.Description

	item.Properties.StartTime = dataset.TimeMin.Format(EDD_TIME_LAYOUT)
	item.Properties.EndTime = dataset.TimeMax.Format(EDD_TIME_LAYOUT)

	item.Properties.Created = dataset.CreatedTime.Format(EDD_TIME_LAYOUT)
	item.Properties.Updated = dataset.ModifiedTime.Format(EDD_TIME_LAYOUT)

	md := stac.Asset{
		Href:  dataset.GetMetadataUri(),
		Type:  stac.STAC_ASSET_MEDIA_TYPE_XML,
		Title: "ISO19115 metadata",
		Roles: []string{},
	}
	md.Roles = append(md.Roles, stac.STAC_ASSET_ROLE_METADATA)
	item.Assets["metadata"] = md

	if dataset.MetadataLink != "" {
		md2 := stac.Asset{
			Href:  dataset.MetadataLink,
			Title: "Further Information",
			Roles: []string{},
		}
		md2.Roles = append(md2.Roles, stac.STAC_ASSET_ROLE_OVERVIEW)
		item.Assets["further_information"] = md2
	} else if dataset.InfoUrl != "" {
		md2 := stac.Asset{
			Href:  dataset.InfoUrl,
			Title: "Further Information",
			Roles: []string{},
		}
		md2.Roles = append(md2.Roles, stac.STAC_ASSET_ROLE_OVERVIEW)
		item.Assets["further_information"] = md2
	}

	d := stac.Asset{
		Href:  dataset.GetDataUri(ERDDAP_FORMAT_JSON),
		Type:  stac.STAC_ASSET_MEDIA_TYPE_JSON,
		Title: "JSON Data",
		Roles: []string{},
	}
	d.Roles = append(d.Roles, stac.STAC_ASSET_ROLE_DATA)
	item.Assets["json_data"] = d

	d2 := stac.Asset{
		Href:  dataset.GetDataUri(ERDDAP_FORMAT_NETCDF),
		Type:  stac.STAC_ASSET_MEDIA_TYPE_NETCDF,
		Title: "NetCDF Data",
		Roles: []string{},
	}
	d2.Roles = append(d2.Roles, stac.STAC_ASSET_ROLE_DATA)
	item.Assets["netcdf_data"] = d2

	thumb := stac.Asset{
		Href:  dataset.GetSummaryMapUri(),
		Type:  stac.STAC_ASSET_MEDIA_TYPE_PNG,
		Title: "Thumbnail map",
		Roles: []string{},
	}
	thumb.Roles = append(thumb.Roles, stac.STAC_ASSET_ROLE_THUMBNAIL)
	item.Assets["thumbnail_map"] = thumb

	item.Properties.Keywords = dataset.Keywords

	//TODO: We can get a better geographic envelope for the dataset by
	//interrogating the dataset on the Erddap server
	item.Geometry.Type = stac.GEOJSON_TYPE_POLYGON
	bottom_left := [2]float32{dataset.BoundingBoxMinLon, dataset.BoundingBoxMinLat}
	top_left := [2]float32{dataset.BoundingBoxMinLon, dataset.BoundingBoxMaxLat}
	top_right := [2]float32{dataset.BoundingBoxMaxLon, dataset.BoundingBoxMaxLat}
	bottom_right := [2]float32{dataset.BoundingBoxMaxLon, dataset.BoundingBoxMinLat}
	item.Geometry.Coordinates = []interface{}{bottom_left, top_left, top_right, bottom_right, bottom_left}

	if dataset.EddDatastType == EDD_TABLE {
		item.Properties.TableColumns = dataset.Variables
		item.Properties.TableRowCount = dataset.GetTableRowCount()
	}

	links := []stac.Link{}
	links = append(links, stac.Link{Href: "./" +
		dataset.CatalogGUID.String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_ROOT,
		Type: stac.STAC_CATALOG_MIME_TYPE})
	links = append(links, stac.Link{Href: "./" +
		dataset.CollectionGUID.String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_COLLECTION,
		Type: stac.STAC_COLLECTION_MIME_TYPE})
	links = append(links, stac.Link{Href: "./" +
		dataset.CollectionGUID.String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_PARENT,
		Type: stac.STAC_COLLECTION_MIME_TYPE})
	links = append(links, stac.Link{Href: baseURL +
		dataset.ItemGUIDs[0].String() + ".json",
		Rel:  stac.STAC_LINK_RELTYPE_SELF,
		Type: stac.STAC_ITEM_MIME_TYPE})

	item.Links = links

	item.Collection = dataset.CollectionGUID.String()

	return item
}
