package config

import "fmt"

var Infos = struct {
	Version     string
	GitRep      string
	Description string
	Short       string
	Usage       string
	VCheckUrl   string
}{
	Version:   "0.0.1",
	GitRep:    "https://github.com/aallali/deepeye/",
	Short:     "CLI for fast/efficient searching queries through files.",
	Usage:     "deepeye <filename> flags",
	VCheckUrl: "https://github.com/aallali/deepeye/raw/main/version.txt",
}

func init() {

	Infos.Description = fmt.Sprintf(`
DeepEye: a CLI that will allow you to run advanced search queries
	through multipe text files, while having customized + detailed output.
	The "deepeye" program was mainly focused on helping terminal users, 
	to quickly search in files in both plain text and regex queries.
		
Willing to contribute? : "%s"

Author: Abdellah Allali <hi@allali.me>
Birth: 24/05/2023
First release: 30/05/2023`, Infos.GitRep)
}
