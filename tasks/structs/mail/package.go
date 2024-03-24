package mail

type Package struct {
	content Content

	abstractSendable
}

func (pkg *Package) GetContent() Content {
	return pkg.content
}
