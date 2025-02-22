// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"
	"runtime/debug"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/isometry/terraform-provider-faws/internal/provider"
	"github.com/isometry/terraform-provider-faws/version"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	flag.Parse()

	logFlags := log.Flags()
	logFlags = logFlags &^ (log.Ldate | log.Ltime)
	log.SetFlags(logFlags)

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		log.Printf("Starting %s@%s (%s)...", buildInfo.Main.Path, version.ProviderVersion, buildInfo.GoVersion)
	}

	serverFactory, _, err := provider.ProtoV5ProviderServerFactory(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf5server.ServeOpt

	if *debugFlag {
		serveOpts = append(serveOpts, tf5server.WithManagedDebug())
	}

	err = tf5server.Serve(
		"registry.terraform.io/isometry/aws",
		serverFactory,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
