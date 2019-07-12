# セキュリティー
<!-- Security -->
## イントロダクション <!-- Introduction -->
LXD は root ユーザで実行するデーモンです。
<!--
LXD is a daemon running as root.
-->

デフォルトではデーモンへのアクセスはローカルの UNIX ソケット経由でのみ
可能です。設定によって、 TLS ソケット上でネットワーク越しに同じ API を
公開することが可能です。
<!--
Access to that daemon is only possible over a local UNIX socket by default.
Through configuration, it's then possible to expose the same API over
the network on a TLS socket.
-->

**警告**: UNIX ソケット経由でのローカルアクセスは LXD へのフルのアクセスを
常に許可します。これはあらゆるファイルシステムのパスやデバイスをコンテナ
にアタッチする能力やコンテナの全てのセキュリティの機能を変更することも
含みます。あなたのシステムに root 権限でアクセスを許可するほど信頼できる人
だけにそのようなアクセスを許可するべきです。
<!--
**WARNING**: Local access to LXD through the UNIX socket always grants
full access to LXD. This includes the ability to attach any filesystem
paths or devices to any container as well as tweaking all security
features on containers. You should only give such access to someone who
you'd trust with root access to your system.
-->

リモート API は TLS クライアント証明書か Candid ベースの認証のどちらかを
使用します。 Canonical RBAC のサポートは Canded ベースの認証と組み合わせて
API クライアントが LXD で何を出来るかを制限するのに使えます。
<!--
The remote API uses either TLS client certificates or Candid based
authentication. Canonical RBAC support can be used combined with Candid
based authentication to limit what an API client may do on LXD.
-->

## TLS configuration
LXD デーモンとのリモートの通信は HTTPS 上の JSON を使って行います。
サポートしているプロトコルは TLS 1.2 以上です。
<!--
Remote communications with the LXD daemon happen using JSON over HTTPS.
The supported protocol must be TLS1.2 or better.
-->

全ての通信は完全な前方秘匿性 (Perfect Forward Secrecy; PFS) を使う必要があり、
暗号は強力な楕円曲線のもの (ECDHE-RSA や ECDHE-ECDSA など)に限定されます。
<!--
All communications must use perfect forward secrecy and ciphers must be
limited to strong elliptic curve ones (such as ECDHE-RSA or ECDHE-ECDSA).
-->

生成されるキーは最低でも 4096 ビットのRSAでなければならず、 EC384 が好ましいです。
署名を使う場合は SHA-2 の署名だけが信頼されます。
<!--
Any generated key should be at least 4096bit RSA, preferably EC384 and
when using signatures, only SHA-2 signatures should be trusted.
-->

LXD を導入する際はクライアントとサーバの両方を管理するので、後方互換性の
ために弱いプロトコルや暗号をサポートする理由はありません。
<!--
Since we control both client and server, there is no reason to support
any backward compatibility to broken protocol or ciphers.
-->

クライアントとサーバの両方が初回起動時に証明書とキーのペアを生成します。
サーバは LXD ソケットとの全ての https 通信にそれを使用し、クライアントは
その証明書をクライアント・サーバ間の通信にクライアント証明書として使用します。
<!--
Both the client and the server will generate a keypair the first time
they're launched. The server will use that for all https connections to
the LXD socket and the client will use its certificate as a client
certificate for any client-server communication.
-->

証明書を再生成するには単に古い証明書を消すだけです。次に接続する際に
新しい証明書が生成されます。
<!--
To cause certificates to be regenerated, simply remove the old ones. On the
next connection a new certificate will be generated.
-->

## Role Based Access Control (RBAC)
LXD は Canonical RBAC サービスとの統合をサポートします。
<!--
LXD supports integrating with the Canonical RBAC service.
-->

これは Candid ベースの認証を用い、 RBAC サービスがユーザ／グループと
ロールの関係を管理します。ロールは個々のプロジェクト、全てのプロジェクト、
あるいは LXD インスタンス全体に割り当てることができます。
<!--
This uses Candid based authentication with the RBAC service maintaining
roles to user/group relationships. Roles can be assigned to individual
projects, to all projects or to the entire LXD instance.
-->

ロールはプロジェクトに割り当てられると以下のような意味を持ちます。
<!--
The meaning of the roles when applied to a project is as follow:
-->

 - auditor: プロジェクトへの読み取り専用のアクセス <!-- Read-only access to the project -->
 - user: 通常のライフサイクルアクション（起動、停止、…）の実行、コンテナ内でのコマンドの実行、コンソールへのアタッチ、スナップショットの管理、… を行う能力 <!-- Ability to do normal lifecycle actions (start, stop, ...),
   execute commands in the containers, attach to console, manage snapshots, ... -->
 - operator: 上記のすべてに加えてコンテナとイメージを作成、再設定、そして削除する能力 <!-- All of the above + the ability to create, re-configure and
   delete containers and images -->
 - admin: 上記のすべてに加えてプロジェクト自体を再構成する能力 <!-- All of the above + the ability to reconfigure the project itself -->

**警告**: これらのロールのうち現状では `auditor` と `user` だけが
ホストへの root 権限のアクセスを渡す信頼が持てないユーザに適した
ロールです。
<!--
**WARNING**: Of those roles, only `auditor` and `user` are currently
suitable for a user whom you wouldn't trust with root access to the
host.
-->

## Container security
LXD コンテナはかなり広い範囲のセキュリティの機能を利用可能です。
<!--
LXD containers can use a pretty wide range of features for security.
-->

デフォルトではコンテナは非特権 (`unprivileged`) です。これはコンテナが
ユーザ名前空間で稼働することを意味し、コンテナ内のユーザの能力をホスト上の
通常ユーザの能力に制限し、コンテナが所有するデバイスにも限定した権限しか
与えないことを意味します。
<!--
By default containers are `unprivileged`, meaning that they operate
inside a user namespace, restricting the abilities of users in the
container to that of regular users on the host with limited privileges
on the devices that the container owns.
-->

コンテナ間のデータ共有が不要であれば、 `security.idmap.isolated` を
有効にすることで各コンテナに対する uid/gid のマッピングをオーバーラップ
しないようにでき、他のコンテナへの潜在的な DoS 攻撃を防ぐことができます。
<!--
If data sharing between containers isn't needed, it is possible to
enable `security.idmap.isolated` which will use non-overlapping uid/gid
maps for each container, preventing potential DoS attacks on other
containers.
-->

もし望む場合は LXD は特権 (`privileged`) コンテナを実行することもできます。
ただし、これらは root 権限を取得しようとする行為に対して安全ではないこと、
特権コンテナ内の root 権限を持つユーザはホストに DoS をすることができ、
コンテナ内への監禁から脱出する方法を見つけるかもしれないことに注意してください。
<!--
LXD can also run `privileged` containers if you so wish, do note that
those aren't root safe and a user with root in such a container will be
able to DoS the host as well as find ways to escape confinement.
-->

コンテナのセキュリティとカーネル機能についてのより詳細は
[LXC のセキュリティページ](https://linuxcontainers.org/ja/lxc/security/).
を参照してください。
<!--
More details on container security and the kernel features we use can be found on the
[LXC security page](https://linuxcontainers.org/lxc/security/).
-->

## TLS クライアント証明書での認証を使ってリモートを追加する <!-- Adding a remote with TLS client certificate authentication -->
デフォルトのセットアップでは、ユーザが `lxd remote add` で新しいサーバを
追加する際、サーバに https で通信し、証明書がダウンロードされ、
フィンガープリントがユーザに表示されます。
<!--
In the default setup, when the user adds a new server with `lxc remote add`,
the server will be contacted over HTTPS, its certificate downloaded and the
fingerprint will be shown to the user.
-->

ユーザは、これが本当にサーバのフィンガープリントなのかの確認を求められます。
これは接続してみて手動で確認するか、既にそのサーバに接続可能になっている
別のユーザに info コマンドを実行してもらい、表示されたフィンガープリント
と比較することで確認できます。
<!--
The user will then be asked to confirm that this is indeed the server's
fingerprint which they can manually check by connecting to or asking
someone with access to the server to run the info command and compare
the fingerprints.
-->

その後ユーザはそのサーバのトラスト・パスワード
(訳注: サーバに接続できる権限があることを確認するためのパスワード) を
入力する必要があります。正しいパスワードを入力すると、クライアント証明書が
サーバのトラスト・ストア (訳注: 信頼済みクライアントストア) に追加され、
今後はクライアントは追加の機密情報を提供することなく、サーバに接続できます。
<!--
After that, the user must enter the trust password for that server, if
it matches, the client certificate is added to the server's trust store
and the client can now connect to the server without having to provide
any additional credentials.
-->

このワークフローは SSH が未知のサーバに初めて接続したときにプロンプトが
表示されるのと非常に似ています。
<!--
This is a workflow that's very similar to that of SSH where an initial
connection to an unknown server triggers a prompt.
-->

## PKI ベースのセットアップで TLS クライアントを使ってリモートを追加する <!-- Adding a remote with a TLS client in a PKI based setup -->
PKI ベースのセットアップではシステム管理者は中心となる PKI を運営します。
その PKI が全ての lxc クライアント用のクライアント証明書と全ての LXD
デーモンのサーバ証明書を発行します。
<!--
In the PKI setup, a system administrator is managing a central PKI, that
PKI then issues client certificates for all the lxc clients and server
certificates for all the LXD daemons.
-->

それらの証明書と鍵はさまざまなマシンに手動で配置され、自動生成された
証明書と鍵を置き換えます。
<!--
Those certificates and keys are manually put in place on the various
machines, replacing the automatically generated ones.
-->

CA 証明書も全てのマシンに追加します。
<!--
The CA certificate is also added to all machines.
-->

このモードでは、 LXD デーモンへの通信は予め配置しておいた CA 証明書を
使って行われます。サーバ証明書が CA によって署名されていなければ、
通信は単に通常の認証機構 (訳注: 上記のデフォルトのセットアップでリモート
を追加する際の手順) を通ることになります。
<!--
In that mode, any connection to a LXD daemon will be done using the
preseeded CA certificate. If the server certificate isn't signed by the
CA, the connection will simply go through the normal authentication
mechanism.
-->

サーバ証明書が有効で CA によって署名されていれば、その証明書について
ユーザにプロンプトを表示することなく接続は続行します。
<!--
If the server certificate is valid and signed by the CA, then the
connection continues without prompting the user for the certificate.
-->

その後、ユーザはそのサーバのトラスト・パスワード
を入力する必要があります。
正しいパスワードを入力すると、クライアント証明書がサーバのトラスト・ストアに追加され、
今後はクライアントは追加の機密情報を提供することなく、サーバに接続できます。
<!--
After that, the user must enter the trust password for that server, if
it matches, the client certificate is added to the server's trust store
and the client can now connect to the server without having to provide
any additional credentials.
-->

PKI モードを有効にするには、クライアントの設定ディレクトリ (`~/.config/lxc`) に
client.ca ファイルを追加し、サーバの設定ディレクトリ (`/var/lib/lxd`) に
server.ca ファイルを追加します。さらにクライアント用にクライアント証明書を
CA によって発行し、サーバ用にサーバ証明書を発行します。それらの証明書で
事前に自動生成されたファイルを置き換える必要があります。
<!--
Enabling PKI mode is done by adding a client.ca file in the
client's configuration directory (`~/.config/lxc`) and a server.ca file in
the server's configuration directory (`/var/lib/lxd`). Then a client
certificate must be issued by the CA for the client and a server
certificate for the server. Those must then replace the existing
pre-generated files.
-->

この後、サーバを再起動すると PKI モードで起動します。
<!--
After this is done, restarting the server will have it run in PKI mode.
-->

## Candid 認証を使ってでリモートを追加する <!-- Adding a remote with Candid authentication -->
LXD を Candid を使うように設定した場合、 LXD はクライアントが
`candid.api.url` の設定に指定した認証サーバから Discharge トークンを
取得して認証を試みるように依頼します。
<!--
When LXD is configured with Candid, it will request that clients trying to
authenticating with it get a Discharge token from the authentication server
specified by the `candid.api.url` setting.
-->

認証サーバの証明書は LXD サーバに信頼される必要があります。
<!--
The authentication server certificate needs to be trusted by the LXD server.
-->

Macaroon 認証を設定された LXD にリモートを追加するには
`lxd remote add REMOTE ENDPOINT --auth-type=candid`
を実行します。クライアントはユーザを検証するために認証サーバに要求される
機密情報を入力するためのプロンプトを表示します。認証が成功したら、
認証サーバから受け取ったトークンを LXD サーバに渡して接続します。
LXD サーバはトークンを検証し、リクエストを認証します。トークンはクッキーとして
保存され、クライアントが LXD にリクエストを送る度にサーバに渡されます。
<!--
To add a remote pointing to a LXD configured with Macaroon auth, run `lxc
remote add REMOTE ENDPOINT \-\-auth-type=candid`.  The client will prompt for
the credentials required by the authentication server in order to verify the
user. If the authentication is successful, it will connect to the LXD server
presenting the token received from the authentication server.  The LXD server
verifies the token, thus authenticating the request.  The token is stored as
cookie and is presented by the client at each request to LXD.
-->

## 信頼された TLS クライントを管理する <!-- Managing trusted TLS clients -->
LXD サーバが信頼している TLS 証明書の一覧は `lxc config trust list` で
取得できます。
<!--
The list of TLS certificates trusted by a LXD server can be obtained with
`lxc config trust list`.
-->

クライアントは `lxc config trust add <file>` を使用して手動で追加できます。
これにより既存の管理者が新しいクライアント証明書を直接トラスト・ストアに
追加することによって共有されたトラスト・パスワードの必要性を無くします。
<!--
Clients can manually be added using `lxc config trust add <file>`,
removing the need for a shared trust password by letting an existing
administrator add the new client certificate directly to the trust store.
-->

クライアントへの信頼を取り消すには `lxc config trust remove FINGERPRINT` を
実行すると証明書が削除されます。
<!--
To revoke trust to a client its certificate can be removed with `lxc config
trust remove FINGERPRINT`.
-->

## TLS 認証でのパスワード・プロンプト <!-- Password prompt with TLS authentication -->
管理者によって事前に信頼関係がセットアップされていない場合に
新しい信頼関係を確立するには、サーバにパスワードを設定し、クライアントが
自身をサーバに登録する際にそのパスワードを送る必要があります。
<!--
To establish a new trust relationship when not already setup by the
administrator, a password must be set on the server and sent by the
client when adding itself.
-->

ですので、リモートを追加する操作は次のようになります。
<!--
A remote add operation should therefore go like this:
-->

 1. API の GET /1.0 を呼びます。 <!-- Call GET /1.0 -->
 2. PKI のセットアップをしていなければ、フィンガープリントを確認するプロンプトがユーザに表示されます。 <!-- If we're not in a PKI setup ask the user to confirm the fingerprint. -->
 3. サーバから返された dict を見て、 "auth" が "untrusted" だった場合、ユーザにサーバのパスワードを入力させ、 `/1.0/certificates` に `POST` のリクエストを送り、その後再び `/1.0` のリクエストを送って本当に信頼されたかを確認します。 <!-- Look at the dict we received back from the server. If "auth" is
    "untrusted", ask the user for the server's password and do a `POST` to
    `/1.0/certificates`, then call `/1.0` again to check that we're indeed
    trusted. -->
 4. これでリモートが準備完了になりました。 <!-- Remote is now ready -->

## 失敗のシナリオ <!-- Failure scenarios -->
### サーバ証明書が変更されていた場合 <!-- Server certificate changes -->
典型的には次の 2 つの場合があるでしょう。
<!--
This will typically happen in two cases:
-->

 * サーバが完全に再インストールされたため証明書が変わった <!-- The server was fully reinstalled and so changed certificate -->
 * 接続がマン・イン・ザ・ミドル (MITM) 攻撃によりインターセプトされた <!-- The connection is being intercepted (MITM) -->

これらのケースでは、サーバ証明書のフィンガープリントが
(訳注: ローカルに保存されていた) このリモート用の設定に含まれる
フィンガープリントと一致しないため、クライアントはサーバへの接続を拒否します。
<!--
In such cases the client will refuse to connect to the server since the
certificate fringerprint will not match that in the config for this
remote.
-->

後はユーザの責任でサーバ管理者に連絡し、サーバ証明書が本当に変更されたのか
確認する必要があります。変更されたのであれば証明書を新しいもので置き換えるか、
リモートを一旦削除して再度追加できます。
<!--
It is then up to the user to contact the server administrator to check
if the certificate did in fact change. If it did, then the certificate
can be replaced by the new one or the remote be removed altogether and
re-added.
-->

### サーバ上の信頼関係が取り消された <!-- Server trust relationship revoked -->
このケースでは、サーバは同じ証明書をまだ使っていますが、全ての API 呼び出しは
クライアントが信頼されていないことを示す 403 エラーを返します。
<!--
In this case, the server still uses the same certificate but all API
calls return a 403 with an error indicating that the client isn't
trusted.
-->

これは別の信頼されたクライアントかローカルのサーバ管理者がサーバ上の
信頼エントリを削除したときに起こります。
<!--
This happens if another trusted client or the local server administrator
removed the trust entry on the server.
-->

## プロダクションのセットアップ <!-- Production setup -->
プロダクション環境のセットアップでは、全てのクライアントを追加した後、
`core.trust_password` の設定を削除することを推奨します。削除することにより
パスワードを推測しようとするブルート・フォース攻撃を防ぐことができます。
<!--
For production setup, it's recommended that `core.trust_password` is unset
after all clients have been added.  This prevents brute-force attacks trying to
guess the password.
-->

さらに `core.https_address` をサーバにアクセス可能な単一のアドレスに設定し
(ホスト上の任意のアドレスではなく) 、許可されたホストやサブネットからのみ
LXD のポートへのアクセスを許可するようにファイアウォールのルールを設定すべきです。
<!--
Furthermore, `core.https_address` should be set to the single address where the
server should be available (rather than any address on the host), and firewall
rules should be set to only allow access to the LXD port from authorized
hosts/subnets.
-->
