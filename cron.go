package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Settings struct {
	url   string
	time  int64
	debug bool
}

func parseEnv() (Settings, bool) {
	setting := Settings{
		time:  300,
		debug: false,
	}

	if value, exist := os.LookupEnv("URL"); exist {
		setting.url = value
	} else {
		fmt.Println("URL requierd")
		return setting, true
	}

	if value, exist := os.LookupEnv("TIME"); exist {
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("CRITICAL: Could not pare env TIME")
			return setting, true
		}
		setting.time = value
	}

	if value, exist := os.LookupEnv("DEBUG"); exist {
		value, err := strconv.ParseBool(value)
		if err != nil {
			fmt.Println("CRITICAL: Could not pare env DEBUG")
			return setting, true
		}

		setting.debug = value
	}

	return setting, false
}

func doRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: Cound't get %v\n", url)
		return
	}
	fmt.Printf("GET %v -- %v\n", url, resp.StatusCode)
}

func main() {
	env, err := parseEnv()
	if err {
		os.Exit(1)
	}

	fmt.Printf("Starting %v with config (URL:%v TIME:%v DEBUG:%v)\n", "go-webcron", env.url, env.time, env.debug)

	for {
		doRequest(env.url)
		time.Sleep(time.Duration(env.time) * time.Second)
	}
}
