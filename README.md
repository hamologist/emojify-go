# Emojify
**Emojify** is a simple Go package for spicing up a blob of text with a random assortment of emojis.

## Installation
To install and start using emojify:
```bash
$ go get github.com/hamologist/emojify-go
```

## Usage
Once emojify has been installed, you can start feeding it payloads for emojification.
Below is a basic example of how a string is fed into the `EmojifyTransformer` for transformation:
```go
package main

import (
	"fmt"
	"github.com/hamologist/emojify/pkg/model"
	"github.com/hamologist/emojify/pkg/transformer"
)

func main() {
	resp, err := transformer.EmojifyTransformer(
		model.EmojifyPayload{
			Message: "Hello world, this is a test!",
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Message)
}
```

The above will get a message like the following printed out:
```
Hello🏋️‍♀️ world,👉 this is👩‍👦🔘 a👾 test!
```

## Tools
Emojify currently provides the following tools for leveraging the power of emojification:
* [emojifying web server](cmd/emojify-server/main.go)
* [command line emojifier](cmd/emojify/main.go)

Please feel free to use them as inspiration for your own works.
You can install them using the following:
```bash
# Installs the command line emojifier
$ go install github.com/hamologist/emojify-go/cmd/emojify

# Installs the emojifying web server
$ go install github.com/hamologist/emojify-go/cmd/emojify-server
```

## Emojifying Web Server
Installing the web server will add an "emojify-server" binary to the system.
If executed, the server will start running on ":3000".
The server takes requests on its "/emojify" endpoint.
Requests must be a POST.

You can hit the server using curl like this:
```bash
curl --location --request POST 'localhost:3000/emojify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "So wait, we are going to emojify this payload now?"
}'
```

**Note:** Additional help can be found by running:
```
$ emojify-server help
```

## Command Line Emojifier
Installing the command line tool will add an "emojify" binary to the system.
If executed, the tool will read text from STDIN until terminal line input has ended.

A common use case for using this would be piping text into the tool:
```bash
echo "Now we are emojifying text from the command line!?" | emojify
```
Likewise, emojify can be started without piping text in.
Just input all of the text you would like to see emojified and send an end terminal line input sequence (typically ctrl-d).

**Note:** Additional help can be found by running:
```
$ emojify help
```

## Related
* [Emojify Sam](https://github.com/hamologist/emojify-sam)
	* AWS SAM project capable of creating lambda and API gateway resources for hosting an Emojify service.
	* Also supports deploying resources for processing Discord slash command requests.
	* Public web API is also provided [here!](https://github.com/hamologist/emojify-sam#public-resource)
* [Meme Machine](https://github.com/hamologist/meme-machine)
    * Python discord bot that supports emojifying messages via `!emojify` [command](https://github.com/hamologist/meme-machine/blob/master/meme_machine/emojify.py)