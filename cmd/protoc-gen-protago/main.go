package main

import (
	"github.com/timmonfette1/protago/internal/plugin"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	features := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL | pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
	minimumEdition := descriptorpb.Edition_EDITION_2023
	maximumEdition := descriptorpb.Edition_EDITION_2024

	opts := []pgs.InitOption{
		pgs.DebugEnv("PROTAGO_DEBUG"),
		pgs.SupportedFeatures(&features),
		pgs.MinimumEdition(&minimumEdition),
		pgs.MaximumEdition(&maximumEdition),
	}

	pgs.Init(opts...).
		RegisterModule(plugin.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
