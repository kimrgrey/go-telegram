Go Telegram
============

Package `go-telegram` allows interact with [Telegram Bot API](https://core.telegram.org/bots/api#available-methods) and parse web hooks from Telegram without getting bored.

## Usage

To user `go-telegram` you should install it:

```
$ get github.com/kimrgrey/go-telegram
```

Now you can add `go-telegram` to the package list, create new client using your bot key and perform call of any method from [this list](https://core.telegram.org/bots/api#available-methods). Notice, that method names have been converted to satisfy go coding style, so, for example, if you need to call `getMe` you should use `client.GetMe()`, like this:

```go
package main

import (
        "fmt"
        "github.com/kimrgrey/go-telegram"
)
 
var client *telegram.Client = telegram.NewClient("YOUR_BOT_KEY")

func main() {
        me := client.GetMe()
        fmt.Println(me.UserName)
}
```

## License

This code has been published under MIT License.