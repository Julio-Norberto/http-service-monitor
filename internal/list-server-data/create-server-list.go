package listserverdata

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Server struct {
	ServerName    string
	ServerUrl     string
	ExecutionTime float64
	StatusCode    int
	FailDate      string
}

func CreateServerList(serverList *os.File) []Server {
	csvReader := csv.NewReader(serverList)
	serverData, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var allServers []Server

	for i, line := range serverData {
		if i > 0 {
			servers := Server{
				ServerName: line[0],
				ServerUrl:  line[1],
			}
			allServers = append(allServers, servers)
		}
	}
	return allServers
}
