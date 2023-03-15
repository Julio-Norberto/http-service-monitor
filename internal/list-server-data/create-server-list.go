package listserverdata

type Server struct {
	ServerName    string
	ServerUrl     string
	ExecutionTime float64
	StatusCode    int
}

func CreateServerList(data [][]string) []Server {
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
