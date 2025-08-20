package username

import (
	"context"
	"fmt"
	"passive/apis/username/facebook"
	"passive/apis/username/twitter"
	"strings"
)

const MISSING_INPUT = `Usage: passive -u "@user01"`

func Exec(parent context.Context, username string) {
	username = strings.TrimPrefix(username, "@")
	nickname, err := facebook.CheckUsername(parent, username)
	if err != nil {
		fmt.Println("Facebook: no")
	} else {
		fmt.Printf("Facebook: yes (%s) \n", nickname)
	}
	nickname, err = twitter.CheckUsername(parent, username)
	if err != nil {
		fmt.Println("Twitter: no")
	} else {
		fmt.Printf("Twitter: yes (%s) \n", nickname)
	}
}
