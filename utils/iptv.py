import os
import threading
import queue
import json
import requests
from lxml import etree

xpath='//*[@id="channel-list"]/div/div[@data-country]'
main_url="https://www.iptv-epg.com/available-channels"
country_url="https://www.iptv-epg.com/channel/country/{}" #i.e. ae,tr

####### Channel Details #######
#chan_xpath='//*[@id="country-{}"]/table/tbody/tr' #i.e. ae,tr
country_name_all='//*[@id="channel-list"]/div/div[1]/text()' #available-channels
country_name='//*[@id="channel-list"]/div/div[@data-target="#{}"]/text()'
img_xpath='//table/tbody/tr/td[1]/img' #attrib["src"]
name_xpath='//table/tbody/tr/td[2]' #.text
tvgid_xpath='//table/tbody/tr/td[3]/code' #.text
####### Channel Details END #######


def dump_countries():
    data=[]
    q=requests.get(main_url)
    dom=etree.HTML(q.text)

    countries=[{country.attrib["data-country"]:dom.xpath(country_name.format(country.attrib["id"]))[0].strip()} for country in dom.xpath(xpath)]

    for country in countries:
        for k,v in country.items():
            ext=requests.get(country_url.format(k)).text
            c=etree.HTML(ext)
            imgs=c.xpath(img_xpath)
            names=c.xpath(name_xpath)
            tvg_ids=c.xpath(tvgid_xpath)

            for i in range(len(imgs)):
                item={"img":imgs[i].attrib["src"],"name":names[i].text,"tvg-id":tvg_ids[i].text,\
                "country":v,"cc":k}
                data.append(item)

    with open("chans.json","w",encoding='utf8') as f:
        json.dump(data,f,ensure_ascii=False)


def download_logos(url):
    local_filename = os.path.join("logos",url.split('/')[-1])
    if os.path.exists(local_filename):
        return
    # NOTE the stream=True parameter below
    with requests.get(url, stream=True) as r:
        try:
            r.raise_for_status()
        except:
            return
        with open(local_filename, 'wb') as f:
            for chunk in r.iter_content(chunk_size=8192): 
                if chunk: # filter out keep-alive new chunks
                    f.write(chunk)
                    # f.flush()
    print(url)
    return local_filename

def worker():
    while True:
        item=q.get()
        if item is None:
            break
        download_logos(item)
        q.task_done()

def download_logos()
    q=queue.Queue()
    threads=[threading.Thread(target=worker) for i in range(4)]

    for i in data:
        q.put(i["img"])    

    [t.start() for t in threads]
    q.join()

    for i in range(4):
        q.put(None)
    for t in threads:
        t.join()
