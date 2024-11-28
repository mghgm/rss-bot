package sender


type Post struct {
    Title string
    Desc string
    Tags []string
    Ref string
}

type Sender interface {
    Send(post Post)
}
