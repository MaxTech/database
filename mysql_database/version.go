package mysql_database

import "fmt"

var Version string

func init() {
    Version = fmt.Sprintf(
        "%s module:\t\t\t%s",
        "mysql database",
        "0.0.1")
}
