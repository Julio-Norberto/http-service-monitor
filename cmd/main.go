package main

import (
	"os"

	listserverdata "github.com/Julio-Norberto/http-service-monitor/internal/list-server-data"
	generatedowntimelist "github.com/Julio-Norberto/http-service-monitor/pkg/generateDownTimeList"
	openfiles "github.com/Julio-Norberto/http-service-monitor/pkg/openFiles"
	returnservers "github.com/Julio-Norberto/http-service-monitor/pkg/returnServers"
)

func main() {
	serverList, downTimeList := openfiles.OpenFiles(os.Args[1], os.Args[2])
	defer serverList.Close()
	defer downTimeList.Close()
	servers := listserverdata.CreateServerList(serverList)
	downServers := returnservers.PrintServerList(servers)
	generatedowntimelist.GenerateDownTimeServers(downTimeList, downServers)

	returnservers.PrintServerList(servers)
}
