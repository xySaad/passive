package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const DOMAIN = "https://x.com/"

func CheckUsername(parent context.Context, username string) (nickname string, err error) {
	ctx, cancel := chromedp.NewContext(parent)
	defer cancel()
	var requestID network.RequestID
	found := make(chan struct{}, 1)
	chromedp.ListenTarget(ctx, func(ev any) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			if strings.Contains(e.Request.URL, "UserByScreenName") && e.Request.Method == "GET" {
				requestID = e.RequestID
			}
		case *network.EventLoadingFinished:
			if requestID == e.RequestID {
				close(found)
			}
		}
	})

	err = chromedp.Run(ctx,
		chromedp.Navigate(DOMAIN+username),

		chromedp.ActionFunc(func(ctx context.Context) error {
			<-found
			body, err := network.GetResponseBody(requestID).Do(ctx)
			if err != nil {
				return fmt.Errorf("error reading body: %s", err)
			}
			var respData Resp
			if err := json.Unmarshal(body, &respData); err != nil {
				return fmt.Errorf("failed JSON unmarshal: %s", err)
			}
			nickname = respData.Data.User.Result.Core.Name
			if nickname == "" {
				return fmt.Errorf("user not found: %s", err)
			}
			return nil
		}),
	)
	return nickname, err
}
