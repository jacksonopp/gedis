package handlers

import (
	"fmt"

	"github.com/jacksonopp/gedis/internal/parser"
)

func (h *GedisHandler) HandleString(p *parser.GedisParser) (string, error) {
	switch p.Command {
	case parser.GET:
		return h.get(p.Args[0])
	case parser.SET:
		err := h.set(p.Args[0], p.Args[1])
		if err != nil {
			return "", err
		}
		return "OK", nil
	case parser.DELETE:
		err := h.delete(p.Args[0])
		if err != nil {
			return "", err
		}
		return "OK", nil
	default:
		return "", fmt.Errorf("error: command %s not found", p.Command)
	}
}

func (h *GedisHandler) get(key string) (string, error) {
	data := h.data[key]
	if data == "" {
		return "", fmt.Errorf("key \"%s\" not found", key)
	}
	return data, nil
}

func (h *GedisHandler) set(key, value string) error {
	h.data[key] = value
	return nil
}

func (h *GedisHandler) delete(key string) error {
	if _, ok := h.data[key]; !ok {
		return fmt.Errorf("key \"%s\" not found", key)
	}
	delete(h.data, key)
	return nil
}
