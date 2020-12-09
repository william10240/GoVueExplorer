# GoVueExplorer

# 图片服务浏览器Go后端Vue前端

指定目录之后就可以方便的在浏览器上观看自己的图库啦,代码超级简单,可随意定制

### Server
```
go run app.go "your image folder"
如果不指定folder则默认为当前路径
```

### Client
```
npm run serve
```

### 修改
#### 后台端口
```
app.go -> http.ListenAndServe(":8000", nil)
```
#### 前台端口
```
App.vue -> serverUrl: "http://localhost:8000/?p=",
```
