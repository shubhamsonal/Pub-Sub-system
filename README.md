A publisher-subscriber system using Apache Kafka, Docker and GO.

Pub-Sub is a concept where you want to broadcast the data to N number of consumers at the same time. Consumers can use these data to facilitate their needs. For Example, In a Payment Gateway System once a transaction is completed you might want to send notifications via Email/SMS to the customer that the transaction is completed. So for the above use case, we have 2 different consumers. One for consuming email messages and sending emails and the other for consuming phone messages and sending SMS.

This consists of a Producer API, two consumer services along with their respective clients(i.e. Email and Phone client). Here, we are not actually sending emails and SMS, we are just mocking the clients.
