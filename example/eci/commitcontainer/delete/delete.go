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
	req := &commitcontainer.DeleteCommitContainerTaskRequest{
		// Fill in the request parameters
		TaskId: "taskId-xxx",
	}
	resp, raw, err := client.CommitContainerTask().DeleteCommitContainerTask(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
