Kirby
===
## Development Environment Setup
Start by installing Go and PostgreSQL on your system if they are not already available:
```
brew update && brew install go postgresql
```

Make sure the database server is running and create the development and test databases:
```
createdb kirby_development
createdb kirby_test
```

The provided `.env.development` file should work out of the box with the default values. If you would like to make changes to your environment (e.g. use a different database name), _DO NOT_ edit `.env.development`. Instead, use `.env.development.local` and/or `.env.test.local` to overwrite specific environment variables.

Next, install the module dependencies:
```
go mod download
```

Run all tests:
```
go test kirby/api/...
```

To run a specific test, provide the package and test name:
```
go test kirby/api/user -run TestAuthenticate
```

Start the server:
```
go run kirby.go
```

At this point, you should be able to start issuing requests to the server:
```
curl --request GET --url http://localhost:5000/health/server
{"success":true,"data":{"status":"ok","message":"Server is accepting connections"}}
```
```
curl --request GET --url http://localhost:5000/health/database
{"success":true,"data":{"status":"ok","message":"Database connection to 'kirby_development' succeeded"}}
```
```
curl --request POST \
  --url http://localhost:5000/users \
  --header 'content-type: application/json' \
  --data '{
        "name": "Example User",
        "email": "user@example.com",
        "password": "topsecret"
}'
{"success":true,"data":{"accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwibmFtZSI6IkV4YW1wbGUgVXNlciIsImVtYWlsIjoidXNlckBleGFtcGxlLmNvbSIsImV4cCI6MTU5NzM0NDU1NSwic3ViIjoidXNlckBleGFtcGxlLmNvbSJ9.XF8MXO6qZj2PqByg0a-MSI8O52BkEYgdl-NVX6_ajQc","refreshToken":"R4nBhewzhAV63sCkS2nIveLZ"}}
```

Remember that your host and port may vary if you're not using the default environment configuration.
