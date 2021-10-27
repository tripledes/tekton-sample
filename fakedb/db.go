package fakedb

import "github.com/tripledes/web-quotes/pkg/types"

type FakeDB struct{}

func (f FakeDB) Close() error {
	return nil
}

func (f FakeDB) FindAll() ([]types.Quote, error) {
	qs := []types.Quote{
		{
			Tags:   []string{"test", "your", "code"},
			Author: "Jane Doe",
			Text:   "The test shall pass!",
		},
		{
			Tags:   []string{"life", "science", "wisdom"},
			Author: "John Doe",
			Text:   "Did they?",
		},
	}
	return qs, nil
}

func (f FakeDB) FindOne() (types.Quote, error) {
	q := types.Quote{
		Tags:   []string{"test", "your", "code"},
		Author: "Jane Doe",
		Text:   "The test shall pass!",
	}
	return q, nil
}
