// ErddapSTAC project main.go
package main

import (
	"ErddapSTAC/src/erddap"
	"ErddapSTAC/src/stac"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
)

// Define a structure type to unmarshal the list of Erddap servers to contact
type ErddapInstance struct {
	Name      string
	ShortName string
	URL       string
	Public    bool
}

func main() {

	var t erddap.EDDJsonTable
	var eddType string
	var startTime time.Time
	var endTime time.Time

	baseURL := ("https://raw.githubusercontent.com/adamml/ErddapSTAC/refs/" +
		"heads/main/json/")

	var baseCatalog stac.Catalog
	baseCatalog.Id = "erddap_stac"
	baseCatalog.Title = "A STAC Catalog of instances of NOAA's Erddap environmental data server"
	baseCatalog.StacVersion = stac.STAC_VERSION
	baseLinks := []stac.Link{}
	baseCatalog.Type = stac.STAC_TYPES_CATALOG
	baseLinks = append(baseLinks, stac.Link{
		Href: baseURL + "catalog.json",
		Rel:  stac.STAC_LINK_RELTYPE_SELF,
		Type: stac.STAC_CATALOG_MIME_TYPE,
	})

	l, e := http.Get("https://raw.githubusercontent.com/" +
		"IrishMarineInstitute/awesome-erddap/refs/heads/master/erddaps.json")

	pwd, _ := os.Getwd()
	var outdir string
	if strings.Index(pwd, "\\") > 0 {
		pwdSlice := strings.Split(pwd, "\\")
		for k := 0; k < len(pwdSlice)-1; k++ {
			outdir = outdir + pwdSlice[k] + "\\"
		}
		outdir = outdir + "json\\"
	} else {
		pwdSlice := strings.Split(pwd, "/")
		for k := 0; k < len(pwdSlice)-1; k++ {
			outdir = outdir + pwdSlice[k] + "/"
		}
		outdir = outdir + "json/"
	}

	if e == nil {
		lr, e := ioutil.ReadAll(l.Body)
		allErddaps := []ErddapInstance{}
		if e == nil {
			err := json.Unmarshal(lr, &allErddaps)
			if err == nil {
				for i := 0; i < len(allErddaps); i++ {
					fmt.Println()
					catalogGUID, _ := uuid.NewV4()
					r, err := http.Get(allErddaps[i].URL + "tabledap/" +
						"allDatasets.json?datasetID%2Cinstitution%2C" +
						"dataStructure%2Ctitle%2CminTime%2C" +
						"maxTime%2CinfoUrl%2Cemail%2Csummary")

					baseLinks = append(baseLinks, stac.Link{
						Href:  "./" + catalogGUID.String() + ".json",
						Rel:   stac.STAC_LINK_RELTYPE_CHILD,
						Type:  stac.STAC_CATALOG_MIME_TYPE,
						Title: allErddaps[i].Name})

					thisCatalog := stac.Catalog
					thisCatalog.Id = catalogGUID.String()
					thisCatalog.Title = allErddaps[i].Name
					thisCatalog.Description = ("A catalog of all datasets on the " +
						allErddaps[i].Name + " Erddap Server")
					thisCatalog.StacVersion = stac.STAC_VERSION
					thisCatalogLinks := []stac.Link{}
					thisCatalogLinks = append(thisCatalogLinks, stac.Link{
						Href: baseURL + catalogGUID.String() + ".json",
						Rel: stac.STAC_LINK_RELTYPE_SELF,
						Type: stac.STAC_CATALOG_MIME_TYPE
					})
					thisCatalogLinks = append(thisCatalogLinks, stac.Link{
						Href: "./catalog.json",
						Rel: stac.STAC_LINK_RELTYPE_PARENT,
						Type: stac.STAC_CATALOG_MIME_TYPE,
						Title = "A STAC Catalog of instances of NOAA's Erddap environmental data server"
					})
					thisCatalogLinks = append(thisCatalogLinks, stac.Link{
						Href: "./catalog.json",
						Rel: stac.STAC_LINK_RELTYPE_ROOT,
						Type: stac.STAC_CATALOG_MIME_TYPE,
						Title = "A STAC Catalog of instances of NOAA's Erddap environmental data server"
					})

					if err == nil {
						c, err := ioutil.ReadAll(r.Body)
						if err == nil {
							err = json.Unmarshal(c, &t)
							if err == nil {
								for j := 0; j < len(t.Table.Rows); j++ {
									if t.Table.Rows[j][2] == erddap.EDD_TABLE {
										eddType = erddap.EDD_TABLE
									} else if t.Table.Rows[j][2] == erddap.EDD_GRID {
										eddType = erddap.EDD_GRID
									}
									startTime, _ = time.Parse(
										erddap.EDD_TIME_LAYOUT,
										t.Table.Rows[j][4])
									endTime, _ = time.Parse(
										erddap.EDD_TIME_LAYOUT,
										t.Table.Rows[j][5])
									fmt.Println(catalogGUID.String() +
										"    Server: " + allErddaps[i].URL +
										"    Dataset: " + t.Table.Rows[j][0])

									edd := erddap.NewEDDDataset(allErddaps[i].URL+"info/"+t.Table.Rows[j][0]+"/index.json",
										t.Table.Rows[j][0],
										t.Table.Rows[j][3],
										t.Table.Rows[j][8],
										eddType,
										catalogGUID,
										t.Table.Rows[j][1],
										t.Table.Rows[j][6],
										startTime,
										endTime)

									stac_c := edd.ToSTACCollection(baseURL)
									stac_i := edd.ToSTACItem(baseURL)
									a, err := json.MarshalIndent(stac_c,
										"", "    ")
									if err == nil {
										ioutil.WriteFile(outdir+edd.CollectionGUID.String()+".json",
											a, os.ModePerm)
									}
									b, err := json.MarshalIndent(stac_i,
										"", "    ")
									if err == nil {
										//TODO: Fix urlencoding prior to writing for download links
										ioutil.WriteFile(outdir+edd.ItemGUIDs[0].String()+".json",
											b, os.ModePerm)
									}
								}
							} else {
								fmt.Println(
									"Server not found: " + allErddaps[i].URL)
							}
						} else {
							fmt.Println(
								"Server not found: " + allErddaps[i].URL)
						}
					} else {
						fmt.Println("Server not found: " + allErddaps[i].URL)
					}
				}
			}
		}
	}
	baseCatalog.Links = baseLinks
	a, err := json.MarshalIndent(baseCatalog, "", "    ")
	if err == nil {
		ioutil.WriteFile(outdir+"catalog.json",
			a, os.ModePerm)
	}
}
