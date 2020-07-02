package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var (
	slackClient *slack.Client
)

func main() {
	os.Setenv("SLACK_API_TOKEN", "xoxb-1232205680049-1243574468144-nmvFEPCGRzvj6l3RTVakVspc")
	slackClient = slack.New(os.Getenv("SLACK_API_TOKEN"))
	fmt.Printf(os.Getenv("SLACK_API_TOKEN"))
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(ev)
		}
	}
}

func handleMessage(ev *slack.MessageEvent) {
	fmt.Printf("%v\n", ev)

	re
}
