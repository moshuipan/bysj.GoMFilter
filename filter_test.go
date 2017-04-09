package main

import (
	//	"fmt"
	"testing"
)

func Test_genWordMap(t *testing.T) {
	text := `HTTP/2 is designed to address many of the failings 
	of HTTP/1.x. Modern web pages use many resources: HTML, stylesheets, scripts, 
	images, and so on. In HTTP/1.x, each of these resources must be requested explicitly. 
	This can be a slow process. The browser starts by fetching the HTML, then learns of 
	more resources incrementally as it parses and evaluates the page. Since the server 
	must wait for the browser to make each request, the network is often idle and underutilized.
	=========================own test========================================================
	products is decrese 12%; and $100.i don't konw anythings.`
	wordmap := genWordMap(text)
	t.Log(wordmap)
}

func Test_CountWordMap(t *testing.T) {
	spam := "E:\\mygo\\src\\bysj.GoMFilter\\train\\spam"
	_, n := CountWordMap(spam)
	t.Log("spam files:", n)
	ham := "E:\\mygo\\src\\bysj.GoMFilter\\train\\ham"
	_, n = CountWordMap(ham)
	t.Log("ham files:", n)
	t.Log("TestTrainWordMap end!")
}

func Test_Filter(t *testing.T) {
	InitTrain()
	datas := `Subject: your membership community charset = iso - 8859 - 1
your membership community & commentary ( july 6 , 2001 )
it ' s all about making money
information to provide you with the absolute
best low and no cost ways of providing traffic
to your site , helping you to capitalize on the power and potential the web brings to every net - preneur .
- - - this issue contains sites who will trade links with you ! - - -
- - - - - - - - - - - - -
in this issue
- - - - - - - - - - - - -
internet success through simplicity
member showcase
win a free ad in community & commentary
| | | = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = > >
today ' s special announcement :
| | | = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = > >
we can help you become an internet service provider within 7 days or we will give you $ 100 . 00 ! !
click here
we have already signed 300 isps on a 4 year contract ,
see if any are in your town at :
click here
you are a member in at least one of these programs
- you should be in them all !
bannersgomlm . com
profitbanners . com
cashpromotions . com
mysiteinc . com
timshometownstories . com
freelinksnetwork . com
myshoppingplace . com
bannerco - op . com
putpeel . com
putpeel . net
sellinternetaccess . com
be - your - own - isp . com
seventhpower . com
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
internet success through simplicity
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
every day of the week , i get questions from people all
over the world , including my no bs gimg members ,
wanting to know some of the most valuable " secrets "
to my on - going internet success .
let me say , above all else , i don ' t believe there are any
* true * " secrets " to success on the net . what you do to
become successful in the online world is not a " secret " ,
in my opinion . most successful people follow simple , clear ,
repeatedly - proven strategies to succeed , whether on the
net or off .
but , when it comes to someone asking for advice ,
consultation , or simply asking , " what ' s your secret ? " ,
i have to blush & say . . .
persistence and personality .
of course , i always follow the advice with my own little
disclaimer : what makes me successful may not work the
same for you . . . & your first lesson is to get over
the deep - seeded idea that success - of any kind , in
my opinion - is somehow an unknown , unattainable secret .
clearly , it is not . it ' s not unknown . it ' s not unattainable .
it ' s not years of digging to find the " secrets " to internet riches .
one thing that " gets to me " so often in my work as an
internet consultant , author and internet success
strategist is that so many people on the net seem to
have this incredibly huge mental block that stands
between themselves and success on the net . it ' s
almost as if they ' ve been barraged by so many claims
of what works and what doesn ' t work , and so many
long , complicated routes to actually succeeding in
their online venture , that " success " is the equivelant of a 100 - foot high brick wall .
it ' s not that difficult , my friends ! it is not that complicated ! !
long - time friend and business associate rick beneteau
has a new ebook out called branding you & breaking
the bank . get it ! !
http : / / www . roibot . com / bybb . cgi ? im 7517 _ bybtb .
but , the reason i mention this is the fact that he talks
so dynamically about the true simplicity of making your
online venture a success .
and , yes , rick & i come from the same school of
" self marketing " - marketing you ! obviously , that ' s
the core of his excellent new ebook , and i couldn ' t
agree with him more .
point being , * you * are everything you do online to
succeed . you are your web site , your business , your
marketing piece , your customer service , your customers '
experiences with your business - - all of it , is you !
read his ebook & you ' ll see more of what i ' m saying .
the matter at hand is that brick wall you might have
standing high as you can see , blocking the path
between you & internet success . listen to me - it is
not real ok ? it doesn ' t exist . there ' s nothing there
to fear to begin with . . . get over it ! !
what i ' m telling you is , the only thing standing between
you and the success you most desire . . . is yourself .
when you realize this , you will tear down that brick
wall by means of complete and instantaneous
disintegration . it will no longer exist * in your mind * ,
which is the only " real " place it ever was anyhow !
yes , " persistence and personality " inherently includes
honesty , integrity , accountability , and many other
qualities but you also have to hone in on your ultimate
goals and realize that probably the most valuable ,
powerful key to your success . . . is you !
that may be the most incredible " secret " we ever
uncover in our lifetime ! and , trust me , that brick wall
won ' t ever get in your way again . . . unless you let it .
talk about simple ! !
bryan is a " veteran " internet consultant , author ,
internet success strategist & marketer . he publishes
mega - success . com chronicles to over 11 , 500 subscribing
members , authors articles which appear all over the
net , and helps hundreds of wealth - hungry people in
their journey to internet success .
bryan is also director of his no bs guerrilla internet
marketing group at http : / / . com
& a fantastic new joint venture partners program
for that site .
bryan hall is a founding member and the development
consultant for the prestigious icop ( tm ) at
http : / / www . i - cop . org / 1016 . htm
you can reach bryan at 877 . 230 . 3267 or by
emailing him directly at bryan . hall @ mega - success . com
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
member showcase
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
examine carefully - those with email addresses included will
trade links with you . . . you are encouraged to contact them .
there are many ways to build a successful business - just look at
these successful sites programs other members are involved
in . . .
get insane amounts of traffic to your website .
purchase 10 , 000 guaranteed visitors to your site
and receive 5 , 000 free . more traffic = more money !
less than 2 cents a visitor . space is limited .
order now ! http : / / www . freepicklotto . com
trade links - businessopps @ aol . com
stop smoking - free lesson ! !
discover the secret to stopping smoking .
to master these powerful techniques , come to
http : / / www . breath - of - life . net
for your free lesson .
act now ! p . s . tell someone you care about .
trade links - jturco 3 @ hotmail . com
celebration sale !
$ 99 . 00 on casinos / sportsbetting sites , lingerie stores ,
gift stores , adult sites toy stores .
mention ad # bmlm 99 to receive this special sale price .
order now !
http : / / www . cyberopps . com / ? = bmlm 99
affiliates of the world !
top rated affiliate programs , excellent business opportunities ,
great marketing resources and free advertising for you !
visit the site to trade links . http : / / www . affiliates . uk . com
trade links - adrianbold @ affiliates . uk . com
just been released ! !
internet marketing guru corey rudl has just released a
brand new version of his # 1 best - selling internet marketing
course , " the insider secret ' s to marketing your business on
the internet " . a must have ! so don ' t hesitate ,
visit . . http : / / www . adminder . com / c . cgi ? startbgmlmezine
we have a 260 page catalog with over 3000 gift items for men ,
women , children - a gift for everyone . we show 100 gift items
on our web site alone , with the catalog you have access to
the rest . we also feel we have the best prices on the web .
visit at http : / / www . . net
trade links - georgel 932 me @ yahoo . com
if you have a product , service , opportunity or quality merchandise
that appeals to people worldwide , reach your targeted audience !
for a fraction of what other large newsletters charge you can exhibit
your website here , and trade links for only $ 8 cpm . compare
that
to the industry average of $ 10 - $ 15 cpm . why ? . . . because as a
valuable member we want you to be successful ! order today -
showcases are limited and published on a first come , first serve
basis .
for our secure order form , click here : http : / / bannersgomlm . com / ezine
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
win a free ad in community & commentary
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
to keep this interesting , how about this , every month we ' ll
draw a name from the replies and that person will win one
sponsorship showcase ad in the community commentary , for free .
that ' s a value of over $ 700 . 00 ! respond to each weekly survey ,
and increase your chances to win with four separate entries .
question
of the week ( 07 / 06 / 01 ) . . .
no right or wrong answers , and just by answering
you are entered to win a showcase ad - free !
~ ~ ~ do you spend more or less time ~ ~ ~
~ ~ ~ online in the summer months ? ~ ~ ~
more
mailto : one @ aeopublishing . com
less
mailto : two @ aeopublishing . com
same
mailto : three @ aeopublishing . com
to make this as easy as possible for you , just click on the
e - mail address that matches your answer - you do not need to
enter any information in the subject or body of the message .
* * add your comments ! follow directions above and
add your comments in the body of the message , and we ' ll
post the best commentaries along with the responses .
you will automatically be entered in our drawing for a free
sponsorship ad in the community commentary . please
respond only one time per question . multiple responses
from the same individual will be discarded .
last weeks ' s results ( 06 / 29 / 01 )
~ ~ ~ what is the goal of your website ? ~ ~ ~
sell 40 %
get leads 20 %
build branding 5 %
provide information 20 %
other 15 %
comments :
- - - - - - - - - - - - - - - - - - - - - - - - - - - -
our web site is initially designed to get leads , build
branding , and provide information . . . . . . . with a 12 month goal
of selling our service more specifically via a shopping cart .
we offer a service and at this time take deposits and payments
via our site .
our site has been up less than 2 months and our expectation
was that we would refer to our site for leads developed in
traditional media and by referral for more information , and
to make a professional impression on someone you may not
meet before providing service .
the growth of our customer base shopping on line has grown
outside of anyone ' s expectations . . . . . . . certainly mine and
i ' ve been in this business for 25 years . the internet is not
dead in the horse business , it is just getting it ' s legs , and
the folks using it want to get all the ancillary services
on - line as well . our site ( the first we ' ve developed ) has
exceeded our expectations , and we aren ' t satisfied with it
yet . . . . . . . we just wanted to get it there for information !
jeff and rebecca marks http : / / www . grand - champion . com
branding . while quality customer service and product
have been and will always be our top priority brand building
zesto is our most challenging task .
zesto . com ranks very high and most often # 1 or 2 on
all major search engines and directories even yahoo entering
the keyword zesto . the problem is simply that , who if anyone
would type the keyword zesto , therefore we must try to
build our brand by ensuring that generic keywords associated with our products ( citrus peel ) are used throughout
our site as well as search engine submissions .
fortunately owning a non generic domain short , easy
to remember and trademarked works in our favor because
the marketability potential is limitless .
arlene turner http : / / www . zesto . com
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
to change your subscribed address ,
send both new and old address to submit
see below for unsubscribe instructions .
please send suggestions and comments to : editor
i invite you to send your real successes and showcase
your strategies and techniques , or yes , even your total bombs ,
" working together we can all prosper . " submit
for information on how to sponsor your membership
community commentary visit : sponsorship
showcase
copyright 2001 aeopublishing . com
email : yourmembership 2 @ aeopublishing . com
voice :
web : http : / / www . aeopublishing . com
this email has been sent to jm @ netnoteinc . com at your
request , by your membership newsletter services .
visit our subscription center to edit your interests or unsubscribe .
http : / / ccprod . roving . com / roving / d . jsp ? p = oo & id = bd 7 n 7877 . 7 giv 5 d 57 & m = bd 7 n 7877 charset = iso - 8859 - 1
in this issue
internet success through simplicity
member showcase
win a free ad in community & commentary
today ' s special announcement :
win a free ad in community & commentaryto keep this interesting , how about this , every month we ' ll
draw a name from the replies and that person will win one
sponsorship showcase ad in the community commentary , for free .
that ' s a value of over $ 700 . 00 ! respond to each weekly survey ,
and increase your chances to win with four separate entries .
question
of the week ( 07 / 06 / 01 ) . . .
no right or wrong answers , and just by answering
you are entered to win a showcase ad - free !
~ ~ ~ do you spend more or less time ~ ~ ~
~ ~ ~ online in the summer months ? ~ ~ ~
more
mailto : one @ aeopublishing . com
less
mailto : two @ aeopublishing . com
same
mailto : three @ aeopublishing . com
to make this as easy as possible for you , just click on the
e - mail address that matches your answer - you do not need to
enter any information in the subject or body of the message .
* * add your comments ! follow directions above and
add your comments in the body of the message , and we ' ll
post the best commentaries along with the responses .
you will automatically be entered in our drawing for a free
sponsorship ad in the community commentary . please
respond only one time per question . multiple responses
from the same individual will be discarded .
last weeks ' s results ( 06 / 29 / 01 )
~ ~ ~ what is the goal of your website ? ~ ~ ~
sell 40 %
get leads 20 %
build branding 5 %
provide information 20 %
other 15 %
comments :
- - - - - - - - - - - - - - - - - - - - - - - - - - - -
our web site is initially designed to get leads , build
branding , and provide information . . . . . . . with a 12 month goal
of selling our service more specifically via a shopping cart .
we offer a service and at this time take deposits and payments
via our site .
our site has been up less than 2 months and our expectation
was that we would refer to our site for leads developed in
traditional media and by referral for more information , and
to make a professional impression on someone you may not
meet before providing service .
the growth of our customer base shopping on line has grown
outside of anyone ' s expectations . . . . . . . certainly mine and
i ' ve been in this business for 25 years . the internet is not
dead in the horse business , it is just getting it ' s legs , and
the folks using it want to get all the ancillary services
on - line as well . our site ( the first we ' ve developed ) has
exceeded our expectations , and we aren ' t satisfied with it
yet . . . . . . . we just wanted to get it there for information !
jeff and rebecca marks http : / / www . grand - champion . com
branding . while quality customer service and product
have been and will always be our top priority brand building
zesto is our most challenging task .
zesto . com ranks very high and most often # 1 or 2 on
all major search engines and directories even yahoo entering
the keyword zesto . the problem is simply that , who if anyone
would type the keyword zesto , therefore we must try to
build our brand by ensuring that generic keywords associated with our products ( citrus peel ) are used throughout
our site as well as search engine submissions .
fortunately owning a non generic domain short , easy
to remember and trademarked works in our favor because
the marketability potential is limitless .
arlene turner http : / / www . zesto . com
to change your subscribed address ,
send both new and old address to submit
see below for unsubscribe instructions .
please send suggestions and comments to : editor
i invite you to send your real successes and showcase
your strategies and techniques , or yes , even your total bombs ,
" working together we can all prosper . " submit
for information on how to sponsor your membership
community commentary visit : sponsorship
showcase
copyright 2001 aeopublishing . com
email us : : visit our site
phone :
this email was sent to jm @ netnoteinc . com , at your request , by your membership newsletter services .
visit our subscription center to edit your interests or unsubscribe .
view our privacy policy .
powered by
`
	datah := `Subject: congrats !
contratulations on the execution of the central maine sos deal ! this is another great example of what we can do when everyone comes together to get something done . this transaction brings both strategic value to the business , nice positions for the book and quite a nice chunk of change as well !
great job guys !
( hey dana , are you paying for the celebration dinner ? ! )`
	t.Log("test filter-spam:", Filter(datas))
	t.Log("test filter-ham:", Filter(datah))
}
