# Go Lang Learning Plan 

- General / Core GoLang (struct, packages, main, some logging)
- How to write Rest APIs (GET, POST apis)
- Load up Yaml data from a string array (later this will be loaded from S3 folder potentially)
- Goroutine - load the data asynchronously

## Learning list:
- Writing tests
- handling errors (panics)
- handling dbs
- writing plugs (mostly interfaces)
- triggering functions (either using reflections, or dont know exactly) …

## Generic List 
- How to debug Go code in VS Code --> https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code#performing-further-debugging-tests

- How to kill a hanging process with blocking port number 
`To list any process listening to the port 8080:lsof -i:8080`
`To kill any process listening to the port 8080: kill $(lsof -t -i:8080)`
`or more violently: kill -9 $(lsof -t -i:8080)`

