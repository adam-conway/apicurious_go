package config

type DBConfig struct {
  Host string
  Port string
  User string
  DBName string
  Password string
}

type Config struct {
  DB *DBConfig
}

func GetConfig() *Config {
  return &Config {
    DB: &DBConfig {
      Host: "localhost",
      Port: "5342",
      User: "go_user",
      DBName: "quantified_self_go",
      Password: "pass",
    }
  }
}
