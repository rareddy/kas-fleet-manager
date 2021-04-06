package config

import (
	"github.com/spf13/pflag"
)

type OSDClusterConfig struct {
	IngressControllerReplicas    int                   `json:"ingress_controller_replicas"`
	OpenshiftVersion             string                `json:"cluster_openshift_version"`
	ComputeMachineType           string                `json:"cluster_compute_machine_type"`
	StrimziOperatorVersion       string                `json:"strimzi_operator_version"`
	ImagePullDockerConfigContent string                `json:"image_pull_docker_config_content"`
	ImagePullDockerConfigFile    string                `json:"image_pull_docker_config_file"`
	DynamicScalingConfig         *DynamicScalingConfig `json:"dynamic_scaling_config"`
}

type DynamicScalingConfig struct {
	Enabled bool `json:"enabled"`
}

func NewOSDClusterConfig() *OSDClusterConfig {
	return &OSDClusterConfig{
		OpenshiftVersion:             "",
		ComputeMachineType:           "m5.4xlarge",
		StrimziOperatorVersion:       "v0.21.3",
		ImagePullDockerConfigContent: "",
		ImagePullDockerConfigFile:    "secrets/image-pull.dockerconfigjson",
		IngressControllerReplicas:    9,
		DynamicScalingConfig: &DynamicScalingConfig{
			Enabled: false,
		},
	}
}

func (s *OSDClusterConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.OpenshiftVersion, "cluster-openshift-version", s.OpenshiftVersion, "The version of openshift installed on the cluster. An empty string indicates that the latest stable version should be used")
	fs.StringVar(&s.ComputeMachineType, "cluster-compute-machine-type", s.ComputeMachineType, "The compute machine type")
	fs.StringVar(&s.StrimziOperatorVersion, "strimzi-operator-version", s.StrimziOperatorVersion, "The version of the Strimzi operator to install")
	fs.StringVar(&s.ImagePullDockerConfigFile, "image-pull-docker-config-file", s.ImagePullDockerConfigFile, "The file that contains the docker config content for pulling MK operator images on clusters")
	fs.IntVar(&s.IngressControllerReplicas, "ingress-controller-replicas", s.IngressControllerReplicas, "The number of replicas for the IngressController")
	fs.BoolVar(&s.DynamicScalingConfig.Enabled, "enable-dynamic-scaling", s.DynamicScalingConfig.Enabled, "Enable Dynamic Scaling functionality")
}

func (s *OSDClusterConfig) ReadFiles() error {
	if s.ImagePullDockerConfigContent == "" && s.ImagePullDockerConfigFile != "" {
		err := readFileValueString(s.ImagePullDockerConfigFile, &s.ImagePullDockerConfigContent)
		if err != nil {
			return err
		}
	}
	return nil
}