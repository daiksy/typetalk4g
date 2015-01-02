# Typetalk4g

## Usage

### post messages
```go
package main

import (
	"fmt"
	typetalk "github.com/daiksy/typetalk4g"
)

func main() {

	client := typetalk.NewClient({topic_id}, "{token}")

	res, err := client.Post("hello! typetalk")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
```

### get messages
```go
package main

import (
	"fmt"
	typetalk "github.com/daiksy/typetalk4g"
)

func main() {

	client := typetalk.NewClient({topic_id}, "{token}")
	client.MessagesSize = 100 // default 20

	res, err := client.GetTopicMessages()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	for _, post := range res.Posts {
		fmt.Println(post.Account.FullName + ":" + post.Message)
	}
}
```