package conf

const (
	LocalJavaAddr = "localhost:8091"
	AppId         = "goApp"
	TxGroup       = "goTxGroup"
)

type Configuration struct {
	SeataAddr string `yaml:"seata_addr"`
	AppId     string `yaml:"app_id"`
	TxGroup   string
}

var defConf = Configuration{
	SeataAddr: LocalJavaAddr,
	AppId:     AppId,
	TxGroup:   TxGroup,
}

var Config = defConf
