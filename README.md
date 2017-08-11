# google-iot-mqtt-publisher

google-iot-mqtt-publisher is a MQTT client used to publish messages to a specific topic on Google Cloud IoT.

### Build Instructions

Fetch the third party libraries:

`$ go get`

Build for x86:

`$ go build -o google-iot-mqtt-publisher`

Build for arm64:

`env GOOS=linux GOARCH=arm64 go build -o google-iot-mqtt-publisher`

This will generate an executable with the filename of `google-iot-mqtt-publisher` in the same directory. Note that the `certs` directory must be distributed alongside with the executable and both must be located in the same directory. Otherwise, the application will not work correctly.

### Usage
```
$ ./google-iot-mqtt-publisher --help
google-iot-mqtt-publisher is a MQTT client used to publish messages to a specific topic on Google Cloud IoT.

Usage:
  google-iot-mqtt-publisher [flags]

Flags:
  -d, --device-id string     Device ID
  -m, --message string       A string which will be sent to a topic.
  -p, --project-id string    Project ID
  -r, --registry-id string   Registry ID
  -t, --topic string         A topic to which the message will be sent.
```

Example:

To publish a message, run the `google-iot-mqtt-publisher` tool with the correct values for flags.

`$ ./google-iot-mqtt-publisher --project-id "eco-world-166017" --registry-id "example-registry" --device-id "example-device" --topic "events" --message "some really cool message containing telemetry data"`

To fetch the published message, you can use the `gcloud` available from Google Cloud SDK.

`$ gcloud beta pubsub subscriptions pull --auto-ack projects/eco-world-166017/subscriptions/my-sub`

### Tutorial

How to use google-iot-mqtt-publisher: https://asciinema.org/a/4qe9gmh5rwa971zflze67b12b

### Library Dependencies

The google-iot-mqtt-publisher depends on following libraries:

| Library Name                | Link                                        | License    |
| --------------------------- | ------------------------------------------- | ---------- |
| Cobra                       | https://github.com/spf13/cobra              | Apache-2.0 |
| Eclipse Paho MQTT Go client | https://github.com/eclipse/paho.mqtt.golang | EPL-1.0    |
| jwt-go                      | https://github.com/dgrijalva/jwt-go         | MIT        |
