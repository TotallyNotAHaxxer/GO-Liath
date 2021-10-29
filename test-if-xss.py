import os, sys, datetime, requests
import time as t
import colorama 
#from imports 
from datetime import datetime 
from colorama import Fore, Back, Style, init 
from bs4 import BeautifulSoup as bs
from urllib.parse import urljoin
from prettytable import PrettyTable
init()
def get_all_forms(url):
    soup = bs(requests.get(url).content, "html.parser")
    return soup.find_all("form")
def get_form_details(form):
    details = {}
    action = form.attrs.get("action").lower()
    method = form.attrs.get("method", "get").lower()
    inputs = []
    for input_tag in form.find_all("input"):
        input_type = input_tag.attrs.get("type", "text")
        input_name = input_tag.attrs.get("name")
        inputs.append({"type": input_type, "name": input_name})
    details["action"] = action
    details["method"] = method
    details["inputs"] = inputs
    return details
def submit_form(form_details, url, value):
    target_url = urljoin(url, form_details["action"])
    inputs = form_details["inputs"]
    data = {}
    for input in inputs:
        if input["type"] == "text" or input["type"] == "search":
            input["value"] = value
        input_name = input.get("name")
        input_value = input.get("value")
        if input_name and input_value:
            data[input_name] = input_value
    if form_details["method"] == "post":
        return requests.post(target_url, data=data)
    else:
        return requests.get(target_url, params=data)
def scan_xss(url):
    sc = "defxss.txt"
    file2 = open('defxss.txt', 'r')
    l1    = file2.readlines()
    count = -0
    for line in l1:
            file = open('defxss.txt', 'r')
            l = file.readlines()
            count =+ 1
            for line in l:
                    forms = get_all_forms(url)
                    js_script = "{}".format(line.strip())
                    is_vulnerable = False
                    for form in forms:
                        form_details = get_form_details(form)
                        content = submit_form(form_details, url, js_script).content.decode()
                        if js_script in content:
                            print(f"\033[34m[+] XSS Detected -> {url}") 
                            ptable = PrettyTable(["Content and form details"])
                            ptable.add_row([form_details])
                            print("\033[34m",ptable)
                            is_vulnerable = True
    if is_vulnerable == False:
        print(Fore.RED+"[-] XSS testing came back false, not XSS injectable")
    else:
        print(f"\033[32m[+] URL {url} came back with a True Value | This MIGHT be injectable")
if __name__ == "__main__":
    url = sys.argv[1]
    print(scan_xss(f'{url}'))