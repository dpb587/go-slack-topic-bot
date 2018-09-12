package main

import (
	"log"
	"os"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/dpb587/go-slack-topic-bot/message/boshio"
	"github.com/dpb587/go-slack-topic-bot/message/github"
	"github.com/dpb587/go-slack-topic-bot/message/pairist"
	"github.com/dpb587/go-slack-topic-bot/slack"
)

func main() {
	msg, err := message.Join(
		" || ",
		message.Prefix(
			"_interrupt_ ",
			message.Join(
				" ",
				pairist.Interrupt{
					Team:          "sf-bosh",
					Interruptible: pairist.InterruptibleHours("12:00", "18:00", "America/Los_Angeles"),
					People: map[string]string{
						"Luan":    "luan",
						"Josh R":  "jrussett",
						"Josh":    "jaresty",
						"Danny":   "dberger",
						"Mike":    "mxu",
						"Jim":     "jfmyers9",
						"Morgan":  "mfine",
						"Belinda": "belinda_liu",
						"Max":     "mpetersen",
					},
				},
			),
		),
		message.Literal("_docs_ <https://bosh.io|bosh.io>"),
		message.Prefix(
			"_latest_ ",
			message.Join(
				" ",
				boshio.Release{Alias: "bosh", Repository: "github.com/cloudfoundry/bosh"},
				github.Release{Alias: "bosh-cli", Owner: "cloudfoundry", Name: "bosh-cli"},
				boshio.Stemcell{Alias: "ubuntu-xenial", Name: "bosh-aws-xen-hvm-ubuntu-xenial-go_agent"},
				boshio.Stemcell{Alias: "ubuntu-trusty", Name: "bosh-aws-xen-hvm-ubuntu-trusty-go_agent"},
			),
		),
	).Message()
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	log.Printf("DEBUG: expected message: %s", msg)

	err = slack.UpdateChannelTopic(os.Getenv("SLACK_CHANNEL"), msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}
}
