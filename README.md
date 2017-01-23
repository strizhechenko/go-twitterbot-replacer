# Setup and run

``` shell
mv config.yaml.tmplt config.yaml
vim config.yaml # put here consumer key/secret + access token/secret + phrases for replacements
go build
./go-twitterbot-replacer
```

# state

Just util generating replaced tweets, no automatical tweeting included.
