# tcping-golang

## Description
This is a simple program to demonstrate how to determine your latency to another machine via TCP.

### Build the app
```
C:\Users\paradox\Desktop\Paping>go build main.go

C:\Users\paradox\Desktop\Paping>main.exe
Usage: ./main.exe <ip> <port> <count> (-1 if forever) <delay ex: 5s, 5m, 2s, 2h> (optional)
```

### Example
```
C:\Users\paradox\Desktop\Paping>main.exe 1.1.1.1 80 4
Connected to 1.1.1.1:80 time=28ms
Connected to 1.1.1.1:80 time=29ms
Connected to 1.1.1.1:80 time=28ms
Connected to 1.1.1.1:80 time=28ms
```
