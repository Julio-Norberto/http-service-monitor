package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	ServerName    string
	ServerUrl     string
	ExecutionTime float64
}

func createServerList(data [][]string) []Server {
	var allServers []Server

	for i, line := range data {
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

func main() {
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

	servers := createServerList(serverData)
	for _, server := range servers {
		now := time.Now()
		get, err := http.Get(server.ServerUrl)
		if err != nil {
			fmt.Println(err)
		}

		status := get.StatusCode
		server.ExecutionTime = float64(time.Since(now).Seconds())

		fmt.Printf("Status: [%d] Tempo de carga: [%f] Server: [%s]\n", status, server.ExecutionTime, server.ServerUrl)
		//fmt.Println(server)
	}
}
