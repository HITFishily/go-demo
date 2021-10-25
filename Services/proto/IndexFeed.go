package proto

const (
	FEED_LIST_DEFAULT = iota
	FEED_LIST_HOTTEST
	FEED_LIST_LATEST
	FEED_LIST_REPLIED
	FEED_LIST_FINISHED
)

type IndexFeedArgs struct {
	Type int
}

type IndexFeedReply struct {
	Title []string
}

type SelectCompaniesByNameArgs struct {
	Name string	`json:"name"`
}

type SelectCompaniesByNameReply struct {
	Id []int `json:"id"`
}