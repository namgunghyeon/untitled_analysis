package writer

import (
  "fmt"
  "lib/model"
  "encoding/json"
  "io/ioutil"
)

func WriteToJson(project string, keyword model.KeywordMap) {
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
  _ = ioutil.WriteFile(project + ".json", jsonString, 0644)
}
