package xsync

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestAsync_Basic(t *testing.T) {
	expected := 99

	// given
	testee := Async(
		func() int {
			return 99
		},
	)

	// when && then
	if actual := <-testee(); actual != expected {
		t.Errorf("Got() = '%v', want '%v'", actual, expected)
	}
}

func TestAsync_NonBlocking(t *testing.T) {
	// given
	timeout := 1 * time.Second

	testee := Async(
		func() int {
			time.Sleep(timeout)
			return 0
		},
	)

	// when
	ch := testee()

	// then
	select {
	case <-ch:
		{
			t.Error(
				"Expected the channel to not be ready immediately, but it already contains a value",
			)
		}
	default:
		{

		}
	}

}

func TestAsync_ChannelClose(t *testing.T) {
	// given
	testee := Async(
		func() int {
			return 99
		},
	)

	// when
	actual := testee()
	if _, open := <-actual; !open {
		t.Error("Channel is closed")
	}

	// then
	_, open := <-actual
	if open {
		t.Error("Channel is open")
	}
}

func TestAsyncOnce_Basic(t *testing.T) {
	expected := 99

	// given
	testee := AsyncOnce(
		func() int {
			return 99
		},
	)

	// when && then
	if actual := <-testee(); actual != expected {
		t.Errorf("Got() = '%v', want '%v'", actual, expected)
	}
}

func TestAsyncOnce_NonBlocking(t *testing.T) {
	// given
	timeout := 1 * time.Second

	testee := AsyncOnce(
		func() int {
			time.Sleep(timeout)
			return 0
		},
	)

	// when
	ch := testee()

	// then
	select {
	case <-ch:
		{
			t.Error(
				"Expected the channel to not be ready immediately, but it already contains a value",
			)
		}
	default:
		{

		}
	}
}

func TestAsyncOnce_ChannelClose(t *testing.T) {
	// given
	testee := AsyncOnce(
		func() int {
			return 99
		},
	)

	// when
	actual := testee()
	if _, open := <-actual; !open {
		t.Error("Channel is closed")
	}

	// then
	_, open := <-actual
	if open {
		t.Error("Channel is open")
	}
}

func TestAsyncOnce_ConcurrentExecution(t *testing.T) {
	// given
	var counter atomic.Int64

	testee := AsyncOnce(
		func() int64 {
			counter.Add(1)
			return counter.Load()
		},
	)

	// when
	for i := 0; i < 100; i++ {
		<-testee()
	}

	// then
	if counter.Load() != 1 {
		t.Errorf(
			"Expected counter was incremented only once, but was incremented %d times",
			counter.Load(),
		)
	}
}
