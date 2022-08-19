# コントリビュート

プロジェクトに貢献する前に、以下のガイドラインを確認してください。

## プルリクエスト
このプロジェクトへの変更は、Githubでプルリクエストとして提案してください。
<https://github.com/lxc/lxd>

そのあと、提案はコードレビューを経て承認され、メインブランチにマージされます。

## コミット構成
コミットを次のように分類する必要があります:

 - API 拡張 (`doc/api-extensions.md` と `shared/version.api.go` を含む変更に対して `api: Add XYZ extension`)
 - ドキュメント (`doc/` 内のファイルに対して `doc: Update XYZ`)
 - API 構造 (`shared/api/` の変更に対して `shared/api: Add XYZ`)
 - Go クライアントパッケージ (`client/` の変更に対して `client: Add XYZ`)
 - CLI (`lxc/` の変更に対して `lxc/<command>: Change XYZ`)
 - スクリプト (`scripts/` の変更に対して `scripts: Update bash completion for XYZ`)
 - LXD デーモン (`lxd/` の変更に対して `lxd/<package>: Add support for XYZ`)
 - テスト (`tests/` の変更に対して `tests: Add test for XYZ`)

同様のパターンが LXD コードツリーの他のツールにも適用されます。
そして複雑さによっては、さらに小さな単位に分けられるかもしれません。

CLI ツール (`lxc/`) 内の文字列を更新する際は、テンプレートを更新してコミットする必要があるでしょう:

    make i18n
    git commit -a -s -m "i18n: Update translation templates" po/

このようにすることで、コントリビューションに対するレビューが容易になり、stable ブランチへバックポートするプロセスが大幅に簡素化されます。

## ライセンスと著作権
デフォルトで、このプロジェクトに対するいかなる貢献も Apache 2.0 ライセンスの下で行われます。

変更の著者は、そのコードに対する著作権を保持します (著作権の譲渡はありません)。

## 開発者の起源の証明
このプロジェクトへの貢献の追跡を改善するために、DCO 1.1を使用しており、ブランチに入るすべての変更に対して「サインオフ」手順を使用しています。

サインオフとは、あなたがそのコミットを書いたことを証明する、そのコミットの説明の最後にある単純な行です。
この行は、自分が書いたものであることを証明したり、オープンソースとして渡す権利があることを証明したりします。

> Developer Certificate of Origin
> Version 1.1
>
> Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
> 660 York Street, Suite 102,
> San Francisco, CA 94110 USA
>
> Everyone is permitted to copy and distribute verbatim copies of this
> license document, but changing it is not allowed.
>
> Developer's Certificate of Origin 1.1
>
> By making a contribution to this project, I certify that:
>
> (a) The contribution was created in whole or in part by me and I
>     have the right to submit it under the open source license
>     indicated in the file; or
>
> (b) The contribution is based upon previous work that, to the best
>     of my knowledge, is covered under an appropriate open source
>     license and I have the right under that license to submit that
>     work with modifications, whether created in whole or in part
>     by me, under the same open source license (unless I am
>     permitted to submit under a different license), as indicated
>     in the file; or
>
> (c) The contribution was provided directly to me by some other
>     person who certified (a), (b) or (c) and I have not modified
>     it.
>
> (d) I understand and agree that this project and the contribution
>     are public and that a record of the contribution (including all
>     personal information I submit with it, including my sign-off) is
>     maintained indefinitely and may be redistributed consistent with
>     this project or the open source license(s) involved.

有効なサインオフラインの例は以下の通りです。

```
Signed-off-by: Random J Developer <random@developer.org>
```

実名と有効な電子メールアドレスを使用してください。
残念ながら、ペンネームや匿名での投稿はできません。

また、それぞれのコミットには作者が個別に署名する必要があります。
大きなセットの一部であってもです。`git commit -s`が役に立つでしょう。

## Code of Conduct

コントリビュートする際には、行動規範を遵守しなければなりません。行動規範は、以下のサイトから入手できます。 [`https://github.com/lxc/lxd/blob/master/CODE_OF_CONDUCT.md`](https://github.com/lxc/lxd/blob/master/CODE_OF_CONDUCT.md)

<!-- Include end contributing -->

## その他の情報

より詳しい情報は、ドキュメントの [Contributing](doc/contributing.md) をご覧ください。
