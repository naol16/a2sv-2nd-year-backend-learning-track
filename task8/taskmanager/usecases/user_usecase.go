package usecases

import ("taskmanager/domain"
  "time"
  "context"

)

type UserUsecase struct {
	useRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(useRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase{
	return &UserUsecase{
		useRepository: useRepository,
		contextTimeout: timeout,
	}
}

func(r*UserUsecase) Loginfunctionality (ctx context.Context,userinfo domain.User)( booleanvalue bool, returnedstring string){
ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.useRepository.Loginfunctionality(ctx,userinfo)


}
func(r*UserUsecase)	CreateUser(ctx context.Context, user  domain.User) string {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.useRepository.CreateUser(ctx,user)



	}
