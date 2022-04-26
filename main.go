package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	// Colors
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	blue   = "\x1b[34m"
	reset  = "\x1b[0m"
)

func main() {
	// Get the file name
	filenameSplit := strings.Split(os.Args[0], "/")
	filename := filenameSplit[len(filenameSplit)-1]

	// Make sure the user has the correct number of arguments
	if len(os.Args) < 4 {
		fmt.Printf("Usage: ./%v <ip> <port> <count> (-1 if forever) <delay ex: 5s, 5m, 2s, 2h> (optional)\n", filename)
		return
	}

	// Get OS args
	var timeout time.Duration
	ip := os.Args[1]
	port := os.Args[2]
	count, _ := strconv.Atoi(os.Args[3])

	// Make timeout optional
	if len(os.Args) == 5 {
		timeout, _ = time.ParseDuration(os.Args[4])
	}

	// Connect to the host for count times
	for i := 0; ; i++ {
		// If count is -1, loop forever
		// Otherwise, break when count is reached
		if count != -1 && i >= count {
			break
		}

		beforeTime := time.Now()                                         // Get current time before connection
		conn, err := net.Dial("tcp", ip+":"+port)                        // Connect to host
		roundedEndTime := time.Since(beforeTime).Round(time.Millisecond) // Round time since the connection was initiated to milliseconds
		conn.Close()                                                     // Close connection

		// If GOOS is windows, do not use colors
		if runtime.GOOS == "windows" {
			red = ""
			green = ""
			yellow = ""
			blue = ""
			reset = ""
		}

		if err == nil {
			fmt.Printf("%vConnected to %v:%v time=%v\n%v", green, ip, port, roundedEndTime, reset)
		} else {
			fmt.Printf("%vFailed to connect to %v:%v time=%v\n%v", red, ip, port, roundedEndTime, reset)
		}

		time.Sleep(timeout | time.Second) // Sleep for the specified timeout or 1 second
	}
}
