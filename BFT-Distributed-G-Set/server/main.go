package main

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/modules"
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"os"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)
	tools.ResetLogFile()
	hosts_filename := "/users/loukis/Thesis/BFT-Distributed-G-Set/hosts"

	peers := config.GetHosts(hosts_filename, "servers")

	config.N = len(peers) + 1
	config.F = (config.N - 1) / 3

	s := server.CreateServer(peers)

	if len(os.Args) < 2 {
		modules.StartNormal(s)
	} else {
		behaviour := os.Args[1]
		switch behaviour {
		case "normal":
			modules.StartNormal(s)
		case "mute":
			modules.StartMute(s)
		case "malicious":
			modules.StartMute(s)
		default:
			panic("Invalid argument")
		}
	}

}
