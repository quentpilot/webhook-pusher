# webhook-pusher
Webhooks handler to avoid pain implementations

# Goal
Stop throwing your webhooks around the Web without any transfer assurance: Blind Mode Avoided.

- Get an advanced retry system in case endpoint fails.
- Follow all your webhooks from a web interface
- Manually inspect and retry each webhook
- Test your webhook integration

# API Integration
## Quick Start
- Create your account here: https://localhost:8001/register then log-in
- Navigate to https://localhost:8001/account/api to manage your API key and quick start tutorial
- Navigate to https://localhost:8001/api/playground to sandbox API features

### API call example
``` bash
curl -X POST -H "Authorization: Bearer your_api_key" --location https://localhost:42001/send -d '{"event", "cutomer.pay", "target": "https://domain.com/api/v2/webhook", "retry": 3, "fallback": "https://domain.com/api/v1/legacy_webhook", "payload": "{"event": "payday", "uuid": "uI4seR56Xs99", "price": 59,99, "currency": "USD"}"}'
```

## Documentation
- Navigate to https://localhost:8001/api/swagger to retrieve a Swagger documentation
- Navigate to https://localhost:8001/api/doc to retrieve complete documentation

## Main Endpoint Documentation
### Send Webhook

<details>
 <summary><code>POST</code> <code><b>/send</b></code> <code>(Persist and send webhook to your target)</code></summary>

#### Headers
> | name            |  type                 | description                              |
> |---------------  |-----------------------|----------------------------------------- |
> | Authorization   | string / Bearer Token |  Bearer token created from your account  |


#### Parameters (Json)

> | name        |  type                | description                                              |  required
> |-----------  |----------------------|----------------------------------------------------------| ----------
> | event       |  string              |  The webhook target type. Used to sort and find them     | ✅
> | target      |  string              |  The user endpoint you want to send data                 | ✅
> | payload     |  json                |  Json object data to transfert                           | ✅
> | throttle    |  integer             |  Maximum send per second. To avoid spam/DDOS end-user    | ❌
> | retry       |  integer             |  Maximum send retry before give up                       | ❌
> | fallback    |  string              |  Fallback endpoint to send if retry gives up             | ❌
> | sentry      |  boolean             |  Wheter to be notified if a webhook totally failed       | ❌


#### Responses (Json)

> | http code     | response                                                        |
> |---------------|-----------------------------------------------------------------|
> | `202`         | `{"message": "Webhook Accepted","uuid": "{uuid}"}`              |
> | `400`         | `{"message": "Invalid parameter", "parameter": "param_name"}`   |
> | `401`         | `{"message": "Invalid auth token"}`                             |
> | `429`         | `{"message": "Too many requests"}`                              |
> | `503`         | `{"message": "Service Unavailable"}`                            |
> | `504`         | `{"message": "Request Timeout"}`                                |
> | `507`         | `{"message": "Plan storage reached"}`                           |
> | `509`         | `{"message": "Plan limit reached"}`                             |


</details>

# UI Monitoring
- Connect with your account here: https://localhost:8001/login
- Navigate to https://localhost:8001/traffic to list all webhooks published
- Navigate to https://localhost:8001/traffic to list all webhooks published

# Self-Hosted Installation
Clone this repository and start with Docker in your development environment.