package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/oxodao/simple-notifier/notification_service"
	"github.com/spf13/cobra"

	// Loading all the available notification services
	_ "github.com/oxodao/simple-notifier/discord"
	_ "github.com/oxodao/simple-notifier/gotify"
)

const VERSION = "v0.1.1"

var rootCmd = &cobra.Command{
	Use:   "sn",
	Short: "Simple notifier",
	Run: func(cmd *cobra.Command, args []string) {
		locationStr, err := cmd.Flags().GetString("location")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		titleStr, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		messageStr, err := cmd.Flags().GetString("message")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		priority, err := cmd.Flags().GetInt("priority")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfg, err := LoadConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		msg := notification_service.Message{
			Title:    titleStr,
			Content:  messageStr,
			Priority: priority,
		}

		for k, v := range cfg.Locations {
			if strings.EqualFold(k, locationStr) {
				if err := v.Send(msg); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return
			}
		}

		fmt.Printf("Location %v not found !\n", locationStr)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
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
	rootCmd.Flags().StringP("title", "t", "", "The title of the message")
	rootCmd.Flags().StringP("message", "m", "", "The message to send")
	rootCmd.Flags().IntP("priority", "p", 5, "The priority of the message (if supported by the location). Defaults to 5 (standard). 0 is silent, 10 is high priority")

	rootCmd.MarkFlagRequired("location")
	rootCmd.MarkFlagRequired("message")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
