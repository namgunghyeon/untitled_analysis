package model

type File struct {
  Path string
  Ext string
  Language string
}

type Scan struct{
  Files []File
}

type Keyword struct {
  Type string
  Path string
  Name string
}

type KeywordMap map[string][]Keyword

type JsonKeyword struct {
  Type string `json:"Type"`
  Path string `json:"Path"`
  Name string `json:"Name"`
}
