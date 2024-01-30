package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile  string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a Todo list CLI app",
	Long:  `Tri is designed to make organizing your tasks fast and easy straight through your terminal`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
    cobra.OnInitialize(initConfig)

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
        "config file (default is $HOME/.tri/tricfg.yaml)")
}

func initConfig() {
    viper.SetConfigName("tricfg")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/etc/tri/")
    viper.AddConfigPath("$HOME/.tri")
    viper.AddConfigPath(".tri")
    viper.SetEnvPrefix("tri")

    // Overrides config file if set
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err == nil {
        log.Println("Using config file:", viper.ConfigFileUsed())
    }
}
