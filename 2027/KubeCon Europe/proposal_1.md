# CNI Demystified: Philosophy, Design, and Plugin Trade-Offs

## Status

### ❔ In Evaluation

## Description

- ベースとなるIssue・Pull Request
  - <https://github.com/youki-dev/youki/issues/3342>
  - <https://github.com/youki-dev/youki/pull/3347> (ほぼすべてのコンテクストはここにある。runc型fallbackの試作。後に方針転換してclose)
    - <https://github.com/youki-dev/youki/pull/3347#issuecomment-4557359874> (runcとyoukiの処理フロー差分の根拠)
    - <https://github.com/youki-dev/youki/pull/3347#issuecomment-4557652166> (同上)
    - <https://github.com/youki-dev/youki/pull/3347#issuecomment-5224550358> (危険パターン・シンプルさの根拠)
  - <https://github.com/youki-dev/youki/pull/3563> (マージ済み)
  - <https://github.com/youki-dev/youki/pull/3564> (proactive inferenceの実装。当時Open。Dev Container / 通常のDocker-in-Docker / KinDでexec成功を確認)
  - 参考: <https://github.com/opencontainers/runc/pull/2416> (runcのfallback導入PR)
    - <https://github.com/opencontainers/runc/pull/2416/changes#r428886985> (危険なパスとして指摘された箇所)
- 問題・再現環境
  - Nested Containers + cgroup v2の組み合わせで`exec`が`Device or resource busy` (EBUSY) で失敗する
  - 発見環境: Dev ContainersのDocker in Docker構成
  - 動作確認環境: Dev Container / 通常のDocker-in-Docker / KinD
  - 聴講者に馴染みがある例としてKindやMinikubeを背景に提示できる。youki固有の問題に閉じずNested Containers一般の問題として提示する
- 技術背景: cgroup v2の制約
  - `no internal process constraint`が直接の原因
  - landlordのinit processがいるcgroupにtenantを追加しようとすると制約に抵触して失敗する
  - cgroup v1では同じ問題は起きない前提で、v2固有の制約として説明する
- 用語: landlord / tenant container
  - youki内部の用語なので主役にしすぎない
  - landlordはexec先のコンテナ、tenantはexecするプロセス
  - exec時はプロセスが「コンテナ内にある」状態にしなければいけないので、landlordのcgroup内にtenantが入る必要がある
  - 講演ではまず一般的な説明 (既存コンテナの中でさらにコンテナを動かし、その中で`exec`したプロセスを既存コンテナの適切なcgroupへ配置する) を置き、必要な場面だけyouki用語を導入する
- runcの設計: リトライ型 (retrying / fallback) アプローチ
  - 指定されたcgroupへの配置に失敗した後、landlordのinit processのcgroupへ移動して再試行する
  - 互換性を優先する代わりに、別の異常までfallbackで覆い隠す可能性がある
    - 元のエラーがEBUSY以外でもfallbackしてしまう可能性がある
    - 本来のcgroup消失など別の異常を隠してしまう
    - init processが別のcgroupへ移動されていた場合、OCIで指定した制限が適用されなくなる
    - `exec`自体は成功してしまうが、本来は失敗させるべき異常を見逃す可能性がある
  - PR上でも互換性のためにretryingを維持する判断が示されており、単純に悪い設計ではない
- youkiの設計: 事前推定 (proactive inference) 型アプローチ
  - エラーを受けて別経路へ逃がすのではなく、exec対象のinit processが実際に所属するsub-cgroupを先に推定し、そのパスを最初から利用する (#3564)
  - 具体的な仕組み: `/proc/<pid>/cgroup`を読む、cgroup v2か判定する、landlordのcgroup pathからsub-cgroupを推定する、tenant作成前に配置先へ反映する
  - 価値・シンプルさの内訳:
    - cgroup操作をcgroup driver経由で完結できる
    - EBUSYかどうかを判定してfallbackする必要がない
    - 失敗した処理を別の場所で成功扱いにするのではなく、最初から適切な配置先を選ぶ
    - 危険なfallback経路を避け、異常は異常として扱える
    - retryingに必要な型変換や複雑なエラー伝播を避けられる
    - 「リトライで拾う」ではなく元々「正しいであろうパス」を見つけることで、失敗をそのまま失敗として扱える
- 設計比較の軸 (講演の主役)
  - runc: 失敗後にfallbackする vs youki: exec開始前に実際の所属先を推定する
  - 互換性を最大化するのか、異常を広く隠さないのか
  - fallbackで救済するのか、最初から正しい配置先を推定するのか
  - 単なる「youkiのほうがきれい」という主観ではなく、互換性・異常の可視性・実装の複雑さ・cgroup driverとの責務分担のトレードオフとして見せる
- 互換性メモ (Issue記載のCompatibility notes for the sub-cgroup approach)
  - 1. landlordのinitがsub-cgroupにいるが、親コンテナの`cgroup.subtree_control`に有効なcontrollerがない場合:
    - 親cgroupへのjoinは成功する (no internal process constraintの例外。kernelドキュメント参照: <https://docs.kernel.org/admin-guide/cgroup-v2.html#:~:text=Note%20that%20the%20restriction%20doesn%E2%80%99t%20get%20in%20the%20way%20if%20there%20is%20no%20enabled%20controller%20in%20the%20cgroup%E2%80%99s%20%E2%80%9Ccgroup.subtree_control%E2%80%9D.>)
    - この場合`runc`は親cgroupへjoinさせるが、youki案は意図したsub-cgroupへjoinさせる
    - 実効リソース制限は同じなので、実害を許容できる差分としてメンテナーと合意したもの
  - 1. container initがまったく別のcgroupへ移動されている場合:
    - `runc`はこのケースの互換性を維持する
    - youki初回実装では対象外とする。KirとAleksaが問題ありと指摘したため
    - 後発プロダクトなので先行実装のworkaroundをすべて踏襲する必要はなく、その状態自体が異常である可能性があるため、複雑な互換処理より対象範囲を限定する明示的なスコープ判断。今後必要になれば対応できる
  - 表現注意: 「前バージョンとの互換性」ではなく「先行実装との互換性」「runcとの互換性」と書く
- 自分の貢献・経緯 (Why You?)
  - 問題の発見・原因調査・実装・設計は全て担当
  - runcの既存実装・PRを調べて初期案 (runc踏襲) を出し、そこから追加調査と実装テストを行うことで新案 (proactive inference) を設計した
  - 最初から独自案を思いついたのではなく、runc踏襲案を試作した後、実装とテストを通じて再設計した泥臭さが強み
  - レビューは全てyoukiのメンテナーによるもの
  - runcとyouki両方の実装・関連PRを調査し、設計比較に関わるレビュー・議論を経験した
- KubeCon Japanでの会合の位置づけ
  - 自分から提案した内容自体は変わっていない
  - youkiのメンテナーを納得させることができた
  - runcのメンテナーから意見・runc側の設計判断を聞いた上で、自分の提案方針を採用することが決まった
  - 実装の話と因果関係がある (採用判断の合意形成) ので残す価値がある。単なる「会合を開いた」紹介に留めない
  - コミュニティとコーディング双方のコントリビューションが繋がった場所
- 参加者の持ち帰り (2〜3個に絞る)
  - Nested Containersのcgroup配置問題を説明できる
  - runcとyoukiの実装差分を追える
  - 後発OSSが独自設計を選ぶ際の判断軸を得る
  - 補足軸: 先行実装をそのまま採用せず設計上の前提を検証する、互換性と異常の扱いをトレードオフとして比較する、実装・テスト・メンテナーとの議論を通じて採用設計を決める
- 30分構成案 (設計判断の比較を主役に)
  - 問題提起: Docker-in-DockerやKindで`exec`がEBUSYになる理由
  - 制約説明: cgroup v2のno internal process constraint
  - runcの設計: 失敗後にinit processのcgroupへfallbackするretrying
  - youkiの設計: init processのsub-cgroupを事前推定するproactive inference
  - 設計判断: 互換性、異常の扱い、実装の複雑さ、cgroup driverとの責務分担、何を維持し何を捨てたか
  - コミュニティでの合意形成: PRレビューとKubeCon Japanでのrunc / youkiメンテナーとの議論
- 伝えたい抽象的なメッセージ
  - 後発プロダクトが先行プロダクトの設計をそのまま踏襲する必要はない。
    - デファクトスタンダードは必ずしも「すべて正解」なわけではない
  - 後発プロダクトが、先行プロダクトが叶えられなかった「きれいな仕様・設計」を目指すことで、それは差別化・エコシステムの充実につながる
- 今回取り上げる具体的な技術話
  - Nested Containers + cgroup v2の組み合わせへの、コンテナランタイム側の対処
  - runcが採用した「リトライ型アプローチ」と、youkiが採用した「事前推定 (proactive inference) 型アプローチ」の違い
    - youkiの方が、起こりえる危険なパターンを排除できる かつ よりシンプルな設計になっている
    - ただし「危険」「シンプル」は具体化が必要。上記のfallbackの異常隠蔽リスクと、cgroup driver完結・エラー伝播削減の観点で説明する
  - 実装する中で新しいやり方を発見し提案するだけでなく、KubeCon Japanの場でruncのメンテナーとyoukiのメンテナーを交えて会合を行う、コミュニティとしての努力
    - これはまさにコミュニティとコーディング双方のコントリビューションが繋がった場所
- 表現上の注意 (レビューで確定した言い回し)
  - 「異なる解決法を選んだ」ではなく「runcの解決法を調査・試作した結果、youkiでは別の解決法を提案し実装・検証へ進めた」と書くのが正確
  - 「正しいパス」とだけ書かず「init processが実際に所属しているsub-cgroupを推定し、execプロセスの配置先として利用する」と書く
  - 「危険なパスを排除する」より「異常をfallbackで覆い隠さず、配置先を先に推定する」と書く (runc批判に見せない)
  - #3564は「マージ済みの完成機能」と断定せず「実装し、Dev Container・Docker-in-Docker・KinDで動作を確認した設計案」と書く (#3563はマージ済み)
  - youki継続利用の当事者性 (Dockerのrunc代替として開発環境で継続利用していたところ問題発見) を冒頭に必ず残す

## Track

## Session format

Session Presentation (30 minutes)

## Level

Intermediate

## Benefits to the ecosystem

## CNCF-hosted software

## Open source projects

## Bio

Takuto is an undergraduate student at Chiba Institute of Technology, Japan, pursuing a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He is the founder of Coding Kubernetes, a tutorial series that rebuilds Kubernetes core components from scratch to teach its core internals. Through this project, he shares practical insights with future cloud-native maintainers.
