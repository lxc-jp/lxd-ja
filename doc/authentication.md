---
discourse: 13114,15142
relatedlinks: https://www.youtube.com/watch?v=6O0q3rSWr8A
---

(authentication)=
# リモートAPI認証

LXDデーモンとのリモート通信は、HTTPS上のJSONを使って行われます。

リモートAPIにアクセスするためには、クライアントはLXDサーバとの間で認証を行う必要があります。
以下の認証方法がサポートされています。

- {ref}`authentication-tls-certs`
- {ref}`authentication-candid`
- {ref}`authentication-rbac`

(authentication-tls-certs)=
## TLSクライアント証明書

```{youtube} https://www.youtube.com/watch?v=4iNpiL-lrXU
```

認証に{abbr}`TLS(Transport Layer Security)`クライアント証明書を使用する場合、クライアントとサーバーの両方が最初に起動したときにキーペアを生成します。
サーバはそのキーペアをLXDソケットへの全てのHTTPS接続に使用します。
クライアントは、その証明書をクライアント証明書として、あらゆるクライアント・サーバ間の通信に使用します。

証明書を再生成させるには、単に古いものを削除します。
次の接続時には、新しい証明書が生成されます。

### 通信プロトコル

通信プロトコルは TLS1.2 以上に対応しています。
すべての通信には完全な前方秘匿を使用し、暗号は強力な楕円曲線（ECDHE-RSA や ECDHE-ECDSA など）に限定してください。

生成される鍵は最低でも4096ビットのRSA、できれば384ビットのECDSAが望ましいです。
署名を使用する場合は、SHA-2署名のみを信頼すべきです。

我々はクライアントとサーバーの両方を管理しているので、壊れたプロトコルや暗号の下位互換をサポートする理由はありません。

(authentication-trusted-clients)=
### 信頼できるTLSクライアント

LXDサーバが信頼するTLS証明書のリストは、`lxc config trust list`で取得できます。

信頼できるクライアントは以下のいずれかの方法で追加できます。

- {ref}`authentication-add-certs`
- {ref}`authentication-trust-pw`
- {ref}`authentication-token`

サーバーとの認証を行うワークフローは、SSHの場合と同様で、未知のサーバーへの初回接続時にプロンプトが表示されます。

1. ユーザーが `lxc remote add` でサーバーを追加すると、HTTPS でサーバーに接続され、その証明書がダウンロードされ、フィンガープリントがユーザーに表示されます。
1. ユーザーは、これが本当にサーバーのフィンガープリントであることを確認するよう求められます。これは、サーバーに接続して手動で確認するか、サーバーにアクセスできる人に info コマンドを実行してフィンガープリントを比較してもらうことで確認できます。
1. サーバーはクライアントの認証を試みます。

   - クライアント証明書がサーバーのトラストストアにある場合は、接続が許可されます。
   - クライアント証明書がサーバーのトラストストアにない場合、サーバーはユーザーにトークンまたはトラストパスワードの入力を求めます。
     提供されたトークンまたはトラストパスワードが一致した場合、クライアント証明書はサーバーのトラストストアに追加され、接続が許可されます。
     そうでない場合は、接続が拒否されます。

クライアントへの信頼を取り消すには、`lxc config trust remove FINGERPRINT`でそのクライアント証明書をサーバーから削除します。

TLSクライアントを1つまたは複数のプロジェクトに制限することが可能です。
この場合、クライアントは、グローバルな構成変更の実行や、アクセスを許可されたプロジェクトの構成（制限、制約）の変更もできなくなります。

アクセスを制限するには、`lxc config trust edit FINGERPRINT`を使用します。
`restricted`キーを`true`に設定し、クライアントのアクセスを制限するプロジェクトのリストを指定します。
プロジェクトのリストが空の場合、クライアントはどのプロジェクトへのアクセスも許可されません。

(authentication-add-certs)=
#### 信頼できる証明書をサーバーに追加する

信頼できるクライアントを追加するには、そのクライアント証明書をサーバーのトラストストアに直接追加するのが望ましい方法です。
これを行うには、クライアント証明書をサーバーにコピーし、`lxc config trust add <file>`で登録します。

(authentication-trust-pw)=
#### トラストパスワードを使ったクライアント証明書の追加

クライアント側から新しい信頼関係を確立できるようにするには、サーバーにトラストパスワード([`core.trust_password`](server-options-core))を設定する必要があります。クライアントは、プロンプト時にトラストパスワードを入力することで、自分の証明書をサーバのトラストストアに追加することができます。

本番環境では、すべてのクライアントが追加された後に、`core.trust_password`の設定を解除してください。
これにより、パスワードを推測しようとするブルートフォース攻撃を防ぐことができます。

(authentication-token)=
#### トークンを使ったクライアント証明書の追加

トークンを使って新しいクライアントを追加することもできます。
トークンは調整可能な時間([`core.remote_token_expiry`](server-options-core))を過ぎるか一度使用すると無効になるため、これはトラストパスワードを使用するよりも安全な方法です。

この方法を使用するには，クライアント名の入力を促す `lxc config trust add` を呼び出して，各クライアント用のトークンを生成します。
その後，クライアントは，トラストパスワードの入力を求められたときに生成されたトークンを提供することで，自分の証明書をサーバのトラストストアに追加することができます。

<!-- Include start NAT authentication -->

```{note}
LXD サーバが NAT の後ろ側にいる場合、クライアント用のリモートを追加する際には外部のパブリックアドレスを指定する必要があります。

    lxc remote add <name> <IP_address>

admin パスワードのプロンプトが表示されたら、生成されたトークンを入力してください。

サーバでトークンを生成する際、 LXD はクライアントがサーバにアクセスするために使える IP アドレスのリストを含めます。
しかし、サーバが NAT の後ろ側にいる場合、これらのアドレスはクライアントが接続できないローカルアドレスの場合があります。
その場合、手動で外部アドレスを指定する必要があります。
```

<!-- Include end NAT authentication -->

あるいは、クライアントはリモートの追加時にトークンを直接提供することもできます。`lxc remote add <name> <token>`.

### PKI システムの使用

{abbr}`PKI (Public key infrastructure)`の設定では、システム管理者が中央のPKIを管理し、すべてのLXDクライアント用のクライアント証明書とすべてのLXDデーモン用のサーバー証明書を発行します。

PKIモードを有効にするには、以下の手順を実行します。

1. すべてのマシンに{abbr}`CA（認証局）`の証明書を追加します。

   - クライアントの設定ディレクトリ（`~/.config/lxc`）に`client.ca`ファイルを配置する。
   - `server.ca`ファイルをサーバの設定ディレクトリ（`/var/lib/lxd`またはsnapユーザの場合は`/var/snap/lxd/common/lxd`）に置く。
1. CAから発行された証明書をクライアントとサーバーに配置し、自動生成された証明書を置き換える。
1. サーバーを再起動します。

このモードでは、LXDデーモンへの接続はすべて、事前に発行されたCA証明書を使って行われます。

もしサーバ証明書がCAによって署名されていなければ、接続は単に通常の認証メカニズムを通過します。
サーバ証明書が有効でCAによって署名されていれば、ユーザに証明書を求めるプロンプトを出さずに接続を続行します。

生成された証明書は自動的には信頼されないことに注意してください。そのため、{ref}`authentication-trusted-clients`で説明している方法のいずれかで、サーバーに追加する必要があります。

(authentication-candid)=
## Candidベースの認証

```{youtube} https://www.youtube.com/watch?v=FebTipM1jJk
```

[`candid.*`](server-options-candid-rbac)サーバオプションにより[Candid](https://github.com/canonical/candid)認証を使うようにLXDを設定できます。
この場合、サーバーで認証を行おうとするクライアントは、[`candid.api.url`](server-options-candid-rbac)で指定された認証サーバーからディスチャージトークンを取得しなければなりません。

認証サーバーの証明書は、LXDサーバーから信頼されていなければなりません。

Candid/Macaroon認証を設定したLXDサーバにリモートポインティングを追加するには、`lxc remote add REMOTE ENDPOINT --auth-type=candid`を実行します。
ユーザーを確認するために、クライアントは認証サーバーが要求する認証情報の入力を求められます。
認証が成功した場合、クライアントはLXDサーバに接続し、認証サーバから受け取ったトークンを提示します。
LXDサーバはトークンを検証し、リクエストを認証します。
トークンはクッキーとして保存され、クライアントがLXDにリクエストするたびに提示されます。

Candidベースの認証を設定する方法については、チュートリアルの[Candid authentication for LXD](https://ubuntu.com/tutorials/candid-authentication-lxd)を参照してください。

(authentication-rbac)=
## 役割ベースのアクセスコントロール(RBAC)

```{youtube} https://www.youtube.com/watch?v=VE60AbJHT6E
```

LXDはCanonicalのRBACサービスとの連携をサポートしています。
Candidベースの認証と組み合わせることで、{abbr}`RBAC (Role Based Access Control)`は、APIクライアントがLXD上でできることを制限するために使うことができます。

このような設定では、認証はCandidを通して行われ、RBACサービスはユーザー/グループの関係に役割を維持します。
ロールは個々のプロジェクトにも、すべてのプロジェクトにも、あるいはLXDインスタンス全体にも割り当てることができます。

プロジェクトに適用された場合のロールの意味は以下の通りです。

- 監査役: プロジェクトへの読み取り専用のアクセス権
- ユーザー: 通常のライフサイクルアクション（開始、停止、...）を実行する能力。
        インスタンスでのコマンド実行、コンソールへのアタッチ、スナップショットの管理など。
- オペレーター: 上記のすべての機能に加え、インスタンスやイメージの作成、再設定、削除を行う機能
            インスタンスとイメージの作成、再設定、削除
- 管理者: 上記の機能に加えて、プロジェクト自体を再構成する機能を持つ

```{important}
制限のないプロジェクトでは、`auditor`と`user`のロールだけが、ホストへのルートアクセスを任せられないユーザーに適しています。

また、{ref}`制限付きプロジェクト <projects-restrictions>` では、適切に設定されていれば、`operator` ロールも安全に使用することができます。
```

LXDサーバでRBACを有効にするには[`rbac.*`](server-options-candid-rbac)サーバオプションを設定してください。これは`candid.*`オプションのスーパーセットで、LXDをRBACサービスに統合できます。

(authentication-server-certificate)=
## TLS サーバ証明書

LXD は {abbr}`ACME (Automatic Certificate Management Environment)` サービス (例えば [Let's Encrypt](https://letsencrypt.org/)) を使ったサーバ証明書の発行をサポートします。

この機能を有効にするには以下の{ref}`サーバ設定 <server-options-acme>`をしてください。

- `acme.domain`: 証明書を発行するドメイン。
- `acme.email`: ACME サービスのアカウントに使用する email アドレス。
- `acme.agree_tos`: ACME サービスの利用規約に同意するためには `true` に設定する必要あり。
- `acme.ca_url`: ACME サービスのディレクトリ URL。デフォルトでは LXD は "Let's Encrypt" を使用。

この機能を利用するには、 LXD は 80 番ポートを開放する必要があります。
これは [HAProxy](http://www.haproxy.org/) のようなリバースプロキシを使用することで実現できます。

以下は `lxd.example.net` をドメインとして使用する HAProxy の最小限の設定です。
証明書が発行された後、 LXD は`https://lxd.example.net/` でアクセスできます。

```
# Global configuration
global
  log /dev/log local0
  chroot /var/lib/haproxy
  stats socket /run/haproxy/admin.sock mode 660 level admin
  stats timeout 30s
  user haproxy
  group haproxy
  daemon
  ssl-default-bind-options ssl-min-ver TLSv1.2
  tune.ssl.default-dh-param 2048
  maxconn 100000

# Default settings
defaults
  mode tcp
  timeout connect 5s
  timeout client 30s
  timeout client-fin 30s
  timeout server 120s
  timeout tunnel 6h
  timeout http-request 5s
  maxconn 80000

# Default backend - Return HTTP 301 (TLS upgrade)
backend http-301
  mode http
  redirect scheme https code 301

# Default backend - Return HTTP 403
backend http-403
  mode http
  http-request deny deny_status 403

# HTTP dispatcher
frontend http-dispatcher
  bind :80
  mode http

  # Backend selection
  tcp-request inspect-delay 5s

  # Dispatch
  default_backend http-403
  use_backend http-301 if { hdr(host) -i lxd.example.net }

# SNI dispatcher
frontend sni-dispatcher
  bind :443
  mode tcp

  # Backend selection
  tcp-request inspect-delay 5s

  # require TLS
  tcp-request content reject unless { req.ssl_hello_type 1 }

  # Dispatch
  default_backend http-403
  use_backend lxd-nodes if { req.ssl_sni -i lxd.example.net }

# LXD nodes
backend lxd-nodes
  mode tcp

  option tcp-check

  # Multiple servers should be listed when running a cluster
  server lxd-node01 1.2.3.4:8443 check
  server lxd-node02 1.2.3.5:8443 check
  server lxd-node03 1.2.3.6:8443 check
```

## 失敗のシナリオ

以下のシナリオでは認証は失敗します。

### サーバ証明書が変更された場合

サーバ証明書は以下の場合に変更されるかも知れません。

- サーバが完全に再インストールされたため新しい証明書に変わった。
- 接続がインターセプトされた ({abbr}`MITM (Machine in the middle)`)。

このような場合、このリモートの設定内のフィンガープリントと証明書のフィンガープリントが一致しないため、クライアントはサーバへの接続を拒否します。

この場合サーバ管理者に連絡して証明書が実際に変更されたのかを確認するのはユーザ次第です。
実際に変更されたのであれば、証明書を新しいものに置き換えるか、リモートを削除して追加し直すことができます。

### サーバとの信頼関係が取り消された場合

別の信頼されたクライアントまたはローカルのサーバ管理者がサーバ上で対象のクライアントの信頼エントリを削除した場合、そのクライアントに対するサーバの信頼関係は取り消されます。

この場合、サーバは引き続き同じ証明書を使用していますが、全ての API 呼び出しは対象のクライアントが信頼されていないことを示す 403 のステータスコードを返します。
