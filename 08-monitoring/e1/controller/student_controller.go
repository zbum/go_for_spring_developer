package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zbum/scouter-agent-golang/scouterx/strace"
	"go_for_spring_developer/08-monitoring/e1/service"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	// TODO 4 메소드 호출까지 추적하려면 이런 코드를 작성해야 합니다..
	ctx := r.Context()
	step := strace.StartMethod(ctx)
	defer strace.EndMethod(ctx, step)

	vars := mux.Vars(r)
	value := vars["id"]
	id, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	student, err := h.studentService.GetStudent(r.Context(), uint(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(student)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(responseBody))
}
