package main

import (
	"context"
	"crypto/tls"
	"fmt"

	cli "github.com/state-cloud/client-go/pkg/client"
	"github.com/state-cloud/client-go/pkg/openapi/config"
	"github.com/state-cloud/statecloud-sdk-go/service/eci"
	containergroup "github.com/state-cloud/statecloud-sdk-go/service/eci/types/containergroup"
)

func main() {
	baseDomain := "https://eci-global.ctapi-test.ctyun.cn:21443"
	config := &config.OpenapiConfig{
		AccessKey: "b1accda96cb74be390009d2144923466",
		SecretKey: "*",
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
		eci.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &containergroup.ExecCommandRequest{
		// Fill in the request parameters
		ContainerGroupId: "eci-xxx",
		ContainerName:    "container-x",
		Command:          []string{"ls", "/"},
		TTY:              false,
		Stdin:            false,
		Sync:             true,
	}
	resp, raw, err := client.ContainerGroup().ExecCommand(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
