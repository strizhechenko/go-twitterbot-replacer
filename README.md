# Setup and run

``` shell
cd ./server
cp -a config.yaml.tmplt config.yaml
vim config.yaml # put here consumer key/secret + access token/secret + phrases for replacements
go build
./server
```

# state

Just util generating replaced tweets, no automatical tweeting included.
