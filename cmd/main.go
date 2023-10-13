package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/koenno/termos-negros/client"
	"github.com/koenno/termos-negros/client/portal"
	"github.com/koenno/termos-negros/parser"
	"golang.org/x/time/rate"
)

func main() {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second), 1)
	reqSender := client.New(rateLimiter)
	req, err := portal.NewRequestFactory().NewRequest(context.Background())
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	bb, headers, err := reqSender.Send(req)
	if err != nil {
		log.Fatalf("failed to send a request: %v", err)
	}

	menuParser := parser.NewMenuParser()
	respParser := portal.NewBodyParser(menuParser)
	menu, err := respParser.Parse(bb, headers)
	if err != nil {
		log.Fatalf("failed to parse menu: %v", err)
	}

	now := time.Now()
	for _, m := range menu {
		if m.Date.Day() == now.Day() && m.Date.Month() == now.Month() {
			fmt.Println(m)
		}
	}
}
