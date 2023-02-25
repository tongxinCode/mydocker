package main

import (
	"fmt"
	"mydocker/cgroups/subsystems"
	"mydocker/container"
	"mydocker/network"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		&cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		&cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		&cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		&cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		&cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		&cli.StringFlag{
			Name:  "net",
			Usage: "container network",
		},
		&cli.StringSliceFlag{
			Name:  "p",
			Usage: "port mapping",
		},
	},
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args().Slice() {
			cmdArray = append(cmdArray, arg)
		}
		tty := context.Bool("ti")
		detach := context.Bool("d")
		if tty && detach {
			return fmt.Errorf("ti and d paramter can not both provided")
		}
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		volume := context.String("v")
		network := context.String("net")
		portmapping := context.StringSlice("p")
		log.Infof("createTty %v", tty)
		containerName := context.String("name")
		Run(tty, cmdArray, resConf, volume, containerName, network, portmapping)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess()
		return err
	},
}

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "commit a container into image",
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container name")
		}
		imageName := context.Args().Get(0)
		//commitContainer(containerName)
		commitContainer(imageName)
		return nil
	},
}

var listCommand = cli.Command{
	Name:  "ps",
	Usage: "list all the containers",
	Action: func(context *cli.Context) error {
		ListContainers()
		return nil
	},
}

var logCommand = cli.Command{
	Name:  "logs",
	Usage: "print logs of a container",
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("please input your container name")
		}
		containerName := context.Args().Get(0)
		logContainer(containerName)
		return nil
	},
}

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "exec a command into container",
	Action: func(context *cli.Context) error {
		//This is for callback
		if os.Getenv(ENV_EXEC_PID) != "" {
			log.Infof("pid callback pid %s", os.Getgid())
			return nil
		}

		if context.Args().Len() < 2 {
			return fmt.Errorf("missing container name or command")
		}
		containerName := context.Args().Get(0)
		var commandArray []string
		for _, arg := range context.Args().Tail() {
			commandArray = append(commandArray, arg)
		}
		ExecContainer(containerName, commandArray)
		return nil
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "stop a container",
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container name")
		}
		containerName := context.Args().Get(0)
		stopContainer(containerName)
		return nil
	},
}

var removeCommand = cli.Command{
	Name:  "rm",
	Usage: "remove unused containers",
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container name")
		}
		containerName := context.Args().Get(0)
		removeContainer(containerName)
		return nil
	},
}

var networkCommand = cli.Command{
	Name:  "network",
	Usage: "container network commands",
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "create a container network",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "driver",
					Usage: "network driver",
				},
				&cli.StringFlag{
					Name:  "subnet",
					Usage: "subnet cidr",
				},
			},
			Action: func(context *cli.Context) error {
				if context.Args().Len() < 1 {
					return fmt.Errorf("Missing network name")
				}
				network.Init()
				err := network.CreateNetwork(context.String("driver"), context.String("subnet"), context.Args().Get(0))
				if err != nil {
					return fmt.Errorf("create network error: %+v", err)
				}
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list container network",
			Action: func(context *cli.Context) error {
				network.Init()
				network.ListNetwork()
				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove container network",
			Action: func(context *cli.Context) error {
				if context.Args().Len() < 1 {
					return fmt.Errorf("Missing network name")
				}
				network.Init()
				err := network.DeleteNetwork(context.Args().Get(0))
				if err != nil {
					return fmt.Errorf("remove network error: %+v", err)
				}
				return nil
			},
		},
	},
}
