# Actual Commands I ran
```
The default interactive shell is now zsh.
To update your account to use zsh, please run `chsh -s /bin/zsh`.
For more details, please visit https://support.apple.com/kb/HT208050.
(base) MacBook-Pro-4:GoStats-AI rohitmishra$ conda deactivate
MacBook-Pro-4:GoStats-AI rohitmishra$ mkdir anscombe-analysis
MacBook-Pro-4:GoStats-AI rohitmishra$ cd anscombe-analysis
MacBook-Pro-4:anscombe-analysis rohitmishra$ go mod init anscombe-analysis
go: creating new go.mod: module anscombe-analysis
MacBook-Pro-4:anscombe-analysis rohitmishra$ # Create the main package directory
mkdir anscombe

# Create the files inside the anscombe directoryMacBook-Pro-4:anscombe-analysis rohitmishra$ mkdir anscombe

# Create the files inside the anscombe directory
touch anscombe/main.go
touch anscombe/main_test.go
touch README.mdMacBook-Pro-4:anscombe-analysis rohitmishra$ 
MacBook-Pro-4:anscombe-analysis rohitmishra$ # Create the files inside the anscombe directory
MacBook-Pro-4:anscombe-analysis rohitmishra$ touch anscombe/main.go
MacBook-Pro-4:anscombe-analysis rohitmishra$ touch anscombe/main_test.go
MacBook-Pro-4:anscombe-analysis rohitmishra$ touch README.md
MacBook-Pro-4:anscombe-analysis rohitmishra$ touch main.go
MacBook-Pro-4:anscombe-analysis rohitmishra$ cd anscombe
MacBook-Pro-4:anscombe rohitmishra$ go test -v
=== RUN   TestLinearRegression
--- PASS: TestLinearRegression (0.00s)
=== RUN   TestRSquared
--- PASS: TestRSquared (0.00s)
PASS
ok      anscombe-analysis/anscombe      0.196s
MacBook-Pro-4:anscombe rohitmishra$ cd anscombe
bash: cd: anscombe: No such file or directory
MacBook-Pro-4:anscombe rohitmishra$ cd ..
MacBook-Pro-4:anscombe-analysis rohitmishra$ cd anscombe
MacBook-Pro-4:anscombe rohitmishra$ go test -bench=.
goos: darwin
goarch: arm64
pkg: anscombe-analysis/anscombe
cpu: Apple M2 Pro
BenchmarkLinearRegression-10            21445708                55.99 ns/op
BenchmarkRSquared-10                      983757              1193 ns/op
PASS
ok      anscombe-analysis/anscombe      3.496s
```