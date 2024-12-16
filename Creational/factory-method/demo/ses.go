package main

type Ses struct {
	provider
}

func newSes() iProvider {
	return &Ses{
		provider{
			name: "Ses",
			key:  "SendSes",
		},
	}
}
