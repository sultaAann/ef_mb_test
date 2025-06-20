package handlers

import "ef_md_test/internal/services"

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

type Handler interface {
}

type handler struct {
	s *services.Service
}

func NewHandler(s *services.Service) Handler {
	return &handler{s: s}
}
