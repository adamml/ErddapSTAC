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

	r, _ := http.Get("https://erddap.marine.ie/erddap/tabledap/allDatasets.json?datasetID%2Cinstitution%2CdataStructure%2Ctitle%2CminTime%2CmaxTime%2CinfoUrl%2Cemail%2Csummary")
	c, _ := ioutil.ReadAll(r.Body)
	var t erddap.EDDJsonTable
	var eddType string
	_ = json.Unmarshal(c, &t)

	if t.Table.Rows[74][2] == "table" {
		eddType = erddap.EDD_TABLE
	} else {
		eddType = erddap.EDD_GRID
	}

	edd := erddap.NewEDDDataset(
		"https://erddap.marine.ie/erddap/info/bunaveela_water_temp_profiles/index.json",
		t.Table.Rows[74][0], t.Table.Rows[74][3], t.Table.Rows[74][8],
		eddType, t.Table.Rows[74][1], t.Table.Rows[74][6])
	stac_c := erddap.EDDDatasetToSTACCollection(edd)
	a, _ := json.Marshal(stac_c)
	fmt.Println(string(a))
}
