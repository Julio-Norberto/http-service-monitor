package main

import (
	"encoding/csv"
	"fmt"
	"os"

	listserverdata "github.com/Julio-Norberto/http-service-monitor/internal/list-server-data"
	openfiles "github.com/Julio-Norberto/http-service-monitor/pkg/openFiles"
	returnservers "github.com/Julio-Norberto/http-service-monitor/pkg/returnServers"
)

func main() {
	serverList, downTimeList := openfiles.OpenFiles(os.Args[1], os.Args[2])
	defer serverList.Close()
	defer downTimeList.Close()

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	serverData, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	servers := listserverdata.CreateServerList(serverData)
	returnservers.PrintServerList(servers)
}
