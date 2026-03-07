# Batch requests

Using Batch requests feature, you can send **up to 500** REST API requests increasing your server performance.&#x20;

{% hint style="info" %}
Requests inside a batch are executed strictly in sequence — FIFO. This means that until the earlier batches are fully processed, the later ones will not start.
{% endhint %}

You can speed up the completion of batch requests by passing the `allowParallelMode=true` query param. In that case, up to **100** requests from your batch will be processed simultaneously.
