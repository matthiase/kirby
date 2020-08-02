## Development Setup
```
# Install PostgreSQL and create the database
brew update && brew install postgresql
createdb kirby

# Create an .env file by copying the provided example
cp example.env .env

# Start the server
go run main.go
```
