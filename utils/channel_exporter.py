import os
import threading
import queue
import json
import requests
from lxml import etree


# //*[@id="channel-list"]

xpath = '//*[@id="channel-list"]/div/div[@data-country]'
main_url = "https://epg.best/available-channels"
country_url = "https://epg.best/channel/country/{}"  # i.e. ae,tr

####### Channel Details #######
# chan_xpath='//*[@id="country-{}"]/table/tbody/tr' #i.e. ae,tr
country_name_all = '//*[@id="channel-list"]/div/div[1]/text()'  # available-channels
country_name = '//*[@id="channel-list"]/div/div[@data-target="#{}"]/text()'
img_xpath = "//table/tbody/tr/td[1]/img"  # attrib["src"]
name_xpath = "//table/tbody/tr/td[2]"  # .text
tvgid_xpath = "//table/tbody/tr/td[3]/code"  # .text
####### Channel Details END #######


def dump_countries():
    data = []
    q = requests.get(
        main_url,
        headers={
            "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
            "referer": "https://epg.best/available-channels",
            "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
            "cookie": "__cfduid=db984cf93a868a0ff79a2faf5166df6e81604261900; XSRF-TOKEN=eyJpdiI6Im9XSmZUdXZMdk16Sm0wTnh0K1Nod2c9PSIsInZhbHVlIjoiN3hFaFlPYm5qQmV3bVE3TkRkXC9KSzBES3NJbUN5cW5WZ2R1V2lWOTRacU14NjBzbW5hMVo0VHNNcENPZjZXcHQiLCJtYWMiOiI2ODgwZWVhNDMwMjc2MjcwNmRlNWE3NDhmZmM5NTVlZjBhZTkzNTBhMzJkYzg4NTBiMzFlZDAzN2I3NDI3MTI2In0%3D; laravel_session=eyJpdiI6IkxHVlhKb3dJTnA4aWZlWHRuUzNLSWc9PSIsInZhbHVlIjoiRlNsVm1qa1VXREdabnYzQ3JnV2xOZjIzdnRIUW1BMjh4UWdBbXVuMGVHRGtSdStVSkZFNzVObTJRa1ZlejZQQSIsIm1hYyI6IjQ5YjRlNDQyMGM0NTQ3NjYwNjY1ZjJiYzY1ZGE5MzdlNGFiM2NkNTMyMTVlNDZjZDIxN2QyYjZhNjFlZTE2MDQifQ%3D%3D",
        },
    )
    dom = etree.HTML(q.text)

    countries = [
        {
            country.attrib["data-country"]: dom.xpath(
                country_name.format(country.attrib["id"])
            )[0].strip()
        }
        for country in dom.xpath(xpath)
    ]

    for country in countries:
        for k, v in country.items():
            ext = requests.get(country_url.format(k)).text
            c = etree.HTML(ext)
            imgs = c.xpath(img_xpath)
            names = c.xpath(name_xpath)
            tvg_ids = c.xpath(tvgid_xpath)

            for i in range(len(imgs)):
                item = {
                    "img": imgs[i].attrib["src"],
                    "name": names[i].text,
                    "tvg-id": tvg_ids[i].text,
                    "country": v,
                    "cc": k,
                }
                data.append(item)

        print(k)
    with open("chans.json", "w", encoding="utf8") as f:
        json.dump(data, f, ensure_ascii=False)


def download_logos(url):
    local_filename = os.path.join("logos", url.split("/")[-1])
    if os.path.exists(local_filename):
        return
    # NOTE the stream=True parameter below
    with requests.get(url, stream=True) as r:
        try:
            r.raise_for_status()
        except:
            return
        with open(local_filename, "wb") as f:
            for chunk in r.iter_content(chunk_size=8192):
                if chunk:  # filter out keep-alive new chunks
                    f.write(chunk)
                    # f.flush()
    print(url)
    return local_filename


def worker():
    while True:
        item = q.get()
        if item is None:
            break
        download_logos(item)
        q.task_done()


def download_logos():
    q = queue.Queue()
    threads = [threading.Thread(target=worker) for i in range(4)]

    for i in data:
        q.put(i["img"])

    [t.start() for t in threads]
    q.join()

    for i in range(4):
        q.put(None)
    for t in threads:
        t.join()


dump_countries()