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

# Istio

- Istio のインストール(※ここでは profile=demo にしているが本番では使わない)
  - `istioctl install --set profile=demo -f common/istio/ingressgateway_NodePort.yaml -y`
- サイドカープロキシの自動挿入設定を有効
  - `kubectl label namespace default istio-injection=enabled`

## Blue/Green Deployment

- `kubectl apply -f frontend/k8s/frontend-v1.yaml`
- `kubectl apply -f frontend/k8s/frontend-v2.yaml`

- ラウンドロビンにすることで、v1 と v2 が交互に出力される

  - `kubectl apply -f frontend/istio/destinationrule.yaml`

- v2 にのみルーティングさせる

  - `kubectl apply -f frontend/istio/virtualservice-v2.yaml`

- トラフィックを v1 に 75、%v2 に 25%流す

  - `kubectl apply -f frontend/istio/virtualservice-25-v2.yaml`

- catalogue に対する HTTP リクエストを 2 秒遅延させる(フィールドインジェクション)

  - `kubectl apply -f catalogue/istio/virtualservice-delay.yaml`

- bff のリクエストに対して 0.5 秒のタイムアウト設定

  - `kubectl apply -f bff/istio/virtualservice-timeout.yaml`

- catalogue サービスのサーキットブレイカー設定
  - `kubectl apply -f catalogue/istio/destinationrule-cb.yaml`
  - catalogue で障害を発生させてみる
    - `kubectl scale deploy catalogue-db --replicas=0`

## Security

- 相互 TLS トラフィックのみ許可するように設定(STRICT モード マイクロサービス間の通信暗号化)

  - `kubectl apply -f common/istio/peer-authentication-mtls.yaml`

- Ingress Gateway に対して JWT 検証によるリクエスト認証を設定(認証設定)

  - `kubectl apply -f common/istio/request-authentication-keycloak.yaml`

- bff に対するアクセスは JWT トークンが付与されている場合のみアクセスを許可(認可設定)

  - `kubectl apply -f common/istio/authorization-policy-keycloak.yaml`

- 認証サーバー(keycloak)デプロイ
  - `kubectl apply -f auth/k8s/keycloak.yaml`
