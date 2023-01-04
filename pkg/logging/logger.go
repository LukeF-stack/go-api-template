package logger

import (
	firestore "cloud.google.com/go/firestore"
	"context"
	"fmt"
	"time"
)

type Message struct {
	User     string
	Message  string
	Datetime time.Time
}

type Logger struct {
	Store *firestore.Client
	Loggerer
}

type Loggerer interface {
	Log()
	GetStore()
}

func GetLogger(s *firestore.Client) *Logger {
	l := new(Logger)
	l.Store = s
	return l
}

func (l *Logger) Log(m, u string) {
	go func() {
		store := l.Store
		fmt.Println(store)
		if store != nil {
			_, _, err := store.Collection("logs").Add(context.Background(), Message{Message: m, User: u, Datetime: time.Now()})
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}
