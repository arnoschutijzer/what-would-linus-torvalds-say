Subject: Re: [PATCH] Are you serious with this kindergarten-level code?

Right out the gate, Arno, you're using gorilla/mux merely for a single http endpoint? That's like hiring a dragon to guard a fishbowl. All of those dependencies for something http.NewServeMux() could handle inexpensively, did you accidentally left your efficiency in your other pants?

Your restructuring of the main() function into createServer() and main() is an unnecessary load of crap. You're adding bloated pretentious code without functionality improvement, you've merely added hops for the call stack. You might take yourself seriously, but ain't nobody else going to look at this oafish attempt of yours and think the same!

Your NULL body check in handleVersion() is a wasteful joke. You send a "hello world" message (a bloody cliché response if I ever saw one) as the default case, and all that does is handle error cases in the most ambigious way possible. Use proper response codes and descriptive error messages if you want your code to handle those cases elegantly.

In terms of testing, you're technically 'testing' the server, but the way you've set it up is like a pre-schooler's drawing; you know what it's supposed to be, but it doesn't really look like that. You veered off into the land of idiocy when you wrote assert.HTTPBodyContains() to try to test its output. Tests should return a carefully crafted response, not printing bloody "hello world" like a toddler showing off to his parents.

To add insult to injury, you've also added quite a number of dependencies in the go.mod file. This is not a kitchen sink Arno. You don't have to dump everything you've found on the internet into this file. Stick to the point, keep it lean and mean. Your go.mod file is now thicker than a bible, and provide less value than a footballer's memoir. We're programmers, not bloody hoarders.