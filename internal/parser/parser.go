package parser

import (
	"fmt"
	"strings"
)

type GedisCommand string

const (
	// string commands
	GET    GedisCommand = "GET"
	SET    GedisCommand = "SET"
	DELETE GedisCommand = "DELETE"

	// list commands
	LPUSH GedisCommand = "LPUSH"
	RPUSH GedisCommand = "RPUSH"
	LPOP  GedisCommand = "LPOP"
	RPOP  GedisCommand = "RPOP"
	LLEN  GedisCommand = "LLEN"
	LMOVE GedisCommand = "LMOVE"
	LTRIM GedisCommand = "LTRIM"
)

type GedisDataType string

const (
	STRING GedisDataType = "string"
	LIST   GedisDataType = "list"
)

type GedisParser struct {
	DataType GedisDataType
	Command  GedisCommand
	Args     []string
	rawData  string
}

func NewGedisParser(data string) *GedisParser {
	return &GedisParser{
		rawData: data,
	}
}

func (p *GedisParser) Parse() error {
	words := strings.Fields(p.rawData)
	if len(words) == 0 {
		return fmt.Errorf("empty command")
	}
	p.Command = GedisCommand(words[0])
	p.Args = words[1:]

	switch p.Command {
	case GET, SET, DELETE:
		p.DataType = STRING
	case LPUSH, RPUSH, LPOP, RPOP, LLEN, LMOVE, LTRIM:
		p.DataType = LIST
	default:
		return fmt.Errorf("unknown command: %s", p.Command)
	}

	return nil
}
