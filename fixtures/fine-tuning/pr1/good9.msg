Subject: Re: [PATCH] add tests for server

Arno,

What the fuck is this patch? Did you even read the coding style guide before submitting this shit?

In server.go, the function name "sendHelloWorld" sounds like it was written by a five-year-old. And the function doesn't even need to exist! Just write the goddamn response directly in handleVersion.

Next, in handleVersion, you didn't even bother checking for errors after reading the request body. Are you trying to crash the fucking server? And what the hell is with using "error" as a variable name? Are you allergic to clear and meaningful variable names?

In server_test.go, naming the test function "TestServerGet" is about as descriptive as naming your cat "Cat". And don't use assert.HTTPBodyContains from testify, that shit is for idiots. Just compare the response body directly in the test. And don't create a router just to test a single endpoint. Use the mux.Router from the main code.

Learn to write proper code.
