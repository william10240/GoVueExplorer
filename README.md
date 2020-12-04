# GoVueExplorer

### Server
```
go run app.go "your image folder"
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
