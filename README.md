# Phoebe Event Manager

Phoebe is a minuscule event manager for the command line, with very basic features. This allows for your great ideas to be added quickly to the terminal, where we are
likely spending all of our time.

## Features

In Phoebe, you can do 4 different things:
- add an event to a weekday
- delete an event from a weekday
- view the week or specific days events
- get documentation on all the above commands

## Usage

| Command Name | Usage | Example |
| ----- | ----- | -----|
| Add | pb add [weekday] "event" | pb add Monday "Coffee and code" |
| Week | pb week [weekday] (optional) | pb week |
| Delete | pb del [weekday] "event" | pb del Monday "Coffee and code"

## Installation

Please be on the lookout for your repository to be added.

### MacOS
You can use hombrew for a macOS install.
`brew tap ap0ll02/phoebe` and `brew install phoebe`

### Arch 
Coming soon: `paru phoebe`

### Debian/Ubuntu
Coming soon.

### From Source
Clone the repo, and enter it. Then `go build -o pb`, where you can move the resultant pb somewhere to your path, example: `mv pb /usr/bin/`


Thanks for reading, this is my first CLI, I am mostly making this for myself so if you stumble across this, thanks for reading!
