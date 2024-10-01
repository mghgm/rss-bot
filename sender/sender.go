package sender


type Post struct {
    title string
    desc string 
    tags []string
    ref string
}

type Sender interface {
    Send(post Post)
}
