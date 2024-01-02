## Docker SetUp

- `docker build -t gihyo-ms-dev-book/catalogue:0.1 catalogue/`
- `docker build -t gihyo-ms-dev-book/bff:0.1 bff/`
- `docker build -t gihyo-ms-dev-book/frontend:0.1 frontend/`

## Kind SetUp

- `kind create cluster`

## Deploy

- `kind load docker-image gihyo-ms-dev-book/catalogue:0.1`
- `kubectl apply -f infra/app/catalogue/catalogue.yaml`
- `kind load docker-image gihyo-ms-dev-book/bff:0.1`
- `kubectl apply -f infra/app/bff/bff.yaml`
- `kind load docker-image gihyo-ms-dev-book/frontend:0.1`
- `kubectl apply -f infra/app/frontend/frontend.yaml`
- - ※Use kind load for local environment. Unnecessary in production because images are placed in the container registry.

## Port Forwarding

- `kubectl port-forward service/frontend 8080:80`
- `kubectl port-forward service/bff 4000:4000`
- ※Use port forwarding for local environment. In production, use ingress.

## Log Checking

- `kubectl logs <pod-name> -c <container-name>`
