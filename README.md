# 🏆 Sort Tournament

![Localisation](https://img.shields.io/badge/Made_in-France-red?labelColor=blue)
![Language](https://img.shields.io/badge/Language-Go-f7d3a2?labelColor=00aed8)

Sort Tournament is a project that aims to compare the performance of different sorting algorithms.

## 📋 Summary

### 1. [Features](#1---features)

### 2. [Installation](#2---installation)

### 3. [How to use](#3---how-to-use)

### 4. [Algorithms](#4---algorithms)

### 5. [Benchmark](#5---benchmark)

### 6. [Improve the project](#6---improve-the-project)

### 7. [Credits](#7---credits)

## 1 - Features

- Compare the performance of different sorting algorithms among 26.
- Test the algorithms on a list of 1M integers in random order, roughly sorted, nerly sorted, already sorted and reversed order.
- Export the results into a CSV file.

> **Note**: Roughly sorted means that the list is globally sorted but each block of 5 elements is shuffled meawhile nearly sorted means that the list is sorted but 10% of the elements are shuffled.

## 2 - Installation

If you want to test the algorithms on your own computer, you first need to have Go installed on your computer.

Then you can download the project as a zip file **[here](https://github.com/Pietot/Sort-Tournament/archive/refs/heads/main.zip)** and extract it.

Or clone the repository:

```bash
git@github.com:Pietot/Sort-Tournament.git
```

## 3 - How to use

To use the tool, you just need to open a terminal in the project directory and run the following command:

```go
go run .
```

If you want to change the number of elements in the lists, go on line 57 in main.go and change the value of the variable `n`.

## 4 - Algorithms

Here's a quick overview of the algorithms used in the project:

- Binary Insertion Sort
- Block Sort
- Bubble Sort (Optimized)
- Bucket Sort
- Circle Sort
- Comb Sort
- Counting Sort (Stable)
- Counting Sort (Unstable)
- Double Selection Sort
- Gnome Sort
- Heap Sort
- Insertion Sort
- Intro Sort
- Merge Sort
- Odd-Even Sort
- Pattern-Defeating Quick Sort (default sort in Go)
- Quick Sort (Middle Pivot)
- Radix LSD Sort (Base 64)
- Red Black Tree Sort
- Selection Sort (Optimized)
- Shaker Sort
- Shell Sort
- Smooth Sort
- Tim Sort
- Tournament Sort
- Tree Sort

## 5 - Benchmark

Here are the algorithms ranked from the fastest to the slowest by average time:

| Rank | Algorithm                                             | Average time | Random array | Roughly sorted array | Nearly sorted array | Sorted array | Reversed array |
| :--: | :---------------------------------------------------- | :----------: | :----------: | :------------------: | :-----------------: | :----------: | :------------: |
|  1   | **Counting Sort (unstable)**                          |    ~ 4 ms    |    ~ 5 ms    |        ~ 4 ms        |       ~ 5 ms        |    ~ 4 ms    |     ~ 4 ms     |
|  2   | **Counting Sort (stable)**                            |   ~ 11 ms    |   ~ 17 ms    |        ~ 8 ms        |       ~ 14 ms       |    ~ 8 ms    |     ~ 6 ms     |
|  3   | **Bucket Sort**                                       |   ~ 28 ms    |   ~ 38 ms    |       ~ 27 ms        |       ~ 26 ms       |   ~ 18 ms    |    ~ 31 ms     |
|  4   | **Block Sort**                                        |   ~ 29 ms    |   ~ 30 ms    |        ~ 30ms        |       ~ 29 ms       |   ~ 29 ms    |    ~ 29 ms     |
|  5   | **Pattern-Defeating Quick Sort (default sort in Go)** |   ~ 32 ms    |   ~ 91 ms    |       ~ 30 ms        |       ~ 37 ms       |    ~ 1 ms    |     ~ 2 ms     |
|  6   | **Intro Sort**                                        |   ~ 33 ms    |   ~ 88 ms    |       ~ 21 ms        |       ~ 30 ms       |   ~ 12 ms    |    ~ 13 ms     |
|  7   | **Quick Sort (middle pivot)**                         |   ~ 45 ms    |   ~ 85 ms    |       ~ 36 ms        |       ~ 60 ms       |   ~ 21 ms    |    ~ 24 ms     |
|  8   | **Radix LSD Sort (base 64)**                          |   ~ 67 ms    |   ~ 68 ms    |       ~ 67 ms        |       ~ 67 ms       |   ~ 67 ms    |    ~ 68 ms     |
|  9   | **Comb Sort**                                         |   ~ 70 ms    |   ~ 120 ms   |       ~ 41 ms        |      ~ 111 ms       |   ~ 36 ms    |    ~ 42 ms     |
|  10  | **Tim Sort**                                          |   ~ 71 ms    |   ~ 126 ms   |       ~ 52 ms        |       ~ 66 ms       |   ~ 45 ms    |    ~ 66 ms     |
|  11  | **Shell Sort**                                        |   ~ 77 ms    |   ~ 162 ms   |       ~ 27 ms        |      ~ 146 ms       |   ~ 20 ms    |    ~ 29 ms     |
|  12  | **Smooth Sort**                                       |   ~ 78 ms    |   ~ 123 ms   |       ~ 64 ms        |       ~ 77 ms       |   ~ 61 ms    |    ~ 66 ms     |
|  13  | **Merge Sort**                                        |   ~ 107 ms   |   ~ 167 ms   |       ~ 90 ms        |      ~ 113 ms       |   ~ 82 ms    |    ~ 83 ms     |
|  14  | **Tournament Sort**                                   |   ~ 180 ms   |   ~ 293 ms   |       ~ 147 ms       |      ~ 176 ms       |   ~ 144 ms   |    ~ 141 ms    |
|  15  | **Heap Sort**                                         |   ~ 282 ms   |   ~ 347 ms   |       ~ 242 ms       |      ~ 269 ms       |   ~ 237 ms   |    ~ 316 ms    |
|  16  | **Red Black Tree Sort**                               |   ~ 301 ms   |   ~ 589 ms   |       ~ 195 ms       |      ~ 321 ms       |   ~ 205 ms   |    ~ 195 ms    |
|  17  | **Binary Insertion Sort**                             |   ~ 1 min    | ~ 1 min 30s  |       ~ 40 ms        |        ~ 17s        |   ~ 36 ms    |  ~ 3 min 30s   |
|  18  | **Insertion Sort**                                    | ~ 1 min 30s  | ~ 2 min 30s  |        ~ 8 ms        |        ~ 36s        |    ~ 1 ms    |    ~ 2 min     |
|  19  | **Circle Sort**                                       | ~ 2 min 30s  | ~ 2 min 30s  |     ~ 2 min 30s      |     ~ 2 min 30s     | ~ 2 min 30s  |  ~ 2 min 30s   |
|  20  | **Double Selection Sort**                             |   ~ 7 min    |   ~ 7 min    |       ~ 7 min        |       ~ 7 min       |   ~ 7 min    |    ~ 7 min     |
|  21  | **Odd-Even Sort**                                     |  ~ 7 min 30  |   ~ 15 min   |        ~ 9 ms        |      ~ 14 min       |    ~ 0 ms    |  ~ 7 min 30s   |
|  22  | **Shaker Sort**                                       |   ~ 8 min    |   ~ 19 min   |       ~ 10 ms        |     ~ 2 min 30s     |    ~ 1 ms    |  ~ 16 min 30s  |
|  23  | **Selection Sort (optimized)**                        |   ~ 8 min    |   ~ 8 min    |       ~ 8 min        |       ~ 8 min       |   ~ 8 min    |    ~ 7 min     |
|  24  | **Gnome Sort**                                        |   ~ 8 min    | ~ 12 min 30s |       ~ 10 ms        |       ~ 3 min       |    ~ 1 ms    |    ~ 25 min    |
|  25  | **Bubble Sort (optimized)**                           |   ~ 10 min   | ~ 26 min 30s |       ~ 10 ms        |    ~ 10 minutes     |    ~ 1 ms    |  ~ 16 min 30s  |
|  26  | **Tree Sort**                                         |   ~ 30 min   |   ~ 469 ms   |       ~ 25 min       |      ~ 491 ms       | ~ 1 h 3 min  |     ~ 1 h      |
<p align="center">
    <a href="benchmark.csv">Download csv here</a>
</p>

> **Note**: For each times, the array contained 1 million element from 1 to 1 000 000.

## 6 - Improve the project

If you like this project and/or want to help or improve it, you can:

- Create an issue if you find a bug or want to suggest a feature or any improvement (no matter how small it is).

- Create a pull request if you want to add a feature, fix a bug or improve the code.

- Contact me if you want to talk about the project or anything else (Discord: pietot).

# 7 - Credits

- **[Original video](https://www.youtube.com/watch?v=N4JVT3eVBP8)**: The video that inspired me to create this project.
- **[Explanatory video](https://www.youtube.com/watch?v=h1Bi0granxM)** The video that's helped me to understand most of the algorithms.
