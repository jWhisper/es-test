package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

//https://www.tizi365.com/archives/858.html
func main() {
	// 创建client
	client, err := elastic.NewClient(
		elastic.SetURL("http://120.78.142.42:9200"),
		elastic.SetSniff(false),
		//elastic.SetBasicAuth("user", "secret")
	)
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 首先检测下weibo索引是否存在
	exists, err := client.IndexExists("report").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := client.CreateIndex("report").BodyString(reportMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	//insert(ctx, client)
	//get(ctx, client)

	//insertReport(ctx, client)
	err = getReport(ctx, client)
	fmt.Println(err)
}
