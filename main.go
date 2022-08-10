package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const URL = "https://www.rajan.com/calendar/nepal_time.asp"
const SELECTOR = "html>body>table>tbody>tr>td.normal>center>table>tbody>tr>td.bold"

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Error accessing user's home directory")
	}
	return homeDir
}

func updateCache(lastCallPath, dateCachePath string) {
	// fetch date and return the date
	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Could not connect to the network")
		os.Exit(1)
	})
	c.OnHTML(SELECTOR, func(h *colly.HTMLElement) {
		fmt.Println("Requesting...")
		allText := h.Text
		// The other bitch is some comment
		if len(allText) < 100 {
			timeAndDate := strings.Split(allText, "\n")[2]
			output := strings.Split(timeAndDate, "<")[0]

			// last called record
			lastCallFile, lcFileError := os.Create(lastCallPath)
			if lcFileError != nil {
				panic("Operation aborted. Please check file permissions.")
			}
			fmt.Fprint(lastCallFile, time.Now().Format("2006.01.02"))

			// write today's nepali date in ~/.nepdate/date
			dateCache, dateCacheError := os.Create(dateCachePath)
			if dateCacheError != nil {
				panic("Operation aborted. Please check file permissions.")
			}
			fmt.Fprint(dateCache, output)
		}
	})
	c.Visit(URL)
}

// Function prints date
func printDate(lastCallPath, dateCachePath string) {
	// foo bar
	cachedDate, err := os.ReadFile(dateCachePath)
	// if can't open cache
	if err != nil {
		updateCache(lastCallPath, dateCachePath)
		printDate(lastCallPath, dateCachePath)
	}
	// print cache
	fmt.Println(string(cachedDate))
}

func main() {
	lastCallPath := filepath.Join(getHomeDir(), ".nepdate", "lastcall")
	dateCachePath := filepath.Join(getHomeDir(), ".nepdate", "date")

	lastCalled, err := os.ReadFile(lastCallPath)

	// can't find when last date it was called; might be first run or deleted for some reason
	if err != nil {
		updateCache(lastCallPath, dateCachePath)
	}

	// if today's date and last accessed is same
	lc := string(lastCalled)
	tn := time.Now().Format("2006.01.02")
	if lc != tn {
		// open the file
		updateCache(lastCallPath, dateCachePath)
	}
	printDate(lastCallPath, dateCachePath)
}
