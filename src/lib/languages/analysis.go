package languages

type Analysis interface {
  ReadValues()
  ReadFunctions()
  ReadClasses()
  ReadInterfaces()
}

func InTheForest(p Analysis) {
  p.ReadValues()
  p.ReadFunctions()
  p.ReadClasses()
  p.ReadInterfaces()
}
