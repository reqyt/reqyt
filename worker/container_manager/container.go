package container_manager

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
	"net"
	"sync"
	"time"
)

const (
	STATUS_COLD = iota
	STATUS_STARTING
	STATUS_WARM
	STATUS_STOPPING
)

func AsDockerEnv(env map[string]string) []string {
	var out []string
	for k, v := range env {
		out = append(out, k + "=" + v)
	}
	return out
}

type Container struct {
	StartTime time.Time
	Status    int
	Image     string
	ID        string
	IP        net.IP
	Started   chan struct{}
	Stopped   chan struct{}
	env       map[string]string
	mutex     sync.Mutex
}

func NewContainer(image string, env map[string]string) *Container {
	return &Container{
		Status:  STATUS_COLD,
		Image:   image,
		env:     env,
		Started: make(chan struct{}),
		Stopped: make(chan struct{}),
	}
}

func (cnt *Container) monitor(client *client.Client) {
	ctx := context.Background()
	for {
		insp, err := client.ContainerInspect(ctx, cnt.ID)
		if err != nil {
			log.Println("Error getting container status")
		}
		if insp.State.Status == "exited" {
			log.Println("Container [" + cnt.Image + "] Now cold")
			cnt.Status = STATUS_COLD
			close(cnt.Stopped)
			break
		} else if cnt.Status != STATUS_WARM && (insp.State.Health == nil || insp.State.Health.Status == types.Healthy) {
			log.Println("Container [" + cnt.Image + "] Now warm")
			cnt.Status = STATUS_WARM
			close(cnt.Started)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func (cnt *Container) Stop(client *client.Client) error {
	cnt.mutex.Lock()
	ctx := context.Background()
	if cnt.Status == STATUS_WARM {
		cnt.Status = STATUS_STOPPING
		log.Println("Container [" + cnt.Image + "] Stopping")
		client.ContainerStop(ctx, cnt.ID, nil)
	}
	cnt.mutex.Unlock()
	return nil
}

func (cnt *Container) Start(client *client.Client) error {
	cnt.mutex.Lock()
	if cnt.Status == STATUS_COLD {
		cnt.Status = STATUS_STARTING
		cnt.StartTime = time.Now()
		ctx := context.Background()
		log.Println("Container [" + cnt.Image + "] Starting")
		resp, err := client.ContainerCreate(ctx, &container.Config{
			Image: cnt.Image,
			Env: AsDockerEnv(cnt.env),
			//Cmd:   []string{"echo", "hello world"},
			//Tty:   true,
		}, nil, nil, "")
		if err != nil {
			return err
		}
		cnt.ID = resp.ID

		if err := client.ContainerStart(ctx, cnt.ID, types.ContainerStartOptions{}); err != nil {
			return err
		}

		insp, err := client.ContainerInspect(ctx, cnt.ID)
		if err != nil {
			return err
		}
		//insp.State.Health.Log
		cnt.IP = net.ParseIP(insp.NetworkSettings.DefaultNetworkSettings.IPAddress)
		go cnt.monitor(client)
	}
	cnt.mutex.Unlock()
	<-cnt.Started
	return nil
}
