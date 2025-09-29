# Módulo Terraform Reutilizable para Despliegues en Kubernetes

Este proyecto implementa un flujo de trabajo de <b>Infraestructura como Código (IaC)</b> para Kubernetes, aplicando principios de <b>Test-Driven Infrastructure (TDI)</b>. El objetivo principal es proporcionar un módulo de Terraform robusto, reutilizable y validado automáticamente para el despliegue de aplicaciones web.

---

## 🎯 Propósito del Proyecto
En lugar de crear manifiestos de Kubernetes manualmente para cada aplicación, este proyecto ofrece:
1. un <b>módulo de Terraform reutilizable</b> que abstrae la complejidad de desplegar un `Deployment` y un `Service` en Kubernetes.
2. Un <b>conjunto de pruebas automatizadas con Terratest</b> que garantiza que el módulo funciona como se espera, validando la infraestructura desde su creación hasta su destrucción.

Este enfoque promueve la consistencia, reduce errores y acelera el ciclo de vida de los despliegues en entornos de Kubernetes.

---

## 🛠️ Stack Tecnológico
* <b>Orquestación de Contenedores:</b> Kubernetes (gestionado con <b>Kind</b> para desarrollo local).
* <b>Infraestructura como Código (IaC):</b> Terraform.
* <b>Lenguaje de Pruebas:</b> Go (Golang).
* <b>Framework de Pruebas:</b> Terratest.
* <b>Contenerización:</b> Docker.

---

## 🚀 Guía de Uso

#### 1. Prerrequisitos
Para ejecutar este proyecto, necesitas tener instalado lo siguiente:
* [Docker Desktop](https://www.docker.com/products/docker-desktop/)
* [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/)
* [Terraform](https://developer.hashicorp.com/terraform)
* [Go](https://go.dev/dl/)

#### 2. Estructura del Repositorio
```
.
├── modules/
│   └── nginx/            # Módulo reutilizable de Terraform para NGINX
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
├── test/
│   └── nginx_test.go     # Pruebas automatizadas con Terratest
└── main.tf               # Ejemplo de cómo usar el módulo
```

---

## ⚙️ Despliegue y Pruebas
### Despliegue Manual de la Infraestructura
Si solo deseas desplegar la infraestructura definida en `main.tf`, sigue estos pasos desde la raíz del repositorio.

#### 1. Crear el Clúster de Kubernetes:
```bash
kind create cluster
```

#### 2. Inicializar y Aplicar Terraform:
```bash
# Inicializa el backend y descarga los providers
terraform init

# Aplica la configuración para crear los recursos en Kubernetes
terraform apply --auto-approve
```
Puedes verificar los recursos creados con `kubectl get deploy,svc my-nginx-app`.

### Ejecución de las Pruebas Automatizadas
Las pruebas de Terratest orquestan todo el ciclo: `apply`, `verify` y `destroy`.

#### 1. Asegúrate de tener un clúster `kind` en ejecución. Si no lo tienes, créalo.
```bash
kind create cluster
```

#### 2. Navega a la carpeta de pruebas y ejecuta los tests:
```bash
cd test

# Inicializa el módulo de Go y sincroniza las dependencias
go mod init terraform-k8s-testing
go mod tidy

# Ejecuta las pruebas (esto puede tardar varios minutos)
go test -v -timeout 30m
```
La salida de la prueba mostrará los logs de Terraform aplicando la infraestructura, las validaciones de la prueba y, finalmente, la destrucción de los recursos, terminando en un `PASS`

---

## 🧹 Limpieza del Entorno
Si realizaste un despliegue manual, puedes destruir los recursos con:
```bash
terraform destroy --auto-approve
```
