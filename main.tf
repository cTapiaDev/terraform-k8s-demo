module "nginx_dev" {
    source = "./modules/nginx"
    app_name = "my-nginx-app"
    replicas = 2
}