# HA K8s Clusterに不可欠だったControl Plane LoadBalancerを亡き者にする!? Cilium 1.18の🔥激アツ🔥新機能

## Status

### ❓ On Evaluation

## 発表概要

複数のControl Planeを備えたKubernetesの構成、HA Cluster。  
マネージドサービスを使わずこの構成を組む際、Control Planeに立っているapiserverへの安定した接続には、特定の要件を満たすロードバランサーが必要でした。少なくとも今までは。

しかし、CNIであるCiliumが最新バージョンの1.18にて、外部コンポーネント無しでこれを解決せんとする🔥激アツ🔥新機能を発表したのです！  
このセッションでは、件の新機能「support for kube-apiserver high availability with kube-proxy replacement」について、なぜロードバランサーが必要だったのかを踏まえて解説します。

## その他質問事項やお問い合わせ

実は去年のCNDWプレイベントの登壇内容と連関する、壮大な伏線回収です。  
<https://cloudnativedays.connpass.com/event/331620/>
