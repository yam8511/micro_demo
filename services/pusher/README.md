1. 註冊
client -> register  host & chatroom -> pusher
pusher -> response  id -> client

2. 監聽
pusher -> subscribe channel -> redis

3. 客戶發送訊息
client -> push channel & message -> pusher
pusher -> encode channel & message -> redis

4. 推送
redis -> publish message -> pusher
pusher -> decode message & push -> client

