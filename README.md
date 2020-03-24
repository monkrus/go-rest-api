# go_rest_api

# REST API helps us to provide communication between separate software components.
# We are creating a simple server to handle HTTP requests.

- Three distinct functions:

1. HomePage function handles all requests to our root URL. 

2. HandleRequests function  matches the URL path hit with a defined function. 

3. Main function kicks off the API (entry point)

4. We run `go run main.go` to test the basic functionality of those 3 functions. (Result: "Homepage Endpoint Hit" message )

5. create a fisrs API that returns a JSON response 

6. Define an Article struct, set an Article array

7. Register our function

8. We run `go run main.go`  and return a JSON response under `http://localhost:8000/articles`. Success !

9. 

