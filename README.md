# Websandbox

## backend
- Go: 1.16.3
- gin: 1.7.1
  - web framework
  - https://github.com/gin-gonic/gin
- gorm: 1.21.8
  - ORM library
  - https://gorm.io/ja_JP/docs/index.html
- gorm mysql driver: 1.0.5
  - https://gorm.io/ja_JP/docs/connecting_to_the_database.html
- docker: 20.10.5
- docker-compose: 1.29.0
  
### Development
```
// change directory to backend
$ cd backend

// setup local database server
$ docker-compose up

// start backend server
$ go run backend
```


## frontend
- npm: 6.14.5
- nodejs: 14.4.0
- vue-cli: 4.5.12
  - https://jp.vuejs.org/index.html
- eslint: 7.25.0

### Development
```
// change directory to frontend
$ cd frontend

// start frontend development server
$ npm run serve
```

