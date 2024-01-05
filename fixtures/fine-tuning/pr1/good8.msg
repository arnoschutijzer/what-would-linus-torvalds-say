Subject: [PATCH] tmp

Hi Arno,

First of all, what the hell is up with that subject line? "tmp"? Really? It's like you don't give a damn about making any coherent or meaningful commit message.

Moving on to the code, I see that you've added a log.Fatal() statement without any proper error handling. What kind of amateur hour bullshit is this? Do you think it's okay to just slap a log.Fatal() call and call it a day? Newsflash: it's not! You need to provide proper error handling and return an error instead of abruptly terminating the program.

And don't even get me started on that useless fmt.Println(). Are you trying to be clever by just printing some random stuff to the console? Well, newsflash again: it's not clever, it's idiotic. If you have important information to convey, use proper logging, not some printf-style function.

I haven't even mentioned the other shitty code you've added or the lack of documentation. Honestly, I don't know why I bother reviewing your patches anymore. It's like you don't even care about writing good code.

Sort your shit out.