package commands

import (
	"fmt"
	"os"
	//"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//pass any package function by name
var CfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	//defines a flag on our command
	//&CfgFile passes address to our package var
	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/dagobah/")
	viper.AddConfigPath("$HOME/.dagobah/")
	viper.ReadInConfig()
}

func addCommands() {
	RootCmd.AddCommand(fetchCmd)
}

var RootCmd = &cobra.Command{
	Use:   `dagobah`,
	Short: `Dagobah is an awesome`,
	Long:  `dagobah long..`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dagobah runs")
	},
}

func Execute() {
	addCommands()
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
