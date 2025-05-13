package usecases

import ("taskmanager/domain"
  "time"
  "context"

)

type userUsecase struct {
	useRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(useRepository domain.UserRepository, timeout time.Duration) domain.UserUscase{
	return &userUsecase{
		useRepository: useRepository,
		contextTimeout: timeout,
	}
}

func(r*userUsecase) Loginfunctionality (ctx context.Context,userinfo domain.User)( booleanvalue bool, returnedstring string){
ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.useRepository.Loginfunctionality(ctx,userinfo)


}
func(r*userUsecase)	CreateUser(ctx context.Context, user  domain.User) string {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.useRepository.CreateUser(ctx,user)



	}
