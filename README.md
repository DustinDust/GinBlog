# GinBlog
A Blog api written with gin &amp; gorm

Quick start:
1. Clone the repo with ```git clone https://github.com/DustinDust/GinBlog.git```
2. Install dependencies with ```go run build```
3. Server on localhost:9090 using [`air`](https://github.com/cosmtrek/air) </br> 
\* If you don't have [air](https://github.com/cosmtrek/air) installed you can try `go run main.go`

API Routes: <br/>

1. `/v1/user/me` method `GET`: fetch current user (required bearer token)
2. `/v1/blog-post` method `GET`: fetch all blogposts
3. `/v1/blog-post/:id` method `GET`: fetch blogpost by `ID`
4. `/v1/blog-post` method `POST`: create new blogpost (required json body and bearer token) <br/>
... and more! <br/>

(I'm lazy)
