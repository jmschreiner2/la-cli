/*
Copyright © 2025 Jacob Schreiner jmschreiner2@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jmschreiner2/la-cli/logger"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	subscriptionID string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "la-cli",
	Short: "A CLI Application that will query logic app information.",
	Long: `Query Logic App information to be able to find runs and other things.
This CLI will access your credentials using azure cli config file.

If you have not setup your az, run az login.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug)
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.la-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&logger.Verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().StringVarP(&subscriptionID, "subscirption-id", "s", "", "Azure Subscription ID")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".la-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".la-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
