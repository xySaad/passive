package telegram

import (
	"context"
	"errors"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func CheckUsername(parent context.Context, username string) (nickname string, err error) {
	ctx, cancel := chromedp.NewContext(parent)
	defer cancel()
	done := make(chan struct{}, 1)

	chromedp.ListenTarget(ctx, func(ev any) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			if e.RedirectResponse != nil {
				err = errors.New("invalid username")
				close(done)
			}
		}
	})

	go func() {
		var pageText string
		err2 := chromedp.Run(ctx,
			chromedp.Navigate("https://t.me/"+username),
			chromedp.OuterHTML(`.tgme_page`, &pageText, chromedp.NodeVisible, chromedp.ByQuery),
		)
		if err2 == context.Canceled {
			return
		}
		defer close(done)
		if err2 != nil {
			err = err2
			return
		}
		if strings.Contains(pageText, `<div class="tgme_page_extra">`) {
			err = chromedp.Run(ctx, chromedp.Text(`.tgme_page_title span`, &nickname, chromedp.NodeVisible, chromedp.ByQuery))
		} else {
			err = errors.New("user not found")
		}
	}()

	<-done
	return nickname, err
}
