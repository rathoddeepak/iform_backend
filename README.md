# iForm | Golang Backend

Demo: [watch video](https://youtu.be/TsRTjGAMUrM) <br />
Postman: [documentation](https://rathoddeepak.github.io/iform](https://documenter.getpostman.com/view/5023754/2sA3Qza8Pe))


## Steps to run project

Step 1: Start development server with `docker-compose up --build`
  - Server will be running at port `http://localhost:8887/`

Step 2: To inspect database use adminer at `http://localhost:8887/`
  - Login with `user=postgres password=postgres dbname=iform_db`
  - Only for test, will be added to secrets later

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
