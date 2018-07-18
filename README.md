# wareki-cli

wareki-cli は、Go言語で作られた和暦と西暦の変換ツールです。

## Installation

```
$ go get -u github.com/chase0213/wareki-cli
```

## Usage

### 西暦から和暦へ変換

西暦から和暦への変換時には、コマンド `w` または `wareki` を指定して、その直後に `/` または `-` 区切りの日付を `YYYY/MM/DD` のフォーマットで入力します。

```bash
$ wareki-cli w 1989/02/13
平成1年2月13日

$ wareki-cli w 1989/01/07
昭和64年1月7日
```

### 和暦から西暦へ変換

和暦から西暦への変換時には、コマンド `s` または `seireki` を指定して、元号、年、月、日の順で入力します。

```bash
$ wareki-cli s 平成1年2月13日
1989/2/13

$ wareki-cli s 昭和六十四年一月七日
1989/1/7
```

## Contributing

このツールは、[wareki](https://github.com/chase0213/wareki)を基に作られており、このツールはコマンドラインパーサー以外のロジックを持ちません。
コマンドラインへの機能追加、バグ修正等のコントリビューション以外は、[wareki](https://github.com/chase0213/wareki)にしていたたくようお願いします。

具体的なコントリビュートの方法は、以下の手順を推奨しています。

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
