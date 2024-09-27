package main

import (
	"encoding/json"
	"log"

	"github.com/tren03/artechles-ReactTS/backend/global"
)
 
// THIS CODE WAS ONLY MEANT TO RUN ONCE TO MIGRATE DATA IN CLOUD POSTGRES INSTANCE TO MY LOCAL DOCKER INSTANCE

func Convert() []global.Post {
	bigData := `[
  {
    "ID": 54,
    "Title": "Go By example",
    "Url": "https://gobyexample.com/",
    "Body": "A site that is a good primer for all of the main concepts covered in go language with code snippets for each.",
    "Date": "2024-09-21T16:22:37.917988Z",
    "Category":"Programming Languages & Tools"
  },
  {
    "ID": 53,
    "Title": "DHH Is Right About Everything",
    "Url": "https://www.youtube.com/watch?v=mTa2d3OLXhg",
    "Body": "David Heinemeier Hansson, The creator of Ruby on Rails, shares insights on why he chose Ruby as a programming language, emphasizing the importance of side projects to keep motivation alive. Prime, Tj and DHH touch upon  evolution of open-source, game programming, and impact of AI on the current generation of junior devs.",
    "Date": "2024-09-20T10:00:47.646591Z",
    "Category": "Tech Culture & Career Development"
  },
  {
    "ID": 52,
    "Title": "Pratt Parsers: Expression Parsing Made Easy",
    "Url": "https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/",
    "Body": "A much more hands on article talking about the actual implementation of a pratt parser for a dummy language.",
    "Date": "2024-09-17T14:10:43.40191Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 51,
    "Title": "Simple but Powerful Pratt Parsing",
    "Url": "https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html",
    "Body": "Helped me a bit in understanding more about pratt parsers and recursive descent parsers  for my go interpreter.",
    "Date": "2024-09-17T07:38:55.041995Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 50,
    "Title": "A fully compliant RISC-V computer inside Terraria",
    "Url": "https://youtu.be/zXPiqk0-zDY",
    "Body": "This project emulates a A fully compliant RISC-V computer inside Terraria.\nThe amount of engineering that went into this is just mind blowing. Check out the Youtube video \nUrl : https://youtu.be/zXPiqk0-zDY",
    "Date": "2024-09-16T08:50:31.438043Z",
    "Category": "Cool Projects & Side Explorations"
  },
  {
    "ID": 49,
    "Title": "The Primeagen on \"Developer Excellence\"",
    "Url": "https://www.youtube.com/watch?v=96VlfN7ViyE",
    "Body": "Truly inspirational stuff. \nLove @ThePrimeagen ~         \n\n\"Always bet on yourself\" ",
    "Date": "2024-09-11T08:29:21.643739Z",
    "Category": "Tech Culture & Career Development"
  },
  {
    "ID": 47,
    "Title": "Classic SysAdmin: The Linux Filesystem Explained",
    "Url": "https://www.linuxfoundation.org/blog/blog/classic-sysadmin-the-linux-filesystem-explained",
    "Body": "This article gives an insight into the folders in the root directory, and takes us through what every folder contains, and how the structure has evolved from the early days of linux.",
    "Date": "2024-09-09T21:45:08.733622Z",
    "Category": "Operating Systems & Low-Level Programming"
  },
  {
    "ID": 43,
    "Title": "Tech Book Reviews: Go",
    "Url": "https://ssmertin.com/articles/tech-book-reviews-go/",
    "Body": "The article highlights key texts like \"The Go Programming Language\" by Donovan & Kernighan, which is recommended for beginners, and \"Concurrency in Go\" by Katherine Cox-Buday, ideal for those seeking a deep dive into Go's concurrency features. \"Writing an interpreter in Go\" by Thorsten Ball is a gem.",
    "Date": "2024-09-09T11:26:23.192172Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 42,
    "Title": "Modern JavaScript Explained for Dinosaurs",
    "Url": "https://peterxjang.com/blog/modern-javascript-explained-for-dinosaurs.html",
    "Body": "This article takes you through the evolution of JavaScript from the early days to modern practices, explaining the need for tools like npm, Babel, and Webpack. This really helped me understand what is going on under the hood and the process of building projects.",
    "Date": "2024-09-07T11:03:33.457574Z",
    "Category": "Tech Culture & Career Development"
  },
  {
    "ID": 41,
    "Title": "Linux Kernel Module Programming Guide (LKMPG)",
    "Url": "https://sysprog21.github.io/lkmpg/",
    "Body": "This guide is for learning how to develop modules for the Linux kernel. It breaks down complex concepts into easy-to-understand examples, making kernel programming more accessible. If you’re interested in Linux development or want to extend kernel functionality, this is a solid resource to start with.\n\n(Would be a fun experience to write and maintain one, maybe one day lmao)",
    "Date": "2024-09-07T11:00:19.274717Z",
    "Category": "Operating Systems & Low-Level Programming"

  },
  {
    "ID": 39,
    "Title": "Cryptopals",
    "Url": "https://topals.com/",
    "Body": "Cryptopals is a set of tography challenges that help you get hands-on with encryption techniques. These challenges start easy and get progressively harder, covering topics like AES, hashing, and public key crypto.\n\n(I kinda look it as a way to explore the language you are using, and try new libs)",
    "Date": "2024-09-07T10:55:37.670913Z",
    "Category": "Security & Cryptography"
  },
  {
    "ID": 38,
    "Title": "Learn Git Branching",
    "Url": "https://learngitbranching.js.org/",
    "Body": "This interactive tool is a fun and visual way to learn Git. It walks you through the basics and advanced features of Git branching, merging, and rebasing, making version control a lot easier to grasp. Great for anyone who wants to master Git and better manage their code.\n\n(Bro I still do not understand how git works L)",
    "Date": "2024-09-07T10:53:59.061674Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 37,
    "Title": "MIT 6.828 OS Engineering (xv6 Book)",
    "Url": "https://github.com/fanweng/MIT-6.828-OS-Engineering/blob/master/resources/xv6-book-rev11.pdf",
    "Body": "This book is an excellent resource for learning about operating system design and implementation, based on MIT’s xv6 operating system course. It provides a deep dive into how OSes function, covering essential topics like memory management, scheduling, and file systems, making it ideal for anyone studying or working on operating systems.",
    "Date": "2024-09-07T10:50:58.557631Z",
    "Category": "Operating Systems & Low-Level Programming"
  },
  {
    "ID": 36,
    "Title": "Protohackers – Server-Side Scripting Challenges",
    "Url": "https://protohackers.com/problems",
    "Body": "Protohackers is packed with cool challenges that really put your server-side coding skills to the test. Each problem pushes you to build servers that handle networking, concurrency, and different protocols, so it’s perfect if you want to sharpen your skills in building efficient, reliable systems while having some fun along the way.",
    "Date": "2024-09-07T10:50:11.074148Z",
    "Category": "Security & Cryptography"
  },
  {
    "ID": 35,
    "Title": "Nand2Tetris",
    "Url": "https://www.nand2tetris.org/",
    "Body": "This course walks you through building a complete computer from the ground up, starting with basic logic gates. It’s a hands-on experience that covers everything from hardware to software, teaching you how to design and implement a computer system, including building a CPU, memory, and an operating system, all the way to running your own programs.",
    "Date": "2024-09-07T10:49:04.821886Z",
    "Category": "Operating Systems & Low-Level Programming"
  },
  {
    "ID": 34,
    "Title": "Awesome Compilers",
    "Url": "https://github.com/aalhour/awesome-compilers",
    "Body": "A curated list of awesome resources, learning materials, tools, frameworks, platforms, technologies and source code projects in the field of Compilers, Interpreters and Runtimes. This list has a bias towards education.\n\n(Check out this rep, too cool)",
    "Date": "2024-09-07T10:47:55.820498Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 33,
    "Title": "Ray Tracing in One Weekend",
    "Url": "https://github.com/RayTracing/raytracing.github.io",
    "Body": "This series of books focuses on teaching ray tracing, a core graphics rendering technique, through a practical, hands-on approach. The project walks you through building your own ray tracer from scratch, making complex concepts like light reflection and image rendering accessible. It’s a great resource for those interested in graphics programming and understanding how computer-generated images are created",
    "Date": "2024-09-07T10:44:23.550443Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 32,
    "Title": "Mastering Embedded Linux Programming (Third Edition)",
    "Url": "https://github.com/PacktPublishing/Mastering-Embedded-Linux-Programming-Third-Edition",
    "Body": "This is the companion repo to the \"Mastering Embedded Linux Programming\" book, filled with practical exercises and code to help you get deep into embedded Linux. You’ll find examples that deal with cross-compiling, real-time processing, and more. A great find if you’re into working with embedded systems.",
    "Date": "2024-09-07T10:42:32.391372Z",
    "Category": "Operating Systems & Low-Level Programming"
  },
  {
    "ID": 31,
    "Title": "LLMs from Scratch",
    "Url": "https://github.com/rasbt/LLMs-from-scratch",
    "Body": "If you're curious about how Large Language Models (like GPT) are built, this GitHub repo breaks it down step-by-step. It dives into the theory behind these models and provides code examples to help you create your own. A must-check if you want to really understand how LLMs work from the ground up.",
    "Date": "2024-09-07T10:41:18.984278Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 29,
    "Title": "Learn X in Y Minutes",
    "Url": "https://learnxinyminutes.com",
    "Body": "This site is great if you want to get a quick handle on a new programming language. Each guide walks you through the language's syntax and key concepts using real code examples. It covers tons of languages—from Python to Rust—making it perfect when you need to get up to speed fast.",
    "Date": "2024-09-07T10:26:20.284313Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 28,
    "Title": "Advent of Code",
    "Url": "https://adventofcode.com",
    "Body": "Advent of Code is a popular annual coding event that presents daily programming puzzles throughout December. These challenges are designed to be fun yet stimulating, encouraging participants to either explore a new programming language or deepen their skills in a language they already know. The puzzles are progressively difficult and often build on previous days' problems, making it an engaging experience for both hobbyist coders and experienced developers alike. Advent of Code has become a tradition for many programmers to sharpen their problem-solving skills in a festive setting.\n\n(Learnt go for building this website by doing a the first few challenges, would 100% recommend doing this for learning new languages. Looking forward to doing it the December)",
    "Date": "2024-09-07T10:23:16.442759Z",
    "Category": "Programming Languages & Tools"
  },
  {
    "ID": 27,
    "Title": "OverTheWire: Wargames",
    "Url": "https://overthewire.org/wargames/",
    "Body": "OverTheWire's Wargames offer a unique series of online security challenges designed to teach concepts in cybersecurity, networking, and system administration. Each wargame presents increasingly complex tasks, helping participants develop skills in exploiting vulnerabilities, reverse engineering, and penetration testing. It’s a hands-on approach to learning, and helped me learn cool stuff about low level concepts.",
    "Date": "2024-09-07T10:18:14.955692Z",
    "Category": "Security & Cryptography"
  },
  {
    "ID": 26,
    "Title": "Unusual Boot Device for This Arch Linux System: Google Drive",
    "Url": "https://www.hackster.io/news/ersei-picks-an-unusual-boot-device-for-this-arch-linux-system-google-drive-1145a0b176a6",
    "Body": "Ersei takes an unconventional approach to booting Arch Linux by using Google Drive as the storage medium. The article delves into the technical challenges faced, detailing how they managed to use Google Drive for booting without relying on a local disk or USB drive. This project serves as an intriguing proof of concept for cloud-based storage solutions, blending creativity and technical ingenuity to challenge traditional methods of system booting",
    "Date": "2024-09-07T10:16:54.809607Z",
    "Category": "Cool Projects & Side Explorations"
  },
  {
    "ID": 25,
    "Title": "Linux From Scratch",
    "Url": "https://www.linuxfromscratch.org/lfs/",
    "Body": "\"Linux From Scratch\" (LFS) is a comprehensive guide for building a custom Linux system entirely from source code. The project encourages users to compile each component manually, allowing them to learn the intricacies of the Linux operating system at every step. By following LFS, users gain deeper knowledge of Linux internals, system structure, and installation, along with skills to troubleshoot and optimize Linux setups. The guide is aimed at developers and enthusiasts who want to tailor-make their Linux system while learning in the process. (This has been on my todo lost for so long ugh)",
    "Date": "2024-09-07T09:40:30.632134Z",
    "Category": "Operating Systems & Low-Level Programming"
  },
  {
    "ID": 24,
    "Title": "ARP Explained",
    "Url": "https://kognise.dev/writing/arp",
    "Body": "This article breaks down the Address Resolution Protocol (ARP), essential in converting IP addresses to MAC addresses on a local network. It covers how ARP works, examples of its practical use, and various vulnerabilities such as ARP spoofing. The post highlights real-world applications, as well as potential security risks, providing insights into mitigation strategies to protect against malicious activity in networked environments. It is a clear, technical guide aimed at networking enthusiasts or professionals.",
    "Date": "2024-09-07T09:39:10.002563Z",
    "Category": "Operating Systems & Low-Level Programming"


  },
  {
    "ID": 23,
    "Title": "Writing an Editor in Less Than 1000 Lines of Code",
    "Url": "http://antirez.com/news/108",
    "Body": "In this blog post, Salvatore Sanfilippo (antirez) describes how he created a simple text editor called \"Kilo\" in under 1000 lines of code, featuring syntax highlighting and search functionality. He shares his inspiration, the coding process, and the nostalgia behind this fun weekend project. The full code is available on GitHub, and the post reflects on the joy of working on small, seemingly purposeless projects",
    "Date": "2024-09-07T09:36:31.304801Z",
    "Category": "Cool Projects & Side Explorations"
  },
  {
    "ID": 22,
    "Title": "Putting the \"You\" in CPU",
    "Url": "https://cpu.land",
    "Body": "The article on CPU.land explores what happens when you run a program on your computer, focusing on key concepts like multiprocessing, system calls, memory management, and executable loading in Linux. It walks through the entire process, from program execution to low-level details such as hardware interrupts, paging, and the exec system call. The goal is to demystify how programs interact with the CPU and operating system, providing an accessible yet comprehensive resource for those curious about computer internals.",
    "Date": "2024-09-07T09:27:23.648887Z",
    "Category": "Operating Systems & Low-Level Programming"
  }
]`

	var posts_to_add []global.Post

	bigDataBytes := []byte(bigData)
	err := json.Unmarshal(bigDataBytes, &posts_to_add)
	if err != nil {
		println(err)
	}
	log.Println("success")
	return posts_to_add

}
