# セキュリティー
<!-- Security -->
## イントロダクション <!-- Introduction -->
LXD は root ユーザーで実行するデーモンです。
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
常に許可します。これはあらゆるファイルシステムのパスやデバイスをインスタンス
にアタッチする能力やインスタンスの全てのセキュリティの機能を変更することも
含みます。あなたのシステムに root 権限でアクセスを許可するほど信頼できる人
だけにそのようなアクセスを許可するべきです。
<!--
**WARNING**: Local access to LXD through the UNIX socket always grants
full access to LXD. This includes the ability to attach any filesystem
paths or devices to any instance as well as tweaking all security
features on instances. You should only give such access to someone who
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

これは Candid ベースの認証を用い、 RBAC サービスがユーザー／グループと
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
 - user: 通常のライフサイクルアクション（起動、停止、…）の実行、インスタンス内でのコマンドの実行、コンソールへのアタッチ、スナップショットの管理、… を行う能力 <!-- Ability to do normal lifecycle actions (start, stop, ...),
   execute commands in the instances, attach to console, manage snapshots, ... -->
 - operator: 上記のすべてに加えてインスタンスとイメージを作成、再設定、そして削除する能力 <!-- All of the above + the ability to create, re-configure and
   delete instances and images -->
 - admin: 上記のすべてに加えてプロジェクト自体を再構成する能力 <!-- All of the above + the ability to reconfigure the project itself -->

**警告**: これらのロールのうち現状では `auditor` と `user` だけが
ホストへの root 権限のアクセスを渡す信頼が持てないユーザーに適した
ロールです。
<!--
**WARNING**: Of those roles, only `auditor` and `user` are currently
suitable for a user whom you wouldn't trust with root access to the
host.
-->

## Container security
LXD コンテナーはかなり広い範囲のセキュリティの機能を利用可能です。
<!--
LXD containers can use a pretty wide range of features for security.
-->

デフォルトではコンテナーは非特権 (`unprivileged`) です。これはコンテナーが
ユーザー名前空間で稼働することを意味し、コンテナー内のユーザーの能力をホスト上の
通常ユーザーの能力に制限し、コンテナーが所有するデバイスにも限定した権限しか
与えないことを意味します。
<!--
By default containers are `unprivileged`, meaning that they operate
inside a user namespace, restricting the abilities of users in the
container to that of regular users on the host with limited privileges
on the devices that the container owns.
-->

コンテナー間のデータ共有が不要であれば、 `security.idmap.isolated` を
有効にすることで各コンテナーに対する uid/gid のマッピングをオーバーラップ
しないようにでき、他のコンテナーへの潜在的な DoS 攻撃を防ぐことができます。
<!--
If data sharing between containers isn't needed, it is possible to
enable `security.idmap.isolated` which will use non-overlapping uid/gid
maps for each container, preventing potential DoS attacks on other
containers.
-->

もし望む場合は LXD は特権 (`privileged`) コンテナーを実行することもできます。
ただし、これらは root 権限を取得しようとする行為に対して安全ではないこと、
特権コンテナー内の root 権限を持つユーザーはホストに DoS をすることができ、
コンテナー内への監禁から脱出する方法を見つけるかもしれないことに注意してください。
<!--
LXD can also run `privileged` containers if you so wish, do note that
those aren't root safe and a user with root in such a container will be
able to DoS the host as well as find ways to escape confinement.
-->

コンテナーのセキュリティとカーネル機能についてのより詳細は
[LXC のセキュリティページ](https://linuxcontainers.org/ja/lxc/security/).
を参照してください。
<!--
More details on container security and the kernel features we use can be found on the
[LXC security page](https://linuxcontainers.org/lxc/security/).
-->

## TLS クライアント証明書での認証を使ってリモートを追加する <!-- Adding a remote with TLS client certificate authentication -->
デフォルトのセットアップでは、ユーザーが `lxd remote add` で新しいサーバを
追加する際、サーバに https で通信し、証明書がダウンロードされ、
フィンガープリントがユーザーに表示されます。
<!--
In the default setup, when the user adds a new server with `lxc remote add`,
the server will be contacted over HTTPS, its certificate downloaded and the
fingerprint will be shown to the user.
-->

ユーザーは、これが本当にサーバのフィンガープリントなのかの確認を求められます。
これは接続してみて手動で確認するか、既にそのサーバに接続可能になっている
別のユーザーに info コマンドを実行してもらい、表示されたフィンガープリント
と比較することで確認できます。
<!--
The user will then be asked to confirm that this is indeed the server's
fingerprint which they can manually check by connecting to or asking
someone with access to the server to run the info command and compare
the fingerprints.
-->

その後ユーザーはそのサーバのトラスト・パスワード
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
ユーザーにプロンプトを表示することなく接続は続行します。
<!--
If the server certificate is valid and signed by the CA, then the
connection continues without prompting the user for the certificate.
-->

その後、ユーザーはそのサーバのトラスト・パスワード
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
を実行します。クライアントはユーザーを検証するために認証サーバに要求される
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
 2. PKI のセットアップをしていなければ、フィンガープリントを確認するプロンプトがユーザーに表示されます。 <!-- If we're not in a PKI setup ask the user to confirm the fingerprint. -->
 3. サーバから返された dict を見て、 "auth" が "untrusted" だった場合、ユーザーにサーバのパスワードを入力させ、 `/1.0/certificates` に `POST` のリクエストを送り、その後再び `/1.0` のリクエストを送って本当に信頼されたかを確認します。 <!-- Look at the dict we received back from the server. If "auth" is
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

後はユーザーの責任でサーバ管理者に連絡し、サーバ証明書が本当に変更されたのか
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
信頼エントリーを削除したときに起こります。
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

## ネットワークのセキュリティ <!-- Network security -->

### bridged NIC のセキュリティ <!-- Bridged NIC security -->

LXD のデフォルトのネットワークのモードはそれぞれのインスタンスが接続する「管理された」プライベートなネットワークブリッジを提供するためのものです。
このモードではホスト上に `lxdbr0` と呼ばれるインターフェースが存在し、それぞれのインスタンスに対してブリッジとして振る舞います。
<!--
The default networking mode in LXD is to provide a 'managed' private network bridge that each instance connects to.
In this mode, there is an interface on the host called `lxdbr0` that acts as the bridge for the instances.
-->

ホストはそれぞれの管理されたブリッジに対して `dnsmasq` のインスタンスを稼働します。
それが IP アドレスを割り当て、 DNS の権威サーバーとキャッシュサーバーの両方のサービスを提供します。
<!--
The host runs an instance of `dnsmasq` for each managed bridge, which is responsible for allocating IP addresses
and providing both authoritative and recursive DNS services.
-->

DHCPv4 を使うインスタンスには IPv4 アドレスが割り当てられ、インスタンス名に対する DNS レコードが作成されます。
これによりインスタンスが DHCP リクエスト内に虚偽のホスト名を含めて DNS レコードをスプーフィングできないようにしています。
<!--
Instances using DHCPv4 will be allocated an IPv4 address and a DNS record will be created for their instance name.
This prevents instances from being able to spoof DNS records by providing false hostname info in the DHCP request.
-->

さらに `dnsmasq` サービスは IPv6 のルーター広告の機能も提供します。
これはインスタンスが SLAAC を使って自身の IPv6 アドレスを自動設定することを意味し、 `dnsmasq` による割り当ては行いません。
しかしインスタンスは同等の SLAAC IPv6 アドレスに対して作成された AAAA DNS レコードを DHCPv4 を使って取得することもできます。
これにはインスタンスが IPv6 アドレスを生成する際に IPv6 のプライバシー拡張を使っていないことが前提となります。
<!--
The `dnsmasq` service also provides IPv6 router advertisement capabilities. This means that instances will auto
configure their own IPv6 address using SLAAC, so no allocation is made by `dnsmasq`. However instances that are
also using DHCPv4 will also get an AAAA DNS record created for the equivalent SLAAC IPv6 address.
This assumes that the instances are not using any IPv6 privacy extensions when generating IPv6 addresses.
-->

このデフォルトの設定では DNS の名前はスプーフィングできませんが、インスタンスは Ethernet ブリッジに接続しており、 Layer 2 の希望するトラフィックを送信できますので、信頼できないインスタンスが実質的にはブリッジ上の MAC アドレスあるいは IP アドレスをスプーフィングできることを意味します。
<!--
In this default configuration, whilst DNS names cannot not be spoofed, the instance is connected to an Ethernet
bridge and can transmit any layer 2 traffic that it wishes, which means an untrusted instance can effectively do
MAC or IP spoofing on the bridge.
-->

デフォルトの設定ではブリッジに接続されたインスタンスがブリッジに （場合によっては悪意のある） IPv6 ルーター広告を送ることで LXD ホストの IPv6 ルーティングテーブルを変更することも可能です。
これは `lxdbr0` インターフェースが `/proc/sys/net/ipv6/conf/lxdbr0/accept_ra` を `2` に設定して作られており、それは LXD ホストが `forwarding` を有効にしているときでさえもルーター広告を受け付けることを意味しているからです（より詳細な情報については https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt を参照）。
<!--
It is also possible in the default configuration for instances connected to the bridge to modify the LXD host's
IPv6 routing table by sending (potentially malicious) IPv6 router advertisements to the bridge. This is because
the `lxdbr0` interface is created with `/proc/sys/net/ipv6/conf/lxdbr0/accept_ra` set to `2` meaning that the
LXD host will accept router advertisements even though `forwarding` is enabled (see
https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt for more info).
-->

しかし LXD はいくつかの `bridged` NIC セキュリティ機能を提供しており、インスタンスがネットワークに送信できるトラフィックの種類を制限するのに使用できます。
これらの NIC 設定はインスタンスが使用するプロファイルに設定するべきですが、以下に示すように個々のインスタンスに追加することもできます。
<!--
However LXD offers several `bridged` NIC security features that can be used to control the type of traffic that
an instance is allowed to send onto the network. These NIC settings should be added to the profile that the
instance is using, or can be added to individual instances, as shown below.
-->

`bridged` NIC に対して以下のセキュリティ機能が利用可能です。
<!--
The following security features are available for `bridged` NICs:
-->

Key                      | Type      | Default           | Required  | Description
:--                      | :--       | :--               | :--       | :--
security.mac\_filtering  | boolean   | false             | no        | インスタンスが別のインスタンスの MAC アドレスを詐称するのを防ぐ <!-- Prevent the instance from spoofing another's MAC address -->
security.ipv4\_filtering | boolean   | false             | no        | インスタンスが別のインスタンスの IPv4 アドレスを詐称するのを防ぐ（これを有効にすると mac\_filtering も有効になります）  <!-- Prevent the instance from spoofing another's IPv4 address (enables mac\_filtering) -->
security.ipv6\_filtering | boolean   | false             | no        | インスタンスが別のインスタンスの IPv6 アドレスを詐称するのを防ぐ（これを有効にすると mac\_filtering も有効になります）  <!-- Prevent the instance from spoofing another's IPv6 address (enables mac\_filtering) -->

プロファイルに設定されたデフォルトの `bridged` NIC 設定を以下のコマンドでインスタンスごとにオーバーライドできます。
<!--
One can override the default `bridged` NIC settings from the profile on a per-instance basis using:
-->

```
lxc config device override <instance> <NIC> security.mac_filtering=true
```

これらの機能を合わせて使うとブリッジに接続されたインスタンスが MAC アドレスや IP アドレスを詐称するのを防ぐことができます。
これらは `xtables` (iptables, ip6tables そして ebtables) あるいは `nftables` のいずれかを使って実装されていて、どちらが使われるかはホストでどちらが利用可能かによって決まります。
<!--
Used together these features can prevent an instance connected to a bridge from spoofing MAC and IP addresses.
These are implemented using either `xtables` (iptables, ip6tables and ebtables) or `nftables`, depending on what is
available on the host.
-->

これらのオプションを使うと、ネストシタコンテナーでは異なる MAC アドレス (例えばブリッジされているか macvlan の NICを使うなど) を持つ親のネットワークを実質的に使えなくなることに注意が必要です。
ネストしたコンテナー、少なくとも親と同じネットワーク上のネストしたコンテナーを実質的に使えなくなることに注意が必要です。
<!--
It's worth noting that those options effectively prevent nested containers from using the parent network with a
different MAC address (i.e using bridged or macvlan NICs).
-->

IP フィルタリング機能は詐称されたソースアドレスを含む全てのパケットをブロックするだけでなく、詐称された IP を含む ARP と NDP 広告もブロックします。
<!--
The IP filtering features block ARP and NDP advertisements that contain a spoofed IP, as well as blocking any
packets that contain a spoofed source address.
-->

`security.ipv4\_filtering` か `security.ipv6\_filtering` が有効で （ `ipvX.address=none`  であるかブリッジで DHCP サービスが有効になっていないために）インスタンスに IP アドレスが割り当てられない場合、そのインスタンスからの（訳注： `security.ipv4\_filtering` なら IPv4 、 `security.ipv6\_filtering` なら IPv6 と）設定に対応するプロトコルの全ての IP トラフィックはブロックされます。
<!--
If `security.ipv4\_filtering` or `security.ipv6\_filtering` is enabled and the instance cannot be allocated an IP
address (because `ipvX.address=none` or there is no DHCP service enabled on the bridge) then all IP traffic for
that protocol is blocked from the instance.
-->

`security.ipv6\_filtering` が有効な場合、インスタンスからの IPv6 ルーター広告はブロックされます。
<!--
When `security.ipv6\_filtering` is enabled IPv6 router advertisements are blocked from the instance.
-->

`security.ipv4\_filtering` か `security.ipv6\_filtering` が有効な場合は ARP、IPv4、IPv6 以外の全ての Ethernet フレームがドロップされます。
これはスタックされた VLAN QinQ (802.1ad) のフレームが IP フィルタリングをバイパスするのを防ぎます。
<!--
When `security.ipv4\_filtering` or `security.ipv6\_filtering` is enabled, any Ethernet frames that are not ARP,
IPv4 or IPv6 are dropped. This prevents stacked VLAN QinQ (802.1ad) frames from bypassing the IP filtering.
-->

### routed NIC のセキュリティ <!-- Routed NIC security -->

`routed` という別のネットワークモードが使用でき、これはコンテナーとホストの間に veth のペアを提供します。
このネットワークモードでは LXD ホストはルーターとして機能し、ホストに静的ルートが追加され、コンテナの IP アドレスへのトラフィックはコンテナーの veth インタフェースに向けられます。
<!--
An alternative networking mode is available called `routed` that provides a veth pair between container and host.
In this networking mode the LXD host functions as a router and static routes are added to the host directing
traffic for the container's IPs towards the container's veth interface.
-->

デフォルトではホスト上に作成される veth インタフェースは `accept_ra` 設定が無効になっています。
これは LXD ホストの IPv6 ルーティングテーブルをコンテナーからのルーター広告で変更されないようにするためです。
さらにホスト上の `rp_filter` は `1` に設定されます。
これはコンテナーが持っているとホストが知らない IP アドレスに対してソースアドレスのスプーフィングを防ぐためです。
<!--
By default the veth interface created on the host has its `accept_ra` setting disabled to prevent router
advertisements from the container modifying the IPv6 routing table on the LXD host. In addition to that the
`rp_filter` on the host is set to `1` to prevent source address spoofing for IPs that the host does not know the
container has.
-->
