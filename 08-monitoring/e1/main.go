package main

import (
	"github.com/gorilla/mux"
	"github.com/zbum/scouter-agent-golang/scouterx"
	scouter_middleware "github.com/zbum/scouter-agent-golang/scouterx/middleware"
	"go_for_spring_developer/08-monitoring/e1/configuration/database"
	"go_for_spring_developer/08-monitoring/e1/controller"
	"go_for_spring_developer/08-monitoring/e1/repository"
	"go_for_spring_developer/08-monitoring/e1/service"
	"log"
	"net/http"
)

func main() {

	// TODO 1 스카우터 에이전트 초기화
	scouterx.Init()
	r := mux.NewRouter()

	// TODO 2 웹요청을 추적하기 위한 미들웨어 추가
	r.Use(scouter_middleware.HttpTracingMiddleware)

	studentHandler := initStudentHandler()

	r.HandleFunc("/students/{id}", studentHandler.GetStudent).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initStudentHandler() *controller.StudentHandler {
	datasource := database.NewDatasource()
	studentRepository := repository.NewStudentRepositoryWithContext()
	studentService := service.NewStudentService(datasource, studentRepository)
	return controller.NewStudentHandler(studentService)
}
