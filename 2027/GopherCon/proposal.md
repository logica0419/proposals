# The Double-Edged Sword of Goroutines: Why runc Needs cgo

## Status

### ❔ In Evaluation

## Description

For nearly a decade, runc has been a de facto standard low-level container runtime.  
Although runc is written in Go, it uses cgo for a critical part of container setup.  
Why not implement this part in pure Go?

Containers are isolated processes. Setting them up requires precise control over process creation and namespaces.  
Go's runtime, however, manages OS threads on behalf of goroutines, abstracting away much of the thread-level control that container setup requires.  
Some namespace operations require a single-threaded process, which is hard to guarantee in a running Go program.  
In runc, this leads to an interesting design choice: a C program registered through cgo performs part of the container setup before the Go runtime starts.

In this talk, I'll explain how the Go runtime schedules goroutines using the GMP model, and why process creation and thread-specific state can challenge these abstractions.  
We'll explore how these constraints shape the design of Go's standard library, including `os/exec` and `syscall`, and trace how runc uses cgo to set up a container's namespaces.

The talk draws on a small container-runtime PoC in Go that I built as part of a hands-on workshop.  
My contributions to youki, a Rust container runtime, also inform my perspective on runtime design across languages.

These challenges aren't unique to container runtimes. Any Go system that needs fine-grained control over OS processes, threads, or thread-specific state can encounter similar constraints.  
By tracing the path from goroutines through the standard library to runc, you'll see where Go's abstractions work well, where they impose constraints, and how to account for these trade-offs when designing Go systems or choosing a language for low-level systems.

## Target Audience

All Attendees

## Outline

1. Introduction (2 min)
   - Raising a question: Why does runc use cgo?
   - Agenda
2. How Goroutines Work: The GMP Model (4 min)
   - How the Go runtime schedules goroutines across OS threads
   - Why goroutines enable lightweight, efficient concurrency
3. What Are Containers? (4 min)
   - Containers are isolated processes
   - runc and its role as a low-level container runtime
4. Goroutine Constraints and the Standard Library (6 min)
   - Why process creation and thread-specific state pose challenges for Go programs
   - The role of `runtime.LockOSThread()`
   - Why Go's `os/exec` APIs favor fork-and-exec patterns over a bare `fork` API
   - Re-exec through a subcommand pattern
5. Why runc Uses cgo (7 min)
   - How runc combines namespace setup, double fork, and handoff to Go
   - How runc uses cgo to perform low-level setup
   - The trade-offs of this approach
6. Conclusion (2 min)
   - Understanding Go's limitations to understand its strengths better
   - Making design or language choices based on the requirements of the task
   - Examples of container runtimes in other languages: youki and crun

## Speaker Biography

Takuto is a graduate student in computer science at Chiba Institute of Technology in Japan, pursuing a career as a cloud platform developer. His expertise includes Go, Kubernetes, and cloud-native infrastructure.  
Drawn to container technology through his exploration of cloud-native infrastructure, he developed Coding Container Runtime, a hands-on workshop to build a PoC container runtime in Go. He also contributes to youki, a container runtime written in Rust.
