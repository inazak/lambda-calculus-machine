# lambda-calculus-machine example

astを拡張して変数を定義し、`100 mod 13` を計算する例。

## 変数の定義

インターフェイス ast.Expression を満たす Variable を定義する。
`SUCC` や `PRED` 以外に、再帰的な計算をするために
不動点コンビネータ `Y` も定義している。

Variable は、そのままでは vm.Compile を通らないので、
コンパイル前に展開する expandAll を用意している。

## 動作の結果

main.goでは `100 mod 13` を実行している。

```
input:
(((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^
y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) x) y))
 n) m)) ((f ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m) n)) n)
) m)) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))))))))
```

結果は `9` となる。

```
output:
^f.^x.(f (f (f (f (f (f (f (f (f x)))))))))
```

