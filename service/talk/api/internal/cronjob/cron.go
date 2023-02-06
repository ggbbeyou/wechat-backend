package cronjob

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
	"wechat-backend/common/utils"
	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/model"
)

/*
* @Author: chuang
* @Date:   2023/2/4 14:22
 */

func StartCronJob(svc svc.ServiceContext) {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	cron := gocron.NewScheduler(timezone)

	cron.Every(2).Seconds().Do(func() {
		go cronJobUpdateLike(svc)
	})
}

func cronJobUpdateLike(svc svc.ServiceContext) {
	//1.从redis中获取点赞数据
	likeRecord, err := svc.Redis.Hgetall(utils.LIKE_KEY)
	if err != nil {
		logx.Error(err)
		return
	}
	//2.删除redis中该key
	_, err = svc.Redis.Del(utils.LIKE_KEY)
	if err != nil {
		logx.Error(err)
		return
	}
	//3.批量写入mysql中
	var data []*model.Likes
	snowFlake, _ := utils.NewSnowFlake(0, 0)

	for key, isLike := range likeRecord {
		snowId, _ := snowFlake.NextId()
		split := strings.Split(key, "::")

		data = append(data, &model.Likes{
			Lid:   fmt.Sprintf("%s%s", utils.LIKE_PREFIX, snowId),
			Tid:   split[0],
			Uid:   utils.Str2Int64(split[1]),
			State: utils.Str2Int64(isLike),
		})
	}
	_, err = svc.LikeModel.BatchInsert(context.Background(), data)
	if err != nil {
		logx.Error(err)
		return
	}
}
