package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.86.2"
	port     = 5432
	user     = "pi"
	password = "Boomer2025"
	dbname   = "incident"
)

type Response struct {
	ObjectIDFieldName string `json:"objectIdFieldName"`
	UniqueIDField     struct {
		Name               string `json:"name"`
		IsSystemMaintained bool   `json:"isSystemMaintained"`
	} `json:"uniqueIdField"`
	GlobalIDFieldName string `json:"globalIdFieldName"`
	Fields            []struct {
		Name         string      `json:"name"`
		Type         string      `json:"type"`
		Alias        string      `json:"alias"`
		SQLType      string      `json:"sqlType"`
		Domain       interface{} `json:"domain"`
		DefaultValue interface{} `json:"defaultValue"`
		Length       int         `json:"length,omitempty"`
	} `json:"fields"`
	ExceededTransferLimit bool `json:"exceededTransferLimit"`
	Features              []struct {
		Attributes struct {
			KeyCrash                  int    `json:"key_crash"`
			Objectid                  int    `json:"OBJECTID"`
			City                      string `json:"City"`
			DLClass                   string `json:"DLClass"`
			DLRestrictions            string `json:"DLRestrictions"`
			DLState                   string `json:"DLState"`
			Cdl                       string `json:"CDL"`
			Age                       string `json:"Age"`
			VehicleSeizure            string `json:"VehicleSeizure"`
			AlcoholSuspected          string `json:"AlcoholSuspected"`
			AlcoholTest               string `json:"AlcoholTest"`
			AlcoholResultType         string `json:"AlcoholResultType"`
			AirbagSwitch              string `json:"AirbagSwitch"`
			AirbagDeployed            string `json:"AirbagDeployed"`
			Ejection                  string `json:"Ejection"`
			Gender                    string `json:"Gender"`
			Race                      string `json:"Race"`
			Injury                    string `json:"Injury"`
			Protection                string `json:"Protection"`
			Trapped                   string `json:"Trapped"`
			PersonType                string `json:"PersonType"`
			VisionObstruction         string `json:"VisionObstruction"`
			ContributingCircumstance1 string `json:"ContributingCircumstance1"`
			ContributingCircumstance2 string `json:"ContributingCircumstance2"`
			ContributingCircumstance3 string `json:"ContributingCircumstance3"`
			VehicleType               string `json:"VehicleType"`
			CrashDate                 int64  `json:"crash_date"`
		} `json:"attributes"`
	} `json:"features"`
}

func main() {

	//-----------------------------------------------------------------------
	//--- Load data from the API
	var ApiURL = "https://utility.arcgis.com/usrsvcs/servers/35922646263e4769b171317965815ec8/rest/services/Crash_Reports/FeatureServer/1/query?where=1%3D1&outFields=*&outSR=4326&f=json"
	//var ApiURL = "https://utility.arcgis.com/usrsvcs/servers/35922646263e4769b171317965815ec8/rest/services/Crash_Reports/FeatureServer/1/query?where=1%3D1&outFields=*&outSR=4326&f=json"
	resp, err := http.Get(ApiURL)
	if err != nil {
		fmt.Println("No response from request")
	}
	//--- Load all the JSON as text into "body"
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	//--- Unmarshal the JSON into the "result" value
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Make and open the database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//--- Are we good?
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Loop through the API results and run the inserts
	for _, rec := range result.Features {

		var KeyCrash = rec.Attributes.KeyCrash
		var City = rec.Attributes.City
		var DLClass = rec.Attributes.DLClass
		var DLRestrictions = rec.Attributes.DLRestrictions
		var DLState = rec.Attributes.DLState
		var Cdl = rec.Attributes.Cdl
		var Age = rec.Attributes.Age
		var VehicleSeizure = rec.Attributes.VehicleSeizure
		var AlcoholSuspected = rec.Attributes.AlcoholSuspected
		var AlcoholTest = rec.Attributes.AlcoholTest
		var AlcoholResultType = rec.Attributes.AlcoholResultType
		var AirbagSwitch = rec.Attributes.AirbagSwitch
		var AirbagDeployed = rec.Attributes.AirbagDeployed
		var Ejection = rec.Attributes.Ejection
		var Gender = rec.Attributes.Gender
		var Race = rec.Attributes.Race
		var Injury = rec.Attributes.Injury
		var Protection = rec.Attributes.Protection
		var Trapped = rec.Attributes.Trapped
		var PersonType = rec.Attributes.PersonType
		var VisionObstruction = rec.Attributes.VisionObstruction
		var ContributingCircumstance1 = rec.Attributes.ContributingCircumstance1
		var ContributingCircumstance2 = rec.Attributes.ContributingCircumstance2
		var ContributingCircumstance3 = rec.Attributes.ContributingCircumstance3
		var VehicleType = rec.Attributes.VehicleType
		var CrashDateMilli = rec.Attributes.CrashDate

		CrashDate := time.UnixMilli(CrashDateMilli)
		//fmt.Println(CrashDate)

		var sql = "CALL add_crash_person ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26);"

		_, err := db.Exec(sql,
			KeyCrash,
			City,
			DLClass,
			DLRestrictions,
			DLState,
			Cdl,
			Age,
			VehicleSeizure,
			AlcoholSuspected,
			AlcoholTest,
			AlcoholResultType,
			AirbagSwitch,
			AirbagDeployed,
			Ejection,
			Gender,
			Race,
			Injury,
			Protection,
			Trapped,
			PersonType,
			VisionObstruction,
			ContributingCircumstance1,
			ContributingCircumstance2,
			ContributingCircumstance3,
			VehicleType,
			CrashDate)
		if err != nil {
			panic(err.Error())
		}
	}
	//-----------------------------------------------------------------------

}
