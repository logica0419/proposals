# EBUSY in Nested Containers: From Runc Prototype to Youki's Different Design

## Status

### ❔ In Evaluation

## Session Title

EBUSY in Nested Containers: From Runc Prototype to Youki's Different Design

## Description

Nested container exec broke with EBUSY, known from KinD.  
The speaker found this bug while using youki as a runc alternative and traced it to the cause: the exec process cannot join the container's cgroup.

This session compares runc's "retrying" method, which falls back to init cgroup, masking non-EBUSY errors, with the new approach youki chose, "infer-first".  
Infer-first reads `/proc/{pid}/cgroup` to infer the sub-cgroup where init actually belongs, choosing placement first without masking errors via fallback.  
The speaker prototyped the runc approach (#3347), proposed inference design (#3564) validated in DevContainer and Docker-in-Docker, scoped runc compat to subtree_control exception, leaving moved-init out of scope, adopted through the discussion with runc and youki maintainers at KubeCon Japan.

OSS contributors and platform engineers leave with a retrying vs. infer-first flowchart, a compat-scope checklist, and steps to decide what compat to keep or drop.

## Submitting for

KubeCon + CloudNativeCon

## Poster Session

No

## Value to the community

Later products do not have to copy precedent to be compatible.  
This session shares the speaker's method for negotiating compat scope with maintainers and choosing a different path: maximizing compatibility with runc versus inferring placement first without masking errors via fallback.  
Attendees leave with a retrying vs. infer-first selection flowchart for a better understanding of the nested containers' cgroup problem, plus a compat-scope checklist covering the subtree_control exception and the moved-init out-of-scope case, and concrete steps for reading `/proc/{pid}/cgroup` to infer the sub-cgroup for exec placement.

## Case study

No

## Talk presented before

No

## Session format

Session Presentation (30 minutes)

## Relevant CNCF/Open Source Projects

youki, runc, OCI runtime-spec

## Bio

Takuto is a graduate student in computer science at Chiba Institute of Technology in Japan, pursuing a career as a cloud platform developer. His expertise includes Go, Kubernetes, and cloud-native infrastructure.  
He organizes Gophers EX, a project that supports the global expansion of the Japanese Go community. With his experience as a speaker at an international conference, he shares practical approaches for challenging and participating in conferences overseas.
