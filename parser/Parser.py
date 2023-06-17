import requests
from bs4 import BeautifulSoup

def parseWords(words,url):
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    word_elements = soup.find_all('a')

    for element in word_elements:
        word = element.text.strip()
        if word.isupper() and word.isalpha() and len(word) == 5:
            word = word.replace(' ','Ё')
            words.append(word)
            
def main():
    words = list()
    # Отправляем GET-запрос к странице со словами
    url = 'https://slovardalja.net/letter.php?charkod='
    for i in range(192,224):
        parseWords(words, url+str(i))
    
    file = open("/../static/FiveLettersWords.txt", "w",encoding='utf-8')
    for i in words:
        file.write(i+"\n")
    file.close()
    
main()