package search

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"wechat-backend/service/search/api/internal/svc"
	"wechat-backend/service/search/api/internal/types"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {

	var r map[string]interface{}
	var buf bytes.Buffer
	var messages []types.MessageInfo

	//matchPhrase := map[string]interface{}{
	//	"content": req.Keyword,
	//}
	//term := map[string]interface{}{
	//	"from": req.Uid,
	//	"to":   req.FriendUid,
	//}
	//srange := map[string]interface{}{
	//	"from": req.EndDate,
	//	"to":   req.StartDate,
	//}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					//"match_phrase": map[string]interface{}{
					//	"content": req.Keyword,
					//},
					////"term": map[string]interface{}{
					////	"from": req.Uid,
					////	"to":   req.FriendUid,
					////},
					//"range": map[string]interface{}{
					//	"from": req.EndDate,
					//	"to":   req.StartDate,
					//},

				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
		return nil, err
	}

	client := l.svcCtx.ElasticClient.Client

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(l.svcCtx.ElasticClient.Index),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
		client.Search.WithFrom(int(req.From)),
		client.Search.WithSize(int(req.Size)),
	)
	if err != nil {
		logx.Errorf("Error getting response: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		m := hit.(map[string]interface{})["_source"].(map[string]interface{})
		messages = append(messages,
			types.MessageInfo{
				Content:  m["content"].(string),
				From:     int64(m["from"].(float64)),
				To:       int64(m["to"].(float64)),
				Type:     int64(m["type"].(float64)),
				SendTime: m["sendtime"].(string),
			})
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	return &types.SearchResponse{
		RecordList: messages,
	}, err
}
