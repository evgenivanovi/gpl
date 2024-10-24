package reflect_test

import (
	"github.com/evgenivanovi/gpl/std/reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type Foo struct {
	text string
}

func (f *Foo) Foo() {}

type IFoo interface {
	Foo()
}

var _ IFoo = &Foo{}

type Bar struct {
	value int
}

func (f *Bar) Bar() {}

type IBar interface {
	Bar()
}

var _ IBar = &Bar{}

func TestAssign(t *testing.T) {

	t.Run("ValueToValue", func(t *testing.T) {
		src := Foo{"text"}
		var dst Foo

		require.True(t, reflect.Assign(src, &dst))
		require.Equal(t, src, dst)
	})

	t.Run("ValueToValueInvalid", func(t *testing.T) {
		src := Foo{"text"}
		var dst Bar

		require.False(t, reflect.Assign(src, &dst))
		require.Equal(t, Bar{}, dst)
	})

	t.Run("ValueToInterface", func(t *testing.T) {
		src := Foo{"text"}
		var dst IFoo

		require.True(t, reflect.Assign(&src, &dst))
		require.NotNil(t, dst)

		v, ok := dst.(*Foo)

		require.True(t, ok)
		require.Equal(t, &src, v)
	})

	t.Run("ValueToInterfaceInvalid", func(t *testing.T) {
		src := Bar{42}
		var dst IFoo

		require.False(t, reflect.Assign(&src, &dst))
		require.Nil(t, dst)
	})

	t.Run("InterfaceToInterface", func(t *testing.T) {
		src := IFoo(&Foo{"text"})
		var dst IFoo

		require.True(t, reflect.Assign(src, &dst))
		require.NotNil(t, dst)
		require.Equal(t, src, dst)

		v, ok := dst.(*Foo)

		require.True(t, ok)
		require.Equal(t, src, v)
	})

	t.Run("InterfaceToInterfaceInvalid", func(t *testing.T) {
		src := IFoo(&Foo{"text"})
		var dst IBar

		require.False(t, reflect.Assign(src, &dst))
		require.Nil(t, dst)
	})

}
