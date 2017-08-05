package writer

import (
  "fmt"
  "lib/model"
  "encoding/json"
  "io/ioutil"
  "gopkg.in/couchbase/gocb.v1"
  "config"
)

type Project struct {
	Project string `json:"project"`
	Version string `json:"version"`
	Type string `json:"type"`
  Path string `json:"path"`
  Name string `json:"name"`
  Count int `json:"count"`
}

type ProjectInfo struct {
	Name string `json:"name"`
	Language string `json:"language"`
	Version string `json:"version"`
}

type KeywordIndex struct {
  Keyword string `json:"keyword"`
}

type Keyword struct {
  Keyword_Index string `json:"keyword_index"`
  Project string `json:"project"`
  Version string `json:"version"`
  Type string `json:"type"`
  Path string `json:"path"`
  Count int `json:"count"`
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

func WriteToProejct(name string, version string, keyword model.KeywordMap) {
  config := config.LoadCouchbase()
  cluster, _ := gocb.Connect("couchbase://" + config.Couchbase.Host)
  bucket, _ := cluster.OpenBucket(config.Couchbase.Project, "")
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      fmt.Println("item", item)
      key := name + "_" + version + "_" +  item.Type + "_" +  item.Name
      bucket.Insert(key, Project{
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

func WriteToKeywordIndex(keyword model.KeywordMap) {
  config := config.LoadCouchbase()
  cluster, _ := gocb.Connect("couchbase://" + config.Couchbase.Host)
  bucket, _ := cluster.OpenBucket("keyword_index", "")
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      fmt.Println("item", item)
      key := item.Name
      bucket.Insert(key, KeywordIndex{
        Keyword: item.Name,
      }, 0)
    }
  }
}

func WriteToKeyword(name string, version string, keyword model.KeywordMap) {
  config := config.LoadCouchbase()
  cluster, _ := gocb.Connect("couchbase://" + config.Couchbase.Host)
  bucket, _ := cluster.OpenBucket("keyword", "")
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      fmt.Println("item", item)
      key := name + "_" + version + "_" +  item.Type + "_" +  item.Name
      bucket.Insert(key, Keyword{
        Keyword_Index: item.Name,
        Project: name,
        Version: version,
        Type: item.Type,
        Path: item.Path,
        Count: 0,
      }, 0)
    }
  }
}


func WriteToProejctInfo(name string, version string, language string) {
  config := config.LoadCouchbase()
  cluster, _ := gocb.Connect("couchbase://" + config.Couchbase.Host)
  bucket, _ := cluster.OpenBucket(config.Couchbase.ProjectInfo, "")
  key := name + "_" + version
  bucket.Insert(key, ProjectInfo{
    Name: name,
    Language: language,
    Version: version,
  }, 0)
}
