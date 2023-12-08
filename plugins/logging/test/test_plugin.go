package test

import (
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/pkg/test"
	"github.com/open-panoptes/opni/plugins/logging/pkg/agent"
	"github.com/open-panoptes/opni/plugins/logging/pkg/gateway"
)

func init() {
	test.EnablePlugin(meta.ModeGateway, gateway.Scheme)
	test.EnablePlugin(meta.ModeAgent, agent.Scheme)
}
