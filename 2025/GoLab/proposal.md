# Fundamentals of Memory Management in Go: Learning Through the History

## Status

### 😞 Declined

## Description

Memory management is crucial for Go performance, but most sessions have been overly complex. If you're new to memory management or have struggled to learn it, this session is for you.  
A solid understanding of memory management helps tackle performance challenges. For example, only a single memory allocation change once led to a 57% reduction in CPU usage and 99% lower memory consumption.  
Yet, past conference sessions are often difficult for learners because they dive straight into technical details like heap escaping and allocation logic, focusing on how they work. But before tackling the "how," it's essential to learn what those functionalities are and why they were developed.  
Heap, stack, and garbage collection—memory management concepts have evolved to solve data management challenges in programming languages. By exploring their motivations, you'll gain a deeper understanding of memory management and the confidence to engage in Go memory management discussions.  
Let's unravel the chronicle of memory management and build a strong foundation for mastering advanced Go memory concepts.

## Session format

Talk (30 mins + Q/A time)

## Level

Introductory and overview

## Abstract

1. Introduction (3min)
   - Memory management is important for performance
     - Example case: one memory allocation change -> 57% reduction in CPU usage and 99% lower memory consumption
   - Sessions about memory management are difficult
     - Contain large technical details of logic and implementation techniques ("how")
     - Lack of structured abstract interpretation of processing and understanding of needs ("what" and "why")
   - To learn advanced Go memory concepts, we need to understand the basics of a broad range of functionalities
     - We can acquire this only by grasping the "what" and "why" of each functionality
     - Let's focus on the "what" and "why", not the "how": the most efficient path from 0 to hero
2. What memory is: Ancient memory management (1min)
   - Memory = temporary data storage for the process
   - Direct data insertion to memory (data management of assembly)
3. Era of variables: the stack and the heap (4min)
   - (Pain) Programmers need to know all the addresses
   - High-level programming language arose
     - Variable = abstraction of data
     - Programmers don't need to know the address, but language needs to manage them
   - Invention of the stack
     - The stack was perfect to represent the nested structure of function calls & variable scope on memory space
   - Illustrated explanation of the stack
     - (Example) Go variable declaration and function call
   - Partial use of direct insertion: the heap and the pointer
     - For big / unknown-sized data
     - For data which is used across the functions
     - (Example) malloc & free function in C
4. Automatic heap management: garbage collection (2min)
   - (Pain) Programmers need to delete data in the heap when they are no longer in use
     - Risk of memory leak & OOM kill
   - Garbage collection (GC)
     - Detection of unused data in the heap
     - Deletion of them
5. Heap escaping in Go (5min)
   - (Pain) Placing many data in the heap = high GC cost
     - (Example) Putting all struct & array in heap -> over half of the variables are GC target
   - Heap escaping
     - Put as much data in the stack as possible
     - Put necessary data in a heap, types don't matter
   - Heap escape analysis
     - Heap escaping is determined by the compiler, not the runtime
     - Determined in the binary build phase
     - A build option to show which variable escapes
   - Zero-allocation libraries
     - Package that no data escapes to the heap in its functions
     - Efficient, basically increases performance
     - Implementation is relatively complex, so difficult to read internal code
6. Flexible stack with stack copying (4min)
   - (Pain) The stack sometimes has a limit = stack overflow happens
   - Preparing a huge stack may be good…
   - (Pain) Always having a huge stack = a small amount of heap, which is also inconvenient
   - Flexible-sized stack = efficient memory usage
     - Segmented stacks (old implementation)
     - Stack copying (current implementation)
7. Explanation of the example case in the introduction (3min)
   - cf. <https://developers.cyberagent.co.jp/blog/archives/54653/> (written in Japanese)
   - Filtering function makes one slice in the heap every time it is called
     - This escaped variable was the bottleneck

    ```go
    func (list targetList) Filter(matcher Matcher) targetList {
        filtered := make(targetList, 0, len(l))  // The bottleneck
        for _, e := range list {
            if matcher.Match(e) {
                filtered = append(filtered, e)
            }
        }
        return filtered
    }
    ```

   - Deleting allocation by updating the existing slice

    ```go
    func (list *targetList) Filter(matcher Matcher) targetList {
        n := 0
        for _, e := range list {
            if matcher.Match(e) {
                [*list](n) = e
                n++
            }
        }
        *list = [*list](:n)
        return list
    }
    ```

   - 99% lower memory consumption
     - Shows how one allocation in a common function affects the performance
   - 57% reduction in CPU usage
     - The reduction of allocation & GC CPU usage combined
8. Conclusion (2min)
   - This session is the start
   - Deep dive into each functionality with existing sessions!
9. Recommendation of the past sessions to see next (if I have extra time)
   - With the basics of a broad range of functionalities in this session as a strong foundation, you can move on to the next step!
   - Heap escape condition and analysis: "Escape Analysis in Go: Understanding and Optimizing Memory Allocation" at Go Conference 2023 Online
   - GC logics: "Garbage Collection Semantics" at GopherCon SG 2019
   - Memory optimizations with pprof: "Memory Management in Go: The good, the bad and the ugly" at GopherCon UK 2023
   - (Advanced) Allocation of heap/stack (how Go allocation works with OS): "Understanding Go's Memory Allocator" at GopherCon UK 2018
   - Memory design over concurrency (memory model): "The Go Memory Model" at GoSF Meetup

## Speaker Biography

Takuto is an undergraduate student at Chiba Institute of Technology in Japan. He pursues a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He organizes Gophers EX, a project to support the overseas expansion of the Japan Go community. With his experience as a speaker at an international conference, he provides practical methods to challenge conferences overseas.
