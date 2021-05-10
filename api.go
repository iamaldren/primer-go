package primer

import (
	"context"
	"fmt"
)

func (c *Client) CreatePayment(ctx context.Context, req *CreatePaymentRequest) (*PaymentResponse, error) {
	path := "/payments"
	return c.postRequest(ctx, req.CreatePayment, path, req.XIdempotencyKey)
}

func (c *Client) CapturePayment(ctx context.Context, req *CapturePaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/capture", paymentId)
	return c.postRequest(ctx, req.CapturePayment, path, req.XIdempotencyKey)
}

func (c *Client) CancelPayment(ctx context.Context, req *CancelPaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/cancel", paymentId)
	return c.postRequest(ctx, req.CancelPayment, path, req.XIdempotencyKey)
}

func (c *Client) RefundPayment(ctx context.Context, req *RefundPaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/refund", paymentId)
	return c.postRequest(ctx, req.RefundPayment, path, req.XIdempotencyKey)
}

func (c *Client) ResumePayment(ctx context.Context, req *ResumePaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/resume", paymentId)
	return c.postRequest(ctx, req.ResumePayment, path, req.XIdempotencyKey)
}

func (c *Client) GetPayment(ctx context.Context, req *GetPaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s", paymentId)

	resp := &PaymentResponse{}
	if err := c.get(ctx, strSafeDeref(req.XIdempotencyKey), path, nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) postRequest(ctx context.Context, request interface{}, path string, idempotencyKey *string) (*PaymentResponse, error) {
	if strSafeDeref(idempotencyKey) == "" {
		return nil, ErrXIdempotencyKeyMissing
	}

	resp := &PaymentResponse{}
	if err := c.post(ctx, strSafeDeref(idempotencyKey), path, request, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func strSafeDeref(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
