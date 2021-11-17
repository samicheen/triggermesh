/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	kind := flag.String("kind", "TestTarget", "Specify the Target kind")
	cfgDir := flag.String("config", "../../",
		"Path of the directory containing the TriggerMesh deployment manifests")
	flag.Parse()
	*cfgDir = path.Clean(*cfgDir)
	temp := &component{}
	temp.LowercaseKind = strings.ToLower(*kind)
	temp.UppercaseKind = strings.ToUpper(*kind)

	// make cmd directory
	path := "cmd/" + temp.LowercaseKind + "-adapter"
	cmdPath := filepath.Join(*cfgDir, path)
	mustMkdirAll(cmdPath)

	// // make adapter directory
	path = "pkg/targets/adapter/" + temp.LowercaseKind
	adapterPath := filepath.Join(*cfgDir, path)
	mustMkdirAll(adapterPath)

	// make reconciler directory
	path = "pkg/targets/reconciler/" + temp.LowercaseKind
	reconcilerPath := filepath.Join(*cfgDir, path)
	mustMkdirAll(reconcilerPath)

	// populate cmd directory
	// read main.go and replace the template variables
	path = *cfgDir + "/cmd/" + temp.LowercaseKind + "-adapter/main.go"
	err := temp.replaceTemplates("scaffolding/cmd/newtarget-adapter/main.go", path)
	if err != nil {
		log.Fatal("failed creating the cmd templates")
		log.Fatal(err)
		return
	}

	// populate adapter directory
	// read adapter.go
	path = *cfgDir + "/pkg/targets/adapter/" + temp.LowercaseKind + "/adapter.go"
	err = temp.replaceTemplates("scaffolding/pkg/targets/adapter/newtarget/adapter.go", path)
	if err != nil {
		log.Fatal("failed creating the adapter templates")
		log.Fatal(err)
		return
	}

	// read newtarget_lifecycle.go
	path = *cfgDir + "/pkg/apis/targets/v1alpha1/" + temp.LowercaseKind + "_lifecycle.go"
	err = temp.replaceTemplates("scaffolding/pkg/apis/targets/v1alpha1/newtarget_lifecycle.go", path)
	if err != nil {
		log.Fatal("failed creating the newtarget_lifecycle.go template")
		log.Fatal(err)
		return
	}

	// read newtarget_types.go
	path = *cfgDir + "/pkg/apis/targets/v1alpha1/" + temp.LowercaseKind + "_types.go"
	err = temp.replaceTemplates("scaffolding/pkg/apis/targets/v1alpha1/newtarget_types.go", path)
	if err != nil {
		log.Fatal("failed creating the newtarget_types.go template")
		log.Fatal(err)
		return
	}

	// populate reconciler directory
	// read adapter.go
	path = *cfgDir + "/pkg/targets/reconciler/" + temp.LowercaseKind + "/adapter.go"
	err = temp.replaceTemplates("scaffolding/pkg/targets/reconciler/newtarget/adapter.go", path)
	if err != nil {
		log.Fatal("failed creating the reconciler templates")
		log.Fatal(err)
		return
	}

	// read controller_test.go
	path = *cfgDir + "/pkg/targets/reconciler/" + temp.LowercaseKind + "/controller_test.go"
	err = temp.replaceTemplates("scaffolding/pkg/targets/reconciler/newtarget/controller_test.go", path)
	if err != nil {
		log.Fatal("failed creating the controller_test.go template")
		log.Fatal(err)
		return
	}

	// read controller.go
	path = *cfgDir + "/pkg/targets/reconciler/" + temp.LowercaseKind + "/controller.go"
	err = temp.replaceTemplates("scaffolding/pkg/targets/reconciler/newtarget/controller.go", path)
	if err != nil {
		log.Fatal("failed creating the controller.go template")
		log.Fatal(err)
		return
	}

	// read reconciler.go
	path = *cfgDir + "/pkg/targets/reconciler/" + temp.LowercaseKind + "/reconciler.go"
	err = temp.replaceTemplates("scaffolding/pkg/targets/reconciler/newtarget/reconciler.go", path)
	if err != nil {
		log.Fatal("failed creating the reconciler.go template")
		log.Fatal(err)
		return
	}

	// populate the config directory
	// read 301-newtarget.yaml.go
	path = *cfgDir + "/config/301-" + temp.LowercaseKind + ".yaml"
	temp.replaceTemplates("scaffolding/config/301-newtarget.yaml", path)

	fmt.Println("done")
	fmt.Println("Next Steps:")
	fmt.Println("Update `cmd/triggermesh-controller/main.go`")
	fmt.Println("Update `config/500-controller.yaml`")
	fmt.Println("Update `pkg/api/targets/v1alpha1/register.go`")
	fmt.Printf("Create kodata symlinks in cmd/%s", temp.LowercaseKind)
	fmt.Println("")
	fmt.Println("Run `make codegen`")
}

type component struct {
	Kind          string
	LowercaseKind string
	UppercaseKind string
}

func (a *component) replaceTemplates(filename, outputname string) error {
	tmp, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(outputname)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmp.Execute(file, a)
}

func mustMkdirAll(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("failed creating directory: %v", err)
	}
}