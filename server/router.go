package server

// import "github.com/gin-gonic/gin"

// type RouterOpts struct {
// 	BookHandler   *handler.BookHandler
// 	UserHandler   *handler.UserHandler
// 	BorrowHandler *handler.BorrowHandler
// }

// func NewRouter(opts RouterOpts) *gin.Engine {
// 	router := gin.New()
// 	router.ContextWithFallback = true

// 	router.Use(gin.Recovery())
// 	router.Use(middleware.WithTimeout)
// 	router.Use(middleware.AuthorizeHandler())
// 	router.Use(middleware.ErrorHandler())
// 	router.Use(middleware.Logger(logger.NewLogger()))

// 	books := router.Group("/books")
// 	books.GET("", opts.BookHandler.HandleGetBooks)
// 	books.POST("", opts.BookHandler.HandleCreateBook)

// 	users := router.Group("/users")
// 	users.GET("", opts.UserHandler.HandleGetUsers)
// 	users.POST("/register", opts.UserHandler.HandleUserRegister)
// 	users.POST("/login", opts.UserHandler.HandleUserLogin)

// 	borrow := router.Group("/borrows")
// 	borrow.GET("", opts.BorrowHandler.HandleGetRecords)
// 	borrow.POST("", opts.BorrowHandler.HandleBorrowBook)
// 	borrow.PUT("", opts.BorrowHandler.HandleReturnBook)
// 	return router
// }
