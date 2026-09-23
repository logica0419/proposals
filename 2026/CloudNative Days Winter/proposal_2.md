# Nested Containersとコンテナランタイム ～いかにしてyoukiはruncと異なる解決法を選んだか

## Status

### ❔ In Evaluation

## 講演内容 - Abstract

- ベースとなるIssue・Pull Request
  - <https://github.com/youki-dev/youki/issues/3342>
  - <https://github.com/youki-dev/youki/pull/3347> (ほぼすべてのコンテクストはここにある)
  - <https://github.com/youki-dev/youki/pull/3563>
  - <https://github.com/youki-dev/youki/pull/3564>
- 伝えたい抽象的なメッセージ
  - 後発プロダクトが先行プロダクトの設計をそのまま踏襲する必要はない。
    - デファクトスタンダードは必ずしも「すべて正解」なわけではない
  - 後発プロダクトが、先行プロダクトが叶えられなかった「きれいな仕様・設計」を目指すことで、それは差別化・エコシステムの充実につながる
    - これはまさに、エコシステムの選択肢が広がっていく「Scaling Together」
- 今回取り上げる具体的な技術話
  - Nested Containers + cgroup v2の組み合わせへの、コンテナランタイム側の対処
  - runcが採用した「リトライ型アプローチ」と、youkiが採用した「事前推定 (proactive inference) 型アプローチ」の違い
    - youkiの方が、起こりえる危険なパターンを排除できる かつ よりシンプルな設計になっている
  - 実装する中で新しいやり方を発見し提案するだけでなく、KubeCon Japanの場でruncのメンテナーとyoukiのメンテナーを交えて会合を行う、コミュニティとしての努力
    - これはまさにコミュニティとコーディング双方のコントリビューションが繋がった場所
    - Scaling Togetherなのでは？

## 主なカテゴリ

Runtime

## 受講者レベル

上級者

## 想定受講者

- developer - システム開発
- app-developer - アプリケーション開発
- その他

## 実行フェーズ

- Other

## 必要とする講演時間 - Session time you need

30min (full session)

## 講演者プロフィール

千葉工業大学 情報科学専攻 修士1年。クラウドプラットフォームの開発に携わるソフトウェアエンジニアを目指し、日々研鑽を積んでいる。  
Kubernetes、Cilium、Traefikなど、クラウドネイティブ分野のOSSへのコントリビューションに取り組む。最近は特に低レベルコンテナランタイム「youki」をDockerと組み合わせて開発環境で利用し、バグの発見・報告と修正に取り組んでいる。
