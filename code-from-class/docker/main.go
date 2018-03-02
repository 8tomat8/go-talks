package main

import (
	"fmt"
	"log"

	"context"

	"github.com/fsouza/go-dockerclient"
	"github.com/pkg/errors"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot create new client"))
	}
	images, err := cli.ListImages(docker.ListImagesOptions{All: true, Context: ctx, Filter: "mysql"})
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot get images list"))
	}

	for i := range images {
		//spew.Println(containers[i])
		fmt.Printf("%#v\n", images[i])
	}

	cli.RemoveContainer(docker.RemoveContainerOptions{
		Force: true,
		ID:    "my-mysql-cont",
	})
	container, err := cli.CreateContainer(docker.CreateContainerOptions{
		"my-mysql-cont",
		&docker.Config{
			Image: images[0].ID,
			Env:   []string{"MYSQL_ROOT_PASSWORD=root"},
		},
		&docker.HostConfig{},
		&docker.NetworkingConfig{},
		ctx,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Printf("%#v\n", container)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	err = cli.StartContainer(container.ID, &docker.HostConfig{
		Memory: 2 << 20,
	})
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ListContainers(docker.ListContainersOptions{
		All:     true,
		Context: ctx,
		Filters: map[string][]string{"name": {"my-mysql-cont", "fabric"}},
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot get images list"))
	}

	fmt.Printf("%#v\n", containers[0])
}
