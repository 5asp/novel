package collyfx

import (
	"github.com/gocolly/colly/v2"
	"go.uber.org/fx"
	// other imports...
)

// ProvideColly is a function that returns a colly collector instance
func ProvideColly() (*colly.Collector, error) {
	return nil, nil
}

// CollyModule provided to fx
var CollyModule = fx.Options(
	fx.Provide(ProvideColly),
)
