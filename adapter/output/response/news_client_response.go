package response

type NewsClientResponse struct {
	Status       string
	TotalResults int
	Articles     []ArticleResponse
}

type ArticleResponse struct {
	Source      SourceResponse
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type SourceResponse struct {
	ID   string
	Name string
}
