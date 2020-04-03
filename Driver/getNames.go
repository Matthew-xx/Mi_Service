package main

var (
	familyNames = []string{"赵","钱","孙","李","周",
		"吴","郑","王","冯","陈","楚","卫","蒋","沈","韩","杨","张","马","钟","彭"}
	middleNamesMap = map[string][]string{}
	lastNames = []string{"闰","余","成","岁","律","吕","调","阳","云","腾","致","雨","露","结","为","霜","墨","悲","丝","染","诗","赞","羔","羊",
		"右","通","广","内","左","达","承","明","盘","溪","伊","尹","佐","时","阿","衡","岳","宗","泰","岱","禅","主","云","亭","毛","施","淑","姿","工","颦","妍","笑"}
)


func GetRandomName() (name string) {
	familyName := familyNames[GetRandomInt(0,len(familyNames)-1)]
	middleName := middleNamesMap[familyName][GetRandomInt(0,len(middleNamesMap[familyName])-1)]
	lastName := lastNames[GetRandomInt(0,len(lastNames)-1)]

	return familyName+middleName+lastName
}

func init()  {
	for _,x := range familyNames{
		middleNamesMap[x] = []string{"承","斯","龙","天","子","单","铭","德","步","廷","腾","世","学","文","正","国","兴","邦","永","绍","汉","应","汝","先","君","定","登","云"}
	}
}

