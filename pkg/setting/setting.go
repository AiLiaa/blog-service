package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// NewSetting 始化本项目的配置的基础属性
// 设定配置文件的名称为 config，配置类型为 yaml，并且设置其配置路径为相对路径 configs/，以此确保在项目目录下执行运行时能够成功启动。
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
