package collector

import (
	"time"
    "log"

	"github.com/mmcdole/gofeed"

    "github.com/mghgm/camelnews/config"
)

const (
    CollectorBufferSize = 5
)


var (
    _ Collector = &rssAgencyCollector{}
)


type rssAgencyCollector struct {
    AgencyCollector
    link string
}

func NewRSSAgencyCollectorFromConfig(cfg config.CollectorConfig) *rssAgencyCollector {
    return NewRSSAgencyCollector(cfg.Title, cfg.Category, cfg.Link, cfg.ScrapeDuration)
}

func NewRSSAgencyCollector(title, category, link string, scrapeDuration time.Duration) *rssAgencyCollector {
    return &rssAgencyCollector{
        AgencyCollector: AgencyCollector{
            title: title,
            category: category,
            lastFetch: time.Time{},
            scrapeDuration: scrapeDuration,
        },
        link: link,
    } 
}

func (ra *rssAgencyCollector) Start() chan News {
    ticker := time.NewTicker(ra.scrapeDuration) 
 
    updates := make(chan News, CollectorBufferSize)
    

    log.Println(ra.scrapeDuration)
    go func() {
        defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                ticker.Stop()
                ra.collect(updates)
                ticker.Reset(ra.scrapeDuration)
            }
        }
    } ()
    
    return updates
}

func (ra *rssAgencyCollector) collect(updates chan <- News) error {
    fp := gofeed.NewParser()
    
    feed, err := fp.ParseURL(ra.link)
    if err != nil {
        return err 
    } 
    
    lastTimestamp := ra.lastFetch
    log.Println(ra.lastFetch)
    log.Println(len(feed.Items))
    
    for _, item := range feed.Items {
        pubDate, err := pubDate2Time(item.Published)
        if err != nil {
            continue
        }

        if pubDate.After(ra.lastFetch) {
            updates <- News{
                Title: item.Title,
                Desc: item.Description,
                Link: item.Link,
            }
        }

        if pubDate.After(lastTimestamp) {
            lastTimestamp = pubDate
        }
    }
    ra.lastFetch = lastTimestamp

    return nil
}

func pubDate2Time(pubDate string) (time.Time, error) {
    layout := "Mon, 02 Jan 2006 15:04:05 -0700"
    t, err := time.Parse(layout, pubDate)
    return t, err
}
