package products

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s}
}

func (h *handler) ListProducts(http.ResponseWriter, *http.Request) {}
