package intelligence

import (
	"github.com/verticalgmbh/intelligence-core/pkg/pb"
)

type Client struct {
	cli pb.ApiClient
}

func (c *Client) NewTxn(set string) *Txn {
	return &Txn{set: set}
}
