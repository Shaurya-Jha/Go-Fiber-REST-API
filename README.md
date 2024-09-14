# REST API with Go fiber framework

- Description - REST API made with fiber framework that manages tasks in real-time. Some more changes will be incorporated in near future.

### Framework used
- Fiber for Go

### External packages used
- google's uuid for creating uuid for each tasks

### Routes

- GET /api/tasks - get all tasks
- GET /api/tasks/:id - get particular task as per the id
- POST /api/tasks - create a new task
- PUT /api/tasks/:id - update a particular task as per the id
- DELETE /api/tasks/:id - delete a task as per the id