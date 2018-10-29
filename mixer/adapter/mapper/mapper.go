package mapper

import "context"

type MyAdapter struct{}

func (MyAdapter) HandleMapper(_ context.Context, req *HandleMapperRequest) (*OutputMsg, error) {
	lookup := map[string]string{
		"hello": "world",
	}
	return &OutputMsg{Value: lookup[req.Instance.Key]}, nil
}
