package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// the only thing i know is what i have to do
	// and the servers in the network
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]

	config_file := os.Args[1]
	data, err := os.ReadFile("../" + config_file)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	// all_servers := []string{}
	// for i := 0; i < N; i++ {
	// 	p := "node" + strconv.Itoa(i) + ":"
	// 	all_servers = append(all_servers, p+config.DEFAULT_PORT)
	// }
	// server.Create(hostname, all_servers)
	fmt.Println("Hello from server", hostname)

}
