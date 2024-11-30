package collector

import (
    "time"
)

type News struct {
    Title string
    Desc string
    Link string
}

type AgencyCollector struct {
    title string 
    category string 
    lastFetch time.Time
    scrapeDuration time.Duration
}

type Collector interface {
    Start() chan News
    collect(updates chan <- News) (error)
}

