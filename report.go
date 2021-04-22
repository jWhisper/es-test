package main

import (
	"encoding/json"
	"fmt"
)

type indexRecord struct {
	Indexname string `json:"indexname,omitempty"`
	Timestamp int64  `json:"timestamp,string,omitempty"` // 毫秒
	Datas     map[string]interface{}
}

// int64问题
type Data struct {
	Appid      json.Number `json:"appid,omitempty"`
	Uid        json.Number `json:"uid,omitempty"`
	Taskid     string      `json:"taskid,omitempty"`
	Streamname string      `json:"streamname,omitempty"`
	Timestamp  json.Number `json:"timestamp,omitempty"`
	Duration   json.Number `json:"duration,omitempty"`
	Mediaurl   string      `json:"mediaurl,omitempty"`
}

func (rec *indexRecord) String() string {
	return ""
}

func (rec *indexRecord) Source() ([]string, error) {
	jsondata, err := json.Marshal(rec.Datas)
	if err != nil {
		fmt.Printf("add index json marshal err %s\n", rec)
		return nil, err
	}
	return []string{fmt.Sprintf(`{"index" : { "_index" : "%s", "_type" : "_doc" }}`, rec.Indexname), string(jsondata)}, nil
}

const reportMapping = `
{
  "settings": {
	"number_of_shards": 15,
	"number_of_replicas": 1, 
	"refresh_interval": "60s" 
	},
  "mappings": { 
	"properties": { 
	  "timestamp": { 
		"format": "strict_date_optional_time||epoch_millis",  
		"type": "date" 
	  } 
	}
  } 
}`

const another = `
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
