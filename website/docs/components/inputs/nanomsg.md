---
title: nanomsg
type: input
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the contents of:
     lib/input/nanomsg.go
-->


Consumes messages via Nanomsg sockets (scalability protocols).


import Tabs from '@theme/Tabs';

<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

import TabItem from '@theme/TabItem';

<TabItem value="common">

```yaml
input:
  nanomsg:
    urls:
    - tcp://*:5555
    bind: true
    socket_type: PULL
    sub_filters: []
```

</TabItem>
<TabItem value="advanced">

```yaml
input:
  nanomsg:
    urls:
    - tcp://*:5555
    bind: true
    socket_type: PULL
    sub_filters: []
    poll_timeout: 5s
```

</TabItem>
</Tabs>

Currently only PULL and SUB sockets are supported.

## Fields

### `urls`

`array` A list of URLs to connect to (or as). If an item of the list contains commas it will be expanded into multiple URLs.

### `bind`

`bool` Whether the URLs provided should be connected to, or bound as.

### `socket_type`

`string` The socket type to use.

Options are: `PULL`, `SUB`.

### `sub_filters`

`array` A list of sub filters to use when consuming from a SUB socket.

### `poll_timeout`

`string` The period to wait until a poll is abandoned and reattempted.


