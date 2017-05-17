package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []article{
	{1, "Article 1", "asdfasdfsadfsadfsadfsadf"},
	{2, "Article 2", "qqqqqqqqqqqqqqqqqqqqqqq"},
	{3, "Article 3", "33333333333333333333333"},
}

func getAllArticles() []article {
	return articles
}

func getArticleById(id int) *article {
	for _, v := range articles {

		if v.ID == id {
			return &v
		}

	}
	return nil
}
