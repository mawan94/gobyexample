<h1>Publish/Subscribe</h1>
<h3>Sending messages to many customers at once </h3>
<p>To illustrate the pattern, we're going to build a simple logging system. It will consist of two programs -- the first will emit log messages and the second will receive and print them.
   
   In our logging system every running copy of the receiver program will get the messages. That way we'll be able to run one receiver and direct the logs to disk; and at the same time we'll be able to run another receiver and see the logs on the screen.
   
   Essentially, published log messages are going to be broadcast to all the receivers.</p>
   
   
<h4>1. Exchanges</h4>
<p>The core idea in the messaging model in RabbitMQ is that the producer never sends any messages directly to a queue. Actually, quite often the producer doesn't even know if a message will be delivered to any queue at all.
   
   Instead, the producer can only send messages to an exchange. An exchange is a very simple thing. On one side it receives messages from producers and the other side it pushes them to queues. The exchange must know exactly what to do with a message it receives. Should it be appended to a particular queue? Should it be appended to many queues? Or should it get discarded. The rules for that are defined by the exchange type.</p>

![](https://www.rabbitmq.com/img/tutorials/exchanges.png)
<p>There are a few exchange types available: direct, topic, headers and fanout. We'll focus on the last one -- the fanout. Let's create an exchange of this type, and call it logs:</p>

<p>The fanout exchange is very simple. As you can probably guess from the name, it just broadcasts all the messages it receives to all the queues it knows. And that's exactly what we need for our logger.</p>

<h4>2. Temporary queues</h4>
<p>We want to hear about all log messages, not just a subset of them. We're also interested only in currently flowing messages not in the old ones. To solve that we need two things.
   
   Firstly, whenever we connect to Rabbit we need a fresh, empty queue. To do this we could create a queue with a random name, or, even better - let the server choose a random queue name for us.
   
   Secondly, once we disconnect the consumer the queue should be automatically deleted.
   
   In the amqp client, when we supply queue name as an empty string, we create a non-durable queue with a generated name:</p>
   
<p>When the method returns, the queue instance contains a random queue name generated by RabbitMQ. For example it may look like amq.gen-JzTY20BRgKO-HjmUJj0wLg.
   
   When the connection that declared it closes, the queue will be deleted because it is declared as exclusive.</p>
   
<h4>3. Bindings</h4>
![](https://www.rabbitmq.com/img/tutorials/bindings.png)
<p>We've already created a fanout exchange and a queue. Now we need to tell the exchange to send messages to our queue. That relationship between exchange and a queue is called a binding</p>

<h4>4. Putting it all together</h4>
![](https://www.rabbitmq.com/img/tutorials/python-three-overall.png)
<p>The producer program, which emits log messages, doesn't look much different from the previous tutorial. The most important change is that we now want to publish messages to our logs exchange instead of the nameless one. We need to supply a routingKey when sending, but its value is ignored for fanout exchanges.</p>