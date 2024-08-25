package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultCheckInterval = 10
	defaultLogFile       = "connection_log.txt"
	testURL              = "http://www.google.com"
)

func main() {
	intervalPtr := flag.Int("interval", defaultCheckInterval, "Check interval in seconds")
	flag.IntVar(intervalPtr, "i", defaultCheckInterval, "Check interval in seconds (shorthand)")

	logFilePtr := flag.String("logfile", defaultLogFile, "Log file path")
	flag.StringVar(logFilePtr, "l", defaultLogFile, "Log file path (shorthand)")

	printOutputPtr := flag.Bool("print", false, "Print output to console")
	flag.BoolVar(printOutputPtr, "p", false, "Print output to console (shorthand)")

	flag.Parse()

	checkInterval := time.Duration(*intervalPtr) * time.Second

	file, err := os.OpenFile(*logFilePtr, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	if *printOutputPtr {
		fmt.Printf("Monitoring internet connection. Check interval: %v, Log file: %s\n", checkInterval, *logFilePtr)
	}

	var lastStatus bool
	for {
		connected := checkConnection()

		if !connected && lastStatus {
			logger.Println("Internet connection lost")

			if *printOutputPtr {
				fmt.Println("Internet connection lost")
			}
		} else if connected && !lastStatus {
			logger.Println("Internet connection restored")
			if *printOutputPtr {
				fmt.Println("Internet connection restored")
			}
		}

		lastStatus = connected
		time.Sleep(checkInterval)
	}
}

func checkConnection() bool {
	_, err := http.Get(testURL)
	return err == nil
}
