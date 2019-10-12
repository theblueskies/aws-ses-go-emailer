# aws-ses-go-emailer

This is an emailer app. It currently sends emails through:
1. AWS SES (Simple Email Service)

It can be further extended to be used with Sendgrid, Mailchimp or other mail clients
by implementing the interface:    
   * ```EmailStore```

2. Build the service by executing:   
   ```make build```
