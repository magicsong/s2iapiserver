package s2irun

import (
	"encoding/json"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ConfigDataKey = "data"
)

func (c *S2iRunControllerImpl) createConfigMap(instance *v1alpha1.S2iRun, config *v1alpha1.S2iConfig) (*corev1.ConfigMap, error) {
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	dataMap := make(map[string]string)
	dataMap[ConfigDataKey] = string(data)
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name + "-s2iconfigmap",
			Namespace: instance.ObjectMeta.Namespace,
		},
		Data: dataMap,
	}
	return configMap, nil
}
func (c *S2iRunControllerImpl) createJob(instance *v1alpha1.S2iRun) *batchv1.Job {
	//create configmap
	cfgString := "config-data"
	configMapName := instance.Name + "-s2iconfigmap"
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name + "-s2ijob",
			Namespace: instance.ObjectMeta.Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "s2irun",
							Image:           "magicsong/s2irun",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Env: []corev1.EnvVar{
								{
									Name:  "S2I_CONFIG_PATH",
									Value: "/etc/data/config.json",
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      cfgString,
									ReadOnly:  true,
									MountPath: "/etc/data",
								},
								{
									Name:      "docker-sock",
									MountPath: "/var/run/docker.sock",
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{
						{
							Name: cfgString,
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: configMapName,
									},
									Items: []corev1.KeyToPath{
										{
											Key:  ConfigDataKey,
											Path: "config.json",
										},
									},
								},
							},
						},
						{
							Name: "docker-sock",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{Path: "/var/run/docker.sock"},
							},
						},
					},
				},
			},
			BackoffLimit: &instance.Spec.BackoffLimit,
		},
	}
	return job
}
