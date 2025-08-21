package telegram

import (
	"context"
	"errors"
	"strings"

	"github.com/chromedp/chromedp"
)

func CheckUsername(parent context.Context, username string) (nickname string, err error) {
	ctx, cancel := chromedp.NewContext(parent)
	defer cancel()

	var pageText string
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://t.me/"+username),
		chromedp.OuterHTML(`.tgme_page`, &pageText, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if strings.Contains(pageText, `<div class="tgme_page_extra">`) {
		err = chromedp.Run(ctx, chromedp.Text(`.tgme_page_title span`, &nickname, chromedp.NodeVisible, chromedp.ByQuery))
		return
	}

	return "", errors.New("user not found")
}
