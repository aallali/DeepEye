package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/aallali/deepeye/src/config"
	"github.com/dlclark/regexp2"
)

// @params query Query: a type Query variable containing the options of our search query
// @return void
func DeepEye(query Query) {

	var line int = 0           // to keep track of lines scanned
	var totalMatchs int = 0    // to count number of matches found
	var scanner *bufio.Scanner // where the scanner gonna be initiated
	var startTime time.Time    // to hold the start time value
	var elapsed time.Duration  // to hold the end time value after finishing the scan
	var rgx *regexp2.Regexp    // the compiled regex expression from query.Regex

	f, err := os.Open(query.FilePath) // get the file descriptor of path given if found

	if err != nil { // if there is an error while opening the file, print it and OUT!
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close() // close the file descriptor when finish the DeepEye() function
	// https://golang.org/pkg/bufio/#Scanner.Scan
	scanner = bufio.NewScanner(f)

	rgx, rgxErr := regexp2.Compile(query.Regex, 0)

	if rgxErr != nil {
		fmt.Println(rgxErr)
		os.Exit(1)
	}
	// snapshot the current time
	startTime = time.Now()

	for scanner.Scan() { // Scan advances the Scanner to the next toke

		var lineStr string = scanner.Text() // ge the String value of current line

		// if query.Regex is present, then run search through regex search
		if query.Regex != "" {
			// run the match with regex expression then remove duplication
			rgxMatchResult := removeDuplicates(regexp2FindAllString(rgx, lineStr))

			// if Silent is false then run the margin generator
			if !query.Silent {
				results := SpotAndMargin(lineStr, rgxMatchResult, query.Range)

				// loop through all formed strings from SpotAndMargin and print them with followng template
				for _, el := range results {
					fmt.Printf("{l:%d}: [%s]\n", line, el)
					totalMatchs++
				}
			} else {
				// if Silent i true, then increment the number of matched keywords
				totalMatchs += len(rgxMatchResult)
			}

			// otherwise get the query.Keyword value
		} else {

			// if Silent false, search for the query.Keyword with SpotAndMargin
			if !query.Silent {
				results := SpotAndMargin(lineStr, []string{query.Keyword}, query.Range)
				// loop through the results, format them, and increat match counter.
				for _, el := range results {
					fmt.Printf("{l:%d}: [%s]\n", line, el)
					totalMatchs++
				}
			} else {
				totalMatchs += strings.Count(lineStr, query.Keyword)
			}
		}
		line++

	}

	// calculate the time duration from line 42
	elapsed = time.Since(startTime)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("-------------------------------")
	fmt.Printf("Target file path    : %s\n", query.FilePath)
	fmt.Printf("Search query        : `%s`\n", ifElse(query.Keyword == "", query.Regex, query.Keyword))
	fmt.Printf("Total matches found : %d\n", totalMatchs)
	fmt.Printf("Total lines scanned : %d\n", line)
	fmt.Printf("File Scan took 	    : %s\n", elapsed.String())

}
