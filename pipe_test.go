package pipe

import (
	"errors"
	"fmt"
	"testing"
)

func TestPipe(t *testing.T) {
	t.Run("error in the middle", func(t *testing.T) {
		err := New().
			Next(func() error { fmt.Println("one"); return nil }).
			Next(func() error { fmt.Println("two"); return nil }).
			Next(func() error { fmt.Println("three"); return errors.New("error on three") }).
			Next(func() error { fmt.Println("four"); return nil }).
			Do()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	})
	t.Run("on error", func(t *testing.T) {
		New().
			Next(func() error { fmt.Println("one"); return nil }).
			Next(func() error { fmt.Println("two"); return nil }).
			Next(func() error { fmt.Println("three"); return nil }).
			Next(func() error { fmt.Println("four"); return nil }).
			OnErr(func(err error) { fmt.Printf("error: %v\n", err) })
	})
	t.Run("no errors", func(t *testing.T) {
		err := New().
			Next(func() error { fmt.Println("one"); return nil }).
			Next(func() error { fmt.Println("two"); return nil }).
			Next(func() error { fmt.Println("three"); return nil }).
			Next(func() error { fmt.Println("four"); return nil }).
			Do()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	})
	t.Run("realistic? not really", func(t *testing.T) {
		f1 := func() error {
			return nil
		}
		f2 := func() error {
			return errors.New("some error")
		}

		New().
			Next(f1).
			Next(f2).
			OnErr(func(err error) {
				fmt.Printf("error: %v\n", err)
			})
	})
}
