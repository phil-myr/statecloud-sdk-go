package main

import (
	"context"
	"crypto/tls"
	"fmt"

	cli "github.com/state-cloud/client-go/pkg/client"
	"github.com/state-cloud/client-go/pkg/openapi/config"
	"github.com/state-cloud/statecloud-sdk-go/service/eci"
	commitcontainer "github.com/state-cloud/statecloud-sdk-go/service/eci/types/commitcontainer"
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
	req := &commitcontainer.CreateCommitContainerTaskRequest{
		// Fill in the request parameters
		RegionId:         "b342b77ef26b11ecb0ac0242ac110002",
		TenantId:         "xxx",
		ContainerGroupId: "eci-xxx",
		ContainerName:    "container-x",
		Image: &commitcontainer.ImageInfo{
			Repository: "docker.io/library/nginx",
			Tag:        "latest",
			Message:    "xxx",
			Author:     "admin@example.org",
		},
		Registry: &commitcontainer.RegistryInfo{
			Registry: "docker.io",
			Username: "xxx",
			// base64 encoded
			Password: "cGFzc3dvcmQ=",
		},
	}
	resp, raw, err := client.CommitContainerTask().CreateCommitContainerTask(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
