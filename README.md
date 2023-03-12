1. 會用到的套件
github.com/joho/godotenv
github.com/gofiber/fiber/v2

2.  config/dev.env 中，HOST 設為自己主機的 IP ，PORT 為提供服務的入口，還有需設定環境變數 APP_ENV=dev

3. 在 gitlab 的 repo 中要設定 webhook 的接收位置，也就是 go 的 webserver，由於一定要用 https，所以可以用 ngrok 來代替
例如設為 https://7e99-140-112-179-52.jp.ngrok.io/webhook

4. api/projects.json 中，key 為 repo 的名稱，value 為當接收到 webhook 的 event 時(目前加了 push)要執行的 command