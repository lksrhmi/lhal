package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lhal",
	Short: "Overview of what is going on, bassicly",
	Long: `lhal - Let's have a look - is a Windows Taskmanager on crack`,
	

	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("this should run the dashboard like standard wise, yk. When Working")
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

func init() {
	// Here you will define your flags and configuration settings. (--IamAFlaaag, --help (-h))

	// ideas:
	// --version -v
	// yea, no ideas :(
}


