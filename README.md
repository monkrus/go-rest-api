<a href="https://github.com/monkrus"><img src="https://monosnap.com/image/nGg2FJacdl0L9qIgZSxBNtx3BibEBa" title="GoLang champion" alt="GoLang champion"></a>

# REST API 
We are creating a simple server to handle HTTP requests which helps us to provide communication between separate software components.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to start the project on a live system.

### Prerequisites
What things you need to install the software and how to install them
```
install GoLang
```

### Features
> Three distinct functions:

- HomePage function handles all requests to our root URL. 
- HandleRequests function  matches the URL path hit with a defined function. 
- Main function kicks off the API (entry point)

## Running

- We run `go run main.go` to test the basic functionality of those 3 functions. (Result: "Homepage Endpoint Hit" message )
- Create an API that returns a JSON response 
- Define an Article struct, set an Article array
- Register our function

# Example
- `$ go run main.go`  
- Return a JSON response under `http://localhost:8000/articles`. Success !
<img src="https://monosnap.com/image/j2YVHoLcSMRNLpejkWfcvJ88qQ4za9">

# Example (Optional)
![Refactor](recordgif.gif)
- Introducing gorilla/mux
- Add lines to handleRequests func
- Test Gorilla/mux,`$ go run main.go`
```go 
 myRouter.HandleFunc("/articles", allArticles).Methods("GET")
 myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
``` 
- Run `$ go run main.go` and compare following lines in Postman
 
 <img src="https://monosnap.com/image/Po8V4upwAgdt7AIEAoQLCjQidMB3pM">
 
## FAQ 

- **How do I do *specifically* so and so?** 
    - No problem! Just do this.
 



