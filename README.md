# go-boilerplate

This project serves as a boilerplate for golang projects. The project has a simple sqlite db connection.

## Dependency Management
Deps is used as the dependency management. 

### Initialize dep
```dep init```

### Add a package
```dep ensure --add github.com/jmoiron/sqlx```

## Configuration
Configuration is setup using cobra/viper. See config.yml 

## Running tests
```make test``` 