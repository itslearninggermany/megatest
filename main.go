package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	awsBrooker "github.com/itslearninggermany/itswizard_m_awsbrooker"
	itswizard_basic "github.com/itslearninggermany/itswizard_m_basic"
	"github.com/jinzhu/gorm"
)

var (
	allDatabases  map[string]*gorm.DB
	dbWebserver         *gorm.DB
	dbClient            *gorm.DB
)


func init() {

	// Datenbanken einlesen
	var databaseConfig []itswizard_basic.DatabaseConfig
	b, _ := awsBrooker.DownloadFileFromBucket("brooker", "admin/databaseconfig.json")
	err := json.Unmarshal(b, &databaseConfig)
	if err != nil {
		fmt.Println("Error by reading database file " + err.Error())
		return
	}
	allDatabases = make(map[string]*gorm.DB)
	for i := 0; i < len(databaseConfig); i++ {
		database, err := gorm.Open(databaseConfig[i].Dialect, databaseConfig[i].Username+":"+databaseConfig[i].Password+"@tcp("+databaseConfig[i].Host+")/"+databaseConfig[i].NameOrCID+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println(err)
			return
		}
		allDatabases[databaseConfig[i].NameOrCID] = database
	}
	dbWebserver = allDatabases["Webserver"]
	dbClient = allDatabases["Client"]
	// Datenbank einlesen ende
}


func main() {
	var user itswizard_basic.LusdPerson
	allDatabases["43"].Find(&user)

	fmt.Println(user)
}
