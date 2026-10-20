package bootstrap

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/spf13/viper"
)

//type Env struct {
//	AppEnv                 string `mapstructure:"APP_ENV"`
//	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
//	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
//	DBHost                 string `mapstructure:"DB_HOST"`
//	DBPort                 string `mapstructure:"DB_PORT"`
//	DBUser                 string `mapstructure:"DB_USER"`
//	DBPass                 string `mapstructure:"DB_PASS"`
//	DBName                 string `mapstructure:"DB_NAME"`
//	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
//	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
//	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
//	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
//}

type EnvPreload struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`

	// Admin
	RootUser  string `mapstructure:"ROOT_USER"`
	RootPass  string `mapstructure:"ROOT_PASS"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	// Postgres
	PGHost      string `mapstructure:"PG_HOST"`
	PGPort      string `mapstructure:"PG_PORT"`
	PGAdmin     string `mapstructure:"PG_ADMIN"`
	PGAdminPass string `mapstructure:"PG_ADMIN_PASS"`
	PGUser      string `mapstructure:"PG_USER"`
	PGPass      string `mapstructure:"PG_PASS"`
	PGDBName    string `mapstructure:"PG_DB_NAME"`
	// storage
	StorageType string `mapstructure:"STORAGE_TYPE"`
	// S3
	S3PartitionID     string `mapstructure:"S3_PARTITION_ID"`
	S3URL             string `mapstructure:"S3_URL"`
	S3SigningRegion   string `mapstructure:"S3_SIGNING_REGION"`
	S3AccessKeyID     string `mapstructure:"S3_ACCESS_KEY_ID"`
	S3SecretAccessKey string `mapstructure:"S3_SECRET_ACCESS_KEY"`

	// DUFS
	DufsURL string `mapstructure:"DUFS_URL"`
}

type Env struct {
	AppEnv         string
	LogLevel       string
	JWTSecret      string
	Admin          Admin
	ServerAddress  string
	ContextTimeout int
	Postgres       Postgres
}

//type Nats struct {
//	Url       string
//	Name      string
//	EnableTls bool
//	Cert      Tls
//	JetStream JetStream
//	Ctx       context.Context
//}

//type Tls struct {
//	CertFile string
//	KeyFile  string
//	CaFile   string
//}
//
//type JetStream struct {
//	Name    string
//	Durable string
//}

type Admin struct {
	Username string
	Password string
}

type Postgres struct {
	DBHost      string
	DBPort      string
	DBAdmin     string
	DBAdminPass string
	DBUser      string
	DBPass      string
	DBName      string
	Ctx         context.Context
}

//type S3 struct {
//	PartitionID     string
//	URL             string
//	SigningRegion   string
//	AccessKeyID     string
//	SecretAccessKey string
//	Ctx             context.Context
//}
//
//type Dufs struct {
//	URL string
//	Ctx context.Context
//}

func NewEnv() *Env {
	envPreload := EnvPreload{}
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("CONTEXT_TIMEOUT", 10)
	// Admin
	viper.SetDefault("ROOT_USER", "wombatadmin")
	viper.SetDefault("ROOT_PASS", "wombatadmin")
	viper.SetDefault("JWT_SECRET", "your-secret-key-change-in-production")
	// Postgres
	viper.SetDefault("PG_HOST", "192.168.1.66")
	viper.SetDefault("PG_PORT", "5432")
	viper.SetDefault("PG_ADMIN", "postgres")
	viper.SetDefault("PG_ADMIN_PASS", "postgres")
	viper.SetDefault("PG_USER", "postgres")
	viper.SetDefault("PG_PASS", "postgres")
	viper.SetDefault("PG_DB_NAME", "ppp_db")
	viper.AutomaticEnv()
	err := viper.Unmarshal(&envPreload)
	if err != nil {
		Fatal("Environment can't be loaded", "err", err)
	}
	if envPreload.AppEnv == "development" {
		slog.Info("The App is running in development env")
	}
	resultJson, err := json.MarshalIndent(envPreload, "", "    ")
	if err != nil {
		Fatal("Error generating JSON", "err", err)
	}
	slog.Debug("envPreload loaded", "envPreload", string(resultJson))
	ctx := context.TODO()
	newEnv := Env{
		AppEnv:         envPreload.AppEnv,
		LogLevel:       envPreload.LogLevel,
		JWTSecret:      envPreload.JWTSecret,
		ServerAddress:  envPreload.ServerAddress,
		ContextTimeout: envPreload.ContextTimeout,
		Admin: Admin{
			Username: envPreload.RootUser,
			Password: envPreload.RootPass,
		},
		Postgres: Postgres{
			DBHost:      envPreload.PGHost,
			DBPort:      envPreload.PGPort,
			DBAdmin:     envPreload.PGAdmin,
			DBAdminPass: envPreload.PGAdminPass,
			DBUser:      envPreload.PGUser,
			DBPass:      envPreload.PGPass,
			DBName:      envPreload.PGDBName,
			Ctx:         ctx,
		},
	}
	resultJson, err = json.MarshalIndent(newEnv, "", "    ")
	if err != nil {
		Fatal("Error generating JSON", "err", err)
	}
	slog.Debug("newEnv loaded", "newEnv", string(resultJson))
	return &newEnv
}
