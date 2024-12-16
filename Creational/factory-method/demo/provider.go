package main

type provider struct {
	name string
	key  string
}

func (p *provider) setName(name string) {
	p.name = name
}

func (p *provider) setKey(key string) {
	p.key = key
}

func (p *provider) getName() string {
	return p.name
}

func (p *provider) getKey() string {
	return p.key
}
