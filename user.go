package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	. "github.com/logrusorgru/aurora"
	"github.com/steelx/extractlinks"
	"golang.org/x/net/html"
)

//structureing and handeling functions

type UrlTitle struct {
	idx   int
	url   string
	title string
}

var (
	wg        sync.WaitGroup
	urlQueue  = make(chan string)
	config    = &tls.Config{InsecureSkipVerify: true}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	hasCrawled = make(map[string]bool)
	netClient  *http.Client
)

func banner() {
	f, err := os.Open("forsextor.txt")
	checkErr(err)
	defer f.Close()
	fo := bufio.NewScanner(f)
	for fo.Scan() {
		fmt.Println(fo.Text())
	}

	if err := fo.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	netClient = &http.Client{
		Transport: transport,
	}
	go sighandel(make(chan os.Signal, 1))
}

// simple functions

func sys() {
	if runtime.GOOS == "linux" {
		fmt.Println(Cyan("\033[32m[*] Detected System -> Linux"))
	}
	if runtime.GOOS == "windows" {
		fmt.Println(Cyan("\033[32m[*] Detected System -> Windows"))
	}
}

func clear() {
	if runtime.GOOS == "linux" {
		ex := "clear"
		cmd := exec.Command(ex)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(Cyan(string(stdout)))
	}
	if runtime.GOOS == "windows" {
		cl := "cls"
		cd := exec.Command(cl)
		stdout, err := cd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(Cyan(string(stdout)))
	}
}

func online() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		fmt.Println(Cyan("\033[32m[+] Connection Good...."))
		return true
	}
	fmt.Println(Cyan("[-] Interface has been disconnected from the network, please connect or set a connection "))
	os.Exit(1)
	return false
}

//// get each URLS title
/// as of now this code is in running and testing
/// will not work as a main function
func isValidUri(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	return err == nil
}

func toUrlList(input string) []string {
	list := strings.Split(strings.TrimSpace(input), "\n")
	urls := make([]string, 0)

	for _, url := range list {
		if isValidUri(url) {
			urls = append(urls, url)
			file, fileErr := os.Create("urls.txt")
			if fileErr != nil {
				fmt.Println("[!] Could not Create a File.......")
				fmt.Println(fileErr)
			}
			fmt.Fprintf(file, "%v\n", url)
		}
	}

	return urls
}

func fetchUrlTitles(urls []string) []*UrlTitle {
	ch := make(chan *UrlTitle, len(urls))
	for idx, url := range urls {
		go func(idx int, url string) {
			doc, err := goquery.NewDocument(url)

			if err != nil {
				ch <- &UrlTitle{idx, url, ""}
			} else {
				ch <- &UrlTitle{idx, url, doc.Find("title").Text()}
			}
		}(idx, url)
	}
	urlsWithTitles := make([]*UrlTitle, len(urls))
	for range urls {
		urlWithTitle := <-ch
		urlsWithTitles[urlWithTitle.idx] = urlWithTitle
	}
	return urlsWithTitles
}

func toMarkdownList(urlsWithTitles []*UrlTitle) string {
	markdown := ""
	for _, urlWithTitle := range urlsWithTitles {
		markdown += fmt.Sprintf("- [%s](%s)\n", urlWithTitle.title, urlWithTitle.url)
	}
	return strings.TrimSpace(markdown)
}

/// get URL ID's

func getHtmlPage(webPage string) (string, error) {
	resp, err := http.Get(webPage)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func parse(text string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var isTd bool
	var n int
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			isTd = t.Data == "td"
		case tt == html.TextToken:
			t := tkn.Token()
			if isTd {

				fmt.Printf("%s ", t.Data)
				n++
			}
			if isTd && n%3 == 0 {
				fmt.Println()
			}
			isTd = false
		}
	}
}

//////////////////////////////////////// complex url shifting //////////////////

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func grabparse() {
	hardurl := "placeholder" // figure out parsing with the command line arguments
	uro := hardurl
	parsedURL, err := url.Parse(uro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------------------- URL PARSED -------------- ")
	fmt.Println("Scheme        --->  " + parsedURL.Scheme)
	fmt.Println("Hostname      --->  " + parsedURL.Host)
	fmt.Println("Path in URL   --->  " + parsedURL.Path)
	fmt.Println("Query Strings --->  " + parsedURL.RawQuery)
	fmt.Println("Fragments     --->  " + parsedURL.Fragment)
}

/////////////////////////////////////////////////////////////////////////////////

/// fucntion banner lolcat

func banlol() {
	prgo := "ruby"
	arg1 := "banner.rb"
	cmd := exec.Command(prgo, arg1)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(Red(string(out)))
}

/// main

func main() {
	clear()
	banlol()
	online()
	sys()
	args := os.Args[1:]  // scrape parsing
	secarg := os.Args[2] //whois IP
	baseUrl := args[0]
	args2 := args[1]
	addr, err := net.LookupIP(secarg)
	resp, err := http.Get(baseUrl)
	t := time.Now()
	fmt.Println("\033[34m[>] Script Started At -> ", t)
	// error argument handeling
	if err != nil {
		fmt.Println(Red("[-] Couldnt Get the hostname? "))
	} else {
		fmt.Println("\033[32m[*]Server IPA -> ", addr)
	}
	if len(args) == 0 {
		fmt.Println(Red("[-] Url seems to be missing? try https://www.google.com"))
		os.Exit(1)
	}
	if len(args2) == 0 {
		fmt.Println(Red("[-] Skipping Complex URL search....."))
		fmt.Println(Red("[-] Complex URL was not parsed as an argument"))
	}

	// argument URO parsing
	if err != nil {
		log.Fatal(err)
	}

	input, _ := clipboard.ReadAll()

	urls := toUrlList(input)

	if len(urls) == 0 {
		fmt.Println("\033[31m[*] Skipping....No URLs found in Copy")
	}
	// cliboard finding titles
	urlsWithTitles := fetchUrlTitles(urls)
	markdown := toMarkdownList(urlsWithTitles)
	fmt.Println(markdown)
	clipboard.WriteAll(markdown)
	//
	fmt.Println("[*] Crawling URL >> ", baseUrl)
	uro := args2
	parsedURL, err := url.Parse(uro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Cyan("─────────────────────────Server Response─────────────────────────────"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mResponse Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mDate Of Request  -> ", resp.Header.Get("date"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mContent-Encoding -> ", resp.Header.Get("content-encoding"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mContent-Type     -> ", resp.Header.Get("content-type"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mConnected-Server -> ", resp.Header.Get("server"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mX-Frame-Options  -> ", resp.Header.Get("x-frame-options"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mScheme        --->  " + parsedURL.Scheme)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mHostname      --->  " + parsedURL.Host)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mPath in URL   --->  " + parsedURL.Path)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mQuery Strings --->  " + parsedURL.RawQuery)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mFragments     --->  " + parsedURL.Fragment)
	for k, v := range resp.Header {
		fmt.Print(Cyan("\033[34m[\033[35m*\033[34m]\033[35m-> " + k))
		fmt.Print(Red(" -> "))
		fmt.Println(v)
	}
	//grab content
	webPage := baseUrl
	data, err := getHtmlPage(webPage)

	if err != nil {
		log.Fatal(err)
	}

	parse(data)
	go func() {
		urlQueue <- baseUrl
	}()

	for href := range urlQueue {
		if !hasCrawled[href] {
			crawlLink(href)
		}
	}

}

func sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			t := time.Now()
			fmt.Println("\n\n\t\033[31m[>] Script Ended At -> ", t)
			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func desk() {
	url := "https://google.com" // desk net inf
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("[-] Sorry will not be able to run this command")
	} else {
		if resp.StatusCode >= 200 {
			out, err := exec.Command("notify-send", "Server responded with code 200 Connection is stable  °˖✧◝(⁰▿⁰)◜✧˖° ✔️").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		} else {
			out, err := exec.Command("notify-send", "Server Responded with a code that is not within the indexed list or range").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		}
	}
}

func crawlLink(baseHref string) {
	// declaring name
	hasCrawled[baseHref] = true
	fmt.Println(Cyan("──────────────────────────────────────────────────────"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mURL Found -> ", baseHref)
	u, err := url.Parse(baseHref)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mDomain Name -> ", domain)
	addr, err := net.LookupIP(domain) //domain IP for each
	if err != nil {
		fmt.Println(Red("[-] Couldnt Get the hostname? is there even one? "))
	} else {
		fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mDomain IPA  -> ", addr)
	}
	resp, err := http.Get(baseHref)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mConnected-Server -> ", resp.Header.Get("server"))
		fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mResponse Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	// finally test the query if it is injectable or not with SQL error //
	if runtime.GOOS == "windows" {
		fmt.Println(Red("[-] This is a windows system, script might be a bit slow"))
		fmt.Println(Yellow("[*] Testing SQLI this might take a while...."))
		prgo := "python3"
		arg1 := "test-if-sql.py"
		arg2 := baseHref
		cmd := exec.Command(prgo, arg1, arg2)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(Red(string(out)))
	} else {
		fmt.Println(Yellow("[*] Testing SQLI this might take a while...."))
		prgo := "python3"
		arg1 := "test-if-sql.py"
		arg2 := baseHref
		cmd := exec.Command(prgo, arg1, arg2)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(Red(string(out)))
	}
	if runtime.GOOS == "windows" {
		fmt.Println(Yellow("[*] Testing XSSI this might take a while...."))
		fmt.Println(Red("[-] This is a windows system, script might be a bit slow"))
		prgo := "python3"
		arg1 := "test-if-xss.py"
		arg2 := baseHref
		cmd := exec.Command(prgo, arg1, arg2)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(Red(string(out)))
	} else {
		fmt.Println(Yellow("[*] Testing XSSI this might take a while...."))
		prgo := "python3"
		arg1 := "test-if-xss.py"
		arg2 := baseHref
		cmd := exec.Command(prgo, arg1, arg2)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(Red(string(out)))
	}
	//
	resp, err = netClient.Get(baseHref)
	checkErr(err)
	defer resp.Body.Close()

	links, err := extractlinks.All(resp.Body)
	checkErr(err)

	for _, l := range links {
		if l.Href == "" {
			continue
		}
		Url := fixedURL(l.Href, baseHref)
		if baseHref != Url {
		}
		go func(url string) {
			urlQueue <- url
		}(Url)
	}
}

func fixedURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" {
		return base
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// 			LIVE MONITOR RUN AND CHECK, EXTRA FLAG ARGUMENT 	///////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CODE COPIED AND MANIPULATED FROM GO-SERVE/OG IN RED RABBIT MAIN

func design() {
	f, err := os.Open("banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(Blue(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func IsOnline() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		return true
	}
	fmt.Println(Cyan("[-] Interface has been disconnected from the network, please connect or set a connection "))
	return false
}

func mndesknot() {
	if runtime.GOOS == "windows" {
		fmt.Println(Cyan("[-] Sorry, but t this time i can not run this command"))
	} else {
		out, err := exec.Command("notify-send", "Testing Server Conn and Node every 20-30 seconds").Output()
		if err != nil {
			log.Fatal(err)
		} else {
			output := string(out[:])
			fmt.Println(output)
		}
	}
}

func resplog() {
	baseUrl := os.Args[0]
	url := baseUrl
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("[-] Sorry will not be able to run this command")
	} else {
		if resp.StatusCode >= 200 {
			out, err := exec.Command("notify-send", "Server responded with code 200 Connection is stable  °˖✧◝(⁰▿⁰)◜✧˖° ✔️").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		} else {
			out, err := exec.Command("notify-send", "Server Responded with a code that is not within the indexed list or range").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		}
	}
}

func logged() {
	if runtime.GOOS == "windows" {
		fmt.Println("This appends to a linux system only command, i will not be able to run it")
	} else {
		out, err := exec.Command("notify-send", "There was an error within the response").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
}

func clsa() {
	if runtime.GOOS == "windows" {
		fmt.Println(Red("[-] I Will not be able to execute this"))
	} else {
		out, err := exec.Command("clear").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
	if runtime.GOOS == "windows" {
		os := "linux"
		fmt.Println("[-] Sorry, this command is system spacific to -> ", os, "Systems")
	} else {
		out, err := exec.Command("pwd").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println("[~] Working Directory ~> ", output)
	}
}

func get() {
	clear()
	url := "https://google.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	design()
	fmt.Println(Cyan("--------------------------Server Response---------------------------"))
	fmt.Println("[+] Response Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("[+] Date Of Request  -> ", resp.Header.Get("date"))
	fmt.Println("[+] Content-Encoding -> ", resp.Header.Get("content-encoding"))
	fmt.Println("[+] Content-Type     -> ", resp.Header.Get("content-type"))
	fmt.Println("[+] Connected-Server -> ", resp.Header.Get("server"))
	fmt.Println("[+] X-Frame-Options  -> ", resp.Header.Get("x-frame-options"))
	fmt.Println(Cyan("--------------------------Server X-Requests-----------------------------"))
	for k, v := range resp.Header {
		fmt.Print(Cyan("[+] -> " + k))
		fmt.Print(Red(" -> "))
		fmt.Println(v)
	}
}

func urlidea() {
	var url string
	fmt.Scanf("%s", &url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("Can not run this string on win32-64")
		os.Exit(1)
	} else {
		fmt.Println("[+] YAY STRING VARIABLES ARE WORKING!!! response code -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
		os.Exit(1)
	}
}

func actualmonitor() {
	banlol()
	IsOnline()
	clear()
	time.Sleep(10 * time.Second)
	mndesknot()
	seconds := "20"
	time.Sleep(1 * time.Second)
	fmt.Println("[~] Testing Connection Every ", seconds, "Seconds")
	time.Sleep(1 * time.Second)
	for {
		time.Sleep(30 * time.Second)
		url := "https://google.com"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode >= 200 {
			fmt.Println("[+] Response Status Given -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println("[+] Response seems good")
			resplog()
			get()
		}
		if resp.StatusCode >= 300 && resp.StatusCode <= 400 {
			fmt.Println("[+] Response Status Given -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println("[~] Response may be laggy")
			logged()
			get()
		}
	}
}
