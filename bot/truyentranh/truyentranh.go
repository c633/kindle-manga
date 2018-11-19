package truyentranh

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/vqhuy/kindle-manga/bot"
)

func init() {
	bot.Register(&collector{
		Bot: &bot.Bot{},
	})
}

type collector struct {
	*bot.Bot
}

var _ bot.Collector = (*collector)(nil)

func (b *collector) Page() string {
	return "http://truyentranh.net/"
}

func (b *collector) GetLink(base string, chap int) string {
	return fmt.Sprintf("%s/Chap-%03d/", base, chap)
}

func (b *collector) Collect(base string, chap int, outputDir string) {
	b.Colly = colly.NewCollector()

	b.Colly.OnHTML(`#viewer`, func(e *colly.HTMLElement) {
		e.ForEach("img", func(i int, ee *colly.HTMLElement) {
			link := ee.Attr("src")
			b.Colly.Visit(strings.TrimSpace(link))
		})
	})
	b.Colly.OnHTML(`div.each-page`, func(e *colly.HTMLElement) {
		e.ForEach("img", func(i int, ee *colly.HTMLElement) {
			link := ee.Attr("src")
			b.Colly.Visit(strings.TrimSpace(link))
		})
	})
	b.Colly.OnHTML(`div.OtherText`, func(e *colly.HTMLElement) {
		e.ForEach("img", func(i int, ee *colly.HTMLElement) {
			link := ee.Attr("src")
			b.Colly.Visit(strings.TrimSpace(link))
		})
	})

	b.Bot.Collect(base, chap, outputDir)

	link := b.GetLink(base, chap)
	b.Colly.Visit(link)
}
