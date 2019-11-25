package worker

import (
	"fmt"
	"log"
	"spider/engine"
	"spider/zhenai/parser"
	"spider_dist/config"
)

// {"ParseCityList", nil}
// {"ParseCity", nil}
// {"ProfileParser", userName}
type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	URL    string
	Parser *SerializedParser
}

type ParseResult struct {
	Requests []*Request
	Items    []*engine.Item
}

func SerializeRequest(r *engine.Request) *Request {
	name, args := r.Parser.Serialize()
	return &Request{
		URL: r.URL,
		Parser: &SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r *engine.ParseResult) *ParseResult {
	result := &ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r *Request) (*engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return nil, err
	}
	return &engine.Request{
		URL:    r.URL,
		Parser: parser,
	}, nil
}

func DeserializeResult(r *ParseResult) *engine.ParseResult {
	result := &engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineR, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineR)
	}

	return result
}

func deserializeParser(p *SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ProfileParser:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %+v", p.Args)
		}
	case config.NilParser:
		return &engine.NilParser{}, nil
	default:
		return nil, fmt.Errorf("unknown parser name")
	}
}
