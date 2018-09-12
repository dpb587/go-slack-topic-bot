package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/pkg/errors"
)

type Release struct {
	Alias string
	Owner string
	Name  string
}

var _ message.Messager = &Release{}

type gitHubReleaseApiV2 []struct {
	Name string `json:"name"`
}

func (m Release) Message() (string, error) {
	res, err := http.DefaultClient.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", m.Owner, m.Name))
	if err != nil {
		return "", err
	}

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "reading response")
	}

	var data gitHubReleaseApiV2

	err = json.Unmarshal(resBodyBytes, &data)
	if err != nil {
		return "", errors.Wrap(err, "unmarshalling")
	}

	if len(data) == 0 {
		return "", nil
	}

	return fmt.Sprintf("%s/%s", m.Alias, data[0].Name), nil
}
