package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jacksonopp/gedis/internal/parser"
)

func (h *GedisHandler) HandleList(p *parser.GedisParser) (string, error) {
	switch p.Command {
	case parser.LPUSH:
		return h.lPush(p.Args[0], p.Args[1])
	case parser.RPUSH:
		return h.rPush(p.Args[0], p.Args[1])
	case parser.LPOP:
		return h.lPop(p.Args[0])
	case parser.RPOP:
		return h.rPop(p.Args[0])
	case parser.LLEN:
		return h.lLen(p.Args[0])
	case parser.LMOVE:
		return "", nil
	case parser.LTRIM:
		return "", nil
	}
	return "", nil
}

func (h *GedisHandler) marshallList(key string, list []string) (string, error) {
	jsonData, err := json.Marshal(list)
	if err != nil {
		return "", err
	}
	h.data[key] = string(jsonData)
	return lenStr(list), nil
}

func (h *GedisHandler) lPush(key, value string) (string, error) {
	list := []string{}
	data, err := h.get(key)

	if err != nil {
		list = append([]string{value}, list...)
		return h.marshallList(key, list)
	}

	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return "", err
	}

	list = append([]string{value}, list...)
	return h.marshallList(key, list)
}

func (h *GedisHandler) rPush(key, value string) (string, error) {
	list := []string{}
	data, err := h.get(key)

	if err != nil {
		list = append(list, value)
		return h.marshallList(key, list)
	}

	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return "", err
	}
	list = append(list, value)
	return h.marshallList(key, list)
}

func (h *GedisHandler) lPop(key string) (string, error) {
	data, err := h.get(key)
	if err != nil {
		return "", err
	}

	list := []string{}
	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return "", err
	}
	if len(list) == 0 {
		return "", fmt.Errorf("list \"%s\" is empty", key)
	}

	first := list[0]
	_, err = h.marshallList(key, list[1:])
	if err != nil {
		return "", err
	}

	return first, nil
}

func (h *GedisHandler) rPop(key string) (string, error) {
	data, err := h.get(key)
	if err != nil {
		return "", err
	}

	list := []string{}
	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return "", err
	}
	if len(list) == 0 {
		return "", fmt.Errorf("list \"%s\" is empty", key)
	}

	lastIndex := len(list) - 1
	last := list[lastIndex]
	_, err = h.marshallList(key, list[:lastIndex])
	if err != nil {
		return "", err
	}

	return last, nil
}

func (h *GedisHandler) lLen(key string) (string, error) {
	data, err := h.get(key)

	if err != nil {
		return toString(0), nil
	}

	list := []string{}
	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return "", err
	}
	return lenStr(list), nil
}

func (h *GedisHandler) lMove(source, destination, from, to string) (string, error) {
	sourceListStr, err := h.get(source)
	if err != nil {
		return "", notFound(source)
	}
	sourceList := []string{}
	err = json.Unmarshal([]byte(sourceListStr), &sourceList)
	if err != nil {
		return "", err
	}

	destList := []string{}
	destListStr, err := h.get(destination)
	if err != nil {

	}
}

func performMove(source, destination []string, from, to string) (string, error) {

}
