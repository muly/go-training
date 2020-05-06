func main() {
	r := mux.NewRouter()
	r.Path("/api/v1/user-login").Handler(http.HandlerFunc(UserLoginHandler))   // /api/v1/user-login will be NOT intercepted by ValidationMiddleWare
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(ValidationMiddleware)
	api.Path("/users").Handler(http.HandlerFunc(UsersHandler)) // api/v1/users will be intercepted by ValidationMiddleWare