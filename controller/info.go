package controller
import(
	"github.com/kataras/iris/v12"
	"tp_iris/model"
)

func Info(ctx iris.Context) {
	books := []model.Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
	ctx.JSON(books)
}