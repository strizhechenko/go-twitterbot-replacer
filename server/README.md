# Config

You should register your app and authorize your account in twitter with it in order to fill consumer/access fields.

Replacements is just a list of key (origin phrase) - value (its replacement) pairs.

# main.go

Config parsing and running server.

# text_processing.go

All the magic of tweet-text-processing such a cleanups and replacements.

# twitter.go

Everything about connect to twitter API and grabbing tweets from timeline.

# web.go

Provides REST-API to fetch tweets without authentification and post selected tweets from client-UI later.
