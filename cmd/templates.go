package cmd

type Category string

const (
	Vue   Category = "vue"
	React Category = "react"
)

type templateType struct {
	name        string
	url         string
	downloadUrl string
	description string
	category    Category
}

var TemplatesSource = []templateType{
	{
		category:    "vue",
		name:        "RainManGO/vue3-composition-admin",
		url:         "https://github.com/RainManGO/vue3-composition-admin",
		downloadUrl: "git@github.com:RainManGO/vue3-composition-admin.git",
		description: "基于vue3 的管理端模板",
	},
	{
		category:    "vue",
		name:        "cmdparkour/vue-admin-box",
		url:         "https://github.com/cmdparkour/vue-admin-box",
		downloadUrl: "git@github.com:cmdparkour/vue-admin-box.git",
		description: "vue3,vite,element-plus中后台管理系统",
	},
	{
		category:    "vue",
		name:        "jackchen0120/vueDataV",
		url:         "https://github.com/jackchen0120/vueDataV",
		downloadUrl: "git@github.com:jackchen0120/vueDataV.git",
		description: "基于Vue + Echarts 构建的数据可视化平台",
	},
	{
		category:    "react",
		name:        "z-9527/react-admin-master",
		url:         "https://github.com/z-9527/react-admin-master",
		downloadUrl: "git@github.com:z-9527/react-admin-master.git",
		description: "基于React-Antd的后台模板",
	},
	{
		category:    "react",
		name:        "hsl947/react-antd-multi-tabs-admin",
		url:         "https://github.com/hsl947/react-antd-multi-tabs-admin",
		downloadUrl: "git@github.com:hsl947/react-antd-multi-tabs-admin.git",
		description: "ts+react+antd-多页签后台模板",
	},
}
