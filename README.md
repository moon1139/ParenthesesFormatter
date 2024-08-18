## Run by online IDE
- Navigate to \ParenthesesFormatter\ForOnlineIDE
- Simply copy whole chunk of code
- Paste it to https://nextleap.app/online-compiler/go-programming
- Fill our input and run
- You can find sample input in the input.txt file

## Run by go command
- Install go first https://go.dev/doc/install
- Navigate to /ParenthesesFormatter
- `go run .` 

## Run by executable file
- Please build the executable file first if you want to use this method
- For windows environment, use powershell or command line `.\ParenthesesFormatter.exe` to run.
- For Mac you should able to run by `./hello` once executable is built
- ParenthesesFormatter will take stdin as input and output to stdout (need to ctrl+c to stop then you can see stdout)
- Support mutiple line input
- For example,
#### input file: 
```
(A*(B+C))
1*(2+(3*(4+5)))
2+(3/-5)
```
#### output file: 
```
A*(B+C)
1*(2+3*(4+5))
2+3/-5
```

- Also if you're using powershell. There will also console messages for human.
#### console output:
```
(A*(B+C)) => A*(B+C)
1*(2+(3*(4+5))) => 1*(2+3*(4+5))
2+(3/-5) => 2+3/-5
```

## Build executable file
- We need to install go first https://go.dev/doc/install
- Change to directory \ParenthesesFormatter
- run `go build ./...` to build whole project. This will create/renew our [ParenthesesFormatter.exe](ParenthesesFormatter.exe) or other executables.

## Run tests via VScode
- After we've installed GO and VScode
- Install extension [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.go)
- In the file [formatter_test.go](formatter_test.go), we should just able to use VScode's integrated testing function

## Some constraints about the function
- In this function, I am assuming all of the input would contains valid parentheses pairs.
- That is we will see something like `(())` but not `(()`.
- Since if we want to auto correct invalid pairs like `(a*(b+c)`.
- The result could be `a*(b+c)` or `(a*b+c)`, causing our `ParenthesesFormatter` will have different output.