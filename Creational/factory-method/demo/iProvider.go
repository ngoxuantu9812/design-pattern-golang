package main

type iProvider interface {
	getName() string
	setName(string)
	getKey() string
	setKey(string)
}
