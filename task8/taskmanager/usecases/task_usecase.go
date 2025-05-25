package usecases
import(
	"taskmanager/domain"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

type taskUsecase struct {
	taskRepository domain.TaskReipository
	contextTimeout time.Duration
}
func   NewTaskUsecase(taskrepository  domain.TaskReipository,timeout time.Duration) domain.TaskUsecase{
	return &taskUsecase{
		taskRepository:taskrepository,
		contextTimeout: timeout,
		
	}
}
func(r*taskUsecase) CreateTasks(ctx  context.Context,task domain.Task)(*mongo.InsertOneResult,error){

ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.taskRepository.CreateTasks(ctx,task)
}
func(r*taskUsecase) DeleteTasks(ctx context.Context,id primitive.ObjectID) (domain.Task){
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.taskRepository.DeleteTasks(ctx,id)

}
func(r*taskUsecase) UpdateTask(ctx context.Context, id primitive.ObjectID, updated domain.Task) (*domain.Task, error){
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.taskRepository.UpdateTask(ctx,id,updated)

	
}
func(r*taskUsecase) GetAllTasks(ctx context.Context,) ([]domain.Task ,error){
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.taskRepository.GetAllTasks(ctx)


}


func(r*taskUsecase)GetTaskByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error){
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.taskRepository.GetTaskByID(ctx,id)


}

 