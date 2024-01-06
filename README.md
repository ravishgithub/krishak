# krishak
This is just a sample go app which has just two APIS login & check_auth
login API authenticate the user with user name and PWD and provides a toekn on successful authentication 
check_auth API expect a token in the header and valids if that's valid token

Assuming the Go server is running locally on http://localhost:8080, and the login endpoint is /login.

Sending a Valid Login Request:
```curl -X POST -H "Content-Type: application/json" -d '{"username":"admin","password":"admin123"}' http://localhost:8080/login```

Sending a Valid Login Request: windows machine example
```curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"admin123\"}" http://localhost:8080/login```
This curl command sends a POST request with the JSON payload containing the username and password to the /login endpoint. If the credentials are valid, the server should respond with a success message and a token (in this example, a placeholder token is used).

Sending an Invalid Login Request:
```curl -X POST -H "Content-Type: application/json" -d '{"username":"invaliduser","password":"wrongpassword"}' http://localhost:8080/login```
This command sends another POST request with incorrect credentials. If the credentials are invalid, the server should respond with an error message, indicating "Invalid username or password."

Sending valid auth request
```curl -X POST 'http://localhost:8080/check_auth' -H 'Authorization: example_token'```
