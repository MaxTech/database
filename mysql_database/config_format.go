package mysql_database

type MySQLConfigFormat struct {
    Address  string `json:"address" yaml:"address"`
    Username string `json:"username" yaml:"username"`
    Password string `json:"password" yaml:"password"`
    DBName   string `json:"db_name" yaml:"db_name"`
}
