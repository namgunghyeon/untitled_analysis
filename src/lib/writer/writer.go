package writer

import (
  "fmt"
  "lib/model"
  "encoding/json"
  "io/ioutil"
  "gopkg.in/couchbase/gocb.v1"
  _ "github.com/lib/pq"
  "database/sql"
  "log"
  "config"
  "strings"
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

func CockroacWriteToKeywordIndex(keyword model.KeywordMap) {
  db, err := sql.Open("postgres", "postgresql://root@104.156.238.187:26257/untitled?sslmode=disable")
  if err != nil {
      log.Fatal("error connecting to the database: ", err)
  }
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      query := "INSERT INTO keyword_index(keyword) VALUES ('"+ item.Name + "') ON CONFLICT (keyword) DO NOTHING";
      fmt.Println("query", query)
      _, err := db.Exec(query);
      if err != nil {
        defer func() {
           if err := recover(); err != nil {
               fmt.Println(err)
           }
       }()
      }
    }
  }
}

func CockroacWriteToKeyword(name string, version string, keyword model.KeywordMap) {
  db, err := sql.Open("postgres", "postgresql://root@104.156.238.187:26257/untitled?sslmode=disable")
  if err != nil {
      log.Fatal("error connecting to the database: ", err)
  }
  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]
      query := "INSERT INTO keyword(keyword_index, project, version, type, path, count) VALUES ('" + item.Name + "', '" + name + "', '" + version + "', '" + item.Type + "', '" + item.Path + "', 0)";
      fmt.Println("query", query)
      _, err := db.Exec(query);
      if err != nil {
        defer func() {
           if err := recover(); err != nil {
               fmt.Println("recover", err)
           }
       }()
      }
    }
  }
}

func concatProejct(project string, name string) string{
    if project != "" {
      project += ", " + name
    } else {
      project += name
    }
    return project
}

func errorHandle(err error) {
    if err != nil {
      fmt.Println("error", err)
    }
    if err := recover(); err != nil {
        fmt.Println("error", err)
    }
}

func CockroacWriteToKeywordIndexProjectMeta(name string, keyword model.KeywordMap) {
  db, err := sql.Open("postgres", "postgresql://root@104.156.238.187:26257/untitled?sslmode=disable")
  if err != nil {
      log.Fatal("error connecting to the database: ", err)
  }

  for k, _ := range keyword {
    for i := 0; i < len(keyword[k]); i++ {
      item := keyword[k][i]

      selectQuery := "SELECT keyword_index, project FROM keyword_index_project_meta where keyword_index = '" + item.Name + "'";
      rows, err := db.Query(selectQuery)
      fmt.Println("query", selectQuery)
      defer errorHandle(err)
      defer rows.Close()
      if (rows == nil) {
        continue
      }

      var keyword_index string
      var project string
      for rows.Next() {
        err := rows.Scan(&keyword_index, &project)
        errorHandle(err)
      }
      isContain := strings.Contains(project, name)
      if isContain {
        continue
      }
      project = concatProejct(project, name)
      insertQuery := "INSERT INTO keyword_index_project_meta(keyword_index, project) VALUES ('" + item.Name + "', '" + project + "') ON CONFLICT (keyword_index) DO UPDATE SET project = '" + project + "'"
      fmt.Println("query", insertQuery)
      _, insertErr := db.Exec(insertQuery);
      errorHandle(insertErr)
    }
  }
}

func CockroacWriteToProject(project string, color string) {
  db, err := sql.Open("postgres", "postgresql://root@104.156.238.187:26257/untitled?sslmode=disable")
  if err != nil {
      log.Fatal("error connecting to the database: ", err)
  }

  query := "INSERT INTO project(name, color) VALUES ('" + project + "', '" + color + "')";
  fmt.Println("query", query)
  _, insertErr := db.Exec(query);
  defer errorHandle(insertErr)
}
