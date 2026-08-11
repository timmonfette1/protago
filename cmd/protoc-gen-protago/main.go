package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/timmonfette1/protago/internal/plugin"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	features := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL | pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)

	pgs.Init(pgs.DebugEnv("PROTAGO_DEBUG"), pgs.SupportedFeatures(&features)).
		RegisterModule(plugin.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
