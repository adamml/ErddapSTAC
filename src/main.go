// ErddapSTAC project main.go
package main

import (
	"ErddapSTAC/src/erddap"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	var t erddap.EDDJsonTable
	var eddType string
	var startTime time.Time
	var endTime time.Time

	baseURL := "https://raw.githubusercontent.com/adamml/ErddapSTAC/json/"

	r, _ := http.Get("https://linkedsystems.uk/erddap/tabledap/allDatasets" +
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

	edd := erddap.NewEDDDataset(
		"https://linkedsystems.uk/erddap/info/Amazon_622_R/index.json",
		t.Table.Rows[9][0], t.Table.Rows[9][3], t.Table.Rows[9][8],
		eddType, t.Table.Rows[9][1], t.Table.Rows[9][6],
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

	ioutil.WriteFile(outdir+stac_c.Id+".json", a, os.ModePerm)
	ioutil.WriteFile(outdir+stac_i.Id+".json", b, os.ModePerm)
}
