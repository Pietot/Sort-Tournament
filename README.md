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

- Compare the performance of different sorting algorithms among 23.
- Test the algorithms on a list of 1M integers in random order, roughly sorted, nerly sorted and reversed order.
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

If you want to change the number of elements in the lists, go on line 52 in main.go and change the value of the variable `n`.

## 4 - Algorithms

Here's a quick overview of the algorithms used in the project:

- Binary Insertion Sort
- Block Sort
- Bubble Sort (Optimized)
- Bucket Sort
- Circle Sort
- Comb Sort
- Counting Sort
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
- Selection Sort (Optimized)
- Shaker Sort
- Shell Sort
- Smooth Sort
- Tim Sort
- Tournament Sort
- Tree Sort

## 5 - Benchmark

Here are the algorithms ranked from the fastest to the slowest by sorting the same list of 1M integers in random order:

| Rank | Algorithm                                             | Time (ms) |
| :--: | :---------------------------------------------------- | :-------: |
|  1   | **Counting Sort**                                     |   ~ 24    |
|  2   | **Bucket Sort**                                       |   ~ 43    |
|  3   | **Block Sort**                                        |   ~ 44    |
|  4   | **Radix LSD Sort (Base 64)**                          |   ~ 69    |
|  5   | **Quick Sort (Middle Pivot)**                         |   ~ 89    |
|  6   | **Intro Sort**                                        |   ~ 93    |
|  7   | **Pattern-Defeating Quick Sort (default sort in Go)** |   ~ 98    |
|  8   | **Smooth Sort**                                       |   ~ 138   |
|  9   | **Comb Sort**                                         |   ~ 142   |
|  10  | **Tim Sort**                                          |   ~ 144   |
|  11  | **Shell Sort**                                        |   ~ 173   |
|  12  | **Merge Sort**                                        |   ~ 184   |
|  13  | **Tournament Sort**                                   |   ~ 311   |
|  14  | **Heap Sort**                                         |   ~ 375   |
|  15  | **Tree Sort**                                         |   ~ 690   |
|  16  | **Binary Insertion Sort**                             | ~ 102640  |
|  17  | **Circle Sort**                                       | ~ 237136  |
|  18  | **Insertion Sort**                                    | ~ 241254  |
|  19  | **Double Selection Sort**                             | ~ 421952  |
|  20  | **Selection Sort (Optimized)**                        | ~ 592139  |
|  21  | **Gnome Sort**                                        | ~ 801793  |
|  22  | **Odd-Even Sort**                                     | ~ 980085  |
|  23  | **Shaker Sort**                                       | ~ 1202780 |
|  24  | **Bubble Sort (Optimized)**                           | ~ 1748075 |

## 6 - Improve the project

If you like this project and/or want to help or improve it, you can:

- Create an issue if you find a bug or want to suggest a feature or any improvement (no matter how small it is).

- Create a pull request if you want to add a feature, fix a bug or improve the code.

- Contact me if you want to talk about the project or anything else (Discord: pietot).

# 7 - Credits

- **[Original video](https://www.youtube.com/watch?v=N4JVT3eVBP8)**: The video that inspired me to create this project.
- **[Explanatory video](https://www.youtube.com/watch?v=h1Bi0granxM)** The video that's helped me to understand most of the algorithms.
