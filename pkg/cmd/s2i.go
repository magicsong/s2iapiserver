// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/scheme"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// s2iCmd represents the s2i command
var s2iCmd = &cobra.Command{
	Use:   "rerun",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: s2irerun,
}

func init() {
	rootCmd.AddCommand(s2iCmd)
	flag.Set("v", "7")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s2iCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s2iCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func s2irerun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Must specify the name of s2irun")
		os.Exit(1)
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("[Load Config]" + err.Error())
		os.Exit(1)
	}

	config.GroupVersion = &v1alpha1.ApiVersion.GroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	client, err := rest.RESTClientFor(config)
	if err != nil {
		fmt.Println("[Build Rest Client]" + err.Error())
		os.Exit(1)
	}
	rerun := &v1alpha1.Rerun{}
	rerun.SetName(args[0])
	err = client.Post().Namespace("default").Resource("s2iruns").Name(args[0]).SubResource("rerun").Body(rerun).Do().Error()
	if err != nil {
		fmt.Println("[ Post Rerun ]" + err.Error())
		os.Exit(1)
	}
	fmt.Printf("try to rerun %s\n", args[0])
}
