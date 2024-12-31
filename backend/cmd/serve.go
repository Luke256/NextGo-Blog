package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	driverMySQL "github.com/go-sql-driver/mysql"

	gormRepo "nextgoBlog/repository/gorm"
	"nextgoBlog/router"
	"nextgoBlog/repository"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DB接続
	// MySQLのユーザー名:パスワードはroot:password、プロトコルはtcp(db:3306)、DB名はnextgo_blog

	engine, err := gorm.Open(mysql.New(mysql.Config{
		DSNConfig: &driverMySQL.Config{
			User: "root",
			Passwd: "root",
			Net: 	"tcp",
			Addr: 	"db:3306",
			DBName: "mydb",
		},
	}), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	repo, init, err := gormRepo.NewGormRepository(engine, true)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e = router.Setup(e, engine, repository.Repository(repo))

	if init {
		e.Logger.Info("DB migration has been completed")
		e.Logger.Info("DB initialization has been completed")
	}

	return &Server{
		Router: e,
	}
}

func (s *Server) Start(port string) {
	s.Router.Logger.Fatal(s.Router.Start(port))
}