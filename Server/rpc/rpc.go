package rpc

import (
	"context"
	"fmt"
	"go-demo/Server/db"
	"go-demo/Services/models"
	"go-demo/Services/proto"
)

//type Edc int
type ComplaintHandler struct {

}

func (t *ComplaintHandler) Feed(ctx context.Context, args *proto.IndexFeedArgs, reply *proto.IndexFeedReply) error {
	switch args.Type {
	case proto.FEED_LIST_HOTTEST:
		sql := `select title from complaints 
			where created_at > 1634540908 && 
			status in (1,3,4,6,7) && exposed = 1 
			order by upvote_amount DESC, created_at DESC`
		res, _ := db.GetDefaultDbEngine().Query(sql)
		for _, v := range res {
			reply.Title = append(reply.Title, string(v["title"]))
		}
	case proto.FEED_LIST_LATEST:
		engine := db.GetDefaultDbEngine()
		var cpl []models.Complaints
		err := engine.
			Where("created_at > ? && exposed = ?", 1634540908, 1).
			In("status", []int{1, 3, 4, 6, 7}).
			Desc("upvote_amount").
			Desc("created_at").
			Find(&cpl)
		if err != nil{
			fmt.Println(err)
		}
		for _, v := range cpl {
			reply.Title = append(reply.Title, v.Title)
		}
	}
	return nil
}

func (t *ComplaintHandler) SelectComplaintBySn(ctx context.Context, args *proto.IndexFeedArgs, reply *proto.IndexFeedReply) error {
	engine := db.GetDefaultDbEngine()
	var cpl []models.Complaints
	engine.SQL("select * from complaints where couid = \"1000002\"").Find(&cpl)
	for _, v := range cpl {
		reply.Title = append(reply.Title, v.Title)
	}
	return nil
}

func (t *ComplaintHandler) SelectCompaniesByName(ctx context.Context,
	args *proto.SelectCompaniesByNameArgs, reply *proto.SelectCompaniesByNameReply) error {

	engine := db.GetDefaultDbEngine()
	codi := fmt.Sprintf("title like \"%%%s%%\" OR show_title like \"%%%s%%\"",  args.Name, args.Name)
	var cpl []models.Companies
	engine.
		Where(codi).
		Desc("valid_amount").
		Find(&cpl)
	for _, v := range cpl {
		reply.Id = append(reply.Id, v.Id)
	}
	return nil
}

