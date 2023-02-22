# Zocket Internship Assignment


## Problem Statement

TASK-2- Write a Go program to parse a CSV file: The intern should write a Go program that reads a CSV file and parses its contents into a data structure (e.g., a slice of structs). The program should then output the data in a formatted way (e.g., as a table).


## Solution

Implemented using both standard encoding/csv package (readCsvFile) as well as a custom one using strings package (myCsvReader).
(Sample csv also added)
## Usage

```go run main.go```

## Output

```
Enter file path: 
test.csv // User input
---------------------------------
school_name                                                                    | count |
School of Mechanical Engineering                                               | 6     |
needs confirmation                                                             | 6     |
School of Computer Engineering and Technology                                  | 6     |
School of Electronics and Communication Engineering                            | 5     |
School of Biology                                                              | 3     |
School of Pharmacy                                                             | 2     |
School of Commerce                                                             | 1     |
School of Computer Science                                                     | 1     |
School of Polytechnic - Applied Science/ Mechanical/ IT/ E&TC/ Computer/ Civil | 1     |
```






