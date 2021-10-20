package main

type IPublication interface {
	setName(value string)
	setPages(value int)
	setPublisher(value string)
	getName() string
	getPages() int
	getPublisher() string
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) setName(value string) {
	p.name = value
}

func (p *publication) setPages(value int) {
	p.pages = value
}

func (p *publication) setPublisher(value string) {
	p.publisher = value
}

func (p *publication) getName() string {
	return p.name
}

func (p *publication) getPages() int {
	return p.pages
}

func (p *publication) getPublisher() string {
	return p.publisher
}
