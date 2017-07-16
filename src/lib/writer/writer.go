package writer

import (
  "fmt"
  "lib/model"
  "encoding/json"
  "io/ioutil"
  "gopkg.in/couchbase/gocb.v1"
)

type Keyword struct {
	Project string `json:"project"`
	Version string `json:"version"`
	Type string `json:"type"`
  Path string `json:path`
  Name string `json:name`
  Count int `json:count`
}

func WriteToJson(name string, version string, keyword model.KeywordMap) {
  fmt.Println()
  fmt.Println("function", len(keyword["function"]))
  fmt.Println("value", len(keyword["value"]))
  fmt.Println("class", len(keyword["class"]))
  fmt.Println("interface", len(keyword["interface"]))

  datas := make(map[int]model.JsonKeyword)
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      datas[i] = model.JsonKeyword{Type: item.Type, Name: item.Name, Path: item.Path}
    }
  }
  v := make([]model.JsonKeyword, 0, len(datas))
  for  _, value := range datas {
     v = append(v, value)
  }
  jsonString, _ := json.Marshal(v)
  _ = ioutil.WriteFile(name + "_" + version + ".json", jsonString, 0644)
}

func WriteToCouchbase(name string, version string, keyword model.KeywordMap) {
  cluster, _ := gocb.Connect("couchbase://192.168.56.213")
  bucket, _ := cluster.OpenBucket("default", "")
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      fmt.Println("item", item)
      key := name + "_" + version + "_" +  item.Type + "_" +  item.Name
      bucket.Insert(key, Keyword{
        Project: name,
        Version: version,
        Name: item.Name,
        Type: item.Type,
        Path: item.Path,
        Count: 0,
      }, 0)
    }
  }
}
