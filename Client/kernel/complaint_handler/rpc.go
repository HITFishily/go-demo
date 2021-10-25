package complaint_handler

import (
	"context"
	"go-demo/Services/proto"

	"github.com/smallnest/rpcx/client"
)

type Client struct {
	c client.XClient
}

func NewClient(endpoint string) *Client {
	d, _ := client.NewPeer2PeerDiscovery(endpoint, "")
	return &Client{
		c: client.NewXClient("ComplaintHandler", client.Failtry, client.RandomSelect, d, client.DefaultOption),
	}
}

func (p *Client) Feed(args *proto.IndexFeedArgs, reply *proto.IndexFeedReply) error {
	return p.c.Call(context.Background(), "Feed", args, reply)
}

func (p *Client) SelectComplaintBySn(args *proto.IndexFeedArgs, reply *proto.IndexFeedReply) error {
	return p.c.Call(context.Background(), "SelectComplaintBySn", args, reply)
}

func (p *Client) SelectCompaniesByName(args *proto.SelectCompaniesByNameArgs, reply *proto.SelectCompaniesByNameReply) error {
	return p.c.Call(context.Background(), "SelectCompaniesByName", args, reply)
}