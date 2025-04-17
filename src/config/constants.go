package config

import "fmt"

// AppInfo exported struct to hold app infos globally
type AppInfo struct {
	Name        string
	Version     string
	GhUrl       string
	VCheckUrl   string
	Description string
	Short       string
	Usage       string
}

// Infos exposed instance of AppInfo
var Infos = AppInfo{
	Name:      "DeepEye",
	Version:   "0.0.3",
	GhUrl:     "https://github.com/aallali/DeepEye",
	VCheckUrl: "https://raw.githubusercontent.com/aallali/DeepEye/main/version.txt",
	Short:     "CLI for fast/efficient searching queries through files.",
	Usage:     "deepeye <filename>",
}

func init() {

	Infos.Description = fmt.Sprintf(`
DeepEye: a CLI that will allow you to run advanced search queries
	through multipe text files, while having customized + detailed output.
	The "deepeye" program was mainly focused on helping terminal users, 
	to quickly search in files in both plain text and regex queries.
		
Willing to contribute? : "%s"

Author: Abdellah Allali <hi+deepeye@allali.me>
Birth: 24/05/2023
First release: 30/05/2023`, Infos.GhUrl)
}
