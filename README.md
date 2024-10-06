# starting-go

## requirements

- `go version go1.21.0 darwin/arm64`


## environments

`~/.zshrc` などに以下を追加

```zsh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## install

[asdf-vm](https://asdf-vm.com/)を利用してインストールを実施

## build

ビルドするには以下の通りで実施

```zsh
go build -o hello hello.go
```

## TODO

### ASDF_GOLANG_MOD_VERSION_ENABLED

```
Notice: Behaving like ASDF_GOLANG_MOD_VERSION_ENABLED=true
        In the future this will have to be set to continue
        reading from the go.mod and go.work files
```

[こちら](https://blog.yukii.work/posts/2023-06-29-asdf-golang-mod-version-enabled/)を参考に対応する。以下を設定しておけばメッセージ自体は表示されなくなる模様。

```zsh
export ASDF_GOLANG_MOD_VERSION_ENABLED=true
```

### gocode

現在はあまり使われていない模様なのでインストールはしない事にしている。

### gopls was not able to find modules in your workspace

[VSCodeでProject Managerを使っている場合にgoplsが動かない現象を解消する](https://qiita.com/y_tochukaso/items/da426190a4563c1dbebd)を参考に対応する。

コーディング中に出力されたメッセージは以下

```plaintext
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go
```

`mkdir go.mod` で対応した。

### gore

<https://motemen.hatenablog.com/entry/2015/02/go-repl-gore>
に記載の内容どおりインストールしてみる。
最新の内容は <https://github.com/x-motemen/gore> にある

```zsh
go install github.com/x-motemen/gore/cmd/gore@latest
go get github.com/nsf/gocode # 入力補完に
go get github.com/k0kubun/pp # プリティプリントに
go get github.com/x/tools/cmd/godoc # ドキュメントに
```

### goコマンド

### `go env`

Goのビルドシステムに関係する環境変数の内容を表示する。

```zsh
go env
```

環境変数 | 値 | 内容
--- | --- | ---
GOARCH | amd64, 386, arm, ppc64 | コンパイラが対象とするCPUアーキテクチャ
GOBIN | ディレクトリ | `go install`によってインストールされるコマンドの格納ディレクトリ。指定がない場合は `$GOPATH/bin`
GOOS | linux, darwin, windows, netbsd | コンパイラが対象とするOS環境
GOPATH | ディレクトリ | パッケージのソースコードとオブジェクトファイル、実行ファイルなどが格納されるディレクトリ
GORACE | 文字列 | レースコンディションの問題を検出するツールに指定するオプション
GOROOT | ディレクトリ | Go本体のインストールディレクトリ

### `go fmt`

Goのソースコードを整形する。内部的には`gofmt`を呼び出している。

```zsh
go fmt [-n] [-x] [packages]
```

オプション | 効果
--- | ---
`-n` | 実行されるコマンドの表示(ファイルは書き換えない)
`-x` | 実行されるコマンドの表示(ファイルは書き換える)

### `go doc`

```zsh
go doc [-u] [-c] [package|[package.]symbol[.methodOrField]]
```

オプション | 効果
--- | ---
`-c` | 識別子のマッチングで「大文字小文字」を厳密に区別する
`-u` | 非公開な識別子やメソッドについてもドキュメントを表示する

利用例

```zsh
go doc math/rand
go doc math/rand.Intn
go doc time.Time.Unix
```

パッケージのドキュメント生成

```zsh
go doc animals
go doc animals.ElephantFeed
```

### `go build`

```zsh
go build [-o output] [-i] [build flags] [packages]
```

オプション | 効果
--- | ---
`-i` | インポートパッケージの依存関係を更新する
`-o` | 出力ファイル名を指定する
`-x` | 実行されるコマンドの表示(ファイルは書き換える)

zooのビルドを実行

```zsh
go build zoo 
go build -x -o zoo.a zoo 
cd zoo
go build main.go app.go f.go # => main.go が生成される
```

### `go install`

```zsh
go install [-i] [build flags] [packages]
```

オプション | 効果
--- | ---
`-i` | インポートパッケージの依存関係を更新する
`-x` | 実行されるコマンドの表示(ファイルは書き換える)

### `go get`

```zsh
go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]
```

オプション | 効果
--- | ---
`-d` | ダウンロードのみを実行する
`-f` | 対象パッケージのパスから推測されるリポジトリへの検証をスキップする(`-u`指定時のみ)
`-fix` | 対象パッケージの依存関係を解決する前にgo fixツールを適用する
`-insecure` | カスタムリポジトリを使用する場合に非セキュアなプロトコルの使用を許可する(例: http)
`-t` | 対象パッケージに付属するテストが依存するパッケージもあわせてダウンロードする
`-u` | 対象のパッケージの更新と依存パッケージの更新を検出して再ダウンロードとインストールを実行する

利用例

```zsh
go get golang.org/x/net/http2
```

### `go test`

```zsh
go test [-c] [-i] [build flags] [packages] [flags for test binary]
```

カバレッジ率の計測

```zsh
go test -cover ./...
.
.
PASS
coverage: 50.0% of statements
ok      github.com/kakikubo/starting-go/zoo/animals     0.495s  coverage: 50.0% of statements
```

組み込みパッケージのテストも行える

```zsh
go test -cover fmt
```

オプション | 効果
--- | ---
`-c` | パッケージのテストのビルドのみを実行し、テストは実行しない
`-i` | インポートパッケージの依存関係を更新する
`-o` | ビルドしたテストの実行ファイルを任意のファイル名で出力する
`-x` | 実行されるコマンドの表示(ファイルは書き換える)
