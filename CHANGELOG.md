### 2.0.0 [February 23, 2025]

- **BREAKING CHANGE**: Updated default host to our new synchronous API subdomain: sync.api.docspring.com. (EU customers should use sync.api-eu.docspring.com). Removed all custom polling code from library since this logic is now handled by the API service running on our sync subdomain

### 1.0.1 [November 28, 2021]

- Added 'password' parameter to Generate PDF and Combine PDF API endpoints. You can now encrypt each generated PDF with a unique password.

### 1.0.0 [Aug 28, 2021]

- Initial release of Go API client
