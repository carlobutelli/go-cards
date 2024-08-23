# GO

function signature: func <receiver> funcName(<args>) <return value(s)> {}
  - receiver is what can the function work on
  - args are the input parameters passed to the function
  - return value(s) is the returning value (it can be skipped meaning the func does not return any value)
  
Use:
	- new() when dealing with value types like structs to allocate memory for a new zeroed value.
	  This is suitable for scenarios where you want a pointer to an initialized structure.
	  new() returns a pointer
	- make() to create an initialized instance for slices, maps, and channels, where initialization involves setting up data structures
	  and internal pointers
	  make() returns a non-zeroed value

## Test
- create a new file ending in _test.go
- To run all tests in a package run go test
  - remember, you need to run ´´´go mod init <packageName>´´´ first