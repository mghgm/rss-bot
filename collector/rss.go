package collector

import (
    "time"

    "github.com/mmcdole/gofeed"
)

type rssAgency struct {
    Agency
    link string
}

func NewRSSAgency(title, category, link string) *rssAgency {
    return &rssAgency{
        Agency: Agency{
            title: title,
            category: category,
            lastFetch: time.Time{},
        },
        link: link,
    } 
}

func (ra rssAgency) Collect() ([]News, error) {
    fp := gofeed.NewParser()
    feed, err := fp.ParseURL(ra.link)
    if err != nil {
        return nil, err 
    } 
    
    news := make([]News, 0)

    for _, item := range feed.Items {
        news = append(news, News{
            Title: item.Title,
            Desc: item.Description,
            Link: item.Link,
        }) 
    }

    return news, nil
}
