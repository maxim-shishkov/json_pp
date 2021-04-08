**Задание**: сделать приложение на Golang, запускающее web-сервер на порту 8010 и отвечающее 
на GET запросы `/json/hackers` присылая список хакеров из [sorted set](https://redis.io/topics/data-types) в Redis

**Детали реализации**:

- На основе библиотеки [fasthttp](https://github.com/valyala/fasthttp)
- Ответ сервера должен включать заголовок `Content-Type: "application/json"`

Конфигурация Redis:

```jsx
redis 127.0.0.1:6379> zadd hackers 1953 "Richard Stallman"  
(integer) 1                                                 
redis 127.0.0.1:6379> zadd hackers 1940 "Alan Kay"          
(integer) 1                                                 
redis 127.0.0.1:6379> zadd hackers 1965 "Yukihiro Matsumoto"
(integer) 1                                                 
redis 127.0.0.1:6379> zadd hackers 1916 "Claude Shannon"    
(integer) 1                                                 
redis 127.0.0.1:6379> zadd hackers 1969 "Linus Torvalds"    
(integer) 1                                                 
redis 127.0.0.1:6379> zadd hackers 1912 "Alan Turing"
(integer) 1
```

Запрос к серверу:

```bash
curl -s http://localhost:8010/json/hackers |json_pp
```

Ожидаемый результат:

```json
[
  {
    "name": "Alan Turing",
    "score": 1912
  },
  {
    "name": "Claude Shannon",
    "score": 1916
  },
  {
    "name": "Alan Kay",
    "score": 1940
  },
  {
    "name": "Richard Stallman",
    "score": 1953
  },
  {
    "name": "Yukihiro Matsumoto",
    "score": 1965
  },
  {
    "name": "Linus Torvalds",
    "score": 1969
  }
]
```