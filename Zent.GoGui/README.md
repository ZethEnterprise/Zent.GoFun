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
| $go get .                                           |
| //Fun fact: calling the "Go Get fyne.io/fyne/v2"    |
| //did not fetch any dependencies. Evedently this    |
| //does.                                             |
|                                                     |
| $apt-get install pkg-config                         |
| $apt-get install xorg-dev                           |
| //aparently I did not have the required elements    |
| //this to run at first.                             |
-------------------------------------------------------
</pre>
I tried just to test out the constant link, which did not make any sense (as it gave a 404). It did however look like it was a TinyGo server and possibly running all the time. Therefore I 
just tried to see if I could get the example here to work out of the box, before I played with it for my own enjoyment.

Evedently it seemed like my environment was not up for the requirements. First off the "get dependencies" of fyne did not work as intented. Next a clean "run" of the main.go proved to be
less than successfull. Well for starters it showed me all the "go get xxxxx" I was lacking; second it complained about some missing pkg-config; fixing that it left me with one last thing.
This thing seems to be a missing "xorg.dev"-thingy. Installing those and I actually got to run the main.go file (while writing this last sentence I just figured out what I am doing right now:
I am doing a "react programming/vlog-thingy" in a Git-repo of this tutorial. How fun is that? :'D ).

Fun thing happened. It failed to decode the image. Oh well, now it is time to make changes and see why it failed to decode the image. So - I think I've found out why this was not working.
It might be that my Go version and Fyne version is too new and therefore it cannot work together. Luckily it seems like this is another project of scraly, which means I can try out to set
up the same server and try to run it with that one as my target.

