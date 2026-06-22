# Discovering Cilium: What eBPF Enabled in K8s Networking

## Status

### 🎉 Accepted (self-nomination)

## Description

[This talk will be delivered in English]

Cilium has become one of the most widely adopted CNI implementations in Kubernetes.  
Yet many engineers use Cilium without understanding how it differs from other CNIs, how eBPF is used inside, or what new capabilities it enables.

In this session, we will explore the following topics using real eBPF code from Cilium:

- Core responsibilities of a Kubernetes CNI
- How other CNIs, such as Flannel and Calico, implement those responsibilities
- How Cilium implements those responsibilities with eBPF
- Features and optimizations enabled by eBPF, including kube-proxy replacement and netkit

Through a deep dive into Cilium's internals and comparisons with other CNIs, we'll explore what eBPF made possible in Kubernetes networking.

## Session format

Session Presentation (15 minutes)

## Audience Level

Intermediate

## Bio

Takuto is a graduate student in computer science at Chiba Institute of Technology in Japan, pursuing a career as a cloud platform developer. His expertise includes Go, Kubernetes, and cloud-native infrastructure.
He has been trying to connect platform and network engineers through his sessions in both cloud-native and network conferences. Communicating across domains, he's accelerating the collaboration needed to build better platforms and networks.
