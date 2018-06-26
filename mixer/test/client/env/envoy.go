// Copyright 2017 Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package env

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"istio.io/istio/tests/util"
)

// Envoy stores data for Envoy process
type Envoy struct {
	cmd   *exec.Cmd
	ports *Ports
}

// NewEnvoy creates a new Envoy struct and starts envoy.
func (s *TestSetup) NewEnvoy(stress bool, filtersBeforeMixer string, mfConf *MixerFilterConf, ports *Ports, epoch int,
	disableHotRestart bool) (*Envoy, error) {
	confPath := filepath.Join(util.IstioOut, fmt.Sprintf("config.conf.%v.yaml", ports.AdminPort))
	log.Printf("Envoy config: in %v\n", confPath)
	if err := s.CreateEnvoyConf(confPath, stress, filtersBeforeMixer, mfConf, ports); err != nil {
		return nil, err
	}

	debugLevel := os.Getenv("ENVOY_DEBUG")
	if len(debugLevel) == 0 {
		debugLevel = "info"
	}

	// Don't use hot-start, each Envoy re-start use different base-id
	args := []string{"-c", confPath,
		"--v2-config-only",
		"--base-id", strconv.Itoa(int(ports.AdminPort) + epoch)}
	if stress {
		args = append(args, "--concurrency", "10")
	} else {
		// debug is far too verbose.
		args = append(args, "-l", debugLevel, "--concurrency", "1")
	}
	if disableHotRestart {
		args = append(args, "--disable-hot-restart")
	}
	if s.EnvoyParams != nil {
		args = append(args, s.EnvoyParams...)
	}
	/* #nosec */
	envoyPath := filepath.Join(util.IstioBin, "envoy")
	if path, exists := os.LookupEnv("ENVOY_PATH"); exists {
		envoyPath = path
	}
	cmd := exec.Command(envoyPath, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return &Envoy{
		cmd:   cmd,
		ports: ports,
	}, nil
}

// Start starts the envoy process
func (s *Envoy) Start() error {
	err := s.cmd.Start()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://localhost:%v/server_info", s.ports.AdminPort)
	WaitForHTTPServer(url)

	return nil
}

// Stop stops the envoy process
func (s *Envoy) Stop() error {
	log.Printf("stop envoy ...\n")
	_, _, _ = HTTPPost(fmt.Sprintf("http://127.0.0.1:%v/quitquitquit", s.ports.AdminPort), "", "")
	done := make(chan error, 1)
	go func() {
		done <- s.cmd.Wait()
	}()

	select {
	case <-time.After(3 * time.Second):
		log.Println("envoy killed as timeout reached")
		if err := s.cmd.Process.Kill(); err != nil {
			return err
		}
	case err := <-done:
		log.Printf("stop envoy ... done\n")
		return err
	}

	return nil
}
