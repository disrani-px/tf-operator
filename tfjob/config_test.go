package tfjob

import (
	"testing"

	"github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	"k8s.io/api/core/v1"
)

func TestConfig_Yaml(t *testing.T) {
	config := NewConfig("job_tf")

	// initialize params
	job := config.TFJob
	job.Name = "example-job"

	TFReplicaSpecs := make(map[v1alpha2.TFReplicaType]*v1alpha2.TFReplicaSpec)

	worker := new(v1alpha2.TFReplicaSpec)
	TFReplicaSpecs[v1alpha2.TFReplicaTypeWorker] = worker
	job.Spec = v1alpha2.TFJobSpec{
		TFReplicaSpecs: TFReplicaSpecs,
	}

	worker.Template = v1.PodTemplateSpec{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "tensorflow",
					Image: "gcr.io/tf-on-k8s-dogfood/tf_sample:dc944ff",
				},
			},
		},
	}

	if err := config.Yaml(); err != nil {
		t.Fatal(err)
	}
}
