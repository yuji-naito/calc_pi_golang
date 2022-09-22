以降に記載するビルド、実行方法は Ubuntu で確認したもの。

## ビルド方法

```bash
$ go build -o calc_pi_golang main.go
```

## 実行方法

```bash
$ ./calc_pi_golang
```

## 出力結果の例

count: 点の数
pi: 円周率の計算結果
time: 処理時間[s]

```bash
count:      10000       pi: 3.1444000000        time: 0.0003130000
count:     100000       pi: 3.1319600000        time: 0.0023140000
count:    1000000       pi: 3.1431120000        time: 0.0227780000
count:   10000000       pi: 3.1408260000        time: 0.2213500000
count:  100000000       pi: 3.1415162400        time: 2.1945760000
count: 1000000000       pi: 3.1415192080        time: 22.1791520000
```
