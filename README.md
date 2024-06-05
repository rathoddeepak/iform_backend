# iForm | Golang Backend

Demo: [visit](https://rathoddeepak.github.io/iform).

## Steps to run project

Step 1: install go packages by running `go mod tidy`

Step 2: Start developmet server with `go run main.go -iform start`

### File Structure

 cmd -> Intial entry point to start application
 boot -> Connecting to database and starting server
 config -> Files related parsing .env config
 internal -> Interal files models and controllers
 pkg -> Files that can be accessed external and internally as helper functions
 routes -> routes for controllers
 main.go -> main file to start server

### Information

- Backend is written in golang to demonstarte use of statically 
  typed language for backend as well for better performance.
- Used gorm ORM with postgres database.
- Code is well structured.
