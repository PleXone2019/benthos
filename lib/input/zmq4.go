// +build ZMQ4

package input

import (
	"github.com/Jeffail/benthos/v3/lib/input/reader"
	"github.com/Jeffail/benthos/v3/lib/log"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/types"
	"github.com/Jeffail/benthos/v3/lib/x/docs"
)

//------------------------------------------------------------------------------

func init() {
	Constructors[TypeZMQ4] = TypeSpec{
		constructor: NewZMQ4,
		Summary: `
Consumes messages from a ZeroMQ socket.`,
		Description: `
ZMQ4 is supported but currently depends on C bindings. Since this is an
annoyance when building or using Benthos it is not compiled by default.

There is a specific docker tag postfix ` + "`-cgo`" + ` for C builds containing
ZMQ support.

You can also build it into your project by getting libzmq installed on your
machine, then build with the tag:

` + "```sh" + `
go install -tags "ZMQ4" github.com/Jeffail/benthos/v3/cmd/...
` + "```" + `

ZMQ4 input supports PULL and SUB sockets only. If there is demand for other
socket types then they can be added easily.`,
		FieldSpecs: docs.FieldSpecs{
			docs.FieldCommon("urls", "A list of URLs to connect to. If an item of the list contains commas it will be expanded into multiple URLs."),
			docs.FieldCommon("bind", "Whether to bind to the specified URLs or connect."),
			docs.FieldCommon("socket_type", "The socket type to connect as.").HasOptions("PULL", "SUB"),
			docs.FieldCommon("sub_filters", "A list of filters to use when connecting as a `SUB` socket."),
			docs.FieldAdvanced("high_water_mark", "The message high water mark to use."),
			docs.FieldAdvanced("poll_timeout", "The poll timeout to use."),
		},
	}
}

//------------------------------------------------------------------------------

// NewZMQ4 creates a new ZMQ input type.
func NewZMQ4(conf Config, mgr types.Manager, log log.Modular, stats metrics.Type) (Type, error) {
	z, err := reader.NewZMQ4(conf.ZMQ4, log, stats)
	if err != nil {
		return nil, err
	}
	return NewAsyncReader(TypeZMQ4, true, reader.NewAsyncPreserver(z), log, stats)
}

//------------------------------------------------------------------------------
