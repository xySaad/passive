package facebook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"
	"sync"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const DOMAIN = "https://web.facebook.com/"

func CheckUsername(parent context.Context, username string) (nickname string, err error) {
	ctx, cancel := chromedp.NewContext(parent)
	defer cancel()

	requestID := make(chan network.RequestID, 3)
	var queue []network.RequestID
	var mx sync.Mutex

	chromedp.ListenTarget(ctx, func(ev any) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			if e.RedirectResponse != nil && e.RedirectResponse.URL == DOMAIN+username {
				username = strings.Split(e.Request.URL, DOMAIN)[1]
			}
			if e.Request.URL == DOMAIN+"ajax/bulk-route-definitions/" && e.Request.HasPostData {
				defer mx.Unlock()
				mx.Lock()
				queue = append(queue, e.RequestID)
			}
		case *network.EventLoadingFinished:
			defer mx.Unlock()
			mx.Lock()
			if slices.Contains(queue, e.RequestID) {
				requestID <- e.RequestID
			}
		}
	})
	err = chromedp.Run(ctx,
		chromedp.Navigate(DOMAIN+username),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for id := range requestID {
				postData, err := network.GetRequestPostData(id).Do(ctx)
				if err != nil {
					fmt.Println(err)
					continue
				}
				expected := "route_urls[0]=%2F" + url.QueryEscape(username)
				if postData[:len(expected)] != expected {
					continue
				}

				// return immediately since the expected request has found
				body, err := network.GetResponseBody(id).Do(ctx)
				if err != nil {
					return fmt.Errorf("error reading body: %s", err)
				}
				text := strings.TrimPrefix(string(body), "for (;;);")
				var respData Resp
				if err := json.Unmarshal([]byte(text), &respData); err != nil {
					return fmt.Errorf("failed JSON unmarshal: %s", err)
				}
				nickname = respData.Payload.Payloads["/"+username].Result.Exports.Meta.Title
				if nickname != "" {
					return nil
				} else {
					return errors.New("user not found in any response")
				}
			}
			return errors.New("unexpected error")
		}),
	)
	return nickname, err
}
