package http

import (
	ticketPG "github.com/rodkevich/gmp-tickets/internal/ticket/repository/postgres"
	"github.com/rodkevich/gmp-tickets/internal/ticket/rest"
)

func (srv *Server) initRoutes() {
	ticketRepository, err := ticketPG.NewDatasource(srv.cfg, srv.database)
	if err != nil {
		panic(err)
	}
	rest.RegisterRoutes(srv.router, srv.validator, rest.NewTicketService(ticketRepository))

	// usersRepository, err := usersPG.NewDatasource(srv.cfg, srv.database)
	// if err != nil {
	// 	panic(err)
	// }
	// rest.RegisterRoutes(srv.router, srv.validator, rest.NewUserService(usersRepository))
}
