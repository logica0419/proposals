# 徹底比較！HA Kubernetes ClusterにおけるControl Plane LoadBalancerの選択肢

## Status

### 😞 Declined

前夜祭にて 🎉 Accepted

## 概要

複数のControl Planeを備え、いくつかのControl Planeが壊れてもクラスタを維持できるKubernetesの構成を「HA Cluster」と呼びます。  
マネージドサービスを使わずこの構成を組む際、それぞれのControl Planeに立っているapiserverへの接続には、「接続時、必ず生きているapiserverに到達する」「常に同じドメイン / IPアドレスで接続できる」という要件を満たすロードバランサーが必要です。

このセッションでは、利点欠点やドキュメント化されていない問題にも触れながら、いくつかのソリューションを紹介します。  
HA Clusterを組みたいけど自分の環境に合うロードバランサーがわからない…というみなさん、必見です！

## カテゴリ

Networking

## トーク時間

40min (full session)

## プロフィール

千葉工業大学 情報科学部 情報ネットワーク学科 所属。Webバックエンドを起点として様々な分野のコンピューターサイエンスを学び、パブリッククラウド基盤の開発者になることを目指して日々修行中。  
ICTSCという学生向けIT大会の運営チームでKubernetes基盤運用リーダーを務める他、大学のサークルにてNekko Cloudという小規模マルチリージョンプライベートクラウド計画の基盤ソフトウェア開発にあたる。
