# 5Letters
the game "5 Letters" is a counterpart to the game World, where players have to guess a 5-letter word within 7 attempts. Each attempt must be a valid word consisting of 5 letters and it should be present in the Dal's Dictionary. If a letter is not in the word, it remains unchanged; if it is in the word but in a different position, it turns orange; if it is in the word and in the correct position, it turns green. If the word is not in the dictionary, the row is colored red.
# Install
There's simple way how to install it on your device:
- Clone repository
- Write in console
```
cd 5Letters
docker-compose -f build/pkg/docker-compose.yml up
```
Req:
- Git
- Docker/docker-compose
