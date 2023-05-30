package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/aallali/deepeye/src/config"
	"github.com/dlclark/regexp2"
	"github.com/fatih/color"
)

// declare the color func that will colorize my matched strings in yello background a black font
var info = color.New(color.FgBlack, color.BgYellow).SprintFunc()

// @params re *regexp2.Regexp : the regex compiled to search with
// @params s string : target string to match the regex expression from
// @return []string : list of strings matching the regex expression from the 's' string
// the standar library of regex2 doesnt have a built in function to return all results of a specefic regex match
// so we have to create one
func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

// @params cond boolean : the condition that i wanna check in order to return X or Y
// @params vTrue any : the return value in case cond == true
// @params vFalse any : the return value in case cond == false
// @return any : return either vTrue or vFalse
// @context:
//
//	I came from a JS background so i got used to inline functions and conditional chaining... easy code
//	GO says NO ! :(
//	but since i have spent some pretty nice years with C language, i enjoyed creating things from scratch,
//	so i wrote this dumb func to make the chaining condition for me in a readable way
//	dont steal my IDEAS Bleaaase
func ifElse(cond bool, vTrue interface{}, vFalse interface{}) interface{} {
	if cond {
		return vTrue
	} else {
		return vFalse
	}
}

// @params s []string: array of strings
// @return []string : remove duplicated values from arg 's'and retun new array
func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if str != "" {
			if _, ok := bucket[str]; !ok {
				bucket[str] = true
				result = append(result, str)
			}
		}
	}
	return result
}

// @params txt string : original text to search and extract from
// @params subs []string : list of strings, each represent a keyword to search with
// @params margin int : a number >= -1
// @return []string : list of strings containing the substring match in 'txt' highlighted in yellow
func SpotAndMargin(txt string, subs []string, margin int) []string {

	var results []string       // init empty str array variable to store results
	var txtSize int = len(txt) // get the length of txt, we will use it many times, so calculate it once to optimize CPU ops.
	var i int
	if margin == -1 || margin >= txtSize {
		for _, ky := range subs {
			txt = strings.ReplaceAll(txt, ky, info(ky)) // info(string) : return string highlighted in color defined above.
		}
		return []string{txt}
	}

	if len(subs) == 0 {
		return results
	}

	for _, k := range subs {
		kSize := len(k)
		prevStartI := 0
		prevEndI := 0
		for true {
			// tmporarry buffer to contain some calculation values
			// that will be used multiple times
			tmp := 0
			// get the index of keyword in txt if exists
			keyWordIndex := strings.Index(txt[i:], k)

			// break if not found
			if keyWordIndex == -1 {
				break
			}

			// increase the global index to where our desired match keyword start
			i = i + keyWordIndex

			// :-------:| calculate the margins index only if margin > 0 |:--------:
			if margin > 0 {
				// check if left margin is valid (not out of slice bound)
				tmp = i - margin
				startI := ifElse(tmp <= 0, 0, tmp).(int)

				// check if right margin is valid (not out of slice bound)

				tmp = i + kSize + margin

				endI := ifElse(tmp >= txtSize, txtSize, tmp).(int)

				// this logic below, is to merge multiple matches in same line and sharing same margin into one line output
				// e.g:

				// 		txt = "Lorem Lepsum 12, Lor 42 Lpesum, and BOOM 1337"
				// 		subs = [12, 42, 1337]
				// 		range = 5

				// normally the output was :
				// 		{L:0}: [...psum 12, Lor...]
				// 		{L:0}: [... Lor 42 Lpes...]
				// 		{L:0}: [...BOOM 1337]

				// notice that '12' and '42' margin and overflowing, to avoid this duplication,
				// i wrote this algorithm to check and compare current margin indexs with previous one
				// if there is a possibility of merge, i remove the previous one from the results array (which contains
				// the previous match with its margin alone) then create new one containg the large margin covering the other matches
				// like that :
				// 		{L:0}: [...psum 12, Lor 42 Lpes...]
				// 		{L:0}: [...BOOM 1337]
				if startI < prevEndI { // a possible merge found
					results = results[:len(results)-1] // remove the prev phrase
					startI = prevStartI                // update the current startIndex with previous one, which starts from the first match
				}
				// update the values of 'prev...I' ofc
				prevEndI = endI
				prevStartI = startI

			}

			// the variable that will contain the value you see on screen:\
			// eg: "...psum 12, Lor 42 Lpes...""
			phrase := ""

			// verify our margins
			switch margin {
			case 0:
				// if 0 means print only the matched string
				phrase = txt[i : i+kSize]
			case -1:
				// print full line
				phrase = txt
			default:
				// format the final string that contains the matched keyword highlighted in yellow, inside the margin of neighbor characters choosen
				// if the a side of margin is not reaching the limits of strings (end or start), we prefix/suffix with 3 dots "..."
				phrase = fmt.Sprintf(
					"%s%s%s",
					ifElse(prevStartI > 0, "...", ""),
					txt[prevStartI:prevEndI],
					ifElse(prevEndI-txtSize < 0, "...", ""))
			}
			// add current formed string to results
			results = append(results, phrase)

			// increase the global index to next character after previous match, so we start searching only in left part of 'txt'
			i = i + kSize

		}
	}
	// if there a margin, its better to highlight the match for better view :)
	if margin != 0 {
		for _, ky := range subs {
			for i := 0; i < len(results); i++ {
				results[i] = strings.ReplaceAll(results[i], ky, info(ky))
			}
		}
	}

	return results
}

// print the current installed version of DeepEye defined in
// /src/config/constants.go
func PrintVersion() {
	fmt.Printf("DeepEye v%s\n", Infos.Version)
	os.Exit(0)
}

// request the version.txt file in root level as raw text, and compare it to installed version
// TODO: write auto updater.
func CheckUpdate() {
	fmt.Println("Checking ...")
	res, err := http.Get(Infos.VCheckUrl)
	if err != nil {
		log.Fatal(err)

	} else {
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		resStrList := strings.Split(string(responseData), "\n")
		v := resStrList[0]
		if v != Infos.Version {
			msg := fmt.Sprintf(`
DeepEye:
- The version installed  : %s
- The Latest version     : %s
Please visit the official repo to download last version : [%s]`, Infos.Version, v,
				Infos.GitRep)
			fmt.Println(msg)
		} else {
			fmt.Println("You are running the latest version of DeepEye:", v)
		}
		os.Exit(0)
	}

}
