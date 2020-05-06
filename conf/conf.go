package conf

import (
	"github.com/freezeChen/studio-library/conf"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/freezeChen/studio-library/util"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/v2/config"
)

type Config struct {
	Name    string
	Version string
	Env     string
	Debug   bool
	Log     *zlog.Config
	Mysql   *mysql.Config
	Redis   *redis.Config
}

func Init() (*Config, error) {

	var (
		Conf = &Config{}
	)
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	source := conf.LoadFileSource(util.GetCurrentDirectory() + "/conf.yaml")
	cfg.Load(source)
	if err := cfg.Scan(Conf); err != nil {
		return nil, err
	}

	return Conf, nil
}
