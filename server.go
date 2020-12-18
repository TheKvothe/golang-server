package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"./repository"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postController controller.PostController = controller.NewPostController(postRepository)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/types", postController.GetTypes)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
