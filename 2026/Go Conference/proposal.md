# なぜruncはCGoを使うのか 〜神器「goroutine」とその制約

## Status

### ❔ In Evaluation

## Description

おおよそ10年間にわたり世界中のコンテナ環境を支えてきた低レベルコンテナランタイム「runc」。  
runcはGoで実装されていますが、コンテナ作成の中核部分では意図的にCGoが使われています。  
なぜPure Goで実装しないのでしょうか。

その背景には、Goの象徴的な機能であるgoroutineが持つ制約があります。  
goroutineは高い効率と使いやすさを実現する一方で、OSプロセスの緻密な操作との相性が良くありません。  
この制約は、runcだけでなく、os/execをはじめとする標準ライブラリの設計にも大きな影響を与えています。

本セッションでは、goroutineの制約を起点にGoランタイムや標準ライブラリの設計を紐解くことで、利点だけからは見えてこない「Goらしさ」を理解し、Goをより深く使いこなすための視点を共有します。  
さらに、この制約がruncにどのような設計上のトレードオフをもたらし、なぜruncがCGoという選択に至ったのかを紐解きます。

## Proposal Type

ロングセッション / Long Session (40 minutes)

## Proposal Level

上級者向け / Advanced

## Categories

- クラウド・インフラ / Cloud Infrastructure
- 標準機能・言語仕様 / Standard Features & Language Specifications
- OSSコントリビューション / OSS Contributions

## 備考/Additional Notes

トークのアウトラインはこちらです。

1. 導入 (3min)
   - コンテナとgoroutineで話の流れの入り口が2つになってしまうので、冒頭でDescriptionに書いた内容を簡単に紹介
   - アジェンダの共有
2. Linux基礎: プロセス・スレッド・syscall (3min)
   - OSプロセスとスレッドの基本的な説明
   - fork、clone、exec、pipe
3. コンテナの正体とruncの役割 (5min)
   - コンテナはプロセスの隔離技術
   - 簡易的なコンテナの実装紹介とその挙動
   - 低レベルコンテナランタイムの役割とrunc
4. goroutineの仕組み: GMPモデル (4min)
   - goroutineの実装方法 (GMPモデル)
   - goroutineの利点 (軽量・効率的な並行処理)
5. goroutineの制約と標準ライブラリ (10min)
   - OSプロセスの緻密な操作との相性の悪さ
   - goroutineはfork/cloneで壊れる - スレッド情報の破壊
   - スレッド固定の難しさ- runtime.LockOSThread()
   - os/execの設計
   - syscallパッケージにfork/cloneは存在しない
   - Goでforkっぽいことをやる - サブコマンド呼び出しパターン
6. なぜruncはCGoを使うのか (10min)
   - Namespaceの制約とDouble Fork・プロセス間通信の必要性
   - Goでの実装を試みた場合のつらさ
   - CGoを使う選択とそのメリット・デメリット
   - runcにおけるCGoの呼び出し方とその役割
   - runcのCGo部分の実装紹介
7. まとめ (2min)
   - goroutineは高効率な一方で明確な制約を持つ
   - その制約は標準ライブラリやOSS設計に現れ、runcのCGo利用はその代表例
   - 制約を理解することでGoをより深く理解できる

## Speaker Biography

千葉工業大学 情報科学研究科 情報科学専攻 修士1年。クラウドプラットフォームの開発に携わるソフトウェアエンジニアを目指し、日々研鑽を積んでいる。  
コンテナランタイムやKubernetesの仕組みを題材に、実際にレプリカを実装しながら学ぶチュートリアル「Coding Kubernetes」を執筆中。GWC2025では、その一部である「低レベルコンテナランタイム自作講座」を主催した。  
Goランタイムやコンテナ技術の内部実装に関心を持ち、OSSのコードリーディングやコントリビューションを通じて理解を深めている。
