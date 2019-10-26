Learning Go with Telegram API

Simple and stupid bot really
##Features:
  go func() to send a message every `n` minutes;
    Message text uses a random chosen value
  Message database is just a simple YAML file and a `map[string][]string`
  Message to be sent is based by one of the `keys` if available;
  If `key` found, message text will be the random `value` from "messages DB"
  Also you can set a chance that message will be sent.
