package sender


type Post struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Tags []string `json:"tags"`
    Ref string `json:"ref"`
}
type Bal interface{}

type Sender interface {
    Send(post Post)
}
