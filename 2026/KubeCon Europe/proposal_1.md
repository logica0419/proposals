# CNI Demystified: Philosophy, Design, and Plugin Trade-Offs

## Status

### 😞 Declined

## Description

CNI powers container networking, such as Kubernetes, but often feels like a black box, leading to a discord between platform and network engineers.  
In this session, we unpack CNI by exploring both its philosophy and practical implementations.

We'll cover:

- The flexible, pluggable design of the CNI specification
- How major OSS plugins (Flannel, Calico, Cilium) implement that design
- The strengths, trade-offs, and pitfalls of each approach

Platform engineers with intermediate networking knowledge will learn how to choose, customize, or even build their own CNI solutions.  
You'll also gain confidence to collaborate effectively with network engineers and design more robust platforms.

Outline:

1. CNI philosophy & spec
2. Anatomy of a plugin
3. Life of a packet: Flannel / Calico / Cilium internals
4. Trade-offs and real-world lessons
5. Guidance for customization and extension

## Track

Connectivity

## Session format

Session Presentation (30 minutes)

## Level

Intermediate

## Benefits to the ecosystem

Many engineers misunderstand CNI as too complex, widening the gap between platform and network engineers.  
This session clarifies how CNI bridges these domains, helping engineers move beyond general-purpose OSS CNI choices.  
Attendees will gain the insight needed to select, tune, or extend CNIs for optimal performance and maintainability in their environments.

## CNCF-hosted software

CNI, Kubernetes, Containerd, Cilium

## Open source projects

Flannel, Calico

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan, pursuing a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He is the founder of Coding Kubernetes, a tutorial series that rebuilds Kubernetes core components from scratch to teach its core internals. Through this project, he shares practical insights with future cloud-native maintainers.
