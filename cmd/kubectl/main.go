package main

import (
	"os"

	"github.com/magicsong/s2iapiserver/pkg/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-ks", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewKSCmd(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
