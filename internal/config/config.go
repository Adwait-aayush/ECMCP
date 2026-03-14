package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}
type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret           string
	ExpiresIn        string
	RefreshExpiresIn string
}

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string
}

type UploadConfig struct {
	Path    string
	MaxSize int64
}
func Load() (*Config,error){
	_=godotenv.Load()
	jwtExpiresIn,_ := time.ParseDuration(getEnv("JWT_EXPIRES_IN","24h"))
	jwtRefreshExpiresIn,_ := time.ParseDuration(getEnv("REFRESH_TOKEN_EXPIRES_IN","720h"))
	maxUploadSize,_:=strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE","10485760"),10,64)

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "release"),
	}, 
	Database: DatabaseConfig{
     Host: getEnv("DB_HOST","localhost"),
	 Port: getEnv("DB_PORT","5432"),
	 User: getEnv("DB_USER","postgres"),
	 Password: getEnv("DB_PASSWORD","password"),
	 Name: getEnv("DB_NAME","ecmcp"),
	 SSLMode: getEnv("DB_SSL_MODE","disable"),
	},
	JWT: JWTConfig{
     Secret: getEnv("JWT_SECRET","key"),
	 ExpiresIn: jwtExpiresIn.String(),
	 RefreshExpiresIn: jwtRefreshExpiresIn.String(),
	},
	AWS: AWSConfig{
	 Region: getEnv("AWS_REGION","us-east-1"),
	 AccessKeyID: getEnv("AWS_ACCESS_KEY_ID","test"),
	 SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY","test"),
	 S3Bucket: getEnv("AWS_S3_BUCKET","ecmcp-uploads"),
	 S3Endpoint: getEnv("AWS_S3_ENDPOINT","http://localhost:9000"),
	},
	Upload: UploadConfig{	
		Path: getEnv("UPLOAD_PATH","./uploads"),
		MaxSize: maxUploadSize,
	},

}, nil
}
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}