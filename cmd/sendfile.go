package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
)

var (
	filePath string
	title    string
)

var sendfileCmd = &cobra.Command{
	Use:   "file",
	Short: "Send file",
	Long:  "Send file to channel",
	Run:   sendFile,
}

func init() {
	rootCmd.AddCommand(sendfileCmd)

	sendfileCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "File path")
	sendfileCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Title of message (Default is filepath)")

	sendfileCmd.MarkPersistentFlagRequired("file")
}

func sendFile(cmd *cobra.Command, args []string) {
	fullpath, _ := filepath.Abs(filePath)
	api := slack.New(SLACK_TOKEN)
	if title == "" {
		title = fullpath
	}

	params := slack.FileUploadParameters{
		Title:    title,
		Filename: "upload.txt",
		File:     filePath,
		Channels: []string{SLACK_CHANNEL},
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)

}
