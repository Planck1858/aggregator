package news

type Topic struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

type Archive map[string][]Topic

func New() Archive {
	return make(map[string][]Topic, 0)
}

func (a Archive) collect(category string) {
	sources := getSources(category)
	topics := getTopics(sources)

	a[category] = topics
}

func (a Archive) result(category string) []Topic {
	return a[category]
}
