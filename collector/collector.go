package collector

import (
    "time"
)

type News struct {
    Title string
    Desc string
    Link string
}

type Agency struct {
    title string 
    category string 
    lastFetch time.Time
}

type Collector interface {
    Collect() []News
}
