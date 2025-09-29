provider "kubernetes" {
    config_path = "~/.kube/config"
}

resource "kubernetes_deployment" "nginx" {
    metadata {
        name = var.app_name
        namespace = var.namespace
        labels = {
            app = var.app_name
        }
    }
    spec {
        replicas = var.replicas
        selector {
            match_labels = {
                app = var.app_name
            }
        }
        template {
            metadata {
                labels = {
                    app = var.app_name
                }
            }
            spec {
                container {
                    image = "nginx.1.21.6"
                    name = "nginx"
                    port {
                        container_port = 80
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "nginx" {
    metadata {
        name = var.app_name
        namespace = var.namespace
    }
    spec {
        selector = {
            app = kubernetes_deployment.nginx.metadata[0].labels.app
        }
        port {
            port = 80
            target_port = 80
        }
        type = "NodePort"
    }
}