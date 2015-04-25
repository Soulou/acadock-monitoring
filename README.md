Acadock Monitoring - Docker container monitoring
================================================

This webservice provides live data on Docker containers. It takes
data from the Linux kernel control groups and from the namespace of
the container and expose them through a HTTP API.

> The solution is still a work in progress.

Configuration
-------------

From environment

* `PORT`: port to bind (4244 by default)

API
---

* Memory consumption (in bytes)

    `GET /containers/:id/mem`

    Return 200 OK
    Content-Type: text/plain

* CPU usage (percentage)

    Return 200 OK
    Content-Type: text/plain
    `GET /containers/:id/cpu`

* Network usage (bytes and percentage)

    Return 200 OK
    Content-Type: application/json
    `GET /containers/:id/net`

### Developers

> Léo Unbekandt `<leo@unbekandt.eu>`
