package cmp

type Comparer interface {
	Compare(other any) int
}
