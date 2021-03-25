<h1>Work queues</h1>
<h3>Distributing tasks among workers (the competing consumers pattern)</h3>

<p>The main idea behind Work Queues (aka: Task Queues) is to avoid doing a resource-intensive task immediately and having to wait for it to complete. Instead we schedule the task to be done later. We encapsulate a task as a message and send it to a queue. A worker process running in the background will pop the tasks and eventually execute the job. When you run many workers the tasks will be shared between them.
   
   This concept is especially useful in web applications where it's impossible to handle a complex task during a short HTTP request window.</p>

![](https://www.rabbitmq.com/img/tutorials/python-two.png)

<h4>1. Message acknowledgment</h4>
<p>Doing a task can take a few seconds .You may wonder that what happens if one of the consumers starts along task and dies with it only partly done. With our current code ,once RabbitMQ delivers a message to the consumer it immediately marks it for deletion.In this case,if you kill a worker we will lose the message it was just processing.We'll aslo lose all the messages that were dispatched to this particular worker but were not yet handled.
<p>But we don't want to lose any tasks. If a worker dies, we'd like the task to be delivered to another worker.</p>
<p>In order to make sure a message is never lost, RabbitMQ supports message acknowledgments.An ack(nowledgement) is sent back by the consumer to tell RabbitMQ that a particular message has been received, processed and that RabbitMQ is free to delete it.</p>
<p>There aren't any message timeouts;RabbitMQ will redeliver the message when the consumer dies. It's fine even if processing a message takes a very, very long time.</p>
<p>In this tutorial we will use manual message acknowledgements by passing a false for the "auto-ack" argument and then send a proper acknowledgment from the worker with d.Ack(false) (this acknowledges a single delivery), once we're done with a task.</p>



<h4>2. Message durability</h4>
<p>We have learned how to make sure that even if the consumer dies, the task isn't lost. But our tasks will still be lost if RabbitMQ server stops.</p>
<p>When RabbitMQ quits or crashes it will forget the queues and messages unless you tell it not to. Two things are required to make sure that messages aren't lost: we need to mark both the queue and messages as durable.</p>
<p>First, we need to make sure that the queue will survive a RabbitMQ node restart. In order to do so, we need to declare it as durable(This durable option change needs to be applied to both the producer and consumer code.)</p>
<p>At this point we're sure that the task_queue queue won't be lost even if RabbitMQ restarts. Now we need to mark our messages as persistent - by using the amqp.Persistent option amqp.Publishing takes.</p>



<h4>3. Fair dispatch</h4>
![](https://www.rabbitmq.com/img/tutorials/prefetch-count.png)
<p>You might have noticed that the dispatching still doesn't work exactly as we want. For example in a situation with two workers, when all odd messages are heavy and even messages are light, one worker will be constantly busy and the other one will do hardly any work. Well, RabbitMQ doesn't know anything about that and will still dispatch messages evenly.</p>
<p>This happens because RabbitMQ just dispatches a message when the message enters the queue. It doesn't look at the number of unacknowledged messages for a consumer. It just blindly dispatches every n-th message to the n-th consumer.</p>
<p>In order to defeat that we can set the prefetch count with the value of 1. This tells RabbitMQ not to give more than one message to a worker at a time. Or, in other words, don't dispatch a new message to a worker until it has processed and acknowledged the previous one. Instead, it will dispatch it to the next worker that is not still busy.</p>