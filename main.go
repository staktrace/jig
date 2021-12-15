package main

import (
	"foxygo.at/pony/serve"
	"github.com/alecthomas/kong"
)

var version = "v0.0.0"

type config struct {
	Version kong.VersionFlag `short:"V" help:"Print version information" group:"Other"`
	Serve   cmdServe         `cmd:"" help:"Serve GRPC services"`
}

type cmdServe struct {
	ProtoSet  string `short:"p" help:"Protoset .pb file containing service and deps" required:""`
	MethodDir string `short:"m" default:"." help:"Directory containing method definitions"`
	Listen    string `short:"l" default:"localhost:8080" help:"TCP listen address"`
}

func main() {
	cli := &config{}
	kctx := kong.Parse(cli, kong.Vars{"version": version})
	err := kctx.Run()
	kctx.FatalIfErrorf(err)
}

func (cs *cmdServe) Run() error {
	s := serve.Server{
		Listen:    cs.Listen,
		MethodDir: cs.MethodDir,
		ProtoSet:  cs.ProtoSet,
	}
	return s.Run()
}