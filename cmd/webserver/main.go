package main

import (
	"errors"
	"time"
	"fmt"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

// Interface contains the behavior of a semaphore that can be acquired and/or released.
type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	fmt.Print("Aquire")
	select {
	case s.sem <- struct{}{}:
		fmt.Println("Aquire struct")
		return nil
	case <-time.After(s.timeout):
		fmt.Println("Aquire timeout")
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	fmt.Print("Release")
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}

	return nil
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

func main() {
	tickets, timeout := 1, 3*time.Second
	s := New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	time.Sleep(4*time.Second)
	if err := s.Acquire(); err != nil {
		panic(err)
	}
	// Do important work
	fmt.Println(":working")
	time.Sleep(10*time.Second)

	if err := s.Release(); err != nil {
		panic(err)
	}
}
