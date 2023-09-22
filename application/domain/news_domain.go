package domain

type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Article struct {
	Source      Source
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type Source struct {
	ID   string
	Name string
}
