package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blueprint",
	Short: "Project initialization system",
	Run: func(cmd *cobra.Command, args []string) {
		if quietFlag {
			logrus.SetLevel(logrus.PanicLevel)
		}
		logrus.Info("Hello, world!")
	},
}

var quietFlag bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&quietFlag, "quiet", "q", false, "quiet mode")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
