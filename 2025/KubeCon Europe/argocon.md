# Mastering Config Management: Exploring and Extending Manifest Generation

## Status

### ❓ On Evaluation

## Description

Inside Argo CD, the argocd-repo-server manages the source repository and generates raw Kubernetes manifests, which define the desired state of an application.  
This process, known as config management, is a cornerstone of GitOps workflows.

Understanding config management is crucial for debugging complex application states and improving deployment efficiency.  
This session explores "native" tools like Kustomize and Helm, offering insights into their implementation.  
It also highlights opportunities for contribution, illustrated by the speaker's enhancements to Helm in Argo CD.

The talk also introduces Configuration Management Plugins (CMP), the Argo CD extension to integrate third-party tools.  
Attendees will learn best practices for designing and implementing CMP, with practical insights from the speaker's experience developing a Helmfile CMP.

Discover how mastering config management and CMP can enhance productivity and foster meaningful contributions to open-source projects!

## Level

Advanced

## Benefits to the ecosystem

Config management is the cornerstone of GitOps workflows, yet it's often overlooked.  
This session fills that gap by providing a foundational understanding of manifest generation in Argo CD, helping attendees better debug complex application states.

Attendees will learn how Argo CD uses config management tools like Kustomize and Helm to define the desired state of applications.  
The session will also highlight opportunities for contribution to Argo projects, using the speaker's work enhancing Helm in Argo CD as a concrete example.

Additionally, introducing Configuration Management Plugins (CMP) will show how Argo CD can be extended to integrate third-party tools, enriching the ecosystem and enabling more flexible deployments.  
This session encourages active participation in improving config management, benefiting the Argo and Kubernetes communities.

## Open source projects

Kubernetes, Argo CD, Helm, Helmfile

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan, pursuing a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He has taken the lead in the software development area of the Nekko Cloud team, a private cloud development team in a CS student community. He usually focuses on building an efficient and user-friendly cloud platform.
