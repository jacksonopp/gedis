package handlers

import (
	"fmt"

	"github.com/jacksonopp/gedis/internal/parser"
)

type GedisHandler struct {
	// Add fields here
	data map[string]string
}

func NewGedisHandler() *GedisHandler {
	// Initialize and return a new instance of GedisHandler
	return &GedisHandler{
		data: make(map[string]string),
	}
}

func (h *GedisHandler) HandleCommand(input string) (string, error) {
	p := parser.NewGedisParser(input)
	err := p.Parse()
	if err != nil {
		return "", err
	}

	switch p.DataType {
	case parser.STRING:
		return h.HandleString(p)
	case parser.LIST:
		return h.HandleList(p)
	default:
		return "", fmt.Errorf("unknown data type: %s", p.DataType)
	}
}
