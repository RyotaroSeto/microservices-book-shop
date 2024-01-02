## Docker SetUp

- `docker build -t gihyo-ms-dev-book/bff:0.1 bff/`
- `docker build -t gihyo-ms-dev-book/catalogue:0.1 catalogue/`
- `docker build -t gihyo-ms-dev-book/catalogue-db:0.1 catalogue-db/`
- `docker build -t gihyo-ms-dev-book/frontend:0.1 frontend/`
- `docker build -t gihyo-ms-dev-book/order:0.1 order/`
- `docker build -t gihyo-ms-dev-book/order-db:0.1 order-db/`
- `docker build -t gihyo-ms-dev-book/shipping:0.1 shipping/`

## Kind SetUp

- `kind create cluster --config infra/common/kind/kind-config.yaml`
- `kind delete cluster`

## Deploy

- `sh ./infra/scripts/deploy_all.sh`

## Port Forwarding

- `kubectl port-forward service/frontend 8080:80`
- `kubectl port-forward service/bff 4000:4000`
- ※Use port forwarding for local environment. In production, use ingress.

## Log Checking

- `kubectl logs <pod-name> -c <container-name>`
