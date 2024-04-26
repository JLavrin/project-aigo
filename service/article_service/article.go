package article_service

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (a *Article) Create() bool {
	return true
}

func (a *Article) Update() bool {
	return true
}

func (a *Article) Get() Article {
	article := Article{
		ID:        1,
		Title:     "Hello World",
		Body:      "This is a test article",
		CreatedAt: "2021-01-01 00:00:00",
		UpdatedAt: "2021-01-01 00:00:00",
	}

	return article
}

func (a *Article) GetAll() []Article {
	articles := []Article{
		{
			ID:        1,
			Title:     "Hello World",
			Body:      "This is a test article",
			CreatedAt: "2021-01-01 00:00:00",
			UpdatedAt: "2021-01-01 00:00:00",
		},
		{
			ID:        2,
			Title:     "Hello World 2",
			Body:      "This is a test article 2",
			CreatedAt: "2021-01-01 00:00:00",
			UpdatedAt: "2021-01-01 00:00:00",
		},
	}

	return articles
}

func (a *Article) Delete() bool {
	return true
}
