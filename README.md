# lambda-calculus-machine

SECDマシンよりもっと単純な、型無しラムダ計算のスタックマシンを作る。
複雑になりすぎないところで、止めておく。

## 構成

まずラムダ式を表現するExpressionを定義する。

- Symbol
- Function
- Application

このExpressionで表現されたASTをコンパイルして、
VMで実行するInstructionに変換する。

- Fetch
- Call
- Close
- Apply
- Return


## 履歴

### v0.1.0

- 動くようにする


## 参考

- アンダースタンディング コンピュテーション 9784873116976

