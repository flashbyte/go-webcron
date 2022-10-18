package main

import (
	"fmt"
	"os"
	"strconv"
	"net/http"
	"time"
//	"io"
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
		value, _ := strconv.ParseInt(value, 10, 64) // TODO: Add error handling
		setting.time = value
	}

	if value, exist := os.LookupEnv("DEBUG"); exist {
		value, _ := strconv.ParseBool(value) // TODO: Add error handling
		setting.debug = value
	}

	return setting, false
}

func doRequest(url string){
	resp, _ := http.Get(url) // TODO: Add error handling
	//body, _ := io.ReadAll(resp.Body)
	//resp.Body.Close()
	fmt.Printf("GET %v -- %v\n", url, resp.StatusCode)
}

func main() {
	env, err := parseEnv()
	if err {
		os.Exit(1)
	}

	fmt.Printf("Starting %v with config (URL:%v TIME:%v DEBUG:%v)\n", "go-webcron", env.url, env.time, env.debug)

	for true {
		doRequest(env.url)
		time.Sleep(time.Duration(env.time) * time.Second) 
	}
}
