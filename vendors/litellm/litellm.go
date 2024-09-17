package litellm

import (
	"github.com/danielmiessler/fabric/vendors/openai"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatible("litellm", "http://localhost:4000/", nil)

	return
}

type Client struct {
	*openai.Client
}
