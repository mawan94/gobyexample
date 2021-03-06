<h1>Topics</h1>
<h3>Receiving messages based on a pattern </h3>


<h4>Topic exchange</h4>
<p>Messages sent to a topic exchange can't have an arbitrary routing_key - it must be a list of words, delimited by dots. The words can be anything, but usually they specify some features connected to the message. A few valid routing key examples: "stock.usd.nyse", "nyse.vmw", "quick.orange.rabbit". There can be as many words in the routing key as you like, up to the limit of 255 bytes.
   
   The binding key must also be in the same form. The logic behind the topic exchange is similar to a direct one - a message sent with a particular routing key will be delivered to all the queues that are bound with a matching binding key. However there are two important special cases for binding keys:
   
   \* (star) can substitute for exactly one word.<br>
  \# (hash) can substitute for zero or more words.<br><br>
   It's easiest to explain this in an example:</p>
   
![](https://www.rabbitmq.com/img/tutorials/python-five.png)
<p>In this example, we're going to send messages which all describe animals. The messages will be sent with a routing key that consists of three words (two dots). The first word in the routing key will describe speed, second a colour and third a species: "<speed>.<colour>.<species>".
   
   We created three bindings: Q1 is bound with binding key "\*.orange.\*" and Q2 with "\*.*.rabbit" and "lazy.#".
   
   These bindings can be summarised as:
   
   Q1 is interested in all the orange animals.<br>
   Q2 wants to hear everything about rabbits, and everything about lazy animals.</p>
   
<p>A message with a routing key set to "quick.orange.rabbit" will be delivered to both queues. Message "lazy.orange.elephant" also will go to both of them. On the other hand "quick.orange.fox" will only go to the first queue, and "lazy.brown.fox" only to the second. "lazy.pink.rabbit" will be delivered to the second queue only once, even though it matches two bindings. "quick.brown.fox" doesn't match any binding so it will be discarded.
   
   What happens if we break our contract and send a message with one or four words, like "orange" or "quick.orange.male.rabbit"? Well, these messages won't match any bindings and will be lost.
   
   On the other hand "lazy.orange.male.rabbit", even though it has four words, will match the last binding and will be delivered to the second queue.</p>
