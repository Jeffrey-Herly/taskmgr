package handler

type AppHandler struct {
	Views *ViewsHandler
	Auth  *AuthHandler
}

func NewAppHandler() *AppHandler {
	return &AppHandler{
		Views: NewViewsHandler(),
		Auth:  NewAuthHandler(),
	}
}
