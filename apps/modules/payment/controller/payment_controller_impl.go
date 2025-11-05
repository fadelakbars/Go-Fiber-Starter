package controller

import (
	"fmt"
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/payment/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
)

type PaymentControllerImpl struct {
	service service.PaymentService
}

func NewPaymentController(service service.PaymentService) PaymentController {
	return &PaymentControllerImpl{service: service}
}

func (c *PaymentControllerImpl) FindAll(ctx *fiber.Ctx) error {
	// Parse query params
	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")
	typeParam := ctx.Query("type")
	status := ctx.Query("status")
	voucherCode := ctx.Query("voucher_code")
	boothID := ctx.Query("booth_id")

	filters := map[string]interface{}{
		"limit":        limit,
		"offset":       offset,
		"start_date":   startDate,
		"end_date":     endDate,
		"type":         typeParam,
		"status":       status,
		"voucher_code": voucherCode,
		"booth_id":     boothID,
	}

	// You need to implement FindAllWithFilters in the service and repository layers
	payments, totalData, totalIncome, err := c.service.FindAll(ctx.Context(), filters)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch payments")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, fiber.Map{
		"payments":     payments,
		"total_data":   totalData,
		"total_income": totalIncome,
	}, "Payments fetched successfully")
}

func (c *PaymentControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid payment ID")
	}
	payment, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Payment not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, payment, "Payment fetched successfully")
}

func (c *PaymentControllerImpl) Create(ctx *fiber.Ctx) error {
	var payment domain.Payment
	if err := ctx.BodyParser(&payment); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	createdPayment, err := c.service.Create(ctx.Context(), payment)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create payment")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdPayment, "Payment created successfully")
}

func (c *PaymentControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid payment ID")
	}
	var payment domain.Payment
	if err := ctx.BodyParser(&payment); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	payment.ID = id
	updatedPayment, err := c.service.Update(ctx.Context(), payment)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update payment")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedPayment, "Payment updated successfully")
}

func (c *PaymentControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid payment ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete payment")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Payment deleted successfully")
}

func generateOrderID() string {
	now := time.Now()
	return fmt.Sprintf("ORD-%s-%s",
		now.Format("20060102"), // YYYYMMDD
		now.Format("150405"),   // HHMMSS
	)
}
func (c *PaymentControllerImpl) CreateSnapPayment(ctx *fiber.Ctx) error {
	type Request struct {
		Amount    float64 `json:"amount"`
		BoothID   string  `json:"booth_id"`
		VoucherID *string `json:"voucher_id,omitempty"` // Optional, can be nil
		Type      int     `json:"type"`                 // e.g., 1 for Mandiri, 2 for Voucher
	}

	fmt.Println("1. CreateSnapPayment called")

	var req Request
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	// 1. Create payment in DB (status: pending, paid_at: null, transaction_reference: req.OrderID)
	payment := domain.Payment{
		PaymentMethod:        "QRIS",
		BoothID:              req.BoothID,
		Amount:               req.Amount,
		Status:               "pending",
		VoucherID:            req.VoucherID, // Optional, can be nil
		Type:                 req.Type,
		TransactionReference: generateOrderID(),
	}
	createdPayment, err := c.service.Create(ctx.Context(), payment)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create payment in database")
	}

	// 2. Set global config Midtrans
	midtrans.ServerKey = viper.GetString("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = viper.GetString("MIDTRANS_CLIENT_KEY")

	midtrans.Environment = midtrans.Production // Ganti ke midtrans.Production saat live

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.TransactionReference,
			GrossAmt: int64(req.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapResp, err := snap.CreateTransaction(snapReq)

	fmt.Printf("snapResp: %+v, err: %v\n", snapResp, err)

	return helpers.WriteJSON(ctx, fiber.StatusOK, fiber.Map{
		"payment":      createdPayment,
		"token":        snapResp.Token,
		"redirect_url": snapResp.RedirectURL,
	}, "Snap payment created successfully")
}

func (c *PaymentControllerImpl) HandleMidtransCallback(ctx *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := ctx.BodyParser(&payload); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid callback payload")
	}

	orderID, _ := payload["order_id"].(string)
	transactionStatus, _ := payload["transaction_status"].(string)

	// Default: status = pending
	status := "pending"
	// paidAt := nil

	if transactionStatus == "settlement" || transactionStatus == "capture" {
		status = "paid"
	} else if transactionStatus == "expire" || transactionStatus == "cancel" || transactionStatus == "deny" {
		status = "failed"
	}

	// Update payment in DB
	err := c.service.UpdatePaymentStatusByOrderID(ctx.Context(), orderID, status)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update payment status")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Callback processed",
	})
}
