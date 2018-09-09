#Gollow  [![Go Report Card](https://goreportcard.com/badge/github.com/sourabh1024/gollow)](https://goreportcard.com/report/github.com/sourabh1024/gollow)

### Network resilient in-memory data distribution framework
#### Setup :
1. Install go 1.8.3
2. Set up GOPATH
3. cd scripts
4. docker-compose -f mysql.yml up : this will launch a docker image of mysql
5. export GOLLOW_CF="$GOPATH/src/github.com/gollow"
6. go run generate_dummy_data.go : this will generate some dummy data