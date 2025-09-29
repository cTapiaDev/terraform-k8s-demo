variable "namespace" {
    description = "El namespace donde se desplegará NGINX"
    type = string
    default = "default"
}

variable "app_name" {
    description = "El nombre para la aplicación."
    type = string
}

variable "replicas" {
    description = "Número de réplicas para el despliegue."
    type = number
    default = 1
}