Acadock Monitoring - Docker container monitoring
================================================

This webservice provides live data on Docker containers.

Configuration from env
-----------------------

* `PORT`: port to bind

API
---

* Memory consumption (in bytes)

    `GET /containers/:id/mem`

* CPU usage (percentage)

    `GET /containers/:id/cpu`

* Network usage (bytes and percentage)

    `GET /containers/:id/net`

### Misc

The service binds the port 4244

### Developers

> Léo Unbekandt `<leo@unbekandt.eu>`

### Reference

This project is used by [Acadock](https://github.com/Soulou/acadock)

It was bootstrapped during the hackathon [Hack Le Chalet 2013](http://hacklechalet.com)
