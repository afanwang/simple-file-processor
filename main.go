package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// LogLevel is the level of logging
type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
)

var (
	infoLogger  *log.Logger
	debugLogger *log.Logger
)

// runMain is the main function
func runMain(args []string) {
	logger := log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Printf("Starting %s", "A Coding Challenge")
	// Open the CSV file for reading
	file, err := os.Open("server_log.csv")
	if err != nil {
		logger.Println("Error:", err)
		return
	}
	defer file.Close()

	// Store Results
	users := make(map[string]int)
	uploadsGT50KB := 0
	jeffUploadOnApril15th2020 := 0

	// Create a new CSV reader
	reader := csv.NewReader(bufio.NewReader(file))

	// Read each line of the file and analyze it
	firstLine := true
	for {
		line, err := reader.Read()
		if err != nil {
			// logger.Printf("Error: %v", err)
			break
		}

		if firstLine {
			firstLine = false
			continue
		}

		// Analyze the line here...
		timestamp, err := time.Parse("Mon Jan 02 15:04:05 MST 2006", line[0])
		if err != nil {
			// fmt.Println("Error parsing timestamp:", err)
			continue
		}

		if users[line[1]] == 0 {
			users[line[1]] = 1
		} else {
			users[line[1]]++
		}

		size, errConv := strconv.Atoi(line[3])
		if errConv != nil {
			logger.Println("Error strconv:", errConv)
		}
		if line[2] == "upload" && size > 50 {
			uploadsGT50KB++
		}

		if line[1] == "jeff22" &&
			line[2] == "upload" &&
			timestamp.Day() == 15 && timestamp.Month() == 4 && timestamp.Year() == 2020 {
			jeffUploadOnApril15th2020++
		}
	}

	logger.Printf("Num of Users Accessed The Server: %d", len(users))
	logger.Printf("Num of Upload > 50KB: %d", uploadsGT50KB)
	logger.Printf("Num of Upload to server by `jeff22` on April 15th, 2020: %d", jeffUploadOnApril15th2020)
}

func main() {
	runMain(os.Args)
}
