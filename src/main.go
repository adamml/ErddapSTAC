// ErddapSTAC project main.go
package main

import (
	"ErddapSTAC/src/erddap"
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

	l, e := http.Get("https://raw.githubusercontent.com/" +
		"IrishMarineInstitute/awesome-erddap/refs/heads/master/erddaps.json")
	if e == nil {
		lr, e := ioutil.ReadAll(l.Body)
		allErddaps := []ErddapInstance{}
		if e == nil {
			err := json.Unmarshal(lr, &allErddaps)
			if err == nil {
				for i := 0; i < len(allErddaps); i++ {
					r, err := http.Get(allErddaps[i].URL + "tabledap/" +
						"allDatasets.json?datasetID%2Cinstitution%2C" +
						"dataStructure%2Ctitle%2CminTime%2C" +
						"maxTime%2CinfoUrl%2Cemail%2Csummary")
					if err == nil {
						c, err := ioutil.ReadAll(r.Body)
						if err == nil {
							err = json.Unmarshal(c, &t)
							if err == nil {

							} else {
								fmt.Println("Server not found: " + allErddaps[i].URL)
							}
						} else {
							fmt.Println("Server not found: " + allErddaps[i].URL)
						}
					} else {
						fmt.Println("Server not found: " + allErddaps[i].URL)
					}
				}
			}
		}
	}

	r, _ := http.Get("https://erddap.marine.ie/erddap/tabledap/allDatasets" +
		".json?datasetID%2Cinstitution%2CdataStructure%2Ctitle%2CminTime%2C" +
		"maxTime%2CinfoUrl%2Cemail%2Csummary")
	c, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(c, &t)

	if t.Table.Rows[9][2] == "table" {
		eddType = erddap.EDD_TABLE
	} else {
		eddType = erddap.EDD_GRID
	}

	startTime, _ = time.Parse(erddap.EDD_TIME_LAYOUT, t.Table.Rows[9][4])
	endTime, _ = time.Parse(erddap.EDD_TIME_LAYOUT, t.Table.Rows[9][5])

	catalogGUID, _ := uuid.NewV4()

	edd := erddap.NewEDDDataset(
		"https://erddap.marine.ie/erddap/info/imipublicunderway/index.json",
		t.Table.Rows[31][0], t.Table.Rows[31][3], t.Table.Rows[31][8],
		eddType, catalogGUID, t.Table.Rows[31][1], t.Table.Rows[31][6],
		startTime, endTime)
	stac_c := edd.ToSTACCollection(baseURL)
	stac_i := edd.ToSTACItem(baseURL)
	a, _ := json.MarshalIndent(stac_c, "", "    ")
	b, _ := json.MarshalIndent(stac_i, "", "    ")

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

	ioutil.WriteFile(outdir+edd.CollectionGUID.String()+".json", a, os.ModePerm)
	//TODO: Fix urlencoding prior to writing for download links
	ioutil.WriteFile(outdir+edd.ItemGUIDs[0].String()+".json", b, os.ModePerm)
}
