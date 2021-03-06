package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "slack-notify",
	Short:   "Slack notify",
	Example: "SLACK_TOKEN=xoxb-... SLACK_CHANNEL=[Channel-name] ./slack-notify",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	SLACK_TOKEN   string
	SLACK_CHANNEL string
)

func init() {
	SLACK_TOKEN = os.Getenv("SLACK_TOKEN")
	SLACK_CHANNEL = os.Getenv("SLACK_CHANNEL")
	if SLACK_CHANNEL == "" || SLACK_TOKEN == "" {
		fmt.Println("Please set your environments\n\nExamples:\nSLACK_TOKEN=xoxb-... SLACK_CHANNEL=[Channel-name] ./slack-notify")
		os.Exit(0)
	}
}
