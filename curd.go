package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
)

func insert(ctx context.Context, cli *elastic.Client) {
	// 创建创建一条微博
	msg1 := Weibo{User: "olivere", Message: "打酱油的一天", Retweets: 0}

	// 使用client创建一个新的文档
	put1, err := cli.Index().
		Index("weibo"). // 设置索引名称
		Id("1").        // 设置文档id
		BodyJson(msg1). // 指定前面声明的微博内容
		Do(ctx)         // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
}

func get(ctx context.Context, cli *elastic.Client) {
	// 根据id查询文档
	get1, err := cli.Get().
		Index("weibo"). // 指定索引名
		Id("1").        // 设置文档id
		Do(ctx)         // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
	}

	//手动将文档内容转换成go struct对象
	msg2 := Weibo{}
	// 提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()
	// 将json转成struct结果
	json.Unmarshal(data, &msg2)
	// 打印结果
	fmt.Println(msg2.Message)
}

func insertReport(ctx context.Context, cli *elastic.Client) {
	id1 := &indexRecord{
		Indexname: "report",
		Timestamp: time.Now().UnixNano() / 1000 / 1000,
		Datas: map[string]interface{}{
			"appid":      23,
			"uid":        12,
			"taskid":     "id1234",
			"streamname": "request.Streamname",
			"timestamp":  213344,
			"duration":   1234,
			"mediaurl":   "request.Mediaurl",
		},
	}

	id1.Datas["servip"] = "localhost"
	id1.Datas["timestamp"] = "1234567"
	res, err := elastic.NewBulkService(cli).Index("report").
		Add(id1).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func getReport(ctx context.Context, cli *elastic.Client) error {
	// 根据id查询文档
	//termQuery := elastic.NewTermQuery("user", "olivere")
	res, err := cli.Search().
		Index("report").
		//Query(termQuery).
		Do(ctx)
	if err != nil {
		return err
	}
	hits := res.Hits.TotalHits.Value
	if hits > 0 {
		fmt.Println("has:", hits)
		for _, hit := range res.Hits.Hits {
			//rp := new(indexRecord)
			rp := new(Data)
			err = json.Unmarshal(hit.Source, rp)
			if err != nil {
				return err
			}
			fmt.Println(err, rp, "xxxx")
		}
	} else {
		err = errors.New("not found sources")
	}
	return err

	// 提取文档内容，原始类型是json数据
	//data, _ := get1.Source.MarshalJSON()
}
