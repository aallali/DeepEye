package config

type Query struct {
	FilePath string // the target file path

	Keyword string // plain text keyword to search with in file
	Regex   string // regex expression to match from the file

	Silent bool // to silent the output of the search query , so only the stats will be printed
	Range  int  // specify the range of margin to print around the match
	Lines  int  // set limit to number of lines to output
	Clean  bool // if you want clean output (for file use case)

	Update  bool // check for any possible updates
	Version bool // print the installed version
}
