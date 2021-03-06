package options

import (
	"github.com/akhripko/dummy/kafka/consumer"
	"github.com/akhripko/dummy/kafka/producer"
	"github.com/akhripko/dummy/providers/grpc/hellosrv"
	"github.com/akhripko/dummy/storage/postgres"
	"github.com/spf13/viper"
)

func ReadEnv() *Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("LOG_LEVEL", "DEBUG")

	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("GRAPHQL_PORT", 8081)
	viper.SetDefault("GRPC_PORT", 8090)

	viper.SetDefault("HEALTH_CHECK_PORT", 8888)
	viper.SetDefault("PROMETHEUS_PORT", 9100)

	viper.SetDefault("HELLO_SRV_TARGET", "localhost:8090")

	viper.SetDefault("SQLDB_HOST", "localhost")
	viper.SetDefault("SQLDB_PORT", 5432)
	viper.SetDefault("SQLDB_USER", "postgres")
	viper.SetDefault("SQLDB_PASS", "")
	viper.SetDefault("SQLDB_DB_NAME", "db")
	viper.SetDefault("SQLDB_MAX_OPEN_CONNS", 10)

	viper.SetDefault("CACHE_ADDR", ":6379")

	viper.SetDefault("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
	viper.SetDefault("KAFKA_SSL_ENABLED", false)
	viper.SetDefault("KAFKA_SSL_CA_LOCATION", "")
	viper.SetDefault("KAFKA_SSL_CERTIFICATE_LOCATION", "")
	viper.SetDefault("KAFKA_SSL_KEY_LOCATION", "")
	viper.SetDefault("KAFKA_PRODUCER_IDEMPOTENCE", false)
	viper.SetDefault("KAFKA_PRODUCER_READ_EVENTS", true)
	viper.SetDefault("KAFKA_PRODUCER_FLUSH_TIMEOUT_MS", 15000)

	viper.SetDefault("KAFKA_TOPIC_TEST", "hello-topic")

	viper.SetDefault("KAFKA_CONSUMER_POLL_TIMEOUT_MS", 300)
	viper.SetDefault("KAFKA_CONSUMER_NAME", "hello")
	viper.SetDefault("KAFKA_CONSUMER_GROUP_ID", "dummy")
	viper.SetDefault("KAFKA_CONSUMER_SESSION_TIMEOUT_MS", 6000)
	viper.SetDefault("KAFKA_CONSUMER_AUTO_OFFSET_RESET", "earliest")

	return &Config{
		LogLevel:        viper.GetString("LOG_LEVEL"),
		HTTPPort:        viper.GetInt("HTTP_PORT"),
		GRPCPort:        viper.GetInt("GRPC_PORT"),
		GraphqlPort:     viper.GetInt("GRAPHQL_PORT"),
		HealthCheckPort: viper.GetInt("HEALTH_CHECK_PORT"),
		PrometheusPort:  viper.GetInt("PROMETHEUS_PORT"),
		SQLDB: postgres.Config{
			Host:         viper.GetString("SQLDB_HOST"),
			Port:         viper.GetInt("SQLDB_PORT"),
			User:         viper.GetString("SQLDB_USER"),
			Pass:         viper.GetString("SQLDB_PASS"),
			DBName:       viper.GetString("SQLDB_DB_NAME"),
			MaxOpenConns: viper.GetInt("SQLDB_MAX_OPEN_CONNS"),
		},
		CacheAddr:    viper.GetString("CACHE_ADDR"),
		HelloSrvConf: hellosrv.Config{Target: viper.GetString("HELLO_SRV_TARGET")},
		KafkaTopic:   KafkaTopic{Hello: viper.GetString("KAFKA_TOPIC_TEST")},
		KafkaProducer: producer.Config{
			Idempotence:      viper.GetBool("KAFKA_PRODUCER_IDEMPOTENCE"),
			ReadEvents:       viper.GetBool("KAFKA_PRODUCER_READ_EVENTS"),
			FlushTimeoutMs:   viper.GetInt("KAFKA_PRODUCER_FLUSH_TIMEOUT_MS"),
			BootstrapServers: viper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
			SSL: producer.SSLConfig{
				Enabled:             viper.GetBool("KAFKA_SSL_ENABLED"),
				KeyLocation:         viper.GetString("KAFKA_SSL_KEY_LOCATION"),
				CertificateLocation: viper.GetString("KAFKA_SSL_CERTIFICATE_LOCATION"),
				CALocation:          viper.GetString("KAFKA_SSL_CA_LOCATION"),
			},
		},
		KafkaConsumer: consumer.Config{
			BootstrapServers: viper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
			PollTimeoutMs:    viper.GetInt("KAFKA_CONSUMER_POLL_TIMEOUT_MS"),
			Name:             viper.GetString("KAFKA_CONSUMER_NAME"),
			GroupID:          viper.GetString("KAFKA_CONSUMER_GROUP_ID"),
			SessionTimeoutMs: viper.GetString("KAFKA_CONSUMER_SESSION_TIMEOUT_MS"),
			AutoOffsetReset:  viper.GetString("KAFKA_CONSUMER_AUTO_OFFSET_RESET"),
			SSL: consumer.SSLConfig{
				Enabled:             viper.GetBool("KAFKA_SSL_ENABLED"),
				KeyLocation:         viper.GetString("KAFKA_SSL_KEY_LOCATION"),
				CertificateLocation: viper.GetString("KAFKA_SSL_CERTIFICATE_LOCATION"),
				CALocation:          viper.GetString("KAFKA_SSL_CA_LOCATION"),
			},
		},
	}
}
