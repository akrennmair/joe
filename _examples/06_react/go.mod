module github.com/go-joe/joe/examples/react

go 1.12

require (
	github.com/go-joe/joe v0.7.0
	github.com/go-joe/slack-adapter v0.0.0
	go.uber.org/zap v1.9.1
)

replace github.com/go-joe/joe => ../..

replace github.com/go-joe/slack-adapter => ../../../slack-adapter
