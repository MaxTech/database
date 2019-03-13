package mysql_database

type MySQLConfigFormat struct {
    Address  string `yaml:"address"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    DBName   string `yaml:"db_name"`
}
