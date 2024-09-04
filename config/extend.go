package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	Apns2 Apns2 // 这里配置对应配置文件的结构即可
}

type Apns2 struct {
	Teamid   string `mapstructure:"teamid" yaml:"teamid"`
	Bundleid string `mapstructure:"bundleid" yaml:"bundleid"`
	CertPath string `mapstructure:"certPath" yaml:"certPath"`
	Password string `mapstructure:"password" yaml:"password"`
	Topic    string `mapstructure:"topic" yaml:"topic"`
	Prod     bool   `mapstructure:"prod" yaml:"prod"`
	Push     bool   `mapstructure:"push" yaml:"push"`
}
