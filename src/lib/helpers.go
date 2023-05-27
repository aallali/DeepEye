package lib

import (
	"fmt"
	"os"
	"strings"

	. "github.com/aallali/deepeye/src/config"
)

func ilc(cond bool, vt interface{}, vf interface{}) interface{} {
	if cond {
		return vt
	} else {
		return vf
	}
}

func SubSWithRange(txt string, subs string, rng int) []string {
	k := subs

	i := 0
	results := []string{}
	if k == "" {
		return results
	}
	for true {
		tmpI := strings.Index(txt[i:], k)

		if tmpI == -1 {
			break
		}
		i = i + tmpI
		startI := ilc(i-rng <= 0, 0, i-rng).(int)
		endI := ilc(i+len(k)+rng >= len(txt), len(txt), i+len(k)+rng).(int)
		phrase := fmt.Sprintf(
			"%s%s%s",
			ilc(startI > 0, "...", "").(string),
			txt[startI:endI], ilc(endI-len(txt) < 0, "...", ""))
		results = append(results, phrase)
		i = i + len(k)

		if i >= len(txt) {
			break
		}
	}
	return results
}
func PrintVersion() {
	fmt.Printf("DeepEye version : %s\n", Infos.Version)
	os.Exit(0)
}

func CheckUpdate() {
	fmt.Printf("%s is not the latest version, updating now...\n", Infos.Version)
	os.Exit(0)
}
