# webhook-server
A naive implementation of webhook server


## Naive Version

The naive version will just contain a minimum of two servers - one is the __API server__ that will broadcast a payload upon certain event (e.g. create, update, delete) and another is one the __Webhook server__ that will receive the payload. It is just a simple `POST` request to the webhook server, and the payload can only be sent to a single client.

## Better Version

A better version is to create a separate API that will allow user's to select which events they can subscribe to, and what callback url the payload will be posted to.

The __API Server__ will be sending the message to a queue too, instead of a direct `POST` request. A worker will subscribe to the queue in order to process the payload, query the list of subcribers (callback urls) and sending them.

- webhook_api: allows users to create new webhook subscriptions
- webhook_server: the webhook will post to this server
- webhook_worker: the api server will send the payload through a queue to the webhook worker which will validate the payload through the webhook api and send the message to the webhook server

## Webhook API

The Webhook API allows users to subscribe to events and provide a callback url where the events will be posted.

| Endpoint | Description | 
|--        |--           |
| GET `/webhooks` | Get a list of webhook subscriptions | 
| POST `/webhooks` | Create a webhook with the list of events to be subscribed, and the callback url | 
| DELETE `/webhooks` | Clear all registered webhooks |
| GET `/webhooks/{id}` | Get the info for a specific webhook |
| PUT `/webhooks/{id}` | Update the info for a specific webhook |
| DELETE `/webhooks/{id}` | Delete a webhook by id |

## Webhook API Model

Webhook:

```json
{
	"id": "1",
	"created_at": "",
	"updated_at": "",
	"user_id": "",
	"is_verified": false, 
	"status": "active|error|stop",
	"callback_urls": [],
	"events": ["books:get", "user:create"], // Do we allow user's to subscribe to different resource topics?
	"invocation_count": 0,
	"version": "0.0.1"
}
```

Webhook Events:

```json
{
	"id": "",
	"name": "books:get",
	"service": "", // The service triggering this
	"count": 0,
	"created_at": "",
	"updated_at": "",
	"callback_url": "",
	"batch_size": 10, // Number of items per batch
	"error_count": 0, // 
	"retry_policy": {},
	"version": "git version"
}
```

Webhook API:

```json
{
	"name": "books api",
	"description": "books that serves api",
	// "events": [
	// 	"books:create",
	// 	"books:update",
	// 	"books:delete"
	// ],
	"events": [
		{
			"name": "books:create",
			"description": "",
			"created_at": "",
			"updated_at": "",
			"enabled": true,
			"payload": {},
			"metadata": {}
		}
	],
	"created_at": "",
	"updated_at": "",
	"version": ""
}
```

## Homogenous/Heteregenous Events

Homogenous events can be different events for the same resource, e.g. `books:get`, `books:create`.

Heteregenous events can be different events and different resources, e.g. `books:create`, `users:create`.

It is preferable to store each event and each resource in a new row to simplify query. Caching can be done through Redis too to reduce calls to the database, and each worker can just point to a redis cluster.

## Internal and External Webhook

If the webhook is open for public to consume (e.g. Slack, Github), then it will require certain authorization. The identity of the creator needs to be embedded during the creation of the webhook too.

## Security

Some thoughts and scenarios that could happen:

- can I register any URL?
- do I need to confirm the URL I am posting at?
- can I post to other user's URL (DDOS)?
- how can I verify once-only delivery?
- what if the server to be posted is down?
- can I subscribe multiple urls for the same topic?
- are there any retry policy?
- batching requests?
- how do I deal with changes to the event name