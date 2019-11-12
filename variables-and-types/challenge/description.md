# Challenge description

#### Read a file from terminal and print out the number of lines it contains

##### Steps:
- Read the filename from terminal
- Check if the filename was provided or not empty
- Open the file using the filename
- Check if the file was successfully opened
- Create a scanner and read each line
- Store the number of lines into a variable on each iteration
- Print the variable to the screen

##### Run:
```bash
go build
./solution -filename="moving-water.txt"
```

##### Resources 📖

- https://golang.org/pkg/flag/
- https://godoc.org/flag
- https://golang.org/pkg/strings/#Trim
- https://golang.org/pkg/os/#Open
- https://golang.org/pkg/bufio/#NewScanner
- https://golang.org/pkg/bufio/#Scanner.Scan
