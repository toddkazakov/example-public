package registry

var modules []string

func RegisterModule(name string) {
	modules = append(modules, name)
}

func GetModules() []string {
	return modules
}
