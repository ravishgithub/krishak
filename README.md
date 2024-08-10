# krishak
This is just a sample go app which has just two APIS login & check_auth
login API authenticate the user with user name and PWD and provides a toekn on successful authentication 
check_auth API expect a token in the header and valids if that's valid token

Project Structure
Krishak/
├── authentication/
│   ├── loginapi.go
│   └── checkauth.go
└── cmd/
    └── kheti/
        ├── main.go
        └── configs/
            └── config.json

Assuming the Go server is running locally on http://localhost:8080, and the login endpoint is /login.

Setting JWT Secret Linux
```export JWT_SECRET="ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc"```
Setting JWT Secret Windows
```set JWT_SECRET=ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc``` 

Sending a Valid Login Request:
```curl -X POST -H "Content-Type: application/json" -d '{"username":"admin","password":"ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc"}' http://localhost:8080/login```

Sending a Valid Login Request: windows machine example
```curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc\"}" http://localhost:8080/login```
This curl command sends a POST request with the JSON payload containing the username and password to the /login endpoint. If the credentials are valid, the server should respond with a success message and a token (in this example, a placeholder token is used).

Sending an Invalid Login Request:
```curl -X POST -H "Content-Type: application/json" -d '{"username":"invaliduser","password":"wrongpassword"}' http://localhost:8080/login```
This command sends another POST request with incorrect credentials. If the credentials are invalid, the server should respond with an error message, indicating "Invalid username or password."


Sending valid auth request (Lnix), please use toekn received from login not the sample 
```curl -X POST "http://localhost:8080/check_auth" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMzODAyNDksInVzZXJuYW1lIjoiYWRtaW4ifQ.xaGzzLBwflGKrmnxn1HdnRxdojgqcOMB85ZM9tqKkKM"```

Sending valid auth request (Windows), please use toekn received from login not the sample
```curl -X POST "http://localhost:8080/check_auth" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMzODAyNDksInVzZXJuYW1lIjoiYWRtaW4ifQ.xaGzzLBwflGKrmnxn1HdnRxdojgqcOMB85ZM9tqKkKM"```

Running inside docker container
```go build -o myapp krishak.tech/kheti/cmd/kheti

```docker build -t my-login-app .```
```docker run -p 8080:8080 my-login-app

