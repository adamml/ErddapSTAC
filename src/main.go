// ErddapSTAC project main.go
package main

import (
	"ErddapSTAC/src/erddap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	var t erddap.EDDJsonTable
	var eddType string
	var startTime time.Time
	var endTime time.Time

	r, _ := http.Get("https://erddap.marine.ie/erddap/tabledap/allDatasets.json?datasetID%2Cinstitution%2CdataStructure%2Ctitle%2CminTime%2CmaxTime%2CinfoUrl%2Cemail%2Csummary")
	c, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(c, &t)

	if t.Table.Rows[74][2] == "table" {
		eddType = erddap.EDD_TABLE
	} else {
		eddType = erddap.EDD_GRID
	}

	fmt.Println(t.Table.Rows[74][6])
	startTime, _ = time.Parse(erddap.EDD_TIME_LAYOUT, t.Table.Rows[74][4])
	fmt.Println(startTime)
	endTime, _ = time.Parse(erddap.EDD_TIME_LAYOUT, t.Table.Rows[74][5])

	edd := erddap.NewEDDDataset(
		"https://erddap.marine.ie/erddap/info/bunaveela_water_temp_profiles/index.json",
		t.Table.Rows[74][0], t.Table.Rows[74][3], t.Table.Rows[74][8],
		eddType, t.Table.Rows[74][1], t.Table.Rows[74][6],
		startTime, endTime)
	stac_c := erddap.EDDDatasetToSTACCollection(edd)
	a, _ := json.Marshal(stac_c)
	fmt.Println(string(a))
}
