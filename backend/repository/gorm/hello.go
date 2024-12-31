package gorm

func (repo *Repository)Hello() string {
	return "Hello, World from Docker compose v2!\n";
}

func (repo *Repository)HelloName(name string) string {
	return "Hello, " + name + "!\n";
}