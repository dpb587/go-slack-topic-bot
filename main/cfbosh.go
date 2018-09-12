package main

import (
	"log"

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
				pairist.PeopleInRole{
					Team:          "sf-bosh",
					Role: "interrupt",
					Interruptible: pairist.InterruptibleHours("12:00", "18:00", "America/Los_Angeles"),
					People: map[string]string{
						"Luan":    "U02R9SUJX",
						"Josh R":  "U8DLATS12",
						"Josh":    "U0YGVGTM1",
						"Danny":   "U0FUK0EBH",
						"Mike":    "U21JVA9F0",
						"Jim":     "U02QZ1E3G",
						"Morgan":  "U04V9L81Y",
						"Belinda": "U5EJ8MQUW",
						"Max":     "U4FFS1UAK",
					},
				},
				pairist.PeopleInRole{
					Team:          "boshto",
					Role:          "Interrupt",
					Interruptible: pairist.InterruptibleHours("06:30", "12:00", "America/Los_Angeles"),
					People: map[string]string{
						"Gaurab":  "U0A0ZUT43",
						"Dale":    "U32RHRLE9",
						"Rebecca": "U8YCN97Q9",
						"Andrew":  "U17K4GAKW",
						"Fred":    "UA3MK3AE7",
						"Jamil":   "U0717EQ04",
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

	err = slack.UpdateChannelTopic("CCR5PN34Z", msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}
}
