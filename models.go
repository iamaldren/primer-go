package primer

import "time"

// TransactionStatus ...
type TransactionStatus string

// TransactionStatus enum
const (
	// TransactionStatusPending -
	TransactionStatusPending TransactionStatus = "PENDING"
	// TransactionStatusFailed -
	TransactionStatusFailed TransactionStatus = "FAILED"
	// TransactionStatusAuthorized -
	TransactionStatusAuthorized TransactionStatus = "AUTHORIZED"
	// TransactionStatusSettling -
	TransactionStatusSettling TransactionStatus = "SETTLING"
	// TransactionStatusPartiallySettled -
	TransactionStatusPartiallySettled TransactionStatus = "PARTIALLY_SETTLED"
	// TransactionStatusSettled -
	TransactionStatusSettled TransactionStatus = "SETTLED"
	// TransactionStatusDeclined -
	TransactionStatusDeclined TransactionStatus = "DECLINED"
	// TransactionStatusCancelled -
	TransactionStatusCancelled TransactionStatus = "CANCELLED"
)

// PaymentInstrumentType ...
type PaymentInstrumentType string

// PaymentInstrumentType enum
const (
	// PaymentInstrumentTypePaymentCard -
	PaymentInstrumentTypePaymentCard PaymentInstrumentType = "PAYMENT_CARD"
	// PaymentInstrumentTypePaypalOrder -
	PaymentInstrumentTypePaypalOrder PaymentInstrumentType = "PAYPAL_ORDER"
	// PaymentInstrumentTypePaypalBillingAgreement -
	PaymentInstrumentTypePaypalBillingAgreement PaymentInstrumentType = "PAYPAL_BILLING_AGREEMENT"
	// PaymentInstrumentTypeGooglePay -
	PaymentInstrumentTypeGooglePay PaymentInstrumentType = "GOOGLE_PAY"
	// PaymentInstrumentTypeGoCardlessMandate -
	PaymentInstrumentTypeGoCardlessMandate PaymentInstrumentType = "GOCARDLESS_MANDATE"
	// PaymentInstrumentTypeKlarnaAuthorizationToken -
	PaymentInstrumentTypeKlarnaAuthorizationToken PaymentInstrumentType = "KLARNA_AUTHORIZATION_TOKEN"
	// PaymentInstrumentTypeKlarnaCustomerToken -
	PaymentInstrumentTypeKlarnaCustomerToken PaymentInstrumentType = "KLARNA_CUSTOMER_TOKEN"
	// PaymentInstrumentTypeApplePay -
	PaymentInstrumentTypeApplePay PaymentInstrumentType = "APPLE_PAY"
)

// TokenType ...
type TokenType string

// TokenType enum
const (
	// TokenTypeMultiUse -
	TokenTypeMultiUse TokenType = "MULTI_USE"
	// TokenTypeSingleUse -
	TokenTypeSingleUse TokenType = "SINGLE_USE"
)

// BaseDTO ...
type BaseDTO struct {
	RequestID       *string `json:"requestId,omitempty"`
	XIdempotencyKey *string `json:"xIdempotencyKey,omitempty"`
}

// PaymentInstrument ...
type PaymentInstrument struct {
	Token                      string                      `json:"token"`
	AnalyticsID                *string                     `json:"analyticsId,omitempty"`
	TokenType                  *string                     `json:"tokenType,omitempty"`
	PaymentInstrumentType      *string                     `json:"paymentInstrumentType,omitempty"`
	PaymentInstrumentData      *interface{}                `json:"paymentInstrumentData,omitempty"`
	ThreeDSecureAuthentication *ThreeDSecureAuthentication `json:"threeDSecureAuthentication,omitempty"`
}

// PaymentCardToken ...
type PaymentCardToken struct {
	Last4Digits        string   `json:"last4Digits"`
	ExpirationMonth    string   `json:"expirationMonth"`
	ExpirationYear     string   `json:"expirationYear"`
	CardHolderName     *string  `json:"cardholderName,omitempty"`
	Network            *string  `json:"network,omitempty"`
	IsNetworkTokenized *bool    `json:"isNetworkTokenized,omitempty"`
	BinData            *BinData `json:"binData,omitempty"`
}

// BinData ...
type BinData struct {
	Network                    string  `json:"network"`
	RegionalRestriction        string  `json:"regionalRestriction"`
	AccountNumberType          string  `json:"accountNumberType"`
	AccountFundingType         string  `json:"accountFundingType"`
	PrepaidReloadableIndicator string  `json:"prepaidReloadableIndicator"`
	ProductUsageType           string  `json:"productUsageType"`
	ProductCode                string  `json:"productCode"`
	ProductName                string  `json:"productName"`
	IssuerCountryCode          *string `json:"issuerCountryCode,omitempty"`
	IssuerName                 *string `json:"issuerName,omitempty"`
	IssuerCurrencyCode         *string `json:"issuerCurrencyCode,omitempty"`
}

// ThreeDSecureAuthentication ...
type ThreeDSecureAuthentication struct {
	ResponseCode    string  `json:"responseCode"`
	ReasonCode      *string `json:"reasonCode,omitempty"`
	ReasonText      *string `json:"reasonText,omitempty"`
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
	ChallengeIssued *bool   `json:"challengeIssued,omitempty"`
}

// Customer ...
type Customer struct {
	ID              string           `json:"id"`
	Email           *string          `json:"email,omitempty"`
	BillingAddress  *BillingAddress  `json:"billingAddress,omitempty"`
	ShippingAddress *ShippingAddress `json:"shippingAddress,omitempty"`
}

// BillingAddress ...
type BillingAddress struct {
	Address
}

// ShippingAddress ...
type ShippingAddress struct {
	Address
}

// Address ...
type Address struct {
	AddressLine1 string  `json:"addressLine1"`
	City         string  `json:"city"`
	CountryCode  string  `json:"countryCode"`
	PostalCode   string  `json:"postalCode"`
	FirstName    *string `json:"firstName,omitempty"`
	LastName     *string `json:"lastName,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	State        *string `json:"state,omitempty"`
}

// Transaction ...
type Transaction struct {
	ID                     string        `json:"id"`
	Processor              string        `json:"processor"`
	TransactionType        string        `json:"type"`
	Status                 string        `json:"status"`
	ProcessorTransactionID *string       `json:"processorTransactionId,omitempty"`
	PaymentError           *PaymentError `json:"paymentError,omitempty"`
}

// PaymentError ...
type PaymentError struct {
	Date             time.Time `json:"date"`
	PaymentErrorType string    `json:"type"`
	DeclineCode      *string   `json:"declineCode,omitempty"`
	DeclineType      *string   `json:"declineType,omitempty"`
	ProcessorMessage *string   `json:"processorMessage,omitempty"`
}

// Payment ...
type Payment struct {
	ID                       string                    `json:"id"`
	Date                     time.Time                 `json:"date"`
	Status                   string                    `json:"status"`
	OrderID                  string                    `json:"orderId"`
	CurrencyCode             string                    `json:"currencyCode"`
	Amount                   int64                     `json:"amount"`
	AmountAuthorized         int64                     `json:"amountAuthorized"`
	AmountCapture            int64                     `json:"amountCaptured"`
	AmountRefunded           int64                     `json:"amountRefunded"`
	PaymentInstrument        PaymentInstrument         `json:"paymentInstrument"`
	Transactions             []Transaction             `json:"transactions"`
	Processor                *string                   `json:"processor,omitempty"`
	RequiredAction           *RequiredAction           `json:"requiredAction,omitempty"`
	StatementDescriptor      *string                   `json:"statementDescriptor,omitempty"`
	VaultedPaymentInstrument *VaultedPaymentInstrument `json:"vaultedPaymentInstrument,omitempty"`
	Customer                 *Customer                 `json:"customer,omitempty"`
	LastPaymentError         *PaymentError             `json:"lastPaymentError,omitempty"`
	Metadata                 *map[string]string        `json:"metadata"`
	WorkflowExecutionError   *WorkflowExecutionError   `json:"workflowExecutionError,omitempty"`
}

// RequiredAction ...
type RequiredAction struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ClientToken *string `json:"clientToken,omitempty"`
}

// VaultedPaymentInstrument ...
type VaultedPaymentInstrument struct {
	Token                      string                      `json:"token"`
	AnalyticsID                string                      `json:"analyticsId,omitempty"`
	TokenType                  string                      `json:"tokenType,omitempty"`
	PaymentInstrumentType      string                      `json:"paymentInstrumentType,omitempty"`
	PaymentInstrumentData      interface{}                 `json:"paymentInstrumentData,omitempty"`
	ThreeDSecureAuthentication *ThreeDSecureAuthentication `json:"threeDSecureAuthentication,omitempty"`
}

// WorkflowExecutionError ...
type WorkflowExecutionError struct {
	Reason string  `json:"reason"`
	StepId *string `json:"stepId,omitempty"`
}

// PaymentResponse ...
type PaymentResponse struct {
	BaseDTO
	Payment
}

func (b *BaseDTO) SetRequestID(id string) {
	b.RequestID = &id
}

// PaymentRequest ...
type CreatePayment struct {
	OrderID             string             `json:"orderId"`
	CurrencyCode        string             `json:"currencyCode"`
	Amount              string             `json:"amount"`
	PaymentInstrument   PaymentInstrument  `json:"paymentInstrument"`
	StatementDescriptor *string            `json:"statementDescriptor,omitempty"`
	Customer            *Customer          `json:"customer,omitempty"`
	Metadata            *map[string]string `json:"metadata,omitempty"`
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	BaseDTO
	CreatePayment
}

// CapturePayment ...
type CapturePayment struct {
	Amount int64 `json:"amount"`
	Final  bool  `json:"final"`
}

// CapturePaymentRequest ...
type CapturePaymentRequest struct {
	BaseDTO
	CapturePayment
}

// CancelPayment ...
type CancelPayment struct {
	Reason string `json:"reason"`
}

// CancelPaymentRequest ...
type CancelPaymentRequest struct {
	BaseDTO
	CancelPayment
}

// RefundPayment ...
type RefundPayment struct {
	Amount  int64   `json:"amount"`
	OrderID *string `json:"orderId,omitempty"`
	Reason  *string `json:"reason,omitempty"`
}

// RefundPaymentRequest ...
type RefundPaymentRequest struct {
	BaseDTO
	RefundPayment
}

type ResumePayment struct {
	ResumeToken string `json:"resumeToken"`
}

// ResumePaymentRequest ...
type ResumePaymentRequest struct {
	BaseDTO
	ResumePayment
}

// SearchPayment ...
type SearchPayment struct {
	Data       []Data  `json:"data"`
	NextCursor *string `json:"nextCursor,omitempty"`
	PrevCursor *string `json:"prevCursor,omitempty"`
}

// Data ...
type Data struct {
	ID           string    `json:"id"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
	OrderID      string    `json:"orderId"`
	CurrencyCode string    `json:"currencyCode"`
	Amount       int64     `json:"amount"`
	Processor    *string   `json:"processor,omitempty"`
}

// SearchPaymentResponse ...
type SearchPaymentResponse struct {
	BaseDTO
	SearchPayment
}

// GetPaymentRequest ...
type GetPaymentRequest struct {
	BaseDTO
}
