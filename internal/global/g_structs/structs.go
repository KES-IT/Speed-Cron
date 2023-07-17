package g_structs

type InitData struct {
	Department   string `json:"department" description:"部门"`
	Name         string `json:"name" description:"姓名"`
	LocalVersion string `json:"local_version" description:"本地版本号"`
}
