package goe

const (
	iceboxRootName = "icebox"
)

type (
	icebox struct {
		config iceboxConfig
	}

	// 冰箱隔层
	iceboxLayer struct {
		name   string            // 隔层名称
		blocks []iceboxLayerRoom // 隔层内所拥有的块 MAX(256)
	}
	// 冰箱层室
	iceboxLayerRoom struct {
		name string // 块名称
		// data []interface{} `json:"-"` // 数据数组(可不完整)
	}
	// 列(索引)
	iceboxColumn struct {
		name   string                       // 列名
		blocks map[string][]iceboxLayerRoom // 该列值所在的块
	}

	iceboxConfig struct {
		Unique string                 `json:"unique"` // 唯一标识符号
		Layers map[string]iceboxLayer `json:"layers"` // 一个配置内所拥有的所有组 MAX(256)
	}
)

var (
	Icebox icebox
)

func (i icebox) Set(group string, data interface{}) {
	// MEM -> LOCAL STO -> PRIME STO -> CENTER STO
	//    MIXED         MIXED        MIXED
	// MEM => STO(n)

	// GROUP
	//
	//
	//

	// QUERY( DATA-1 )
	// HTML5 -> ACCESS -> SERVER ( DATA-1 )
	//                 -> SERVER ( DATA-2 )
	// HTML5 <- ACCESS <- SERVER
	//      DATA      DATA
}

func (icebox) Get(group string, where map[string]string) {
	//// MEM <- MEM | MEM <- STO
	//if _, ok := memCache[group]; ok {
	//
	//} else {
	//
	//}
}

// 更新冰箱配置
func (conf iceboxConfig) update() {
	// 如果根文件夹不存在,创建
	IO.MkDir(iceboxRootName)
	//
	confPath := iceboxRootName + "/.config"
	if IO.Exist(confPath) {
		IO.ReadJson(confPath, &Icebox.config)
		Icebox.config.Layers = conf.layers()
	} else {
		Icebox.config = iceboxConfig{
			Unique: currentServer.mark,
			Layers: conf.layers(),
		}
	}
	IO.WriteJson(confPath, Icebox.config)
}

/**
 * 获取组群
 */
func (iceboxConfig) layers() map[string]iceboxLayer {
	// 从每个独自的GROUP中还原.sector
	layerSuspect := IO.Dirs(iceboxRootName)
	layers := make(map[string]iceboxLayer)
	for _, dir := range layerSuspect {
		sectorPath := iceboxRootName + "/" + dir + "/.sector"
		if IO.Exist(sectorPath) {
			layers[dir] = iceboxLayer{}
		}
	}
	return layers
}
