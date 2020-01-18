---
title: redis_pubsub
type: output
---

```yaml
redis_pubsub:
  channel: benthos_chan
  max_in_flight: 1
  url: tcp://localhost:6379
```

Publishes messages through the Redis PubSub model. It is not possible to
guarantee that messages have been received.

This output will interpolate functions within the channel field, you
can find a list of functions [here](/docs/configuration/interpolation#functions).

This output benefits from sending multiple messages in flight in parallel for
improved performance. You can tune the max number of in flight messages with the
field `max_in_flight`.

