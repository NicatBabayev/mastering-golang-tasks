package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type config struct {
	Database string  `xml:"database"`
	Username string  `xml:"username"`
	Password string  `xml:"password"`
	Options  options `xml:"options"`
}

type options struct {
	Auto_backup     bool `xml:"auto_backup"`
	Max_connections int  `xml:"max_connections"`
}

func writeFile(fileName string, content []byte) error {
	fileName = fmt.Sprintf("task/xml/%s", fileName)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func createXML(config config, fileName string) error {
	content, err := xml.Marshal(&config)
	if err != nil {
		return err
	}
	err = writeFile(fileName, content)
	if err != nil {
		return err
	}
	return nil
}

func RunTask6() {
	config := config{
		Database: "mydb",
		Username: "admin",
		Password: "secret",
		Options: options{
			Auto_backup:     true,
			Max_connections: 100,
		},
	}
	err := createXML(config, "config.xml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config is created succesfully.")
	}

}
