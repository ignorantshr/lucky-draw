# lucky draw
简单的抽奖系统。
技术栈：
- vue
- beego
- gorm
- mysql


## build
### UI
```
cd lucky-draw-ui
npm install
npm run build
```

### server
```
cd ..
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go build
```


## setup
### init database
```
mysql > source conf/sql/install.sql
```
### nginx
1. copy lucky-draw-ui/dist directory
2. add conf file according to `conf/nginx.conf`

## docker
### build image
```
docker build -t lucky-draw:v1.1 .
```

### run image
env:
- mysql_host: default 'localhost'
- mysql_port: a number, default 3306
- mysql_db: the database name, default lucky_draw
- mysql_user: default 'root'
- mysql_passwd: default ''

e.g.
```
docker run -d --name draw -p 8081:80 --env mysql_host="10.10.5.7" --env mysql_passwd="xxx" lucky-draw:v1.1
```
