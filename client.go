package ctoai

// Client is the top-level client for Ops Platform services
type Client struct {
	Prompt *Prompt
	Ux     *Ux
	Sdk    *Sdk
}

// NewClient creates an Ops Platform client with all services included
func NewClient() Client {
	return Client{
		Prompt: NewPrompt(),
		Ux:     NewUx(),
		Sdk:    NewSdk(),
	}
}
