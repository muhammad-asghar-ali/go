package movies

type (
	Movie struct {
		ID       string    `json:"id"`
		ISBN     string    `json:"isbn"`
		Title    string    `json:"title"`
		Director *Director `json:"director"`
	}

	Director struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)

var movies []*Movie

func dummy() {
	movies = append(movies, &Movie{
		ID:       "1",
		ISBN:     "543216",
		Title:    "Movie One",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	})

	movies = append(movies, &Movie{
		ID:       "2",
		ISBN:     "5432165",
		Title:    "Movie Two",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	})
}
