package openfiles

import (
	"fmt"
	"os"
)

func OpenFiles(serverListFile, downTimeFile string) (*os.File, *os.File) {
	serverList, err := os.OpenFile(serverListFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// .O_APPEND adiciona dados em vez de sobreescrever | O_CREATE caso o arquivo não exista ele cria
	downTimeList, err := os.OpenFile(downTimeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return serverList, downTimeList
}
