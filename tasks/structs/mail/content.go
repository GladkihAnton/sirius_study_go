package mail

type Content struct {
	content string
	price   int
}

func (content *Content) GetContent() string {
	return content.content
}

func (content *Content) GetPrice() int {
	return content.price
}
