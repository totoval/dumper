package facade

type Dumper interface {
	Dump(v ...interface{})
	DD(v ...interface{})
}
