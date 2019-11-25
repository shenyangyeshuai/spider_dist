package worker

import (
	"spider/engine"
)

type SpiderService struct {
}

func (s *SpiderService) Process(r *Request, result *ParseResult) error {
	engineR, err := DeserializeRequest(r)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineR)
	if err != nil {
		return err
	}

	*result = *SerializeResult(engineResult) // 尼玛, 这个坑啊!
	return nil
}
