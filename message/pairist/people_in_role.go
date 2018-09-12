package pairist

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dpb587/go-pairist/api"
	"github.com/dpb587/go-pairist/denormalized"
	"github.com/dpb587/go-slack-topic-bot/message"
)

type PeopleInRole struct {
	Team          string
	Role          string
	People        map[string]string
	Interruptible func() bool
}

var _ message.Messager = &PeopleInRole{}

func (m PeopleInRole) Message() (string, error) {
	if !m.Interruptible() {
		return "", nil
	}

	curr, err := api.DefaultClient.GetTeamCurrent(m.Team)
	if err != nil {
		return "", err
	}

	var handles []string

	for _, lane := range denormalized.BuildLanes(curr).ByRole(m.Role) {
		for _, person := range lane.People {
			if handle, ok := m.People[person.Name]; ok {
				handles = append(handles, fmt.Sprintf("<@%s>", handle))
			} else {
				handles = append(handles, person.Name)
			}
		}
	}

	if len(handles) == 0 {
		return "", nil
	}

	sort.Strings(handles)

	return strings.Join(handles, " "), nil
}

func InterruptibleHours(start, stop, tzName string) func() bool {
	tz, err := time.LoadLocation(tzName)
	if err != nil {
		panic(err)
	}

	return func() bool {
		now := time.Now().In(tz)

		dt := now.Format("Mon")

		if dt == "Sat" || dt == "Sun" {
			return false
		}

		ts := now.Format("15:04")

		return ts >= start && ts < stop
	}
}
