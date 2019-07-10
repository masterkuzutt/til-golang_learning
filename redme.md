# はまった点
## unexpected nul in input
- https://stackoverflow.com/questions/47347191/go-error-unexpected-nul-in-input?rq=1
- windowsで　echo ""> file.go でファイルを作るとエンコーディングがUTF-8になっていないために発生。
# Web Application Framework 
### gin
#### hot-reload
- https://stackoverflow.com/questions/50780008/can-golang-hot-swap-in-develop-mode
##### [gin](https://github.com/codegangsta/gin)
- これがシンプルそうに見える
- powershellでは動かない
- main.goと同ディレクトリがカレントディレクトリじゃないと動かない。
#### form
- flask_wtfのようアドオンは見つけられてない。
- https://qiita.com/CST_negi/items/99c7f2a2bd82bfb0fafd
#### middleware
- 何それ？
- ルーティングの前後に処理を行う。
- https://qiita.com/tobita0000/items/d2309cc3f0dd32006ead
### セッション管理用
#### デフォ？のセッション管理用Middleware
- https://github.com/gin-contrib/sessions
- Redisってなに？
- セッションとcookieの話
    - https://stackoverflow.com/questions/14347540/what-is-the-difference-between-a-cookie-and-redis-session-store
- save忘れがち
# vendor tool
- 何それ？

# template
- https://blog-asnpce.com/technology/913
- flaskというかjinja的な使い方はこれが必要なよう。
## [pongo2](https://github.com/flosch/pongo2)
- こっちもいいかも
- https://docs.djangoproject.com/en/dev/topics/templates/

# Object Relational Mapping
## gorm
- https://gorm.io/


# Webアプリの話
## リダイレクトの時のステータスコードの話
- https://stackoverflow.com/questions/2839585/what-is-correct-http-status-code-when-redirecting-to-a-login-page
