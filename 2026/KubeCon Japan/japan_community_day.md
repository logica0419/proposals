# Discovering Cilium: What eBPF Enabled in K8s Networking

## Status

### 🎉 Accepted (self-nomination)

## Description

KubernetesのネットワークにeBPFを持ち込んだ代表的なプロジェクトであるCilium。広く利用されている一方、実際にどの部分でeBPFが活用され、他のCNIと何が異なるのかを理解する機会はほとんどありません。

このセッションでは、以下のトピックについて、Ciliumの具体的なeBPFコードを交えながら解説します。

- K8sのCNIが実現・実装すべき機能
- Cilium以外のCNI (Flannel・Calico) の実装アプローチ
- Ciliumの実装方法とeBPFの活用箇所
- eBPFだからこその効率化 (netkitやkube-proxy置き換えなど)

Ciliumの内部実装を読み解きながら、他CNIとの違いを比較し、eBPFの導入によって何が可能になったのかを学びましょう。

## Session format

Session Presentation (15 minutes)

## Audience Level

Intermediate

## Bio

Takuto is a graduate student in computer science at Chiba Institute of Technology in Japan, pursuing a career as a cloud platform developer. His expertise includes Go, Kubernetes, and cloud-native infrastructure.
He has been trying to connect platform and network engineers through his sessions in both cloud-native and network conferences. Communicating across domains, he's accelerating the collaboration needed to build better platforms and networks.
