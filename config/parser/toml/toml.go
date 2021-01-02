package toml

import (
	"github.com/pelletier/go-toml"

	"github.com/go-kratos/kratos/v2/config/parser"
)

type tomlParser struct{}

// NewParser new a toml parser.
func NewParser() parser.Parser {
	return &tomlParser{}
}

func (p *tomlParser) Format() string {
	return "toml"
}

func (p *tomlParser) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (p *tomlParser) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}
