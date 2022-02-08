package config

type (
	Config struct {
		App      App      `yaml:"app"`
		Mysql    Mysql    `yaml:"mysql"`
		Postgres Postgres `yaml:"postgres"`
		Account  Account  `yaml:"account"`
	}

	App struct {
		Name    string `yaml:"name" envconfig:"SNAPP_CLEAN_NAME"`
		Address string `yaml:"address" envconfig:"SNAPP_CLEAN_ADDRESS"`
	}

	Mysql struct {
		Username string `yaml:"username" envconfig:"MYSQL_USERNAME"`
		Password string `yaml:"password" envconfig:"MYSQL_PASSWORD"`
		DBName   string `yaml:"db_name" envconfig:"MYSQL_DBNAME"`
		Host     string `yaml:"host" envconfig:"MYSQL_HOST"`
		Port     string `yaml:"port" envconfig:"MYSQL_PORT"`
	}

	Postgres struct {
		Username string `yaml:"username" envconfig:"POSTGRES_USERNAME"`
		Password string `yaml:"password" envconfig:"POSTGRES_PASSWORD"`
		DBName   string `yaml:"db_name" envconfig:"POSTGRES_DBNAME"`
		Host     string `yaml:"host" envconfig:"POSTGRES_HOST"`
		Port     string `yaml:"port" envconfig:"POSTGRES_PORT"`
	}

	Account struct {
		MinUsernameLength int `yaml:"min_username_length"`
	}
)
