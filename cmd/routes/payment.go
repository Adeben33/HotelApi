package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers/paymentHandler"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(route *gin.Engine) {
	//Retrieve a list of all payments. (Admin Only)
	route.GET("/payments", paymentHandler.GetAllPayment)

	////Retrieve a specific payment by ID.
	//route.GET("/payments/:id", paymentHandler.GetPayment())
	//
	////	Update an existing payment by ID.
	//route.PUT("/payments/:id", paymentHandler.UpdatePayment())
	//
	////Delete an existing payment by ID. (Admin Only)
	//
	//route.DELETE("/payments/:id", paymentHandler.DeletePayment())
	//
	////	Retrieve the booking associated with a specific payment.
	//route.GET("/payments/:id/booking", paymentHandler.GetPaymentwithBooking())
	//
	////	Retrieve the user associated with a specific payment.
	//route.GET("/payments/:id/user", paymentHandler.GetPaymentUser())
	//
	////	Retrieve the amount associated with a specific payment.
	//route.GET("/payments/:id/amount", paymentHandler.GetPaymentAmount())
	//
	////	Retrieve the payment method for a specific payment.
	//route.GET("/payments/:id/amount", paymentHandler.GetPaymentAmount())
	//
	////	Retrieve the status  for a specific payment.
	//route.GET("/payments/:id/status", paymentHandler.GetPaymentstatus())
	//
	////Refund a specific payment.
	//route.GET("/payments/:id/refund", paymentHandler.PaymentRefund())
	//
	////Cancel a specific payment.
	//route.GET("/payments/:id/cancel", paymentHandler.PaymentCancel())
	//
	////Retrieve the invoice for a specific payment.
	//route.GET("/payments/:id/invoice", paymentHandler.PaymentInvoice())
	//
	//// Retrieve the transaction ID for a specific payment.
	//route.GET("/payments/:id/transaction_id", paymentHandler.GetPaymentTranscationId())
	//
	////Retrieve all payments for a specific user.
	//route.GET("/payments/user/:userId", paymentHandler.GetAllpaymentByAUser())
	//
	////Retrieve all payments for a specific user.
	//route.GET("/payments/booking/:booking_id", paymentHandler.GetAllpaymentByBookingID())
	//
	////Retrieve all payments for a specific status.
	//route.GET("/payments/status/:status", paymentHandler.GetAllpaymentByStatus())
	//
	////Retrieve all payments for a specific date.
	//route.GET("/payments/date/:date", paymentHandler.GetAllPaymentByDate())

}
