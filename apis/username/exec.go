package username

import (
	"context"
	"fmt"
	"passive/apis/username/meta"
	"passive/apis/username/twitter"
	"strings"
)

const MISSING_INPUT = `Usage: passive -u "@user01"`

func Exec(parent context.Context, username string) {
	username = strings.TrimPrefix(username, "@")
	meta.CheckUsername(parent, username)
	nickname, err := twitter.CheckUsername(parent, username)
	if err != nil {
		fmt.Println("Twitter: no")
	} else {
		fmt.Printf("Twitter: yes (%s) \n", nickname)
	}
}
