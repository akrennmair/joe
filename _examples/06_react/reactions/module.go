package reactions

import (
	"github.com/go-joe/joe"
	"go.uber.org/zap"
)

const (
	ThumbsUp = "ThumbsUp"
)

type Module struct {
	adapter ReactionAwareAdapter
	logger  *zap.Logger
}

type ReactionAwareAdapter interface {
	React(reaction string, msg joe.Message) error
}

func New() *Module {
	return &Module{logger: zap.NewNop()}
}

func (m *Module) Module() joe.Module {
	return joe.ModuleFunc(func(joeConf *joe.Config) error {
		joeConf.RegisterHandler(m.Init)
		m.logger = joeConf.Logger("reactions")
		return nil
	})
}

func (m *Module) Init(evt joe.InitEvent) {
	m.adapter, _ = evt.Adapter().(ReactionAwareAdapter)
}

func (m *Module) Add(reaction string, msg joe.Message) error {
	if m.adapter == nil {
		m.logger.Info("Adapter does not support reactions")
		return nil
	}

	return m.adapter.React(reaction, msg)
}
