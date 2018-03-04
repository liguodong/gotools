package workflow

var c *Container

func init() {
	c = new(Container)
}

func New() {
	initServices()
}

func initServices() {
	c.RuntimeService = new(RuntimeService)
	c.TaskService = new(TaskService)
	c.IdentityService = new(IdentityService)
	c.RepositoryService = new(RepositoryService)

	c.RuntimeService.LoadWorkFlowsFromDB()
}

func addTask() {

}

func GetRuntimeService() *RuntimeService {
	return c.RuntimeService
}

func GetTaskService() *TaskService {
	return c.TaskService
}

func GetIdentityService() *IdentityService {
	return c.IdentityService
}

func GetRepositoryService() *RepositoryService {
	return c.RepositoryService
}

type Container struct {
	*RuntimeService
	*TaskService
	*IdentityService
	*RepositoryService
}

