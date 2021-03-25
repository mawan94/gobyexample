<h1>Routing</h1>
<h3>Receiving messages selectively </h3> 
<p>In this tutorial we're going to add a feature to it - we're going to make it possible to subscribe only to a subset of the messages. For example, we will be able to direct only critical error messages to the log file (to save disk space), while still being able to print all of the log messages on the console.</p>

<h4>1. Bindings</h4>
<p>A binding is a relationship between an exchange and a queue. This can be simply read as: the queue is interested in messages from this exchange.
   
   Bindings can take an extra routing_key parameter. To avoid the confusion with a Channel.Publish parameter we're going to call it a binding key. This is how we could create a binding with a key:</p>
<p>The meaning of a binding key depends on the exchange type. The fanout exchanges, which we used previously, simply ignored its value.</p>

<h4>2. Direct exchange</h4>
<p>Our logging system from the previous tutorial broadcasts all messages to all consumers. We want to extend that to allow filtering messages based on their severity. For example we may want the script which is writing log messages to the disk to only receive critical errors, and not waste disk space on warning or info log messages.
   
   We were using a fanout exchange, which doesn't give us much flexibility - it's only capable of mindless broadcasting.
   
   We will use a direct exchange instead. The routing algorithm behind a direct exchange is simple - a message goes to the queues whose binding key exactly matches the routing key of the message.
   
   To illustrate that, consider the following setup:</p>
![](https://www.rabbitmq.com/img/tutorials/direct-exchange.png)
<p>In this setup, we can see the direct exchange X with two queues bound to it. The first queue is bound with binding key orange, and the second has two bindings, one with binding key black and the other one with green.
   
   In such a setup a message published to the exchange with a routing key orange will be routed to queue Q1. Messages with a routing key of black or green will go to Q2. All other messages will be discarded.</p>

<h4>3. Multiple bindings</h4>
![](https://www.rabbitmq.com/img/tutorials/direct-exchange-multiple.png)
<p>It is perfectly legal to bind multiple queues with the same binding key. In our example we could add a binding between X and Q1 with binding key black. In that case, the direct exchange will behave like fanout and will broadcast the message to all the matching queues. A message with routing key black will be delivered to both Q1 and Q2.</p>

