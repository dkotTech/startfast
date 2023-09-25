package startfast

import "sync"

type nonblocking[T any] struct {
	once sync.Once
	val  T
	err  error
	fn   func() (T, error)
}

// NewEager immediately returns a new Lazy[T], initializing it in the background.
func NewEager[T any](f func() (T, error)) *nonblocking[T] {
	nb := &nonblocking[T]{fn: f}
	go func() {
		nb.Get()
	}()
	return nb
}

// NewLazy a new Lazy[T], initializing it on the first call to Get(), MustGet(), or NonblockingInit().
func NewLazy[T any](f func() (T, error)) *nonblocking[T] {
	return &nonblocking[T]{fn: f}
}

// Get returns the value of the Lazy[T], initializing it if necessary.
func (nb *nonblocking[T]) Get() (T, error) {
	nb.once.Do(func() {
		// - handle panics
		// - time the result
		// - use the `runtime` package to get the file:line of the CALLER of Get(), not the file:line of the Get() call,
		// - and associate that with the timed result
		nb.val, nb.err = nb.fn()
	})
	return nb.val, nb.err
}

// MustGet is as Get, but panics on error.
func (nb *nonblocking[T]) MustGet() T {
	v, err := nb.Get()
	if err != nil {
		panic(err)
	}
	return v
}
