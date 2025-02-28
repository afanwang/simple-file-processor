# Welcome to Simple File Processor
This simple Golang program read the CSV file in the folder, where column data is ordered: `timestamp`, `username`, `operation`, `size`, and output several simple analytical use cases:
1. How many users accessed the server?
2. How many uploads were larger than `50kB`?
3. How many times did `jeff22` upload to the server on April 15th, 2020?

I included a small custom logger in the program to have a better logger than `fmt`. The result I got:
```
DEBUG: 2023/03/25 23:21:28 main.go:31: Starting A Coding Challenge
DEBUG: 2023/03/25 23:21:28 main.go:90: Num of Users Accessed The Server: 6
DEBUG: 2023/03/25 23:21:28 main.go:91: Num of Upload > 50KB: 140
DEBUG: 2023/03/25 23:21:28 main.go:92: Num of Upload to server by `jeff22` on April 15th, 2020: 3
```

## Future Improvements
Apparently the written code snippet is far from production ready. 
1. The hard-coded metrics are hard-coded, it will be better if they can be passed in by a config file. So every time when Admin wants a different metric he does not need to change the source code.
2. Unit-test is needed to test all the combinations of the metrics, driven from a config file.
3. If the file size is greater than the machine's memory, for example I am using 16GB mem macbook and if the file size is 32GB. Using `csv.NewReader(bufio.NewReader(fileName)`, Go's `csv` package automatically read the file in chunks using the underlying `bufio.Scanner`. `bufio.Scanner` reads a file in small chunks of data, buffering the data in memory until it reaches a newline character or a specified delimiter. At that point, it returns the buffered data as a line or field, and continues scanning from where it left off.
4. Even when we have no issue with processing a large CSV file, the processing speed may be a concern.   Possible improvements could be:
* Using `fastcsv` package, which is faster than the standard `csv` package and can handle large files more efficiently.
* Parallelization: If you have a multi-core CPU, you can parallelize the process by dividing the CSV file into smaller chunks and processing each chunk in parallel. This can be done using Go's goroutines and channels to distribute the work across multiple cores.
* A more advanced technique will be devide and conquer, which divides the large CSV file into multiple chunks and processes each chunk in different processors, after the processing of each chunk aggregates the results.

