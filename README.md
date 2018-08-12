# Golang Mysql Convenience Package

I use this package to centralize and simplify mysql functionality in my Golang projects. It's provided as is and if you improve it please submit a pull request.

DB connection is lazy and occurs upon the first query call. After the first query or database action, all calls share the same connection.

## Configuration

Configuration takes place using environment variables.

```
DB_HOST
DB_PORT
DB_NAME
DB_USERNAME
DB_PASSWORD
```

## Usage

```
import "github.com/ChristopherBrand/golang-mysql"

func main() {
  results, err := mysql.Query("SELECT * FROM table")
}
```

## Return

All implemented functions return the same as their corrolaries as the go-sql-driver/mysql package. Please see further documentation at https://github.com/go-sql-driver/mysql.

## Further Development

I would like to wrap most of the functionality in https://github.com/go-sql-driver/mysql as well as provide for multiple database connections.