package processEvents

import (
	"bufio"
	"example/bookAPI/internal/models/event"
	"example/bookAPI/pkg/queue/process"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Listen(c *fiber.Ctx) error {
	ctx := c.Context()

	ctx.SetContentType("text/event-stream")
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.Response.Header.Set("Transfer-Encoding", "chunked")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Cache-Control")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		process.Process(func(q *process.Queue) { callback(q, w) })
	}))
	return nil
}

func callback(q *process.Queue, w *bufio.Writer) {
	var newEvent event.Event
	query := q.DB.First(&newEvent)
	if query.Error == nil {
		if query.RowsAffected > 0 {
			fmt.Println("processing event: " + newEvent.Name)
			fmt.Fprintf(w, "data: %v\n\n", newEvent.Name)
			err := w.Flush()
			if err != nil {
				fmt.Println("ERROR " + err.Error())
			}
			q.DB.Delete(&newEvent)
		}
	}
}
