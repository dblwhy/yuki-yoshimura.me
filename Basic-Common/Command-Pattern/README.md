# Practical Command Pattern in Web Backend Development

While practical examples of the Command Pattern in web backend development are scarce, I've come up with an approach to utilize it within a Pub/Sub Architecture.

## Note

This Golang repository is structured based on [Layered Architecture](https://www.yuki-yoshimura.me/tech-stack-insights/theory-and-real-world-triumphs) and follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Please note that this code is not fully implemented and won't run as-is. Its primary purpose is to share insights and conceptual understanding.

## Explanation

Let's imagine we need to send an registration completion email to a user after processing `POST /user` request, or handle any kind of asynchronous event message typically managed in `Application Logic (a.k.a Workflow Logic) Layer`.

The goal is to decouple event handling logic from the backend server and process it within AWS Lambda, triggered when AWS SQS receives a specific message. This approach can reduce memory pressure on the backend server and decrease overall complexity.

### Core Concept: Command Service

The core of this implementation is the `Command Service`. All asynchronous event logic is expected to reside in the `service/command` package, implementing the `Command` interface. This setup allows the `Domain layer` to create the various command objects and pass it to an Invoker, which then publishes the message to SQS. Subsequently, Lambda is triggered, unmarshaling the SQS message into a Command object and simply executing its `Execute` method.
