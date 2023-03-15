package returnservers

import (
	"fmt"
	"net/http"
	"time"

	listserverdata "github.com/Julio-Norberto/http-service-monitor/internal/list-server-data"
)

func PrintServerList(servers []listserverdata.Server) {
	for _, server := range servers {
		now := time.Now()
		get, err := http.Get(server.ServerUrl)
		if err != nil {
			fmt.Println(err)
		}

		server.StatusCode = get.StatusCode
		server.ExecutionTime = float64(time.Since(now).Seconds())

		fmt.Printf("Status: [%d] Tempo de carga: [%f] Server: [%s]\n", server.StatusCode, server.ExecutionTime, server.ServerUrl)
		//fmt.Println(server)
	}
}
