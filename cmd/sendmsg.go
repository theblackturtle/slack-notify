package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
)

var (
	msgContent  string
	msgFilePath string
	sendAsCode  bool
)

var msgCmd = &cobra.Command{
	Use:   "msg",
	Short: "Send messages",
	Long:  "Send messages to channel",
	Run:   sendMsg,
}

func init() {
	rootCmd.AddCommand(msgCmd)

	msgCmd.PersistentFlags().StringVarP(&msgContent, "text", "t", "", "Content of messages")
	msgCmd.PersistentFlags().StringVarP(&msgFilePath, "file", "f", "", "Send content of file as message.")
	msgCmd.PersistentFlags().BoolVarP(&sendAsCode, "code-block", "b", false, "Send as code block")
}

func sendMsg(cmd *cobra.Command, args []string) {
	if msgContent == "" && msgFilePath == "" {
		fmt.Println("Please check your input again")
		os.Exit(1)
	}

	api := slack.New(SLACK_TOKEN)
	if msgFilePath != "" {
		b, err := ioutil.ReadFile(msgFilePath)
		if err != nil {
			fmt.Print(err)
			return
		}
		msgContent = string(b)
	}

	if sendAsCode {
		msgContent = fmt.Sprintf("```\n%s\n```", msgContent)
	}

	channelID, timestamp, err := api.PostMessage(SLACK_CHANNEL, slack.MsgOptionText(msgContent, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
