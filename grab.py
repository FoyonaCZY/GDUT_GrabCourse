import time

import requests

cookie = 'JSESSIONID=cookies'
headers = {
    "Host": "jxfw.gdut.edu.cn"
    , "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "
                    "Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0"
    , "Accept": "application/json, text/javascript, */*; q=0.01"
    , "Accept-Language": "en-US,en;q=0.5"
    , "Accept-Encoding": "gzip, deflate, br"
    , "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
    , "X-Requested-With": "XMLHttpRequest"
    , "Content-Length": "36"
    , "Origin": "https://jxfw.gdut.edu.cn"
    , "DNT": "1"
    , "Connection": "keep-alive"
    , "Referer": "https://jxfw.gdut.edu.cn/xskjcjxx!kjcjList.action"
    , "Cookie": f"{cookie}"
    , "Sec-Fetch-Dest": "empty"
    , "Sec-Fetch-Mode": "cors"
    , "Sec-Fetch-Site": "same-origin"
}


def qk(kcrwdm: str, kcmc: str):
    url = 'https://jxfw.gdut.edu.cn/xsxklist!getAdd.action'
    data = f'kcrwdm={kcrwdm}&kcmc={kcmc}'
    res = requests.post(url, headers=headers, data=data.encode('utf-8'))
    print(res.text)


if __name__ == "__main__":
    while 1:
        qk('id', 'name')
        time.sleep(0.5)
