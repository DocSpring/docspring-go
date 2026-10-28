package docspring

// Checks webhook requests' X-DocSpring-Signature header (https://docspring.com/docs/webhooks/verifying-signatures/).

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// DefaultWebhookTolerance is how old (or early) a webhook request's timestamp can be (VerifyWebhookSignature).
const DefaultWebhookTolerance = 5 * time.Minute

// WebhookSignatureOptions changes how VerifyWebhookSignatureWithOptions checks the timestamp.
type WebhookSignatureOptions struct {
	// How old (or early) the timestamp can be. Zero means DefaultWebhookTolerance.
	Tolerance time.Duration
	// The current time, for tests. Zero means time.Now().
	Now time.Time
}

// VerifyWebhookSignature checks a webhook request's X-DocSpring-Signature header. It returns true when one of its v1
// signatures is the HMAC-SHA256 of "{t}.{payload}" with the webhook's secret, and its timestamp t is at most 5 minutes
// from now. It returns false for anything else, including a missing or malformed header.
//
// payload is the request body exactly as it was received (io.ReadAll(r.Body)), not parsed JSON.
func VerifyWebhookSignature(payload []byte, signatureHeader, secret string) bool {
	return VerifyWebhookSignatureWithOptions(payload, signatureHeader, secret, WebhookSignatureOptions{})
}

// VerifyWebhookSignatureWithOptions is VerifyWebhookSignature with another tolerance or current time.
func VerifyWebhookSignatureWithOptions(payload []byte, signatureHeader, secret string, options WebhookSignatureOptions) bool {
	if secret == "" {
		return false
	}

	timestamp := ""
	foundTimestamp := false
	var signatures []string
	for _, part := range strings.Split(signatureHeader, ",") {
		key, value, _ := strings.Cut(strings.TrimSpace(part), "=")
		switch {
		case key == "t" && !foundTimestamp:
			timestamp = value
			foundTimestamp = true
		case key == "v1" && value != "":
			signatures = append(signatures, value)
		}
	}
	if !isDigits(timestamp) || len(signatures) == 0 {
		return false
	}
	seconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false
	}

	tolerance := options.Tolerance
	if tolerance == 0 {
		tolerance = DefaultWebhookTolerance
	}
	now := options.Now
	if now.IsZero() {
		now = time.Now()
	}
	age := now.Unix() - seconds
	if age < 0 {
		age = -age
	}
	if age > int64(tolerance/time.Second) {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "."))
	mac.Write(payload)
	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	valid := false
	for _, signature := range signatures {
		// Constant time, so a timing attack can't find a valid signature byte by byte
		if hmac.Equal([]byte(signature), expected) {
			valid = true
		}
	}
	return valid
}

func isDigits(value string) bool {
	if value == "" {
		return false
	}
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
