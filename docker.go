package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	_ "github.com/docker/docker/pkg/stdcopy"
	"io"
	"log"
	"os"
	"reflect"
	"text/template"
)

func expandTemplateArgs(items []string) []string {
	m := matchSpec{
		Workloads: []string{"the-service"},
	}
	for _, item := range items {
		t := template.Must(template.New("item").Parse(item))
		err := t.Execute(os.Stdout, m)
		log.Println()
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	res := []string{}
	for _, item := range items {
		res = append(res, item)
	}
	log.Println("result: ", res)
	return res
}

func runDocker(spec dockerSpec) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, spec.Image, types.ImagePullOptions{ /*RegistryAuth: authStr*/ })
	if err != nil {
		log.Println("image pull failed", err)
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)

	log.Println("command: ", reflect.TypeOf(spec.Command), len(spec.Command))
	config := container.Config{
		Image:      spec.Image,
		Cmd:        expandTemplateArgs(spec.Args),
		Entrypoint: spec.Command,
	}
	log.Println("Entrypoint:  ", config.Entrypoint)
	hostConfig := container.HostConfig{}
	log.Println("attempting to create container", spec.Image)
	create, err := cli.ContainerCreate(ctx, &config, &hostConfig, nil, "")
	if err != nil {
		log.Println("container create error", err)
		panic(err)
	}

	log.Println("attempting to start", create.ID, spec.Image)
	if err = cli.ContainerStart(ctx, create.ID, types.ContainerStartOptions{}); err != nil {
		log.Println("container start panic", err)
		panic(err)
	}
	log.Println("started container: ", create.ID)

	statusCh, errCh := cli.ContainerWait(ctx, create.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	log.Println("finished running container")

	out, err = cli.ContainerLogs(ctx, create.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)

	return nil
}
