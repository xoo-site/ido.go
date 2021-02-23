package model

type file_parser struct {
	Path, Size string
	Count      int
}

type ModelJob struct {
	...
}

type Reader interface {
	Read(string) ([]byte, error)
}

type Writer interface {
	Write(interface{}, string) error
}
