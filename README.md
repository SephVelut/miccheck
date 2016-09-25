**MicCheck**

MicCheck asserts expectations produced by consumers for producers.

On the consumer end a `specifier` is used to set expectations and how and when
those expectation will be fullfilled. Example: for a mock specifier, an expectation
key value map is assigned to an expected method call. If the method is called on the
mock, then it will consider that expectation fulfilled.

There are two `expectations`. One is the request expectation. The other is the response.
On the consumer end, only the request expectation is validated. On the producer side,
only the response expectation is validated.

When all expectation are fullfilled, a `contract writer` will serialize the expectations.
On the consumer end, only the request expectations are validated and the response expctations
will be serialized along with th requests for future use by the producer. The producer end behaves
in the inverse.

Specifiers can also come in th form of live working servers. They will behave in the same way as
a mock specifier, asserting certain behavior and the expectations to set.
