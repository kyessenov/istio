package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/grpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"istio.io/istio/pilot/proxy/envoy/v2"
)

func main() {
	flag.Parse()
	stop := make(chan struct{})

	if id == "" {
		if name, exists := os.LookupEnv("POD_NAME"); exists {
			id = name
		}
		if namespace, exists := os.LookupEnv("POD_NAMESPACE"); exists {
			id = namespace + "/" + id
		}
	}
	if domain == "" {
		if namespace, exists := os.LookupEnv("POD_NAMESPACE"); exists {
			domain = namespace + ".svc.cluster.local"
		}
	}

	config := cache.NewSimpleCache(v2.Hasher{}, nil /* TODO */)
	server := xds.NewServer(config)
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	server.Register(grpcServer)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			glog.Error(err)
		}
	}()

	// run envoy process
	envoy := exec.Command(binPath,
		"-c", "bootstrap.json",
		"-l", debug,
		"--drain-time-s", "1")
	envoy.Stdout = os.Stdout
	envoy.Stderr = os.Stderr
	envoy.Start()

	var generator *v2.Generator
	if validate {
		generator = v2.NewGenerator(config, map[string]string{
			"services.json":  readFile("testdata/services.json"),
			"instances.json": readFile("testdata/instances.json"),
			"context.json":   readFile("testdata/context.json"),
		})
	} else {
		generator, err = v2.NewKubeGenerator(config, id, domain, kubeconfig)
		if err != nil {
			glog.Fatal(err)
		}
	}
	// expose profiling endpoint
	go http.ListenAndServe(":15005", nil)

	go generator.Run(stop)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	close(stop)
	glog.Flush()
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(bytes)
}

var (
	kubeconfig string
	id         string
	domain     string
	binPath    string
	debug      string
	port       int

	validate bool
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "",
		"Use a Kubernetes configuration file instead of in-cluster configuration")
	flag.StringVar(&id, "id", "", "Workload ID (e.g. pod namespace/name)")
	flag.StringVar(&domain, "domain", "", "Workload domain (e.g. default.svc.cluster.local)")
	flag.IntVar(&port, "port", 15003,
		"ADS port")
	flag.BoolVar(&validate, "valid", false,
		"Validate only (for testing and debugging)")
	flag.StringVar(&binPath, "bin", "/usr/local/bin/envoy", "Envoy binary")
	flag.StringVar(&debug, "l", "info", "Envoy debug level")
}
