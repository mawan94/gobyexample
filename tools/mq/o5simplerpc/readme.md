<h1>RPC</h1>
<h3>Request / reply pattern </h3>
<p>
In this tutorial we're going to use RabbitMQ to build an RPC system: a client and a scalable RPC server. As we don't have any time-consuming tasks that are worth distributing, we're going to create a dummy RPC service that returns Fibonacci numbers.</p>


<h4>Summary</h4>
![](https://www.rabbitmq.com/img/tutorials/python-six.png)

<div>Our RPC will work like this:</div>
<ul>
    <li>When the Client starts up, it creates an anonymous exclusive callback queue.
    
</li>
    <li>For an RPC request, the Client sends a message with two properties: reply_to, which is set to the callback queue and correlation_id, which is set to a unique value for every request.
</li>
    <li>The request is sent to an rpc_queue queue.
</li>
    <li>The RPC worker (aka: server) is waiting for requests on that queue. When a request appears, it does the job and sends a message with the result back to the Client, using the queue from the reply_to field.</li>
    <li>The client waits for data on the callback queue. When a message appears, it checks the correlation_id property. If it matches the value from the request it returns the response to the application.
</li>
</ul>