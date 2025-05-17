package backup

import v1 "k8s.io/client-go/kubernetes/typed/batch/v1"

type KubeClient interface {
	BatchV1() v1.BatchV1Interface
}
