# The Double-Edged Sword of Goroutines: Why runc Needed CGo

## Status

### ❔ In Evaluation

## Description

Goroutine is one of the magics that Go provides to developers.
However, there are a few cases where the design of goroutines makes implementation complicated; one of them is the container runtime.
