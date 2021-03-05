package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CliName holds name of the CLI executable
var (
	CliName   = os.Args[0]
	Language  string
	Overwrite bool
)

// init - Initialize the root command
func init() {
	// Add a global language flag
	rootCmd.PersistentFlags().StringVarP(
		&Language,
		"language",
		"l",
		"julia",
		"programming language used to start and solve problems (default: julia)",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&Overwrite,
		"overwrite",
		"o",
		false,
		"overwrite existing template or target files (default: false)",
	)
	rootCmd.PersistentFlags().Lookup("overwrite").NoOptDefVal = "true"

	// Bind with viper flag to enable reading (and writing) config
	viper.BindPFlag("language", rootCmd.PersistentFlags().Lookup("language"))
	viper.SetDefault("language", "julia")

	// Read in config file if present
	workingDir, err := os.Getwd()
	cobra.CheckErr(err)
	viper.AddConfigPath(workingDir)
	viper.SetConfigName(".euler")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// RootCmd is cobra command executed when the CLI is called without any arguments
var rootCmd = &cobra.Command{
	Use:     "euler",
	Version: "0.1",
	Short:   fmt.Sprintf("%s is a CLI for working on Project Euler problems", CliName),
	Long: fmt.Sprintf(`%s is a CLI for working on Project Euler problems

Use it to create templated solutions to new problems, execute solutions, and check answers.`, CliName),
}

// Execute calls the eponymous function on the root command
func Execute() {
	rootCmd.Execute()
}
