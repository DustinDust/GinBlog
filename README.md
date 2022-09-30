# GinBlog
## A Blog api written with gin &amp; gorm

### Quick start:
1. Clone the repo with ```git clone https://github.com/DustinDust/GinBlog.git```
1.1. Setup database (you can use the docker-compose file in the repo)
<br />
1.2. Setup `.env` for run time
2. Serve on localhost:9090 using [`air`](https://github.com/cosmtrek/air) </br> 
\* If you don't have [air](https://github.com/cosmtrek/air) installed you can try `go run main.go`
<br />
\* Make sure you have all the required packages installed. You can install or update with `go get -u <package-name>` 

### API Routes: <br/>

1. &nbsp; `/v1/user/me` | `GET` : Fetch current user (requires bearer token)
2. &nbsp; `/v1/blog-post?page=<int>` | `GET` : Fetch all blogposts (additional query support)
3. &nbsp; `/v1/blog-post/:id` | `GET` : Fetch blogpost by `ID`
4. &nbsp; `/v1/blog-post` | `POST` : Create new blogpost (require JSON body and bearer token) <br/>
5. &nbsp; `/v1/blog-post/:id` | `PUT` : Update blogpost by `ID` and JSON body (requires bearer token)
6. &nbsp; `/v1/blog-post/:id` | `DELETE` : Delete blogpost by `ID`
7. &nbsp; `/v1/tag?page=<int>` | `GET` : Fetch all Tags(additional query support)
... and more! <br/>
8. &nbsp; `/v1/tag/:id` | `GET` : Fetch one Tag by `ID`
9. &nbsp; `/v1/tag` | `POST` : Create new Tag (requires JSON body and bearer token)
10. &nbsp; `/v1/tag/:id` | `PUT` : Update Tag by `ID` and `json body` (requires bearer token)
11. &nbsp; `/v1/tag/:id` | `DELETE` : Delete token by `ID`
12. &nbsp; `/v1/auth/login` | `POST` : Login with username & password (requires JSON body)
13. &nbsp; `/v1/auth/register` | `POST` : Register a new user account (requires JSON body)
14. &nbsp; `/v1/auth/refresh` | `POST` : Refresh access token (requires refresh bearer token)
<br />
<br />

## Todo:
* Write testcases
* Cleaner code (in handling responses and errors)
* Cleaner architecture (folder structures, increase cohesion and looseness between modules)
* Dependency injection
