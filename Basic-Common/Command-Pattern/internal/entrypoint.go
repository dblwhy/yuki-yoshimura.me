package internal

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"yuki-yoshimura.me/command-pattern/internal/controller"
	"yuki-yoshimura.me/command-pattern/internal/domain"
	internallambda "yuki-yoshimura.me/command-pattern/internal/lambda"
	pkghttp "yuki-yoshimura.me/command-pattern/internal/pkg/server/http"
	"yuki-yoshimura.me/command-pattern/internal/service/command"
)

func StartHttpServer() {
	invoker := command.NewInvoker()
	userDomain := domain.NewUserDomain(invoker)
	userController := controller.NewUserController(userDomain)

	router := pkghttp.NewRouter()
	router.Handle("POST", "/user", userController.Create)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func StartLambda() {
	handler := internallambda.NewLambdaCommandHandler()
	lambda.Start(handler)
}
