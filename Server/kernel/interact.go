package kernel

import (
	"go-demo/Client/kernel/complaint_handler"
)

var proxies Proxies

type Proxies struct {
	ComplaintProxy *complaint_handler.Client
}

func Init(endPoint string) {
	proxies.ComplaintProxy = complaint_handler.NewClient(endPoint)
}

func GetEdcProxy() *complaint_handler.Client {
	return proxies.ComplaintProxy
}
