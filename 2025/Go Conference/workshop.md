# 低レベルコンテナランタイム自作講座 ～コンテナ技術の地盤を理解する～

## Status

### ❓ On Evaluation

## ワークショップ概要 / Overview

Go言語でWebバックエンドを開発している方のほとんどが、デプロイ環境にコンテナ仮想化技術を利用していることかと思います。  
そんなコンテナ技術のこと、きちんと理解していますか？「コンテナの実態とは何か？」「Docker、containerd、runcの違いは何か？」といった問いに、自信を持って答えられるでしょうか？

低レベルコンテナランタイム自作講座は、「docker run」コマンドを実行してからアプリケーションのプロセスが動き始めるまでにどんなソフトウェアがどんな役割を果たしているか、参加者の皆さんが説明できるようになることを目指したワークショップです。  
コンテナ技術を構成するソフトウェア群についての解説に加え、その中で最も低いレイヤで地盤の役割を果たす「低レベルコンテナランタイム」のMVPをGo言語で自作することで、みなさんも冒頭の問いに答えられるようになるはずです。  
グループで協力してコンテナ技術を支える3つのカーネル機能を学び、コンテナの魔術を自分の手で現実にしてみましょう。

## ワークショップから学べること / What you will learn?

- コンテナ技術を構成・利用するソフトウェア群
- 高レベル/低レベルコンテナランタイムの役割の違いを理解する
  - 高レベル = containerd/cri-o/podmanなど。イメージとネットワークの管理、デーモン化
  - 低レベル = runc/crun/youkiなど。コンテナの実行とプロセスの管理
- コンテナの実態が何かを理解する
  - 単なるカーネル機能の組み合わせ
  - Linux Namespace、cgroup、chroot(pivot_root)
  - 3人グループで1人1機能を担当し、最後にそれを組み合わせてコンテナとする

## ワークショップの規模 / Number of people

50人以上 / 50 people or more

## ワークショップの対象者 / Target Audience

最低限Goの知識がないと厳しい / Minimal knowledge of Go required
