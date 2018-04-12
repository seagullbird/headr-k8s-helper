# k8s-helper

[![wercker status](https://app.wercker.com/status/2ac78bf6a9ffede5bb13758a616d4791/s/master "wercker status")](https://app.wercker.com/project/byKey/2ac78bf6a9ffede5bb13758a616d4791)

Project k8s-helper is the consumer of several k8s client related events.

It consumes these events from a rabbitMQ server and execute corresponding tasks.

These events include:

- New Site: Create a new caddy deployment together with its service.
