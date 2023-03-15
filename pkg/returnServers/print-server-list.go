package returnservers

import (
	"fmt"
	"net/http"
	"time"

	listserverdata "github.com/Julio-Norberto/http-service-monitor/internal/list-server-data"
)

func PrintServerList(servers []listserverdata.Server) []listserverdata.Server {
	var downServers []listserverdata.Server

	for _, server := range servers {
		now := time.Now()
		get, err := http.Get(server.ServerUrl)
		if err != nil {
			fmt.Printf("Server %s is down %s", server.ServerName, err.Error())
			server.StatusCode = 0
			server.FailDate = now.Format("02/01/2023 15:04:05")
			downServers = append(downServers, server)
			continue
		}

		server.StatusCode = get.StatusCode
		if server.StatusCode != 200 {
			server.FailDate = now.Format("02/01/2023 15:04:05")
			downServers = append(downServers, server)
		}

		server.ExecutionTime = float64(time.Since(now).Seconds())

		fmt.Printf("Status: [%d] Tempo de carga: [%f] Server: [%s]\n", server.StatusCode, server.ExecutionTime, server.ServerUrl)
		//fmt.Println(server)
	}

	return downServers
}
