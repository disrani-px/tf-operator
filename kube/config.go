package kube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	key string
	Job *v1alpha2.TFJob
}

func NewConfig(key string) *Config {
	config := new(Config)
	config.key = key
	config.Job = new(v1alpha2.TFJob)
	config.Job.Kind = Kind
	config.Job.APIVersion = APIVersion

	return config
}

func (c *Config) Yaml() error {
	if c == nil || c.key == "" {
		return fmt.Errorf("invalid key or nil Config struct")
	}

	b, err := json.Marshal(c.Job)
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
