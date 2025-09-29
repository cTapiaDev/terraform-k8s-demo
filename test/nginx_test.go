package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestNginxDeployment(t *testing.T) {
	t.Parallel()

	// Configuración inicial de terraform
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	// Sirve para destruir la infraestructura luego de usarla
	defer terraform.Destroy(t, terraformOptions)

	// Desplegar
	terraform.InitAndApply(t, terraformOptions)

	// Configurar opciones de kubectl
	kubectlOptions := k8s.NewKubectlOptions("", "", "default")

	// Verificar si el servicio existe
	k8s.WaitUntilServiceAvailable(t, kubectlOptions, "my-nginx-app", 10, 1*time.Second)
	serviceName := "my-nginx-app"
	t.Logf("Servicio %s disponible", serviceName)

	// túnel al servicio
	tunnel := k8s.NewTunnel(kubectlOptions, k8s.ResourceTypeService, serviceName, 8080, 80)
	defer tunnel.Close()
	tunnel.ForwardPort(t)
	endpoint := tunnel.Endpoint()

	// petición HTTP y verificar el status
	url := fmt.Sprintf("http://%s", endpoint)

	http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, 30, 3*time.Second, func(statusCode int, body string) bool {
		return statusCode == 200
	})
	t.Logf("Petición HTTP a %s exitosa con status 200", url)

}
