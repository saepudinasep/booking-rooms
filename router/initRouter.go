package router

import (
	"BookingRoom/src/auth/authDelivery"
	"BookingRoom/src/auth/authRepository"
	"BookingRoom/src/auth/authUsecase"
	"BookingRoom/src/employees/employeeDelivery"
	"BookingRoom/src/employees/employeeRepository"
	"BookingRoom/src/employees/employeeUsecase"
	"BookingRoom/src/report/reportDelivery"
	"BookingRoom/src/report/reportRepository"
	"BookingRoom/src/report/reportUsecase"
	"BookingRoom/src/rooms/roomsDelivery"
	"BookingRoom/src/rooms/roomsRepository"
	"BookingRoom/src/rooms/roomsUseCase"
	"BookingRoom/src/transactions/transactionsDelivery"
	"BookingRoom/src/transactions/transactionsRepository"
	"BookingRoom/src/transactions/transactionsUseCase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authRepo := authRepository.NewAuthRepository(db)
	employeeRepo := employeeRepository.NewEmployeeRepository(db)
	//transactionRepo := transactionRepository.NewTransactionRepo(db)
	reportRepo := reportRepository.NewReportRepo(db)
	transactionsRepo := transactionsRepository.NewTransactionsRepository(db)
	roomRepo := roomsRepository.NewRoomsRepository(db)

	// usecase
	authUC := authUsecase.NewAuthUsecase(authRepo)
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepo)
	//transactionUC := transactionUsecase.NewTransactionUsecase(transactionRepo)
	reportUC := reportUsecase.NewReportUsecase(reportRepo)
	transactionUC := transactionsUseCase.NewTransactionsUseCase(transactionsRepo, employeeRepo)
	roomUC := roomsUseCase.NewRoomsUseCase(roomRepo)

	// delivery
	authDelivery.NewAuthDelivery(v1Group, authUC)
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
	//transactionDelivery.NewTransactionDelivery(v1Group, transactionUC)
	reportDelivery.NewReportDelivery(v1Group, reportUC)
	transactionsDelivery.NewTransactionsDelivery(v1Group, transactionUC)
	roomsDelivery.NewRoomDelivery(v1Group, roomUC)
}
