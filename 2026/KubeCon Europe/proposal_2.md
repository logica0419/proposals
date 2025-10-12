# Process, Goroutine, and Runtime: Why Go Struggles With Low-Level Runtimes

## Status

### ❔ In Evaluation

## Description

Go powers most cloud-native software, including container runtimes like runc.  
However, Go's concurrency model introduces challenges for low-level OS process handling, which is critical for the low-level runtime.  
Core runtime features such as namespace creation rely on CGo, highlighting Go's constraints for runtime development.

This session explores:

- OS processes and syscalls essential to runtimes
- Container runtime architecture
- How Goroutines affect process forking
- Alternative runtimes: youki, crun

Through live demos, we compare minimal runtimes implemented in C and Go, highlighting trade-offs.  
Attendees will understand runtime design, Go's limitations, and how to choose the appropriate language depending on projects.

Outline:

1. OS processes & syscalls
2. Container runtime anatomy
3. Demo: minimal runtime in C
4. Go concurrency & Goroutines explained
5. Demo: minimal runtime in Go
6. runc's CGo solution
7. Alternative approaches: youki & crun

## Track

Platform Engineering

## Session format

Session Presentation (30 minutes)

## Level

Intermediate

## Benefits to the ecosystem

Container runtimes are core to the cloud-native ecosystem, yet engineers often lack a deep understanding of their low-level operation.  
Many projects adopt Go by default, without awareness of its limitations.  
This session clarifies how Go's concurrency model interacts with system-level processes and guides engineers on language and design choices.

Attendees will leave with:

- A solid grasp of container runtime architecture
- Insight into Go's runtime behavior and process limitations
- Practical knowledge of alternative runtime implementations
- Guidance for designing or contributing to container runtimes

## CNCF-hosted software

Kubernetes, Containerd

## Open source projects

runc, youki, crun

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan, pursuing a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He is the founder of Coding Kubernetes, a tutorial series that rebuilds Kubernetes core components from scratch to teach its internals and cloud-native design. Through this project, he shares practical insights with future cloud-native maintainers.
