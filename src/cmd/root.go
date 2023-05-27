/*
Copyright © 2023 Abdellah Allali <hi@allali.me>
*/
package cmd

import (
	"os"

	. "github.com/aallali/deepeye/src/config"
	"github.com/aallali/deepeye/src/lib"

	"github.com/spf13/cobra"
)

var query Query

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   Infos.Usage,
	Short: Infos.Short,
	Long:  Infos.Description,
	// Uncomment the following line if your bare application
	// has an action associated with it:

	Args: func(cmd *cobra.Command, args []string) error {
		if (query.Version || query.Update) && len(args) == 0 {
			return nil
		}
		if len(args) == 1 {
			if query.Regex == "" && query.Keyword == "" {
				cmd.Usage()
				os.Exit(1)
			}
			return nil
		}
		return cmd.Usage()

	},
	Run: func(cmd *cobra.Command, args []string) {
		if query.Update || query.Version {
			if query.Version {
				lib.PrintVersion()
			}
			if query.Update {
				lib.CheckUpdate()
			}
			return
		}

		query.FilePath = args[0]

		lib.DeepEye(query)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var author string

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.Flags().StringVarP(&query.Regex, "regex", "r", "", "regex expression to match in file.")
	rootCmd.Flags().StringVarP(&query.Keyword, "keyword", "k", "", "Keyword to match in file.")

	rootCmd.Flags().BoolVarP(&query.Silent, "silent", "s", false, "if you want to silent the comand, only resume will be printed.")
	rootCmd.Flags().BoolVarP(&query.Update, "update", "u", false, "check for updates.")
	rootCmd.Flags().BoolVarP(&query.Version, "version", "v", false, "output the current installed version of DeepEye CLI.")

}
