# Resizing Animated GIFs Without CGO or Third-Party Libraries

## Status

### 😞 Declined

## Description

GIF animation icons enable users to express themselves much better. Resizing animated GIFs is necessary to support them.  
What would you use when implementing GIF resizing in your Go web application? ImageMagick as we did? Another CLI? Or a C library wrapper?

There is little knowledge of GIF resizing in Pure Go. There are articles about basic image resizing, but they don't cover animated GIF-specific processing. Therefore, popular methods to implement it are using CLIs or C library wrappers.  
However, those invoke several unignorable issues. Calling CLIs directly from the code causes security concerns. Image-processing CLIs and C libraries are typically large files that don't fit into a container environment. Also, they cause performance overheads that affect app latency.

Stop using them by implementing GIF resizing with pure Go! It requires some specific knowledge, but it's so simple that we can implement it with only standard libraries and sub-repositories (golang.org/x libraries)!  
Learn how to implement GIF resizing with pure Go step by step and the cautionary points you may face during the implementation.

## Target Audience

 Advanced

## Outline

- Why we needed to implement GIF resizing in pure Go (5 min)
  - Self Introduction
  - Introducing traQ (<https://github.com/traPtitech/traQ>), a web application I had maintained
  - Explain why we wanted to get rid of ImageMagick
    - Wanted to replace base image with distroless due to security concerns, but blocked by ImageMagick library installation
    - Docker image size was large (over 600MiB), but we don't utilize most of the library
- (1) Basic resizing implementation (5 min)
  - Image resizing with the `golang.org/x/image/draw` package
  - Basic structure of `*(image/gif).GIF` struct
  - Implementation: <https://github.com/logica0419/proposals/blob/main/2024/GopherCon/step1.go>
- (2) Handling frame optimization (5 min)
  - What is frame optimization?
  - How frame optimization affects the`*(image/gif).GIF` struct
    - Each frame is trimmed
    - Frame size differs in each frame
  - Calculating frame size for each frame
    - Calculate the resize ratio for width & height
    - Calculate the resized frame size for each frame
  - Implementation: <https://github.com/logica0419/proposals/blob/main/2024/GopherCon/step2.go>
- (3) Handling black glitchy noise (10 min)
  - Simple explanation of image resizing algorithm and its implementation in Go
  - Problem of transparent pixel handling
  - Frame stacking before resizing
  - Implementation: <https://github.com/logica0419/proposals/blob/main/2024/GopherCon/step3.go>
- (4) Handling frame disposal (10 min)
  - GIF frame disposal specification
    - None: keep stacked frame to the next frame
    - Background: fill the canvas with background color after the current frame rendering
    - Previous: put back canvas to the previous condition
  - Handling frame disposal in Go
    - After-resize processing for the temporary canvas
  - Implementation: <https://github.com/logica0419/proposals/blob/main/2024/GopherCon/step4.go>
- Conclusion (10 min)
  - Summary of the implementation
  - A little explanation of parallelization of resizing
    - The resizing step is the slowest
    - Frame stacking can't be parallelized (must be in order)
    - Wrapping only resizing step with goroutine
  - Effect on our app
    - Docker image shrinks down (40~50% in size)
    - Able to move on to distroless
    - Performance improvement (x3 faster in same CPU usage)
    - Able to select the resizing kernel
  - Introducing resigif (<https://github.com/logica0419/resigif>) library as a result of work.

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan. He drives himself to study computer science, dreaming of becoming a cloud platform developer. His core domain includes web development (backend) and cloud infrastructure. He had been a backend core maintainer of traQ, a communication web service for the university tech club. While in the position, he achieved several significant improvements in performance.

## Additional Notes

I'm a university student who has been willing to visit GopherCon for three years, but I couldn't because of financial problems. I applied to the scholarship program, but I couldn't get after all.  
I hope this proposal leads me to the first time GopherCon!  
Thank you,
