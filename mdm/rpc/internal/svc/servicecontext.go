package svc

import (
	"context"
	"database/sql"
	"fmt"
	"his/mdm/model/brands/ent"
	"his/mdm/rpc/internal/config"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

type ServiceContext struct {
	Config        config.Config
	MdmBrandModel *ent.MdmBrandsClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	PostgresClient, err := NewEntClientWithPool(c)

	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:        c,
		MdmBrandModel: PostgresClient.MdmBrands,
	}
}

func NewEntClientWithPool(c config.Config) (*ent.Client, error) {
	// 打开数据库连接
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", c.PostgresConf.Host, c.PostgresConf.Port, c.PostgresConf.User, c.PostgresConf.DbName, c.PostgresConf.Password, c.PostgresConf.SslMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(c.PostgresConf.MaxOpenConns)
	db.SetMaxIdleConns(c.PostgresConf.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(c.PostgresConf.ConnMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(c.PostgresConf.ConnMaxIdleTime) * time.Second)

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// 创建 Ent 驱动和客户端
	driver := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	return client, nil
}
