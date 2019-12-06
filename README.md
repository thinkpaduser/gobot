## Learning Go with Telegram API

Simple and stupid bot really

## Requirements:

-  [Telegram Bot Api by Syfaro](https://github.com/Syfaro/telegram-bot-api).
-  [Yaml v2](https://gopkg.in/yaml.v2).
-  You can `go get github.com/Syfaro/telegram-bot-api` and `go get gopkg.in/yaml.v2` to your `GOPATH/src/`

## Features:

- go func() to send a message every `n` minutes;
  
- Message text uses a random chosen value
    
-  Message database is just a simple YAML file and a `map[string][]string`
 
-  Message to be sent is based by one of the `keys` if available;
 
-  If `key` found, message text will be the random `value` from "messages DB"
  
-  Also you can set a chance that message will be sent.
