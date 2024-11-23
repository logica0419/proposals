# Nekko Cloud Gen.2: Constructing a Small Cloud With "Dependency Inversion"

## Status

### ❓ On Evaluation

## Description

Nekko Cloud is an experimental small-scale private cloud developed by a CS student community, connecting the home labs of its members.

Currently, the platform combines Proxmox for IaaS with a Kubernetes cluster for PaaS.  
However, with containerized applications as its primary use case, dual virtualization creates inefficiencies, while challenges like fragmented authentication and complex observability limit usability.

To address these issues, we propose "dependency inversion," shifting from traditional "PaaS on IaaS" to "IaaS on PaaS" using ecosystem tools like KubeVirt.  
By making bare-metal Kubernetes the backbone of all cloud functionalities, we simplify the architecture, eliminate redundancy, and optimize for small-scale private cloud needs.  
This ongoing project is called "Nekko Cloud Gen.2."

Discover how Nekko Cloud Gen.2 with "dependency inversion" redefines small-scale cloud design, solving inefficiencies and enhancing usability!

## Level

Advanced

## Benefits to the ecosystem

This presentation offers a paradigm shift in cloud architecture, specifically for small-scale private clouds.  

While some previous experiments have explored placing IaaS on Kubernetes, they are typically executed on large platforms, so the benefits are relatively small. In contrast, Nekko Cloud Gen.2 demonstrates how "dependency inversion" can streamline small-scale cloud design by eliminating inefficiencies and redundancy.  
This approach is valuable for the environment with limited resources where traditional cloud models are inefficient or unnecessarily complex. Focusing on small-scale private clouds, this project introduces a scalable, resource-efficient model that reduces the overhead of managing large virtualized environments. It provides practical insights for developers and communities working with constrained resources, offering solutions that are easy to implement, effective, and tailored to specific use cases.

Ultimately, Nekko Cloud Gen.2 showcases how to build a lightweight cloud platform that enhances sustainability and accessibility, making cloud computing more efficient for small, resource-constrained setups.

## CNCF-hosted software

Kubernetes, Knative, KubeVirt

## Open source projects

Proxmox, Netmaker

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan, pursuing a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He has taken the lead in the software development area of the Nekko Cloud team, a private cloud development team in a CS student community. He usually focuses on building an efficient and user-friendly cloud platform.
