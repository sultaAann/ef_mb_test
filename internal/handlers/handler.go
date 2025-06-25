package handlers

import (
	"ef_md_test/internal/custom_errors"
	"ef_md_test/internal/models"
	"ef_md_test/internal/services"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

/*
1 Для получения данных с различными фильтрами и пагинацией
2 Для удаления по идентификатору
3 Для изменения сущности
4 Для добавления новых людей в формате
```json
{
"name": "Dmitriy",
"surname": "Ushakov",
"patronymic": "Vasilevich" // необязательно
}
```
*/

var (
	People       = regexp.MustCompile(`^/people/*$`)
	PeopleWithID = regexp.MustCompile(`^/people/([1-9][0-9]*)$`)
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	s services.Service
}

func NewHandler(s services.Service) Handler {
	return &handler{s: s}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodDelete && PeopleWithID.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	case r.Method == http.MethodGet && PeopleWithID.MatchString(r.URL.Path):
		h.GetById(w, r)
		return
	case r.Method == http.MethodPost && People.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPut && People.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodGet && People.MatchString(r.URL.Path):
		h.GetAll(w, r)
		return
	}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	page := 1
	pageSize := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
		pageSize = ps
	}
	offset := (page - 1) * pageSize

	result, err := h.s.GetAll(pageSize, offset)
	result["page"] = page

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	matches := PeopleWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	person, err := h.s.GetById(uint(id))
	if err != nil {
		if errors.As(err, &custom_errors.NotFoundError{}) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var person models.CreateDTO

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	id, err := h.s.Create(person)

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	j := map[string]float64{"id": float64(id)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(j)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var updatePerson models.UpdateDTO

	if err := json.NewDecoder(r.Body).Decode(&updatePerson); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	err := h.s.Update(updatePerson)
	if err != nil {
		if errors.As(err, &custom_errors.NotFoundError{}) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	matches := PeopleWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	err = h.s.DeleteById(uint(id))
	if err != nil {
		if errors.As(err, &custom_errors.NotFoundError{}) {
			if errors.As(err, &custom_errors.NotFoundError{}) {
				NotFoundHandler(w, r)
				return
			}
			InternalServerErrorHandler(w, r)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found\nCheck Path"))
}
