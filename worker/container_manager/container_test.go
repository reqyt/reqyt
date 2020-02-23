package container_manager_test

import (
	" reqyt.run/worker/container_manager"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestManager(t *testing.T) {
	cm, err := container_manager.NewContainerManager()
	defer cm.Close()
	if err != nil {
		t.Fail()
		return
	}
	c := container_manager.NewContainer("echo")
	if err := c.Start(cm.Client); err != nil {
		t.Fail()
	}
	<-c.Started
	res, err := http.Post("http://"+c.IP.String()+":8000", "text/plain", strings.NewReader("testje"))
	io.Copy(os.Stdout, res.Body)
}
