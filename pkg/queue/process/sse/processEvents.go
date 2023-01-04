package processEvents

import (
	"bytes"
	"encoding/json"
	"example/bookAPI/internal/models/event"
	"example/bookAPI/pkg/queue/process"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
	"time"
)

func Listen(c *fiber.Ctx) error {
	return process.Process(callback)
}

func callback(q *process.Queue) {
	var newEvent event.Event
	query := q.DB.First(&newEvent)
	if query.Error == nil {
		if query.RowsAffected > 0 {
			fmt.Println("processing event: " + newEvent.Name)
			q.DB.Delete(&newEvent)
		}
	}
}

type Client struct {
	name   string
	events chan *DashBoard
}
type DashBoard struct {
	User uint
}

func Handler(f http.HandlerFunc) http.Handler {
	return f
}
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	client := &Client{name: r.RemoteAddr, events: make(chan *DashBoard, 10)}
	time.Sleep(5 * time.Second)
	go updateDashboard(client)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(5 * time.Second)
	select {
	case ev := <-client.events:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(ev)
		fmt.Fprintf(w, "data: %v\n\n", buf.String())
		fmt.Printf("data: %v\n", buf.String())
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func updateDashboard(client *Client) {
	for {
		db := &DashBoard{
			User: uint(rand.Uint32()),
		}
		client.events <- db
	}
}
