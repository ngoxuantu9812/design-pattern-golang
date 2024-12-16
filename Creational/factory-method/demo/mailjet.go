package main

type Mailjet struct {
	provider
}

func newMailjet() iProvider {
	return &Mailjet{
		provider{
			name: "Mailjet",
			key:  "SendMailjet",
		},
	}
}
