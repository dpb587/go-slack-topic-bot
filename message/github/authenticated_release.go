package github

import (
	"context"
	"fmt"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/pkg/errors"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type AuthenticatedRelease struct {
	Token string
	Alias string
	Owner string
	Name  string
}

var _ message.Messager = &AuthenticatedRelease{}

func (m AuthenticatedRelease) Message() (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: m.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user

	rels, _, err := client.Repositories.ListReleases(ctx, m.Owner, m.Name, nil)
	if err != nil {
		return "", errors.Wrap(err, "listing releases")
	}

  for _, rel := range rels {
	  return fmt.Sprintf("%s/%s", m.Alias, *rel.Name), nil
  }

  return "", nil
}
