package observer

import (
	"errors"
)

type Subject interface {
	Attach(o Observer) (bool, error)
	Detach(o Observer) (bool, error)
	Notify() (bool, error)
}

type Observer interface {
	Update(string)
}

type EmailNotification struct {
	// name string
}

func (s *EmailNotification) Update(bookingId string) {
	println("booking done for bookingId", bookingId)
}

type BookingSubject struct {
	ticker string

	observers []Observer
}

func (s *BookingSubject) Attach(o Observer) (bool, error) {

	for _, observer := range s.observers {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}
	s.observers = append(s.observers, o)
	return true, nil
}

func (s *BookingSubject) Detach(o Observer) (bool, error) {

	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}

func (s *BookingSubject) Notify() (bool, error) {
	for _, observer := range s.observers {
		observer.Update(s.String())
	}
	return true, nil
}

func (s *BookingSubject) String() string {

	return s.ticker + " $"
}
