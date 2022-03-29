package initialize

import (
	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/service"
)

func initTimer() {
	if global.CONFIG.Timer.Start {
		go func() {
			global.Timer.AddTaskByFunc("BOOK_VIEW_COUNT", global.CONFIG.Timer.Spec, service.UpdateBooksView)
			global.Timer.AddTaskByFunc("DOC_VIEW_COUNT", global.CONFIG.Timer.Spec, service.UpdateDocsView)
		}()
	}
}
