package main

import (
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
)

func main() {
	plugin.Serve(Contract)
}

type Cryptonauts struct {
}

func (e *Cryptonauts) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "Cryptonauts",
		Version: "0.0.1",
	}, nil
}

func (e *Cryptonauts) Init(ctx contract.Context, req *plugin.Request) error {
	return nil
}

var Contract plugin.Contract = contract.MakePluginContract(&Cryptonauts{})
