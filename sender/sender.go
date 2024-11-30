package sender

import (
    "github.com/mghgm/camelnews/collector"
)


type Post struct {
    Title string
    Desc string
    Tags []string
    Ref string
}

type Sender interface {
    Start() chan collector.News
    send(chatID int64, post Post) error
}
