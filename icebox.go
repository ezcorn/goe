package goe

const (
	iceboxRootName = "icebox"
)

type (
	icebox struct{}

	iceboxConfig struct {
		Hash string `json:"hash"` // 唯一标识符号
	}

	// 冰箱隔层
	iceboxLayer struct {
		Hash  string           `json:"hash"`  // 隔层名称
		Rooms iceboxLayerRooms `json:"rooms"` // 隔层内所拥有的块 MAX(256)
	}
	// 冰箱隔层组
	iceboxLayers []iceboxLayer
	// 冰箱层室
	iceboxLayerRoom struct {
		Hash string `json:"hash"` // 块名称
		// data []interface{} `json:"-"` // 数据数组(可不完整)
	}
	// 冰箱层室组
	iceboxLayerRooms []iceboxLayerRoom
	//// 列(索引)
	//iceboxColumn struct {
	//	name   string                       // 列名
	//	blocks map[string][]iceboxLayerRoom // 该列值所在的块
	//}
)

var (
	Icebox           icebox
	iceConfig        iceboxConfig                    // 用来记录冰箱唯一标识码
	iceLayers        = make(map[string]iceboxLayers) // 一个配置内所拥有的所有组 MAX(256)
	iceLayersHash    string                          // 冰箱隔层哈希,用来校验是否变更
	iceLayersHashMap = make(map[string]string)       // 冰箱隔层分组哈希,用来校验是否变更
)

func (i icebox) Set(name string, data interface{}) {
	// MEM -> LOCAL STO -> PRIME STO -> CENTER STO
	//    MIXED         MIXED        MIXED
	// MEM => STO(n)
	if _, exist := iceLayers[name]; !exist {
		iceLayers[name] = iceboxLayers{{
			Hash: uniqueHash(),
			Rooms: iceboxLayerRooms{{
				Hash: uniqueHash(),
			}},
		}}
	}
	iceConfig.update()
	logPrintln(iceboxRootName, "operating: "+iceLayers[name].last().Rooms.last().Hash)
	// fmt.Println(iceConfig.Layers[group].last().Rooms.last())
	// file struct : json{}\n
	//

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

func (icebox) Get(name string, where map[string]string) {
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
	IO.ReadJson(confPath, &iceConfig)
	if iceConfig.Hash == "" {
		iceConfig.Hash = currentServer.mark
		Crypto.MD5(IO.WriteJson(confPath, iceConfig))
	}
	// 如果隔层哈希为空, 从磁盘读取隔层
	if iceLayersHash == "" {
		// 从每个独自的Layer中还原.sector
		layerSuspect := IO.Dirs(iceboxRootName)
		for _, name := range layerSuspect {
			sectorPath := iceboxRootName + "/" + name + "/.sector"
			if IO.Exist(sectorPath) {
				var iceboxLayers iceboxLayers
				IO.ReadJson(sectorPath, &iceboxLayers)
				iceLayers[name] = iceboxLayers
				iceLayersHashMap[name] = Crypto.MD5(jsonEncode(iceboxLayers))
			}
		}
		iceLayersHash = Crypto.MD5(jsonEncode(iceLayers))
		logPrintln(iceboxRootName, "init iceLayers")
	} else {
		newIceLayersHash := Crypto.MD5(jsonEncode(iceLayers))
		// 匹配哈希,如果新旧哈希不同,说明有变化
		if iceLayersHash != newIceLayersHash {
			// 逐层查找隔层中,有差异的层
			for name, layer := range iceLayers {
				// 如果在隔层哈希集中不存在,初始化为空字符串
				if _, exist := iceLayersHashMap[name]; !exist {
					iceLayersHashMap[name] = ""
				}
				// 去现在的数据做哈希
				newHash := Crypto.MD5(jsonEncode(layer))
				// 如果新旧哈希不一致, 说明更新过了,刷新文件
				if iceLayersHashMap[name] != newHash {
					sectorPath := iceboxRootName + "/" + name
					IO.MkDir(sectorPath)
					sectorPath += "/.sector"
					IO.WriteJson(sectorPath, iceLayers[name])
					iceLayersHashMap[name] = newHash
				}
			}
			iceLayersHash = newIceLayersHash
			logPrintln(iceboxRootName, "update iceLayers")
		}
	}
}

/**
 * 最后一个层
 */
func (layers iceboxLayers) last() iceboxLayer {
	return layers[len(layers)-1]
}

/**
 * 最后一个层室
 */
func (layerRooms iceboxLayerRooms) last() iceboxLayerRoom {
	return layerRooms[len(layerRooms)-1]
}
