package tfjob

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"k8s.io/api/core/v1"

	"github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	"gopkg.in/yaml.v2"
)

const ContainerName = "tensorflow"

type Config struct {
	key   string
	TFJob *v1alpha2.TFJob
}

func NewConfig(key string) *Config {
	config := new(Config)
	config.key = key
	config.TFJob = new(v1alpha2.TFJob)
	config.TFJob.Kind = v1alpha2.Kind
	config.TFJob.APIVersion = filepath.Join(v1alpha2.GroupName, v1alpha2.GroupVersion)

	return config
}

func GetPodTemplateSpec(b []byte) (*v1.PodTemplateSpec, error) {
	pts := new(v1.PodTemplateSpec)
	if err := json.Unmarshal(b, pts); err != nil {
		return nil, err
	}

	for i := range pts.Spec.Containers {
		pts.Spec.Containers[i].Name = ContainerName
	}

	return pts, nil
}

func (c *Config) Yaml() error {
	if c == nil || c.key == "" {
		return fmt.Errorf("invalid key or nil Config struct")
	}

	b, err := json.Marshal(c.TFJob)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})

	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	b, err = yaml.Marshal(m)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.key+".yaml", b, 0644)
}
