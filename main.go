package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const VERSION = "v0.1.1"


var rootCmd = &cobra.Command{
	Use: "sn",
	Short: "Simple notifier",
	Run: func(cmd *cobra.Command, args []string) {
		locationStr, err := cmd.Flags().GetString("location")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		messageStr, err := cmd.Flags().GetString("message")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfg, err := LoadConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for k, v := range cfg.Locations {
			if strings.ToLower(k) == strings.ToLower(locationStr) {
				if strings.ToLower(v.Type) == strings.ToLower("discord") {
					if err := SendDiscordMessage(v.Webhook, v.BotName, messageStr); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				return
			}
		}

		fmt.Printf("Location %v not found !\n", locationStr)
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Display the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Simple-notifier by Oxodao")
		fmt.Println("https://github.com/oxodao/simple-notifier")
		fmt.Println()
		fmt.Println("Version " + VERSION)
	},
}

func main() {
	rootCmd.Flags().StringP("location", "l", "", "The location to send the message to")
	rootCmd.Flags().StringP("message", "m", "", "The message to send")

	rootCmd.MarkFlagRequired("location")
	rootCmd.MarkFlagRequired("message")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
