# Resizing Animated GIFs without CGO nor Third-Party Libraries

## Description

What would you use when you implement GIF resizing in your web application? ImageMagick as we did? CLI? Or a C library wrapper? Stop using them and get out from the hell of overhead by implementing GIF resizing with pure Go!

There is little knowledge of GIF resizing in Go. However, it is so simple that we can implement it with only standard libraries and sub-repositories (golang.org/x libraries)!
Learn how to implement GIF resizing with pure Go step by step and the cautionary points you may face during the implementation.

## Target Audience

 Advanced

## Outline

- Why we needed to implement GIF resizing in pure Go (5 min)
  - Self Introduction
  - Introducing traQ (<https://github.com/traPtitech/traQ>), a web application I had maintained
  - Explain why we wanted to get rid of ImageMagick
- (1) Basic resizing implementation (5 min)
  - Image resizing with the `golang.org/x/image/draw` package
  - Basic structure of `*(image/gif).GIF` struct
  - Implementation: <https://github.com/logica0419/proposals/blob/main/GopherCon2024/step1.go>
- (2) Handling frame optimization (5 min)
  - What is frame optimization?
  - How frame optimization affects the`*(image/gif).GIF` struct
  - Calculating frame size for each frame
  - Implementation: <https://github.com/logica0419/proposals/blob/main/GopherCon2024/step2.go>
- (3) Handling black glitchy noise (10~15 min)
  - Simple explanation of image resizing algorithm and its implementation in Go
  - Problem of transparent pixel handling
  - Frame stacking before resizing
  - Implementation: <https://github.com/logica0419/proposals/blob/main/GopherCon2024/step3.go>
- (4) Handling frame disposal (10~15 min)
  - GIF frame disposal specification
  - Handling frame disposal in Go
  - Implementation: <https://github.com/logica0419/proposals/blob/main/GopherCon2024/step4.go>
- Conclusion (5 min)
  - Summary of the implementation
  - A little explanation of parallelization of resizing
  - Introducing resigif (<https://github.com/logica0419/resigif>) library as an advanced product

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan. He drives himself to study computer science, dreaming of becoming a cloud platform developer. His core domain includes web development (backend) and cloud infrastructure. He had been a backend core maintainer of traQ, a communication web service for the university tech club. While in the position, he achieved several significant improvements in performance.
