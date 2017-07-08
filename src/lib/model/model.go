package model

type File struct {
  Path string
  Ext string
  Language string
}

type Scan struct{
  Files []File
}
