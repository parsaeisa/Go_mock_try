package doer

type Doer interface {
	DoSomething(i int, s string) error
}
