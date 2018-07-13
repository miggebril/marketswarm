package settings

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

var settings Settings = Settings{}

var AppRoot string

func Init(appRootPath string) {
	AppRoot = appRootPath

	settingsFilePath := filepath.Join(appRootPath, "settings.xml")
	settingsFile, errOpen := os.Open(settingsFilePath)
	settingsData, errRead := ioutil.ReadAll(settingsFile)
	defer settingsFile.Close()

	if errOpen == nil && errRead == nil {
		xml.Unmarshal(settingsData, &settings)
	} else {
		fmt.Println("Could not load settings.  Using defaults.", errOpen, errRead)
		settings = Settings{
			DocumentStoragePath: "~/emapfiles",
			Server: SettingsServer{
				Host: "localhost",
				Key:  "abc123",
			},
			Database: SettingsDB{
				Engine:   "sqlserver",
				Host:     "127.0.0.1",
				Port:     "8080",
				DBName:   "Markets",
				Username: "sa",
				Password: "password",
			},
		}
	}
}

func Get() Settings {
	if &settings == nil {
		Init("settings.xml")
	}
	return settings
}

type SettingsDB struct {
	XMLName  xml.Name `xml:"database"`
	Engine   string   `xml:"engine"`
	Host     string   `xml:"host"`
	Port     string   `xml:"port"`
	DBName   string   `xml:"dbname"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

type SettingsServer struct {
	XMLName xml.Name `xml:"server"`
	Host    string   `xml:"host"`
	Port    string   `xml:"port"`
	Key     string   `xml:"key"`
}

type Settings struct {
	XMLName             xml.Name       `xml:"settings"`
	DocumentStoragePath string         `xml:"documentstorage"`
	Server              SettingsServer `xml:"server"`
	Database            SettingsDB     `xml:"database"`

	Notifications		SettingsNotifications `xml:"notifications"`
}
