package generatedowntimelist

import (
	"encoding/csv"
	"fmt"
	"os"

	listserverdata "github.com/Julio-Norberto/http-service-monitor/internal/list-server-data"
)

func GenerateDownTimeServers(downTimeList *os.File, downServers []listserverdata.Server) {
	csvWriter := csv.NewWriter(downTimeList)
	for _, server := range downServers {
		line := []string{server.ServerName, server.ServerUrl, server.FailDate, fmt.Sprintf("%f", server.ExecutionTime), fmt.Sprintf("%d", server.StatusCode)}
		csvWriter.Write(line)
	}

	csvWriter.Flush()
}
