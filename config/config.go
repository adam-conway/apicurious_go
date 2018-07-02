package config

type DBConfig struct {
  Host string
  Port string
  User string
  DBName string
  Password string
  SSLMode string
}

type Config struct {
  DB *DBConfig
}

func GetConfig() *Config {
  return &Config {
    DB: &DBConfig {
      Host: "localhost",
      Port: "5432",
      User: "adamconway",
      DBName: "adamconway",
      SSLMode: "disable",
    },
  }
}
