# go-slack-topic-bot

Some utilities for making a [Slack](https://slack.com/) bot that manages a channel's topic.


## Usage

Use this programmatically from your own programs. See the ([examples](examples/cfbosh.go)) for how you can declare topics you care about. The following environment variables are used in the examples:

 * `SLACK_CHANNEL` - the channel ID where the topic will be updated (not the name; e.g. `CA1B2C3D4`)
 * `SLACK_TOKEN` - an API token for accessing Slack (you may want to create one [here](https://apps.slack.com/apps/A0F7YS25R-bots))


## License

[MIT License](LICENSE)
