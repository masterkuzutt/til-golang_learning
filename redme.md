# はまった点
## unexpected nul in input
- https://stackoverflow.com/questions/47347191/go-error-unexpected-nul-in-input?rq=1
- windowsで　echo ""> file.go でファイルを作るとエンコーディングがUTF-8になっていないために発生。
# 気になった点
- ginのrun.goのmainの中の実行はいつ走るのか？初回リクエスト受けるまでは知らないように見えた。

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

# clean architecture
- http://nakawatch.hatenablog.com/entry/2018/07/11/181453
- https://github.com/nakabonne/cleanarchitecture-sample/tree/master/src/cleanarchitecture-sample
- (右下の図)[https://nrslib.com/clean-flow-of-control/]
# test 
- https://re-engines.com/2018/10/16/go-testify%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%83%86%E3%82%B9%E3%83%88%E4%BD%9C%E6%88%90/

# Webアプリの話
## リダイレクトの時のステータスコードの話
- https://stackoverflow.com/questions/2839585/what-is-correct-http-status-code-when-redirecting-to-a-login-page

# セキュアコーディング
- https://golang.shop/post/go-scp-010-introduction-ja/