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

- `sh ./scripts/deploy_all.sh`

## Log Checking

- `kubectl logs <pod-name> -c <container-name>`

## Istio

- Istio のインストール(※ここでは profile=demo にしているが本番では使わない)
  - `istioctl install --set profile=demo -f common/istio/ingressgateway_NodePort.yaml -y`
- サイドカープロキシの自動挿入設定を有効
  - `kubectl label namespace default istio-injection=enabled`
