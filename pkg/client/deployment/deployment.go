package deployment

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeking.io/kubesphere/pkg/client/global"
	"kubeking.io/kubesphere/pkg/client/types"
	"kubeking.io/kubesphere/pkg/client/utils"
)

type Deployment struct {
	Host       string
	HttpClient *types.HttpClient
}

const (
	DeploymentApi = "/apis/clusters/{{cluster}}/apps/v1/namespaces/{{namespace}}/deployments/{{deployment}}"
)

func (c *Deployment) Update(cluster, namespace, deployment string, deploy v1.Deployment) error {

	deploy.ManagedFields = []v12.ManagedFieldsEntry{}

	yamlData, err := json.Marshal(deploy)
	if err != nil {
		return err
	}

	_, err = c.HttpClient.Request.R().SetHeader("Content-Type", "application/json").SetAuthToken(global.AuthToken.AccessToken).SetBody(string(yamlData)).Put(utils.ParseUrl(c.Host, DeploymentApi, map[string]string{
		"cluster":    cluster,
		"namespace":  namespace,
		"deployment": deployment,
	}))
	if err != nil {
		return err
	}

	fmt.Println(1)
	return nil
}

func (c *Deployment) Get(cluster, namespace, deployment string) (v1.Deployment, error) {

	d := v1.Deployment{}
	_, err := c.HttpClient.Request.R().SetAuthToken(global.AuthToken.AccessToken).SetResult(&d).Get(utils.ParseUrl(c.Host, DeploymentApi, map[string]string{
		"cluster":    cluster,
		"namespace":  namespace,
		"deployment": deployment,
	}))
	if err != nil {
		return v1.Deployment{}, err
	}

	return d, nil
}
