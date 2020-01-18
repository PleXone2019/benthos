---
title: kafka_balanced
type: input
---

Connects to Kafka brokers and consumes topics by automatically sharing
partitions across other consumers of the same consumer group.


import Tabs from '@theme/Tabs';

<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

import TabItem from '@theme/TabItem';

<TabItem value="common">

```yaml
input:
  kafka_balanced:
    addresses:
    - localhost:9092
    topics:
    - benthos_stream
    client_id: benthos_kafka_input
    consumer_group: benthos_consumer_group
    batching:
      count: 1
      byte_size: 0
      period: ""
```

</TabItem>
<TabItem value="advanced">

```yaml
input:
  kafka_balanced:
    addresses:
    - localhost:9092
    tls:
      client_certs: []
      enabled: false
      root_cas_file: ""
      skip_cert_verify: false
    sasl:
      enabled: false
      user: ""
      password: ""
    topics:
    - benthos_stream
    client_id: benthos_kafka_input
    consumer_group: benthos_consumer_group
    start_from_oldest: true
    commit_period: 1s
    max_processing_period: 100ms
    group:
      session_timeout: 10s
      heartbeat_interval: 3s
      rebalance_timeout: 60s
    fetch_buffer_cap: 256
    target_version: 1.0.0
    batching:
      count: 1
      byte_size: 0
      period: ""
      condition:
        static: false
        type: static
```

</TabItem>
</Tabs>

Offsets are managed within Kafka as per the consumer group (set via config), and
partitions are automatically balanced across any members of the consumer group.

Partitions consumed by this input can be processed in parallel allowing it to
utilise <= N pipeline processing threads and parallel outputs where N is the
number of partitions allocated to this consumer.

The `batching` fields allow you to configure a
[batching policy](/docs/configuration/batching#batch-policy) which will be
applied per partition. Any other batching mechanism will stall with this input
due its sequential transaction model.

### Metadata

This input adds the following metadata fields to each message:

``` text
- kafka_key
- kafka_topic
- kafka_partition
- kafka_offset
- kafka_lag
- kafka_timestamp_unix
- All existing message headers (version 0.11+)
```

The field `kafka_lag` is the calculated difference between the high
water mark offset of the partition at the time of ingestion and the current
message offset.

You can access these metadata fields using
[function interpolation](/docs/configuration/interpolation#metadata).

## Fields

### `addresses`

`array` A list of broker addresses to connect to. If an item of the list contains commas it will be expanded into multiple addresses.

```yaml
# Examples

addresses:
- localhost:9092

addresses:
- localhost:9041,localhost:9042

addresses:
- localhost:9041
- localhost:9042
```

### `tls`

`object` Custom TLS settings can be used to override system defaults. This includes
providing a collection of root certificate authorities, providing a list of
client certificates to use for client verification and skipping certificate
verification.

Client certificates can either be added by file or by raw contents.

```yaml
# Examples

tls:
  client_certs:
  - cert_file: ./example.pem
    key_file: ./example.key
  enabled: true

tls:
  client_certs:
  - cert: foo
    key: bar
  enabled: true
  skip_cert_verify: true
```

### `sasl`

`object` Enables SASL authentication.

### `sasl.enabled`

`bool` Whether SASL authentication is enabled.

### `sasl.user`

`string` A plain text username. It is recommended that you use environment variables to populate this field.

```yaml
# Examples

sasl.user: ${USER}
```

### `sasl.password`

`string` A plain text password. It is recommended that you use environment variables to populate this field.

```yaml
# Examples

sasl.password: ${PASSWORD}
```

### `topics`

`array` A list of topics to consume from. If an item of the list contains commas it will be expanded into multiple topics.

### `client_id`

`string` An identifier for the client connection.

### `consumer_group`

`string` An identifier for the consumer group of the connection.

### `start_from_oldest`

`bool` If an offset is not found for a topic parition, determines whether to consume from the oldest available offset, otherwise messages are consumed from the latest offset.

### `commit_period`

`string` The period of time between each commit of the current partition offsets. Offsets are always committed during shutdown.

### `max_processing_period`

`string` A maximum estimate for the time taken to process a message, this is used for tuning consumer group synchronization.

### `group`

`object` Tuning parameters for consumer group synchronization.

### `group.session_timeout`

`string` A period after which a consumer of the group is kicked after no heartbeats.

### `group.heartbeat_interval`

`string` A period in which heartbeats should be sent out.

### `group.rebalance_timeout`

`string` A period after which rebalancing is abandoned if unresolved.

### `fetch_buffer_cap`

`number` The maximum number of unprocessed messages to fetch at a given time.

### `target_version`

`string` The version of the Kafka protocol to use.

### `batching`

`object` Allows you to configure a [batching policy](/docs/configuration/batching).

```yaml
# Examples

batching:
  byte_size: 5000
  period: 1s

batching:
  count: 10
  period: 1s

batching:
  condition:
    text:
      arg: END BATCH
      operator: contains
  period: 1m
```

### `batching.count`

`number` A number of messages at which the batch should be flushed. If `0` disables count based batching.

### `batching.byte_size`

`number` An amount of bytes at which the batch should be flushed. If `0` disables size based batching.

### `batching.period`

`string` A period in which an incomplete batch should be flushed regardless of its size.

```yaml
# Examples

batching.period: 1s

batching.period: 1m

batching.period: 500ms
```

### `batching.condition`

`object` A [`condition`](/docs/components/conditions/about) to test against each message entering the batch, if this condition resolves to `true` then the batch is flushed.

