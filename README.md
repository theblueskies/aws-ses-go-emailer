# aws-ses-go-emailer

<img src="assets/mailman-gopher.png" width="30%">

## This is an emailer app. 

It currently sends emails through:
1. AWS SES (Simple Email Service)

Note: It expects FORM data and not JSON data.

It can be further extended to be used with Sendgrid, Mailchimp or other mail clients
by implementing the interface:    
   * ```EmailStore```

2. Build the service by executing:   
   ```make build```
