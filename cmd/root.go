// Package cmd provides the command line interface (CLI) for the application.
package cmd

import (
	"os"

	"github.com/MarvellEmbeddedProcessors/google-iot-mqtt-publisher-marvell/mqtt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "google-iot-mqtt-publisher",
	Short: "Publish a message to a topic using MQTT protocol.",
	Long:  "google-iot-mqtt-publisher is a MQTT client used to publish messages to a specific topic on Google Cloud IoT.",
	Run:   rootRun,
}

var projectId string
var registryId string
var deviceId string
var topic string
var message string

// init initializes the root command and all its flags.
func init() {
	rootCmd.Flags().StringVarP(&projectId, "project-id", "p", "", "Project ID")
	rootCmd.Flags().StringVarP(&registryId, "registry-id", "r", "", "Registry ID")
	rootCmd.Flags().StringVarP(&deviceId, "device-id", "d", "", "Device ID")
	rootCmd.Flags().StringVarP(&topic, "topic", "t", "", "A topic to which the message will be sent.")
	rootCmd.Flags().StringVarP(&message, "message", "m", "", "A string which will be sent to a topic.")
}

func rootRun(cmd *cobra.Command, args []string) {
	// Check for required flags.
	checkFlags(cmd)
	// Publish a message to MQTT broker.
	mqtt.PublishMessage(projectId, registryId, deviceId, topic, message)
}

// Execute function executes the root command which is the application binary itself.
// This function is called from main.main().
func Execute() {
	// Just execute the root command and ignore the returned error since the Cobra
	// will emit the help output for invalid flags and commands.
	_ = rootCmd.Execute()
}

// checkFlag checks the required flags.
func checkFlags(cmd *cobra.Command) {
	errorNumber := 0
	if projectId == "" {
		errorNumber++
		cmd.Println("Error: --project-id or -p flag is required.")
	}
	if registryId == "" {
		errorNumber++
		cmd.Println("Error: --registry-id or -r flag is required.")
	}
	if deviceId == "" {
		errorNumber++
		cmd.Println("Error: --device-id or -d flag is required.")
	}
	if topic == "" {
		errorNumber++
		cmd.Println("Error: --topic or -t flag is required.")
	}
	if message == "" {
		errorNumber++
		cmd.Println("Error: --message or -m flag is required.")
	}

	if errorNumber != 0 {
		cmd.Println()
		cmd.Help()
		os.Exit(1)
	}
}
