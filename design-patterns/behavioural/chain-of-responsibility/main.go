package main

import (
	"fmt"
	"strings"
)

// Usage:
// GUI, Web Middleware, Logging System, Approval Process

// Pros:
// Add / remove handlers dynamically
// Seperate of concerns
// Handlers dont need to know about each other

// Cons:
// Testing, Debugging through chain is hard
// Maintaining chain can be complex in large system

func main() {
	query := "billing"
	handler := &GeneralHandler{}
	handler.SetNext(&TechnicalHandler{})
	handler.SetNext(&BillingHandler{})

	handler.Handle(query)
}

type Handler interface {
	Handle(query string)
	SetNext(handler Handler)
}

type QueryHandler struct {
	next Handler
}

func (h *QueryHandler) SetNext(handler Handler) {
	h.next = handler
}

type GeneralHandler struct {
	QueryHandler
}

func (g *GeneralHandler) Handle(query string) {
	if strings.Contains(query, "general") {
		fmt.Println("Handling general query")
	}
	if g.next != nil {
		g.next.Handle(query)
	}
}

type TechnicalHandler struct {
	QueryHandler
}

func (t *TechnicalHandler) Handle(query string) {
	if strings.Contains(query, "technical") {
		fmt.Println("Handling technical query")
	}
	if t.next != nil {
		t.next.Handle(query)
	}
}

type BillingHandler struct {
	QueryHandler
}

func (b *BillingHandler) Handle(query string) {
	if strings.Contains(query, "billing") {
		fmt.Println("Handling billing query")
	}
	if b.next != nil {
		b.next.Handle(query)
	}
}
