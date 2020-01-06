// 参考: https://blog.51cto.com/xingej/2115258

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

type KafkaCluster struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml: "kind"`
	Metadata   Metadata `yaml: "metadata"`
	Spec       Spec     `yaml: "spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
	//map类型
	Labels map[string]*NodeServer `yaml:"labels"`
}

type NodeServer struct {
	Address  string `yaml: "address"`
	Id       string `yaml: "id"`
	Name     string `yaml: "name"`
	NodeName string `yaml:"nodeName"`
	Role     string `yaml: "role"`
}

type Spec struct {
	Replicas int    `yaml: "replicas"`
	Name     string `yaml: "name"`
	Image    string `yaml: "iamge"`
	Ports    int    `yaml: "ports"`
	//slice类型
	Conditions []Conditions `yaml: "conditions"`
}

type Conditions struct {
	ContainerPort string   `yaml:"containerPort"`
	Requests      Requests `yaml: "requests"`
	Limits        Limits   `yaml: "limits"`
}

type Requests struct {
	CPU    string `yaml: "cpu"`
	MEMORY string `yaml: "memory"`
}

type Limits struct {
	CPU    string `yaml: "cpu"`
	MEMORY string `yaml: "memory"`
}

func Test_yaml2struct(t *testing.T) {
	conf := new(KafkaCluster)
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.UnmarshalStrict(yamlFile, conf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("--- conf:%+v\n", conf)
}

func Test_struct2yaml(t *testing.T) {
	yfile := "./temp_aaa.yaml"
	cond := &Spec{
		Replicas: 123,
		Name:     "hello",
		Image:    "world",
		Ports:    456,
		Conditions: []Conditions{
			{
				ContainerPort: "6307",
				Requests: Requests{
					CPU:    "cpu01",
					MEMORY: "money01",
				},
				Limits: Limits{
					CPU:    "cpu02",
					MEMORY: "money02",
				},
			},
			{
				ContainerPort: "6308",
				Requests: Requests{
					CPU:    "cpu03",
					MEMORY: "money03",
				},
				Limits: Limits{
					CPU:    "cpu04",
					MEMORY: "money04",
				},
			},
		},
	}

	bytes, err := yaml.Marshal(cond)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(yfile, bytes, os.ModePerm)
}
