package workflow

type RuntimeService struct {

}

func (p *RuntimeService) LoadWorkFlowsFromDB() {

}

type TaskService struct {
	tasks []Task
}

func (p *TaskService) AddTask(task Task)  {
	p.tasks = append(p.tasks, task)
}

type IdentityService struct {

}

type RepositoryService struct {

}

