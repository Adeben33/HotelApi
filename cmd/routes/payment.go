package routes

import "github.com/gin-gonic/gin"

func PaymentRoutes(route *gin.Engine) {
	//Retrieve a list of all payments. (Admin Only)
	route.GET("/payments", handlers.GetAllPayment())

	//Retrieve a specific payment by ID.
	route.GET("/payments/:id", handlers.GetPayment())

	//	Update an existing payment by ID.
	route.PUT("/payments/:id", handlers.UpdatePayment())

	//Delete an existing payment by ID. (Admin Only)

	route.DELETE("/payments/:id", handlers.DeletePayment())

	//	Retrieve the booking associated with a specific payment.
	route.GET("/payments/:id/booking", handlers.GetPaymentwithBooking())

	//	Retrieve the user associated with a specific payment.
	route.GET("/payments/:id/user", handlers.GetPaymentUser())

	//	Retrieve the amount associated with a specific payment.
	route.GET("/payments/:id/amount", handlers.GetPaymentAmount())

	//	Retrieve the payment method for a specific payment.
	route.GET("/payments/:id/amount", handlers.GetPaymentAmount())

	//	Retrieve the status  for a specific payment.
	route.GET("/payments/:id/status", handlers.GetPaymentstatus())

	//Refund a specific payment.
	route.GET("/payments/:id/refund", handlers.PaymentRefund())

	//Cancel a specific payment.
	route.GET("/payments/:id/cancel", handlers.PaymentCancel())

	//Retrieve the invoice for a specific payment.
	route.GET("/payments/:id/invoice", handlers.PaymentInvoice())

	// Retrieve the transaction ID for a specific payment.
	route.GET("/payments/:id/transaction_id", handlers.GetPaymentTranscationId())

	//Retrieve all payments for a specific user.
	route.GET("/payments/user/:userId", handlers.GetAllpaymentByAUser())

	//Retrieve all payments for a specific user.
	route.GET("/payments/booking/:booking_id", handlers.GetAllpaymentByBookingID())

	//Retrieve all payments for a specific status.
	route.GET("/payments/status/:status", handlers.GetAllpaymentByStatus())

	//Retrieve all payments for a specific date.
	route.GET("/payments/date/:date", handlers.GetAllpaymentByDate())

}
