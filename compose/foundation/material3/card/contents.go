package card

type CardContentContainer struct {
	composableCount int
	children        []*CardChild
}

func newCardContentContainer() *CardContentContainer {
	return &CardContentContainer{
		composableCount: 0,
		children:        []*CardChild{},
	}
}

func (c *CardContentContainer) addChild(child *CardChild) {

	c.children = append(c.children, child)
}

func CardContents(children ...*CardChild) CardContentContainer {

	container := newCardContentContainer()
	for _, child := range children {
		container.addChild(child)
	}
	return *container
}

func Content(composable Composable) *CardChild {
	return &CardChild{
		composable:  composable,
		contentType: CardContent,
	}
}

func ContentCover(composable Composable) *CardChild {
	return &CardChild{
		cover:       true,
		composable:  composable,
		contentType: CardContentCover,
	}
}

type CardChild struct {
	cover       bool
	composable  Composable
	contentType CardContentType
}
