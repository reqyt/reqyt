package container_manager

import (
	"github.com/docker/docker/client"
	"sync"
)

type ImageConfig struct {
	Name    string
	Secrets map[string]string
}
type Config []ImageConfig
func (c Config) ForImage(s string) *ImageConfig {
	for _, v := range c {
		if v.Name == s {
			return &v
		}
	}
	return nil
}

type ContainerManager struct {
	mutex      sync.Mutex
	Containers map[string]*Container
	Client     *client.Client
	config     Config
}

func NewContainerManager(config Config) (*ContainerManager, error) {
	cm := &ContainerManager{}
	cm.config = config
	cm.Containers = make(map[string]*Container)
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	cm.Client = cli
	return cm, nil
}

func (cm *ContainerManager) GetContainer(image string) *Container {
	cm.mutex.Lock()
	cnt, ok := cm.Containers[image]
	if !ok {
		cnt = NewContainer(image, cm.config.ForImage(image).Secrets)
		cm.Containers[image] = cnt
	}
	cm.mutex.Unlock()
	cnt.Start(cm.Client)
	return cnt
}

func (cm *ContainerManager) Close() {
	wg := sync.WaitGroup{}
	for _, cnt := range cm.Containers {
		wg.Add(1)
		go func() {
			cnt.Stop(cm.Client)
			<- cnt.Stopped
			wg.Done()
		}()
	}
	wg.Wait()
}