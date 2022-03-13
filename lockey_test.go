package lockey_test

import (
	"sync"
	"testing"
	"time"

	"github.com/barrydevp/lockey"
)

func TestKeyMutex(t *testing.T) {
	keyMutex := lockey.NewRWLockKey()

	var count = 0

	var wg sync.WaitGroup

	var num = 10000

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keyMutex.Lock("a")
			count += i
			keyMutex.Unlock("a")
		}(i)
	}

	wg.Wait()

	expected := 50005000

	if count != expected {
		t.Fatalf("exptected %d and actual %d", expected, count)
	}
}

func TestLockeyWithLock(t *testing.T) {

	ch := make(chan struct{}, 1)

	go func() {
		keyMutex := lockey.NewRWLockKey()

		var wg sync.WaitGroup

		var num = 10

		for i := 1; i <= num; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				keyMutex.Lock("a")
				time.Sleep(time.Second)
				keyMutex.Unlock("a")
			}(i)
		}

		wg.Wait()
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		t.Fatal("no serialization")
	case <-time.After(time.Second * 2):
	}
}

func TestNewLockeyWithoutLock(t *testing.T) {

	ch := make(chan struct{}, 1)

	go func() {
		keyMutex := lockey.NewRWLockKey()

		var wg sync.WaitGroup

		var num = 10

		for i := 1; i <= num; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				s := string((rune)((i-1)%26 + 'a'))
				keyMutex.Lock(s)
				time.Sleep(time.Second)
				keyMutex.Unlock(s)
			}(i)
		}

		wg.Wait()
		ch <- struct{}{}
	}()

	select {
	case <-ch:
	case <-time.After(time.Second * 2):
		t.Fatal("no serialization")
	}
}

func BenchmarkLockey(b *testing.B) {
	keyMutex := lockey.NewRWLockKey()

	var wg sync.WaitGroup

	for j := 0; j < 255; j++ {
		k := string((rune)(j))
		var count = 0
		for i := 0; i <= b.N; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				keyMutex.Lock(k)
				count += i
				keyMutex.Unlock(k)
			}(i)
		}
	}

	wg.Wait()
}
