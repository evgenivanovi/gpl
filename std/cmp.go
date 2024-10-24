package std

type Comparer interface {
	Compare(other any) int
}
