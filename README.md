# GO-Liath
GO-Liath is a Very VERY inteligent spider and web crawlers for subdomains and domains written from 90% lines of code


```
                                                        | | |  /
                                                         \|_|_/
                       _.-/   _   \-._                    .'|
                     .'::(_,-' `-._)::`.                  |:|
                    (:::::::::::::::::::)                .':|
                     \_:::;;;::::;;;:::/    /            |::|
             \        ,---'..\::/..`-.'    /             |::|
              \       \_;:....|'...:_ )   /             .'=||
               \.       )---. )_.--< (   /'             ||=||
                \\     //|:: /--\:::\\\ //             _||= |
                 \\   ||::\:|----|:/:||/--.______,--==' \ - /
          -._     \`.  \\:|:|-- -|:\:/-.,,\\  .----'//'_.`-'
      \.     `-.   \ \ _|:|:|-- -|::||::\,,||-'////---' |/'
       \\       `._)\ / |\/:|-/|--\:/|. :\_,'---'       /
        \\_      /,,\/:.'\\/-.'`-.-//  \ |
        /`\-    //,,,' |-.\-'\--/|-/ ./| |             /
         /||-   ||, /| |\. |.-==-.| . /| |            | /
 __  |    /||-  ||,,\- | .  \;;;;/ .  .':/         _,-=/-'
/  \//    /||-  ' `,-|::\ . \,..,/   /: /         /.-'
,--||      /||-/.-.'  \  `._ `--' _.' .'|        //
.--||`.    /||//\ '   |-'.___`___' _,'|//       /;
  /\| \     ///\ /     \\_`-.`--`-'_==|'       /;'
 / |'\ \.   //\ /       \_\__)\`==-_'_|       / /
  .'  \.=`./|\ /          \`-- \--._/_/------( /
       \.=| `_/|-          |\`-| -/| `--------'
        \\` ./|-/-         |\`-| |-|     ________
         `--\ |=|'        _|\`-| |-|----'.-._ ..\`-.
             -|-|-     .-':`-;-| |./.-.-( | ||..|=-\\
             `'= /'   / ,--.:|-| ||_|_|_|_|-'__ .`-._)
              /|-|- .' /\ \ \|-` \\ ____,---'  `-. ..|
               /\=\/..'\ \_>-'`-\ \'              \ .|
               `,-':/\`.>' |\/ \/\ \              `- |
               //::/\ \/_/|-' \/| \ `.            / ||
              / (:|\ \/) \ \|.'-'  `-\\          |;:|\_
             || |:`-/:.'-|-' \|       \\          `;_\-`-._
             \\=/:_/::/\| \|          |\\            `-._ =`-._
              \_)' |:|                | //               `--.__`-.
                   |:|                                         )\|
                   /;/                                         / (\_
                  / /                                         |\\;;_`-.
                _/ /                                          ' `---\.-\
               /::||       
              /:::/
             //;;'
                              Go-Liath Version 1.0
---------------------------------------------------------------------------

```

# installs and usages 

Install linux<br>
`git clone https://github.com/ArkAngeL43/GO-Liath.git ; cd GO-Liath ; clear ; chmod +x ./install.sh ; ./install.sh `

usage   | go run user.go <http url> <domain> <https url> pfg
example | go run user.go http://example.com/ www.example.com https://www.example.com pfg

Install Windows <br> 
install the fucking golang first<br>
go run user.go http://example.com/ www.example.com https://www.example.com pfg



# What is Go-Liath?

Go-Liath is a very inteligent spider that can crawl domains, this main script is written from 90% golang, and not only is really fast but can crawl and gather any subdomain within a main URL

# what will it get? 

# First it will gather the Information about the URL or website itself 

```
[+] Connection Good....
[*] Detected System -> Linux
[>] Script Started At ->  2021-10-28 19:30:59.205574366 -0400 EDT m=+0.365378413
[*]Server IPA ->  [2606:2800:220:1:248:1893:25c8:1946 93.184.216.34]
[*] Skipping....No URLs found in Copy

[*] Crawling URL >>  http://example.com/
─────────────────────────Server Response─────────────────────────────
[*] Response Status  ->  200 OK
[*] Date Of Request  ->  Fri, 29 Oct 2021 03:30:41 GMT
[*] Content-Encoding ->  
[*] Content-Type     ->  text/html; charset=UTF-8
[*] Connected-Server ->  ECS (mic/9A8A)
[*] X-Frame-Options  ->  
[*] Scheme        --->  
[*] Hostname      --->  
[*] Path in URL   --->  www.example.com
[*] Query Strings --->  
[*] Fragments     --->  
[*]-> Content-Type -> [text/html; charset=UTF-8]
[*]-> Etag -> ["3147526947"]
[*]-> Expires -> [Fri, 05 Nov 2021 03:30:41 GMT]
[*]-> Last-Modified -> [Thu, 17 Oct 2019 07:18:26 GMT]
[*]-> X-Cache -> [HIT]
[*]-> Accept-Ranges -> [bytes]
[*]-> Age -> [489702]
[*]-> Cache-Control -> [max-age=604800]
[*]-> Date -> [Fri, 29 Oct 2021 03:30:41 GMT]
[*]-> Server -> [ECS (mic/9A8A)]
[*]-> Vary -> [Accept-Encoding]

```

# Then it wil actually start scraping and gather the URL with an addition of XSSI testing, SQLI testing, Name server, and response code 

```
[*] URL Found ->  http://example.com/
[*] Domain Name ->  example.com
[*] Domain IPA  ->  [2606:2800:220:1:248:1893:25c8:1946 93.184.216.34]
[*] Connected-Server ->  ECS (mic/9B14)
[*] Response Status  ->  200 OK
[*] Testing SQLI this might take a while....
[+] Detected 0 forms on http://example.com/.
[-] Might NOT be SQL injectable
[*] Testing XSSI this might take a while....
[-] XSS testing came back false, not XSS injectable
None
──────────────────────────────────────────────────────
[*] URL Found ->  https://www.iana.org/domains/example
[*] Domain Name ->  iana.org
[*] Domain IPA  ->  [2001:500:88:200::8 192.0.43.8]
[*] Connected-Server ->  Apache
[*] Response Status  ->  200 OK
[*] Testing SQLI this might take a while....
[+] Detected 0 forms on https://www.iana.org/domains/example.
[-] Might NOT be SQL injectable
[*] Testing XSSI this might take a while....
[-] XSS testing came back false, not XSS injectable
None
─────────────────────────────────────────────────────

```
