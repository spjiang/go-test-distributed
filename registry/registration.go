package registry

type ServiceName string

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
	// 服务依赖三方服务
	RequiredServices []ServiceName
	ServiceUpdateURL string
}

const (
	LogService   = ServiceName("LogService")
	GradeService = ServiceName("GradeService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
