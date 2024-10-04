# Q1 Project structure

## 1. go-template for go project structure
Using Gin framework (https://github.comn/gin-tonic/gin) for API server and Google wire (https://github.com/google/wire) for injection dependency management.
## 2. nest-template for nest project strucutre
Using NestJS https://nestjs.com default project structure.
## 3. deployment for deployment docker compose
Using docker compose for deploy docker stack. Use docker NGINX for manage proxy for both app.<br>
> Note: to deploy app, copy .env.example and rename it to .env and change content before use. 
<br>
> `NGINX_HOST`: will be use as proxy web server name. If provided `localhost`, you can access `nest-app` with `nest.localhost` and `go-app` with `go.localhost`.
<br>
> `NGINX_CONF`: If you have other config eg. remote, you can add it. but do not forget to change nginx/docker-compose.yaml for mapping key or use certbot for auto key renew.
<br>
> `TAG`: will be use as docker tag.
