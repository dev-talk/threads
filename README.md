# threads
A small discord bot allowing to pull discussions out into dedicated threads with some management capabilities
ome easy interface including at least some spam protection should be included.

## Why

On our discord we saw the lack of threading features like available in Slack. Therefor an idea arose to at least to some extent cover this by a small and pretty simple bot.

## Features

What we want to achieve is a quite easy solution to pull certain discussions out of the current channel into a fresh one which is dedicated to the topic.
After the discussion is finished the channel should be archived and possibly deleted after a longer time of inactivity.

Some more detailed features:
* Creating thread channels by asking the bot to do so
* Grouping the channels under a certain category and possibly keep them muted or hidden for everyone
* Create a card at channel creation in both the source channel and a discussion list channel to allow people to join
* Joining a channel unhides or subscribes to the channel
* The information cards about discussions get updated with details like the participants
* Users can easily create forum threads from the discord channel 
* After X days without activity the channel gets moved to an archive category
* Participants can (based on activity) vote to archive the channel
* After another X days the channel gets deleted entirely

Further features are planned but not mentioned for the first version/mvp.

## Dependencies
All dependencies inside this project are being managed by [dep](https://github.com/golang/dep) and are checked in.
After pulling the repository, it should not be required to do any further preparations aside from `make deps` to prepare the dev tools (once).

If new dependencies get added while coding, make sure to add them using `dep ensure --add "importpath"` and to check them into git.
We recommend adding your vendor changes in a separate commit to make reviewing your changes easier and faster.

## Coding and Style

Coding is done using pull requests and code reviews.
Our code is always checked by Travis using `make test check` therefor all Golang rules on syntax and formating have to be met for pull requests to be merged.
While this might incur more work for possible contributors, we see the code produced here as production critical once finished and therefor strive for high code quality.

The team is developing this mostly using TDD and BDD. If you don't know what this is, we recommend this [video](https://www.youtube.com/watch?v=uFXfTXSSt4I) for starters.

Please do reasonable commit sizes.

## Testing
To run tests you can use:
```bash
make test
```

## Contributing

Feel free to create Issues and provide Feedback.
To directly contribute to the project, please get in touch and base your work around issues.
Pull requests without a respective issue may be denied until the matter is discussed with the team.
Communication is key for OpenSource projects, so feel free to discuss everything related to this either on [Discord](https://discord.gg/8rsEPNe) or inside GitHub Issues.

## Attributions

* [Kolide for providing `kit`](https://github.com/kolide/kit)

## License

This Project and all it's code is licensed using a GPL V3 License found in [LICENSE](LICENSE).
Code and Ideas contributed to this project are provided for free and without any implied rights and therefor property of the project itself.
Contributors are named in the Git history which acts as a license conform mention.