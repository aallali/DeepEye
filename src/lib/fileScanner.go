package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/aallali/deepeye/src/config"
	"github.com/dlclark/regexp2"
	"github.com/fatih/color"
)

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func DeepEye(query Query) {
	starTime := time.Now()

	f, err := os.Open(query.FilePath)
	// time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 0
	totalMatchs := 0
	info := color.New(color.FgBlack, color.BgYellow).SprintFunc()

	// https://golang.org/pkg/bufio/#Scanner.Scan
	r, rerr := regexp2.Compile(query.Regex, 0)

	if rerr != nil {
		println(rerr.Error())
		os.Exit(1)
	}

	for scanner.Scan() {

		var lineStr string = scanner.Text()

		if query.Regex != "" {
			rgxMatchResult := regexp2FindAllString(r, lineStr)
			if !query.Silent {
				for _, singleMatch := range rgxMatchResult {
					if singleMatch != "" {
						totalMatchs++
						results := SubSWithRange(lineStr, singleMatch, 10)
						for _, el := range results {
							if isMatch, _ := r.MatchString(el); isMatch {
								rgxRepl, _ := r.Replace(el, info(singleMatch), 10, -1)
								fmt.Printf("[L:%d]: [%s]\n", line, rgxRepl)
							}
						}
					}
				}
			} else {
				totalMatchs += len(rgxMatchResult)
			}
		} else {
			if !query.Silent {
				results := SubSWithRange(lineStr, query.Keyword, 30)
				for _, r := range results {
					fmt.Printf("[L:%d]: [%s]\n", line, strings.Replace(r, query.Keyword, info(query.Keyword), -1))
					totalMatchs++
				}
			} else {
				totalMatchs += strings.Count(lineStr, query.Keyword)
			}
		}
		line++
	}
	elapsed := time.Since(starTime)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("-------------------------------")
	fmt.Printf("Target file path    : %s\n", query.FilePath)
	fmt.Printf("Search query        : `%s`\n", ilc(query.Keyword == "", query.Regex, query.Keyword))
	fmt.Printf("Total matches found : %d\n", totalMatchs)
	fmt.Printf("Total lines scanned : %d\n", line)
	if elapsed.Milliseconds() >= 1000 {
		fmt.Printf("%s took 	    : %.3f s\n", "File Scan", elapsed.Seconds())
	} else {
		fmt.Printf("%s took 	    : %d ms\n", "File Scan", elapsed.Milliseconds())
	}

}
