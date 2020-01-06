# logrus-demo

Sample logging demo using [logrus](https://github.com/sirupsen/logrus)

## Output
```
{
  "cluster": "mycluster",
  "data": "mydata",
  "file": "/Users/vikram/Projects/GO_HOME/src/github.com/vikramjakhr/logrus-demo/main.go:46",
  "func": "main.main",
  "function": "CreateProject",
  "level": "info",
  "msg": "A sample info message",
  "package": "api",
  "project": "test-project",
  "time": "2020-01-06T13:49:17+05:30"
}
{
  "cluster": "mycluster",
  "data": "mydata",
  "file": "/Users/vikram/Projects/GO_HOME/src/github.com/vikramjakhr/logrus-demo/main.go:53",
  "func": "main.main",
  "function": "CreateProject",
  "level": "warning",
  "msg": "A sample warning message",
  "package": "api",
  "project": "test-project",
  "time": "2020-01-06T13:49:17+05:30"
}
{
  "cluster": "mycluster",
  "data": "mydata",
  "file": "/Users/vikram/Projects/GO_HOME/src/github.com/vikramjakhr/logrus-demo/main.go:60",
  "func": "main.main",
  "function": "CreateProject",
  "level": "debug",
  "msg": "A sample debug message",
  "package": "api",
  "project": "test-project",
  "time": "2020-01-06T13:49:17+05:30"
}
{
  "cluster": "mycluster",
  "data": "mydata",
  "error": "Hola!!!",
  "file": "/Users/vikram/Projects/GO_HOME/src/github.com/vikramjakhr/logrus-demo/main.go:68",
  "func": "main.main",
  "function": "CreateProject",
  "level": "error",
  "msg": "A sample error message",
  "package": "api",
  "project": "test-project",
  "time": "2020-01-06T13:49:17+05:30"
}
```