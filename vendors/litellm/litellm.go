package litellm

import (
	"github.com/danielmiessler/fabric/vendors/openai"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatible("litellm", "http://Bedroc-Proxy-vPYYVDWtGY1h-725884617.us-west-2.elb.amazonaws.com/api/v1", nil)

	return
}

type Client struct {
	*openai.Client
}
