Hello fellow Nerdies!

So... Let's have fun with Go once more. This time it is with creating a Desktop app! I've found this fun tutorial, which I will be following - well... Sort of anyway. The link is here: 
https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1 (you should always give credit where credit due + plus there might be more details at the source - so give it a click => it is a 10 part series worth exploring).

Back to business! (okay - back up. I really like that source! Seems like going the same direction as me => having fun for fun's sake and document it. A link to the Github repo as well. https://github.com/scraly/learning-go-by-examples)
Ehrm... Now I think I am ready to start messing things up. :D

Some of the few interesting commands to follow
<pre>
-------------------------------------------------------
|                                                     |
| $go mod init Zent.GoGui/havingFun                   |
| //It creates our module file for dependencies etc.  |
|                                                     |
| $go get fyne.io/fyne/v2/cmd/fyne                    |
| //The guide said that we should require only one    |
| //thing, whih is fyneV2 v2.0.4                      |
| //I got 17 dependencies and fyne is on v2.3.5       |
| //on the other hand the guide is working with       |
| //Go v1.16, we I think a lot of contributers have   |
| //joined the universe.                              |
|                                                     |
-------------------------------------------------------
</pre>
