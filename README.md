# k8s-helper

Project k8s-helper is the consumer of several k8s client related events.

It consumes these events from a rabbitMQ server and execute corresponding tasks.

These events include:

- New Site: Create a new caddy deployment together with its service.
