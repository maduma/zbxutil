package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	zabbix_server_url string
	zabbi_api_token   string
	logger            *log.Logger
)

func get_config() {
	zabbix_server_url = os.Getenv("ZABBIX_SERVER_URL")
	zabbi_api_token = os.Getenv("ZABBIX_API_TOKEN")
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	if zabbix_server_url == "" || zabbi_api_token == "" {
		panic(errors.New("Environment variable ZABBIX_SERVER_URL or ZABBIX_API_TOKEN not set"))
	}
}

// set usergoupid rights to read-only for all groups
func setUgrpRoRights() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: zbxutil usergroupid1,usergroupid2,...")
	} else {
		rights := allGrpRoRights()
		ugrpids := strings.Split(os.Args[1], ",")
		for _, ugrpid := range ugrpids {
			usergroupid, err := strconv.Atoi(ugrpid)
			if err != nil {
				fmt.Printf("usergroupid \"%s\" sould be a integer\n", ugrpid)
			} else {
				setUgrpRight(usergroupid, rights)
			}
		}
	}
}

func main() {
	get_config()
	setUgrpRoRights()
}
