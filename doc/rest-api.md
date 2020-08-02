# REST API
## イントロダクション <!-- Introduction -->
LXD とクライアントの間の全ての通信は HTTP 上の RESTful API を使って
行います。リモートの操作は SSL で暗号化して通信し、ローカルの操作は
Unix ソケットを使って通信します。
<!--
All the communications between LXD and its clients happen using a
RESTful API over http which is then encapsulated over either SSL for
remote operations or a unix socket for local operations.
-->

全ての REST インターフェースが認証を必要とするわけではありません。
<!--
Not all of the REST interface requires authentication:
-->

 * `/` への `GET` は認証なしで誰でも実行可能です (API エンドポイント一覧を返します) <!-- `GET` to `/` is allowed for everyone (lists the API endpoints) -->
 * `/1.0` への GET は認証なしで誰でも実行可能です (ですが結果は認証ありの場合と異なります) <!-- `GET` to `/1.0` is allowed for everyone (but result varies) -->
 * `/1.0/certificates` への `POST` はクライアント証明書があれば誰でも実行可能です <!-- `POST` to `/1.0/certificates` is allowed for everyone with a client certificate -->
 * `/1.0/images/*` への `GET` は認証なしで誰でも実行可能ですが、その場合認証なしのユーザーに対して公開されているイメージだけを返します。 <!-- `GET` to `/1.0/images/*` is allowed for everyone but only returns public images for unauthenticated users -->

以下では認証なしで利用できるエンドポイントはそのように明記します。
<!--
Unauthenticated endpoints are clearly identified as such below.
-->

## API のバージョニング <!-- API versioning -->
サポートされている API のメジャーバージョンのリストは `GET /` を使って
取得できます。
<!--
The list of supported major API versions can be retrieved using `GET /`.
-->

後方互換性を壊す場合は API のメジャーバージョンが上がります。
<!--
The reason for a major API bump is if the API breaks backward compatibility.
-->

後方互換性を壊さずに追加される機能は `api_extensions` の追加という形になり、
特定の機能がサーバでサポートされているかクライアントがチェックすることで
利用できます。
<!--
Feature additions done without breaking backward compatibility only
result in addition to `api_extensions` which can be used by the client
to check if a given feature is supported by the server.
-->

## 戻り値 <!-- Return values -->
次の 3 つの標準的な戻り値の型があります。
<!--
There are three standard return types:
-->

 * 標準の戻り値 <!-- Standard return value -->
 * バックグラウンド操作 <!-- Background operation -->
 * エラー <!-- Error -->

#### 標準の戻り値 <!-- Standard return value -->
標準の同期的な操作に対しては以下のような dict が返されます。
<!--
For a standard synchronous operation, the following dict is returned:
-->

```js
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "metadata": {}                          // リソースやアクションに固有な追加のメタデータ
}
```

<!--
```js
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "metadata": {}                          // Extra resource/action specific metadata
}
```
-->

HTTP ステータスコードは必ず 200 です。
<!--
HTTP code must be 200.
-->

#### バックグラウンド操作 <!-- Background operation -->
リクエストの結果がバックグラウンド操作になる場合、 HTTP ステータスコードは 202 (Accepted)
になり、操作の URL を指す HTTP の Location ヘッダが返されます。
<!--
When a request results in a background operation, the HTTP code is set to 202 (Accepted)
and the Location HTTP header is set to the operation URL.
-->

レスポンスボディは以下のような構造を持つ dict です。
<!--
The body is a dict with the following structure:
-->

```js
{
    "type": "async",
    "status": "OK",
    "status_code": 100,
    "operation": "/1.0/instances/<id>",                     // バックグラウンド操作の URL
    "metadata": {}                                          // 操作のメタデータ (下記参照)
}
```

<!--
```js
{
    "type": "async",
    "status": "OK",
    "status_code": 100,
    "operation": "/1.0/instances/<id>",                     // URL to the background operation
    "metadata": {}                                          // Operation metadata (see below)
}
```
-->

操作のメタデータの構造は以下のようになります。
<!--
The operation metadata structure looks like:
-->

```js
{
    "id": "a40f5541-5e98-454f-b3b6-8a51ef5dbd3c",           // 操作の UUID
    "class": "websocket",                                   // 操作の種別 (task, websocket, token のいずれか)
    "created_at": "2015-11-17T22:32:02.226176091-05:00",    // 操作の作成日時
    "updated_at": "2015-11-17T22:32:02.226176091-05:00",    // 操作の最終更新日時
    "status": "Running",                                    // 文字列表記での操作の状態
    "status_code": 103,                                     // 整数表記での操作の状態 (status ではなくこちらを利用してください。訳注: 詳しくは下記のステータスコードの項を参照)
    "resources": {                                          // リソース種別 (container, snapshots, images のいずれか) の dict を影響を受けるリソース
      "containers": [
        "/1.0/instances/test"
      ]
    },
    "metadata": {                                           // 対象となっている (この例では exec) 操作に固有なメタデータ
      "fds": {
        "0": "2a4a97af81529f6608dca31f03a7b7e47acc0b8dc6514496eb25e325f9e4fa6a",
        "control": "5b64c661ef313b423b5317ba9cb6410e40b705806c28255f601c0ef603f079a7"
      }
    },
    "may_cancel": false,                                    // (REST で DELETE を使用して) 操作がキャンセル可能かどうか
    "err": ""                                               // 操作が失敗した場合にエラー文字列が設定されます
}
```

<!--
```js
{
    "id": "a40f5541-5e98-454f-b3b6-8a51ef5dbd3c",           // UUID of the operation
    "class": "websocket",                                   // Class of the operation (task, websocket or token)
    "created_at": "2015-11-17T22:32:02.226176091-05:00",    // When the operation was created
    "updated_at": "2015-11-17T22:32:02.226176091-05:00",    // Last time the operation was updated
    "status": "Running",                                    // String version of the operation's status
    "status_code": 103,                                     // Integer version of the operation's status (use this rather than status)
    "resources": {                                          // Dictionary of resource types (container, snapshots, images) and affected resources
      "containers": [
        "/1.0/instances/test"
      ]
    },
    "metadata": {                                           // Metadata specific to the operation in question (in this case, exec)
      "fds": {
        "0": "2a4a97af81529f6608dca31f03a7b7e47acc0b8dc6514496eb25e325f9e4fa6a",
        "control": "5b64c661ef313b423b5317ba9cb6410e40b705806c28255f601c0ef603f079a7"
      }
    },
    "may_cancel": false,                                    // Whether the operation can be canceled (DELETE over REST)
    "err": ""                                               // The error string should the operation have failed
}
```
-->

対象の操作に対して追加のリクエストを送って情報を取り出さなくても、
何が起こっているかユーザーにとってわかりやすい形でボディは構成されています。
ボディに含まれる全ての情報はバックグラウンド操作の URL から取得する
こともできます。
<!--
The body is mostly provided as a user friendly way of seeing what's
going on without having to pull the target operation, all information in
the body can also be retrieved from the background operation URL.
-->

#### エラー <!-- Error -->
さまざまな状況によっては操作を行う前に直ぐに問題が起きる場合があり、
そういう場合には以下のような値が返されます。
<!--
There are various situations in which something may immediately go
wrong, in those cases, the following return value is used:
-->

```js
{
    "type": "error",
    "error": "Failure",
    "error_code": 400,
    "metadata": {}                      // エラーについてのさらなる詳細
}
```

<!--
```js
{
    "type": "error",
    "error": "Failure",
    "error_code": 400,
    "metadata": {}                      // More details about the error
}
```
-->

HTTP ステータスコードは 400, 401, 403, 404, 409, 412, 500 のいずれかです。
<!--
HTTP code must be one of of 400, 401, 403, 404, 409, 412 or 500.
-->

## ステータスコード <!-- Status codes -->
LXD REST API はステータス情報を返す必要があります。それはエラーの理由だったり、
操作の現在の状態だったり、 LXD が提供する様々なリソースの状態だったりします。
<!--
The LXD REST API often has to return status information, be that the
reason for an error, the current state of an operation or the state of
the various resources it exports.
-->

デバッグをシンプルにするため、ステータスは常に文字列表記と整数表記で
重複して返されます。ステータスの整数表記の値は将来に渡って不変なので
API クライアントが個々の値に依存できます。文字列表記のステータスは
人間が API を手動で実行したときに何が起きているかをより簡単に判断
できるように用意されています。
<!--
To make it simple to debug, all of those are always doubled. There is a
numeric representation of the state which is guaranteed never to change
and can be relied on by API clients. Then there is a text version meant
to make it easier for people manually using the API to figure out what's
happening.
-->

ほとんどのケースでこれらは `status` と `status_code` と呼ばれ、前者は
ユーザーフレンドリーな文字列表記で後者は固定の数値です。
<!--
In most cases, those will be called status and `status_code`, the former
being the user-friendly string representation and the latter the fixed
numeric value.
-->

整数表記のコードは常に 3 桁の数字で以下の範囲の値となっています。
<!--
The codes are always 3 digits, with the following ranges:
-->

 * 100 to 199: リソースの状態 (started, stopped, ready, ...) <!-- resource state (started, stopped, ready, ...) -->
 * 200 to 399: 成功したアクションの結果 <!-- positive action result -->
 * 400 to 599: 失敗したアクションの結果 <!-- negative action result -->
 * 600 to 999: 将来使用するために予約されている番号の範囲 <!-- future use -->

### 現在使用されているステータスコード一覧 <!-- List of current status codes -->

コード <!-- Code -->  | 意味 <!-- Meaning -->
:---  | :------
100   | 操作が作成された <!-- Operation created -->
101   | 開始された <!-- Started -->
102   | 停止された <!-- Stopped -->
103   | 実行中 <!-- Running -->
104   | キャンセル中 <!-- Cancelling -->
105   | ペンディング <!-- Pending -->
106   | 開始中 <!-- Starting -->
107   | 停止中 <!-- Stopping -->
108   | 中断中 <!-- Aborting -->
109   | 凍結中 <!-- Freezing -->
110   | 凍結された <!-- Frozen -->
111   | 解凍された <!-- Thawed -->
200   | 成功 <!-- Success -->
400   | 失敗 <!-- Failure -->
401   | キャンセルされた <!-- Cancelled -->

## 再帰 <!-- Recursion -->
巨大な一覧のクエリを最適化するために、コレクションに対して再帰が実装されています。
コレクションに対するクエリの GET リクエストに `recursion` パラメータを指定できます。
<!--
To optimize queries of large lists, recursion is implemented for collections.
A `recursion` argument can be passed to a GET query against a collection.
-->

デフォルト値は 0 でコレクションのメンバーの URL が返されることを意味します。
1 を指定するとこれらの URL がそれが指すオブジェクト (通常は dict 形式) で
置き換えられます。
<!--
The default value is 0 which means that collection member URLs are
returned. Setting it to 1 will have those URLs be replaced by the object
they point to (typically a dict).
-->

再帰はジョブへのポインタ (URL) をオブジェクトそのもので単に置き換えるように
実装されています。
<!--
Recursion is implemented by simply replacing any pointer to an job (URL)
by the object itself.
-->

## フィルタ <!-- Filtering -->
検索結果をある値でフィルタするために、コレクションにフィルタが実装されています。
コレクションに対する GET クエリに `filter` 引数を渡せます。
<!--
To filter your results on certain values, filter is implemented for collections.
A `filter` argument can be passed to a GET query against a collection.
-->

フィルタはインスタンスとイメージのエンドポイントに提供されています。
<!--
Filtering is available for the instance and image endpoints.
-->

フィルタにはデフォルト値はありません。これは見つかった全ての結果が返されることを意味します。
フィルタの引数には以下のような言語を設定します。
<!--
There is no default value for filter which means that all results found will
be returned. The following is the language used for the filter argument:
-->

?filter=field_name eq desired_field_assignment

この言語は REST API のフィルタロジックを構成するための OData の慣習に従います。
フィルタは下記の論理演算子もサポートします。
not(not), equals(eq), not equals(ne), and(and), or(or)
フィルタは左結合で評価されます。
空白を含む値はクォートで囲むことができます。
ネストしたフィルタもサポートされます。
例えば config 内のフィールドに対してフィルタするには以下のように指定します。
<!--
The language follows the OData conventions for structuring REST API filtering
logic. Logical operators are also supported for filtering: not(not), equals(eq),
not equals(ne), and(and), or(or). Filters are evaluated with left associativity.
Values with spaces can be surrounded with quotes. Nesting filtering is also supported. 
For instance, to filter on a field in a config you would pass:
-->

?filter=config.field_name eq desired_field_assignment

device の属性についてフィルタするには以下のように指定します。
<!--
For filtering on device attributes you would pass:
-->

?filter=devices.device_name.field_name eq desired_field_assignment

以下に上記の異なるフィルタの方法を含む GET クエリをいくつか示します。
<!--
Here are a few GET query examples of the different filtering methods mentioned above:
-->

containers?filter=name eq "my container" and status eq Running

containers?filter=config.image.os eq ubuntu or devices.eth0.nictype eq bridged

images?filter=Properties.os eq Centos and not UpdateSource.Protocol eq simplestreams

## 非同期操作 <!-- Async operations -->
完了までに 1 秒以上かかるかもしれない操作はバックグラウンドで実行しなければ
なりません。そしてクライアントにはバックグラウンド操作 ID を返します。
<!--
Any operation which may take more than a second to be done must be done
in the background, returning a background operation ID to the client.
-->

クライアントは操作のステータス更新をポーリングするか long-poll API を使って
通知を待つことが出来ます。
<!--
The client will then be able to either poll for a status update or wait
for a notification using the long-poll API.
-->

## 通知 <!-- Notifications -->
通知のために Websocket ベースの API が利用できます。クライアントへ送られる
トラフィックを制限するためにいくつかの異なる通知種別が存在します。
<!--
A websocket based API is available for notifications, different notification
types exist to limit the traffic going to the client.
-->

リモート操作の状態をポーリングしなくて済むように、リモート操作を開始する
前に操作の通知をクライアントが常に購読しておくのがお勧めです。
<!--
It's recommended that the client always subscribes to the operations
notification type before triggering remote operations so that it doesn't
have to then poll for their status.
-->

## PUT と PATCH の使い分け <!-- PUT vs PATCH -->
LXD API は既存のオブジェクトを変更するのに PUT と PATCH の両方をサポートします。
<!--
The LXD API supports both PUT and PATCH to modify existing objects.
-->

PUT はオブジェクト全体を新しい定義で置き換えます。典型的には GET で現在の
オブジェクトの状態を取得した後に PUT が呼ばれます。
<!--
PUT replaces the entire object with a new definition, it's typically
called after the current object state was retrieved through GET.
-->

レースコンディションを避けるため、 GET のレスポンスから ETag ヘッダを読み取り
PUT リクエストの If-Match ヘッダに設定するべきです。こうしておけば GET と
PUT の間にオブジェクトが他から変更されていた場合は更新が失敗するようになります。
<!--
To avoid race conditions, the Etag header should be read from the GET
response and sent as If-Match for the PUT request. This will cause LXD
to fail the request if the object was modified between GET and PUT.
-->

PATCH は変更したいプロパティだけを指定することでオブジェクト内の単一の
フィールドを変更するのに用いられます。キーを削除するには通常は空の値を
設定すれば良いようになっていますが、 PATCH ではキーの削除は出来ず、代わりに
PUT を使う必要がある場合もあります。
<!--
PATCH can be used to modify a single field inside an object by only
specifying the property that you want to change. To unset a key, setting
it to empty will usually do the trick, but there are cases where PATCH
won't work and PUT needs to be used instead.
-->

## インスタンス、コンテナーと仮想マシン <!-- instances, containers and virtual-machines -->
このドキュメントでは `/1.0/instances/...` のようなパスを常に示します。
これらはかなり新しく、仮想マシンがサポートされた LXD 3.19 で導入されました。
<!--
This documentation will always show paths such as `/1.0/instances/...`.
Those are fairly new, introduced with LXD 3.19 when virtual-machine support.
-->

コンテナーのみをサポートする古いリリースでは全く同じ API を `/1.0/containers/...` で利用します。
<!--
Older releases that only supported containers will instead use the exact same API at `/1.0/containers/...`.
-->

後方互換性の理由で LXD は `/1.0/containers` API を引き続き公開しサポートしますが、簡潔さのため以下では両方をドキュメントはしないことにしました。
<!--
For backward compatibility reasons, LXD does still expose and support
that `/1.0/containers` API, though for the sake of brevity, we decided
not to double-document everything below.
-->

`/1.0/virtual-machines` に追加のエンドポイントも存在し、 `/1.0/containers` とほぼ同様ですが、仮想マシンのタイプのインスタンスのみを表示します。
<!--
An additional endpoint at `/1.0/virtual-machines` is also present and
much like `/1.0/containers` will only show you instances of that type.
-->

## API 構造 <!-- API structure -->
 * [`/`](#)
   * [`/1.0`](#10)
 * [`/1.0/certificates`](#10certificates)
   * [`/1.0/certificates/<fingerprint>`](#10certificatesfingerprint)
 * [`/1.0/instances`](#10instances)
   * [`/1.0/instances/<name>`](#10instancesname)
     * [`/1.0/instances/<name>/console`](#10instancesnameconsole)
     * [`/1.0/instances/<name>/exec`](#10instancesnameexec)
     * [`/1.0/instances/<name>/files`](#10instancesnamefiles)
     * [`/1.0/instances/<name>/snapshots`](#10instancesnamesnapshots)
     * [`/1.0/instances/<name>/snapshots/<name>`](#10instancesnamesnapshotsname)
     * [`/1.0/instances/<name>/state`](#10instancesnamestate)
     * [`/1.0/instances/<name>/logs`](#10instancesnamelogs)
     * [`/1.0/instances/<name>/logs/<logfile>`](#10instancesnamelogslogfile)
     * [`/1.0/instances/<name>/metadata`](#10instancesnamemetadata)
     * [`/1.0/instances/<name>/metadata/templates`](#10instancesnamemetadatatemplates)
     * [`/1.0/instances/<name>/backups`](#10instancesnamebackups)
     * [`/1.0/instances/<name>/backups/<name>`](#10instancesnamebackupsname)
     * [`/1.0/instances/<name>/backups/<name>/export`](#10instancesnamebackupsnameexport)
 * [`/1.0/events`](#10events)
 * [`/1.0/images`](#10images)
   * [`/1.0/images/<fingerprint>`](#10imagesfingerprint)
     * [`/1.0/images/<fingerprint>/export`](#10imagesfingerprintexport)
     * [`/1.0/images/<fingerprint>/refresh`](#10imagesfingerprintrefresh)
     * [`/1.0/images/<fingerprint>/secret`](#10imagesfingerprintsecret)
   * [`/1.0/images/aliases`](#10imagesaliases)
     * [`/1.0/images/aliases/<name>`](#10imagesaliasesname)
 * [`/1.0/networks`](#10networks)
   * [`/1.0/networks/<name>`](#10networksname)
   * [`/1.0/networks/<name>/state`](#10networksnamestate)
 * [`/1.0/operations`](#10operations)
   * [`/1.0/operations/<uuid>`](#10operationsuuid)
     * [`/1.0/operations/<uuid>/wait`](#10operationsuuidwait)
     * [`/1.0/operations/<uuid>/websocket`](#10operationsuuidwebsocket)
 * [`/1.0/profiles`](#10profiles)
   * [`/1.0/profiles/<name>`](#10profilesname)
 * [`/1.0/projects`](#10projects)
   * [`/1.0/projects/<name>`](#10projectsname)
 * [`/1.0/storage-pools`](#10storage-pools)
   * [`/1.0/storage-pools/<name>`](#10storage-poolsname)
     * [`/1.0/storage-pools/<name>/resources`](#10storage-poolsnameresources)
     * [`/1.0/storage-pools/<name>/volumes`](#10storage-poolsnamevolumes)
       * [`/1.0/storage-pools/<name>/volumes/<type>`](#10storage-poolsnamevolumestype)
         * [`/1.0/storage-pools/<pool>/volumes/<type>/<name>`](#10storage-poolspoolvolumestypename)
           * [`/1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots`](#10storage-poolspoolvolumestypenamesnapshots)
             * [`/1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`](#10storage-poolspoolvolumestypevolumesnapshotsname)
 * [`/1.0/resources`](#10resources)
 * [`/1.0/cluster`](#10cluster)
   * [`/1.0/cluster/members`](#10clustermembers)
     * [`/1.0/cluster/members/<name>`](#10clustermembersname)

## API 詳細 <!-- API details -->
### `/`
#### GET
 * 説明: サポートされている API の一覧 <!-- Description: List of supported APIs -->
 * 認証: guest <!-- Authentication: guest -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: サポートされている API エンドポイントの URL の一覧 <!-- Return: list of supported API endpoint URLs -->

戻り値 <!-- Return value: -->

```json
[
    "/1.0"
]
```

### `/1.0/`
#### GET
 * 説明: サーバの設定と環境情報 <!-- Description: Server configuration and environment information -->
 * 認証: guest, untrusted, trusted のいずれか <!-- Authentication: guest, untrusted or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: サーバの状態を表す dict <!-- Return: Dict representing server state -->

戻り値 (trusted の場合) <!-- Return value (if trusted): -->

```js
{
    "api_extensions": [],                           // stable とマークされた API 以降に追加された API 拡張の一覧
    "api_status": "stable",                         // API の実装状態 (development, stable, deprecated のいずれか)
    "api_version": "1.0",                           // 文字列表記での API バージョン
    "auth": "trusted",                              // 認証状態 ("guest", "untrusted", "trusted" のいずれか)
    "config": {                                     // ホストの設定
        "core.trust_password": true,
        "core.https_address": "[::]:8443"
    },
    "environment": {                                // ホストの様々な情報 (OS, カーネル, ...)
        "addresses": [
            "1.2.3.4:8443",
            "[1234::1234]:8443"
        ],
        "architectures": [
            "x86_64",
            "i686"
        ],
        "certificate": "PEM certificate",
        "driver": "lxc",
        "driver_version": "1.0.6",
        "kernel": "Linux",
        "kernel_architecture": "x86_64",
        "kernel_version": "3.16",
        "server": "lxd",
        "server_pid": 10224,
        "server_version": "0.8.1"}
        "storage": "btrfs",
        "storage_version": "3.19",
    },
    "public": false,                                // クライアントにとってサーバを公開された (読み取り専用の) リモートとして扱うべきかどうか
}
```

<!--
```js
{
    "api_extensions": [],                           // List of API extensions added after the API was marked stable
    "api_status": "stable",                         // API implementation status (one of, development, stable or deprecated)
    "api_version": "1.0",                           // The API version as a string
    "auth": "trusted",                              // Authentication state, one of "guest", "untrusted" or "trusted"
    "config": {                                     // Host configuration
        "core.trust_password": true,
        "core.https_address": "[::]:8443"
    },
    "environment": {                                // Various information about the host (OS, kernel, ...)
        "addresses": [
            "1.2.3.4:8443",
            "[1234::1234]:8443"
        ],
        "architectures": [
            "x86_64",
            "i686"
        ],
        "certificate": "PEM certificate",
        "driver": "lxc",
        "driver_version": "1.0.6",
        "kernel": "Linux",
        "kernel_architecture": "x86_64",
        "kernel_version": "3.16",
        "server": "lxd",
        "server_pid": 10224,
        "server_version": "0.8.1"}
        "storage": "btrfs",
        "storage_version": "3.19",
    },
    "public": false,                                // Whether the server should be treated as a public (read-only) remote by the client
}
```
-->

戻り値 (guest または untrusted の場合) <!-- Return value (if guest or untrusted): -->

```js
{
    "api_extensions": [],                   // stable とマークされた API 以降に追加された API 拡張の一覧
    "api_status": "stable",                 // API の実装状態 (development, stable, deprecated のいずれか)
    "api_version": "1.0",                   // 文字列表記での API バージョン
    "auth": "guest",                        // 認証状態 ("guest", "untrusted", "trusted" のいずれか)
    "public": false,                        // クライアントにとってサーバを公開された (読み取り専用の) リモートとして扱うべきかどうか
}
```

<!--
```js
{
    "api_extensions": [],                   // List of API extensions added after the API was marked stable
    "api_status": "stable",                 // API implementation status (one of, development, stable or deprecated)
    "api_version": "1.0",                   // The API version as a string
    "auth": "guest",                        // Authentication state, one of "guest", "untrusted" or "trusted"
    "public": false,                        // Whether the server should be treated as a public (read-only) remote by the client
}
```
-->

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: サーバ設定や他の設定を置き換えます <!-- Description: Replaces the server configuration or other properties -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (既存の全ての設定を指定された設定で置き換えます)
<!--
Input (replaces any existing config with the provided one):
-->

```json
{
    "config": {
        "core.trust_password": "my-new-password",
        "core.https_address": "1.2.3.4:8443"
    }
}
```

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: サーバ設定や他の設定を更新します <!-- Description: Updates the server configuration or other properties -->
 * 導入: `patch` API 拡張により <!-- Introduced: with API extension `patch` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (指定されたキーだけを更新し、残りの既存の設定はそのまま残ります)
<!--
Input (updates only the listed keys, rest remains intact):
-->

```json
{
    "config": {
        "core.trust_password": "my-new-password"
    }
}
```

### `/1.0/certificates`
#### GET
 * 説明: 信頼された証明書の一覧を返します <!-- Description: list of trusted certificates -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 信頼された証明書の URL の一覧 <!-- Return: list of URLs for trusted certificates -->

戻り値
<!-- Return: -->

```json
[
    "/1.0/certificates/3ee64be3c3c7d617a7470e14f2d847081ad467c8c26e1caad841c8f67f7c7b09"
]
```

#### POST
 * 説明: 信頼された証明書を追加します <!-- Description: add a new trusted certificate -->
 * 認証: trusted または untrusted <!-- Authentication: trusted or untrusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```js
{
    "type": "client",                       // 証明書の種別 (keyring)、現在は client のみ
    "certificate": "PEM certificate",       // 提供される場合は有効な x509 形式の証明書。提供されない場合は接続のクライアント証明書が使用される
    "name": "foo",                          // 証明書の名前を指定可能。指定しない場合はリクエストの TLS ヘッダーのホスト名が使用される。
    "password": "server-trust-password"     // そのサーバのトラスト・パスワード (untrusted の場合にのみ必須)
}
```

<!--
```js
{
    "type": "client",                       // Certificate type (keyring), currently only client
    "certificate": "PEM certificate",       // If provided, a valid x509 certificate. If not, the client certificate of the connection will be used
    "name": "foo",                          // An optional name for the certificate. If nothing is provided, the host in the TLS header for the request is used.
    "password": "server-trust-password"     // The trust password for that server (only required if untrusted)
}
```
-->

### `/1.0/certificates/<fingerprint>`
#### GET
 * 説明: 信頼された証明書の情報 <!-- Description: trusted certificate information -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 信頼された証明書を表す dict <!-- Return: dict representing a trusted certificate -->

出力
<!--
Output:
-->

```json
{
    "type": "client",
    "certificate": "PEM certificate",
    "name": "foo",
    "fingerprint": "SHA256 Hash of the raw certificate"
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: 証明書のプロパティを置き換えます <!-- Description: Replaces the certificate properties -->
 * 導入: `certificate_update` API 拡張により <!-- Introduced: with API extension `certificate_update` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "type": "client",
    "name": "bar"
}
```

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: 証明書のプロパティを更新します <!-- Description: Updates the certificate properties -->
 * 導入: `certificate_update` API 拡張により <!-- Introduced: with API extension `certificate_update` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "baz"
}
```


#### DELETE
 * 説明: 信頼された証明書を削除します <!-- Description: Remove a trusted certificate -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input:
-->

```json
{
}
```

レスポンスの HTTP ステータスコードは 202 (Accepted)。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/instances`
#### GET
 * 説明: インスタンスの一覧 <!-- Description: List of instances -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: このサーバがホストしているインスタンスの URL の一覧 <!-- Return: list of URLs for instances this server hosts -->

戻り値
<!--
Return value:
-->

```json
[
    "/1.0/instances/blah",
    "/1.0/instances/blah1"
]
```

#### POST (`?target=<member>` を任意で指定可能) <!-- POST (optional `?target=<member>`) -->
 * 説明: 新しいインスタンスを作成します <!-- Description: Create a new instance -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 ("ubuntu/devel" というエイリアスを持つローカルイメージをベースとするインスタンス)
<!--
Input (instance based on a local image with the "ubuntu/devel" alias):
-->

```js
{
    "name": "my-new-instance",                                          // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                            // プロファイルの一覧
    "ephemeral": true,                                                  // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                      // 設定のオーバーライド
    "type": "container",                                                // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                        // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "instance_type": "c2.micro",                                        // リミットのベースとして使用するための任意で指定可能なインスタンスタイプ
    "source": {"type": "image",                                         // "image", "migration", "copy", "none" のいずれかを指定可能
               "alias": "ubuntu/devel"},                                // エイリアスの名前
}
```

<!--
```js
{
    "name": "my-new-instance",                                          // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                            // List of profiles
    "ephemeral": true,                                                  // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                      // Config override.
    "type": "container",                                                // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                        // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "instance_type": "c2.micro",                                        // An optional instance type to use as basis for limits
    "source": {"type": "image",                                         // Can be: "image", "migration", "copy" or "none"
               "alias": "ubuntu/devel"},                                // Name of the alias
}
```
-->

入力 (フィンガープリントで識別されるローカルのイメージをベースとするインスタンス)
<!--
Input (instance based on a local image identified by its fingerprint):
-->

```js
{
    "name": "my-new-instance",                                          // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                            // プロファイルの一覧
    "ephemeral": true,                                                  // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                      // 設定のオーバーライド
    "type": "container",                                                // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                        // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // "image", "migration", "copy", "none" のいずれかを指定可能
               "fingerprint": "SHA-256"},                               // フィンガープリント
}
```

<!--
```js
{
    "name": "my-new-instance",                                          // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                            // List of profiles
    "ephemeral": true,                                                  // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                      // Config override.
    "type": "container",                                                // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                        // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // Can be: "image", "migration", "copy" or "none"
               "fingerprint": "SHA-256"},                               // Fingerprint
}
```
-->

入力 (指定したイメージのプロパティに対して最も最近マッチしたイメージをベースとするインスタンス)
<!--
Input (instance based on most recent match based on image properties):
-->

```js
{
    "name": "my-new-instance",                                          // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                            // プロファイルの一覧
    "ephemeral": true,                                                  // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                      // 設定のオーバーライド
    "type": "container",                                                // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                        // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // "image", "migration", "copy", "none" のいずれかを指定可能
               "properties": {                                          // プロパティ
                    "os": "ubuntu",
                    "release": "18.04",
                    "architecture": "x86_64"
                }},
}
```

<!--
```js
{
    "name": "my-new-instance",                                          // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                            // List of profiles
    "ephemeral": true,                                                  // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                      // Config override.
    "type": "container",                                                // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                        // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // Can be: "image", "migration", "copy" or "none"
               "properties": {                                          // Properties
                    "os": "ubuntu",
                    "release": "18.04",
                    "architecture": "x86_64"
                }},
}
```
-->

入力 (事前に作成済みの rootfs を除いたインスタンス、既存のインスタンスにアタッチする際に有用)
<!--
Input (instance without a pre-populated rootfs, useful when attaching to an existing one):
-->

```js
{
    "name": "my-new-instance",                                          // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                            // プロファイルの一覧
    "ephemeral": true,                                                  // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                      // 設定のオーバーライド
    "type": "container",                                                // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                        // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "none"},                                         // "image", "migration", "copy", "none" のいずれかを指定可能
}
```

<!--
```js
{
    "name": "my-new-instance",                                          // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                            // List of profiles
    "ephemeral": true,                                                  // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                      // Config override.
    "type": "container",                                                // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                        // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "none"},                                         // Can be: "image", "migration", "copy" or "none"
}
```
-->

入力 (公開されたリモートのイメージを使用)
<!--
Input (using a public remote image):
-->

```js
{
    "name": "my-new-instance",                                          // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                            // プロファイルの一覧
    "ephemeral": true,                                                  // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                      // 設定のオーバーライド
    "type": "container",                                                // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                        // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // "image", "migration", "copy", "none" のいずれかを指定可能
               "mode": "pull",                                          // "local" (デフォルト) か "pull" のいずれか
               "server": "https://10.0.2.3:8443",                       // リモートサーバ (pull モードのときのみ)
               "protocol": "lxd",                                       // プロトコル (lxd か simplestreams のいずれか、デフォルトは lxd)
               "certificate": "PEM certificate",                        // PEM 証明書を指定可能。未指定の場合はシステムの CA が使用される。
               "alias": "ubuntu/devel"},                                // エイリアスの名前
}
```

<!--
```js
{
    "name": "my-new-instance",                                          // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                            // List of profiles
    "ephemeral": true,                                                  // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                      // Config override.
    "type": "container",                                                // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                        // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // Can be: "image", "migration", "copy" or "none"
               "mode": "pull",                                          // One of "local" (default) or "pull"
               "server": "https://10.0.2.3:8443",                       // Remote server (pull mode only)
               "protocol": "lxd",                                       // Protocol (one of lxd or simplestreams, defaults to lxd)
               "certificate": "PEM certificate",                        // Optional PEM certificate. If not mentioned, system CA is used.
               "alias": "ubuntu/devel"},                                // Name of the alias
}
```
-->

入力 (プライベートなリモートのイメージをそのイメージのシークレットを取得した後に使用)
<!--
Input (using a private remote image after having obtained a secret for that image):
-->

```js
{
    "name": "my-new-instance",                                          // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                            // プロファイルの一覧
    "ephemeral": true,                                                  // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                      // 設定のオーバーライド
    "type": "container",                                                // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                        // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // "image", "migration", "copy", "none" のいずれかを指定可能
               "mode": "pull",                                          // "local" (デフォルト) か "pull" のいずれか
               "server": "https://10.0.2.3:8443",                       // リモートサーバ (pull モードのときのみ)
               "secret": "my-secret-string",                            // イメージを取得するために使用するシークレット (pull モードのときのみ)
               "certificate": "PEM certificate",                        // PEM 証明書を指定可能。未指定の場合はシステムの CA が使用される。
               "alias": "ubuntu/devel"},                                // エイリアスの名前
}
```

<!--
```js
{
    "name": "my-new-instance",                                          // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                            // List of profiles
    "ephemeral": true,                                                  // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                      // Config override.
    "type": "container",                                                // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                        // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "image",                                         // Can be: "image", "migration", "copy" or "none"
               "mode": "pull",                                          // One of "local" (default) or "pull"
               "server": "https://10.0.2.3:8443",                       // Remote server (pull mode only)
               "secret": "my-secret-string",                            // Secret to use to retrieve the image (pull mode only)
               "certificate": "PEM certificate",                        // Optional PEM certificate. If not mentioned, system CA is used.
               "alias": "ubuntu/devel"},                                // Name of the alias
}
```
-->

入力 (マイグレーション・ウェブソケットで送られるリモートのインスタンスを使用)
<!--
Input (using a remote instance, sent over the migration websocket):
-->

```js
{
    "name": "my-new-instance",                                                      // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                                        // プロファイルの一覧
    "ephemeral": true,                                                              // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                                  // 設定のオーバーライド
    "type": "container",                                                            // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                                    // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "migration",                                                 // "image", "migration", "copy", "none" のいずれかを指定可能
               "mode": "pull",                                                      // 現状 "pull" と "push" がサポートされる
               "operation": "https://10.0.2.3:8443/1.0/operations/<UUID>",          // リモート操作への完全な URL
               "certificate": "PEM certificate",                                    // PEM 証明書を指定可能。未指定の場合はシステムの CA が使用される。
               "base-image": "<fingerprint>",                                       // 任意で指定可能。インスタンスが作られたベースのイメージ
               "instance_only": true,                                               // スナップショットなしでインスタンスだけをマイグレーションするかどうか。 "true" か "false" のいずれか。
               "secrets": {"control": "my-secret-string",                           // マイグレーションのソースと通信する際に使用するシークレット
                           "criu":    "my-other-secret",
                           "fs":      "my third secret"}
    }
}
```

<!--
```js
{
    "name": "my-new-instance",                                                      // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                                        // List of profiles
    "ephemeral": true,                                                              // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                                  // Config override.
    "type": "container",                                                            // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                                    // optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "migration",                                                 // Can be: "image", "migration", "copy" or "none"
               "mode": "pull",                                                      // "pull" and "push" is supported for now
               "operation": "https://10.0.2.3:8443/1.0/operations/<UUID>",          // Full URL to the remote operation (pull mode only)
               "certificate": "PEM certificate",                                    // Optional PEM certificate. If not mentioned, system CA is used.
               "base-image": "<fingerprint>",                                       // Optional, the base image the instance was created from
               "instance_only": true,                                               // Whether to migrate only the instance without snapshots. Can be "true" or "false".
               "secrets": {"control": "my-secret-string",                           // Secrets to use when talking to the migration source
                           "criu":    "my-other-secret",
                           "fs":      "my third secret"}
    }
}
```
-->

入力 (ローカルのインスタンスを使用)
<!--
Input (using a local instance):
-->

```js
{
    "name": "my-new-instance",                                                      // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "profiles": ["default"],                                                        // プロファイルの一覧
    "ephemeral": true,                                                              // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                                  // 設定のオーバーライド
    "type": "container",                                                            // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                                    // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "copy",                                                      // "image", "migration", "copy", "none" のいずれかを指定可能
               "instance_only": true,                                               // スナップショットなしでインスタンスだけをマイグレーションするかどうか。 "true" か "false" のいずれか。
               "source": "my-old-instance"}                                         // 作成元のインスタンスの名前
}
```

<!--
```js
{
    "name": "my-new-instance",                                                      // 64 chars max, ASCII, no slash, no colon and no comma
    "profiles": ["default"],                                                        // List of profiles
    "ephemeral": true,                                                              // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                                  // Config override.
    "type": "container",                                                            // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                                    // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "copy",                                                      // Can be: "image", "migration", "copy" or "none"
               "instance_only": true,                                               // Whether to copy only the instance without snapshots. Can be "true" or "false".
               "source": "my-old-instance"}                                         // Name of the source instance
}
```
-->

入力 (クライアントプロキシ経由でマイグレーションウェブソケット越しに push モードで送られるリモートインスタンスを使用)
<!--
Input (using a remote instance, in push mode sent over the migration websocket via client proxying):
-->

```js
{
    "name": "my-new-instance",                                                      // 最大 64 文字、 ASCII が使用可、スラッシュ、コロン、カンマは使用不可
    "architecture": "x86_64",
    "profiles": ["default"],                                                        // プロファイルの一覧
    "ephemeral": true,                                                              // シャットダウン時にインスタンスを破棄するかどうか
    "config": {"limits.cpu": "2"},                                                  // 設定のオーバーライド
    "type": "container",                                                            // "virtual-machine" か "container" を任意で指定可能、デフォルト値は "container"
    "devices": {                                                                    // インスタンスが持つデバイスの任意で指定可能なリスト
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "migration",                                                 // "image", "migration", "copy", "none" のいずれかを指定可能
               "mode": "push",                                                      // "pull" と "push" がサポートされている
               "base-image": "<fingerprint>",                                       // 任意で指定可能。インスタンスが作られたベースのイメージ
               "live": true,                                                        // マイグレーションが live で実行されるかどうか
               "instance_only": true}                                               // スナップショットなしでインスタンスだけをマイグレーションするかどうか。 "true" か "false" のいずれか。
}
```

<!--
```js
{
    "name": "my-new-instance",                                                      // 64 chars max, ASCII, no slash, no colon and no comma
    "architecture": "x86_64",
    "profiles": ["default"],                                                        // List of profiles
    "ephemeral": true,                                                              // Whether to destroy the instance on shutdown
    "config": {"limits.cpu": "2"},                                                  // Config override.
    "type": "container",                                                            // Optional Can be: "virtual-machine", "container" by default it set to "container"
    "devices": {                                                                    // Optional list of devices the instance should have
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        },
    },
    "source": {"type": "migration",                                                 // Can be: "image", "migration", "copy" or "none"
               "mode": "push",                                                      // "pull" and "push" are supported
               "base-image": "<fingerprint>",                                       // Optional, the base image the instance was created from
               "live": true,                                                        // Whether migration is performed live
               "instance_only": true}                                               // Whether to migrate only the instance without snapshots. Can be "true" or "false".
}
```
-->

入力 (バックアップを使用)
<!--
Input (using a backup):
-->

バックアップダウンロードにより提供される生の圧縮された tarball

<!--
Raw compressed tarball as provided by a backup download.
-->

### `/1.0/instances/<name>`
#### GET
 * 説明: インスタンスの情報 <!-- Description: Instance information -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: インスタンスの設定と現在の状態の dict `<!-- Return: dict of the instance configuration and current state. -->

出力
<!--
Output:
-->

```js
{
    "architecture": "x86_64",
    "config": {
        "limits.cpu": "3",
        "volatile.base_image": "97d97a3d1d053840ca19c86cdd0596cf1be060c5157d31407f2a4f9f350c78cc",
        "volatile.eth0.hwaddr": "00:16:3e:1c:94:38"
    },
    "created_at": "2016-02-16T01:05:05Z",
    "devices": {
        "root": {
            "path": "/",
            "type": "disk"
        }
    },
    "ephemeral": false,
    "expanded_config": {    // プロファイルを展開したものにインスタンスのローカルの設定を追加した結果
        "limits.cpu": "3",
        "volatile.base_image":  "97d97a3d1d053840ca19c86cdd0596cf1be060c5157d31407f2a4f9f350c78cc",
        "volatile.eth0.hwaddr":: "00:16:3e:1c:94:38"
    },
    "expanded_devices": {   // プロファイルを展開したものにインスタンスのローカルのデバイスを追加した結果
        "eth0": {
            "name": "eth0",
            "nictype": "bridgedd",
            "parent": "lxdbr0",,
            "type": "nic"
        },
        "root": {
            "path": "/",
            "type": "disk"
        }
    },
    "last_used_at": "2016-02-166T01:05:05Z",
    "name": "my-instance",
    "profiles": [
        "default"
    ],
    "stateful": false,      // true の場合はインスタンスがスタートアップ時に復元できる何らかの保管された状態を持つことを意味する
    "status": "Running",
    "status_code": 103
}
```

<!--
```js
{
    "architecture": "x86_64",
    "config": {
        "limits.cpu": "3",
        "volatile.base_image": "97d97a3d1d053840ca19c86cdd0596cf1be060c5157d31407f2a4f9f350c78cc",
        "volatile.eth0.hwaddr": "00:16:3e:1c:94:38"
    },
    "created_at": "2016-02-16T01:05:05Z",
    "devices": {
        "root": {
            "path": "/",
            "type": "disk"
        }
    },
    "ephemeral": false,
    "expanded_config": {    // the result of expanding profiles and adding the instance's local config
        "limits.cpu": "3",
        "volatile.base_image": "97d97a3d1d053840ca19c86cdd0596cf1be060c5157d31407f2a4f9f350c78cc",
        "volatile.eth0.hwaddr": "00:16:3e:1c:94:38"
    },
    "expanded_devices": {   // the result of expanding profiles and adding the instance's local devices
        "eth0": {
            "name": "eth0",
            "nictype": "bridged",
            "parent": "lxdbr0",
            "type": "nic"
        },
        "root": {
            "path": "/",
            "type": "disk"
        }
    },
    "last_used_at": "2016-02-16T01:05:05Z",
    "name": "my-instance",
    "profiles": [
        "default"
    ],
    "stateful": false,      // If true, indicates that the instance has some stored state that can be restored on startup
    "status": "Running",
    "status_code": 103
}
```
-->

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: インスタンスの設定を置き換えるかスナップショットをリストアします <!-- Description: replaces instance configuration or restore snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (インスタンスの設定を更新します)
<!--
Input (update instance configuration):
-->

```json
{
    "architecture": "x86_64",
    "config": {
        "limits.cpu": "4",
        "volatile.base_image": "97d97a3d1d053840ca19c86cdd0596cf1be060c5157d31407f2a4f9f350c78cc",
        "volatile.eth0.hwaddr": "00:16:3e:1c:94:38"
    },
    "devices": {
        "root": {
            "path": "/",
            "type": "disk"
        }
    },
    "ephemeral": true,
    "profiles": [
        "default"
    ]
}
```

GET の戻り値と同じ構造を持つが、名前の変更は許されず (以下の POST 参照)、
status の sub-dict への変更も許されません (status の sub-dict は読み取り
専用のため)。
<!--
Takes the same structure as that returned by GET but doesn't allow name
changes (see POST below) or changes to the status sub-dict (since that's
read-only).
-->

入力 (スナップショットをリストアします)
<!--
Input (restore snapshot):
-->

```json
{
    "restore": "snapshot-name"
}
```

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: インスタンスの設定を更新します <!-- Description: update instance configuration -->
 * 導入: `patch` API 拡張によって <!-- Introduced: with API extension `patch` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "limits.cpu": "4"
    },
    "devices": {
        "root": {
            "path": "/",
            "pool": "default",
            "size": "5GB"
        }
    },
    "ephemeral": true
}
```

#### POST (`?target=<member>` を任意で指定可能) <!-- POST (optional `?target=<member>`) -->
 * 説明: インスタンスをリネーム／マイグレーションするのに用いられます <!-- Description: used to rename/migrate the instance -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

既に存在する名前にリネームしようとすると 409 (Conflict) という HTTP ステータスコードを返します。
<!--
Renaming to an existing name must return the 409 (Conflict) HTTP code.
-->

入力 (単純なリネーム)
<!--
Input (simple rename):
-->

```json
{
    "name": "new-name"
}
```

入力 (lxd インスタンスまたは lxd クラスタメンバ間でのマイグレーション)
<!--
Input (migration across lxd instances or lxd cluster members):
-->

```json
{
    "name": "new-name",
    "migration": true,
    "live": "true"
}
```

誰か (つまり他の lxd インスタンス) が全てのウェブソケットに接続してソースと
交渉を始めるまでは、マイグレーションは実際には開始されません。
<!--
The migration does not actually start until someone (i.e. another lxd instance)
connects to all the websockets and begins negotiation with the source.
-->

クラスタメンバ間でマイグレーションするには `?target=<member>` オプションが
必要です。
<!--
To migrate between cluster members the `?target=<member>` option is required.
-->

メタデータセクション内の出力 (マイグレーションの場合)
<!--
Output in metadata section (for migration):
-->

```js
{
    "control": "secret1",       // マイグレーション制御ソケット
    "criu": "secret2",          // 状態転送ソケット (ライブマイグレーションのときのみ)
    "fs": "secret3"             // ファイルシステム転送ソケット
}
```

<!--
```js
{
    "control": "secret1",       // Migration control socket
    "criu": "secret2",          // State transfer socket (only if live migrating)
    "fs": "secret3"             // Filesystem transfer socket
}
```
-->

これらは作成の呼び出し時に渡すべきシークレットです。
<!--
These are the secrets that should be passed to the create call.
-->

#### DELETE
 * 説明: インスタンスを削除します <!-- Description: remove the instance -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

この操作に対する HTTP レスポンスのステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/instances/<name>/console`
#### GET
 * 説明: インスタンスのコンソールログの内容を返します <!-- Description: returns the contents of the instance's console  log -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 該当なし <!-- Operation: N/A -->
 * 戻り値: コンソールログの内容 <!-- Return: the contents of the console log -->

#### POST
 * 説明: インスタンスのコンソールデバイスにアタッチします <!-- Description: attach to an instance's console devices -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: 標準のエラー <!-- Return: standard error -->

入力 (/dev/console にアタッチします)
<!--
Input (attach to /dev/console):
-->

```js
{
    "width": 80,                    // 端末の初期の幅 (任意で指定可能)
    "height": 25,                   // 端末の初期の高さ (任意で指定可能)
    "type": "console"               // 接続タイプ ("console" か "vga")
}
```

<!--
```js
{
    "width": 80,                    // Initial width of the terminal (optional)
    "height": 25,                   // Initial height of the terminal (optional)
    "type": "console"               // Connection type ("console" or "vga").
}
```
-->

"vga" 接続タイプは仮想マシンでのみサポートされます。
<!--
The "vga" connection type is supported only for virtual machines.
-->

制御用ウェブソケットがコンソールセッションの out-of-band メッセージの送信に使用されます。
現状ではウィンドウサイズの変更に使われています。
<!--
The control websocket can be used to send out-of-band messages during a console session.
This is currently used for window size changes.
-->

制御 (ウィンドウサイズの変更)
<!--
Control (window size change):
-->

```json
{
    "command": "window-resize",
    "args": {
        "width": "80",
        "height": "50"
    }
}
```

#### DELETE
 * 説明: インスタンスのコンソールログを空にします <!-- Description: empty the instance's console log -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: 空のレスポンスまたは標準のエラー <!-- Return: empty response or standard error -->

### `/1.0/instances/<name>/exec`
#### POST
 * 説明: リモートコマンドを実行します <!-- Description: run a remote command -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作 + 任意で指定可能な websocket 情報あるいは標準のエラー <!-- Return: background operation + optional websocket information or standard error -->

入力 (bash を実行する例です)
<!--
Input (run bash):
-->

```js
{
    "command": ["/bin/bash"],       // コマンドと引数
    "environment": {},              // 追加で設定する任意で指定可能な環境変数
    "wait-for-websocket": false,    // プロセスを開始する前に接続を待つかどうか
    "record-output": false,         // 標準出力と標準エラー出力を記録するかどうか (wait-for-websocket=false のときのみ有効) (container_exec_recording API 拡張が必要)
    "interactive": true,            // PIPE の代わりに pty デバイスを割り当てるかどうか
    "width": 80,                    // 端末の初期の幅 (任意で指定可能)
    "height": 25,                   // 端末の初期の高さ (任意で指定可能)
    "user": 1000,                   // コマンドを実行するユーザー (任意で指定可能)
    "group": 1000,                  // コマンドを実行するグループ (任意で指定可能)
    "cwd": "/tmp"                   // 現在の作業ディレクトリ (任意で指定可能)
}
```

<!--
```js
{
    "command": ["/bin/bash"],       // Command and arguments
    "environment": {},              // Optional extra environment variables to set
    "wait-for-websocket": false,    // Whether to wait for a connection before starting the process
    "record-output": false,         // Whether to store stdout and stderr (only valid with wait-for-websocket=false) (requires API extension container_exec_recording)
    "interactive": true,            // Whether to allocate a pty device instead of PIPEs
    "width": 80,                    // Initial width of the terminal (optional)
    "height": 25,                   // Initial height of the terminal (optional)
    "user": 1000,                   // User to run the command as (optional)
    "group": 1000,                  // Group to run the command as (optional)
    "cwd": "/tmp"                   // Current working directory (optional)
}
```
-->

`wait-for-websocket` は操作をブロックしウェブソケットの接続が
(オプショナルである `control` を除く) 全ての利用可能なファイルディスクリプタに対して
開始するのを待つか、あるいは即座に開始するかを指示します。
これによりユーザーが標準入力を渡したり、標準出力や標準エラー出力の出力をバイト列として読み取る選択肢を提供します。
<!--
`wait-for-websocket` indicates whether the operation should block and wait for
a websocket connection to start for all the available file descriptors (except `control`, which is optional), or start immediately.
This gives the possibility to pass stdin inputs and read stdout/stderr outputs as bytes.
-->

即座に開始する場合、 /dev/null が標準入力、標準出力、標準エラー出力に
使われます。これは record-output が true に設定されない場合です。
true に設定される場合は、標準出力と標準エラー出力はログファイルに
リダイレクトされます。
<!--
If starting immediately, /dev/null will be used for stdin, stdout and
stderr. That's unless record-output is set to true, in which case,
stdout and stderr will be redirected to a log file.
-->

interactive が true に設定される場合は、 1 つのウェブソケットが返され、
それが実行されたプロセスの標準入力、標準出力、標準エラー出力用の pty
デバイスにマッピングされます。
<!--
If interactive is set to true, a single websocket is returned and is mapped to a
pty device for stdin, stdout and stderr of the execed process.
-->

interactive が false (デフォルト) に設定される場合は、標準入力、標準出力、
標準エラー出力に 1 つずつ、合計 3 つのパイプがセットアップされます。
<!--
If interactive is set to false (default), three pipes will be setup, one
for each of stdin, stdout and stderr.
-->

interactive フラグの状態によって、 1 つまたは 3 つのウェブソケットと
シークレットの組が返され、それはこの操作の /websocket エンドポイントに
接続するのに有効です。
<!--
Depending on the state of the interactive flag, one or three different
websocket/secret pairs will be returned, which are valid for connecting to this
operations /websocket endpoint.
-->


実行セッションの間、制御用のウェブソケットが out-of-band メッセージを送るのに
利用できます。これは現状はウィンドウサイズの変更とシグナルのフォワーディングに
使われています。
<!--
The control websocket can be used to send out-of-band messages during an exec session.
This is currently used for window size changes and for forwarding of signals.
-->

制御 (ウィンドウサイズの変更)
<!--
Control (window size change):
-->

```json
{
    "command": "window-resize",
    "args": {
        "width": "80",
        "height": "50"
    }
}
```

制御 (SIGUSR1 シグナル)
<!--
Control (SIGUSR1 signal):
-->

```json
{
    "command": "signal",
    "signal": 10
}
```

戻り値 (wait-for-websocket=true で interactive=false の場合)
<!--
Return (with wait-for-websocket=true and interactive=false):
-->

```json
{
    "fds": {
        "0": "f5b6c760c0aa37a6430dd2a00c456430282d89f6e1661a077a926ed1bf3d1c21",
        "1": "464dcf9f8fdce29d0d6478284523a9f26f4a31ae365d94cd38bac41558b797cf",
        "2": "25b70415b686360e3b03131e33d6d94ee85a7f19b0f8d141d6dca5a1fc7b00eb",
        "control": "20c479d9532ab6d6c3060f6cdca07c1f177647c9d96f0c143ab61874160bd8a5"
    }
}
```

戻り値 (wait-for-websocket=true で interactive=true の場合)
<!--
Return (with wait-for-websocket=true and interactive=true):
-->

```json
{
    "fds": {
        "0": "f5b6c760c0aa37a6430dd2a00c456430282d89f6e1661a077a926ed1bf3d1c21",
        "control": "20c479d9532ab6d6c3060f6cdca07c1f177647c9d96f0c143ab61874160bd8a5"
    }
}
```

戻り値 (interactive=false で record-output=true の場合)
<!--
Return (with interactive=false and record-output=true):
-->

```json
{
    "output": {
        "1": "/1.0/instances/example/logs/exec_b0f737b4-2c8a-4edf-a7c1-4cc7e4e9e155.stdout",
        "2": "/1.0/instances/example/logs/exec_b0f737b4-2c8a-4edf-a7c1-4cc7e4e9e155.stderr"
    },
    "return": 0
}
```

実行コマンドが終了した時は、終了ステータスが操作のメタデータに
含まれます。
<!--
When the exec command finishes, its exit status is available from the
operation's metadata:
-->

```json
{
    "return": 0
}
```

### `/1.0/instances/<name>/files`
#### GET (`?path=/path/inside/the/instance`)
 * 説明: ファイルかディレクトリの内容をインスタンスからダウンロードします <!-- Description: download a file or directory listing from the instance -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ファイルの種別がディレクトリの場合、戻り値はメタデータにディレクトリの内容の一覧を
   含んだ同期的なレスポンスになり、それ以外の種別の場合はファイルの生の内容になります。 <!-- Return: if the type of the file is a directory, the return is a sync
   response with a list of the directory contents as metadata, otherwise it is
   the raw contents of the file. -->

次のヘッダがセットされます (標準のサイズと MIME タイプのヘッダに加えて)
<!--
The following headers will be set (on top of standard size and mimetype headers):
-->

 * `X-LXD-uid`: 0
 * `X-LXD-gid`: 0
 * `X-LXD-mode`: 0700
 * `X-LXD-type`: `directory` か `file` のいずれか <!-- one of `directory` or `file` -->

これはコマンドラインあるいはウェブブラウザからでさえ簡単に使えるように
設計されています。
<!--
This is designed to be easily usable from the command line or even a web
browser.
-->

#### POST (`?path=/path/inside/the/instance`)
 * 説明: インスタンスにファイルをアップロードします <!-- Description: upload a file to the instance -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->
 * 標準的な HTTP のファイルアップロード <!-- Standard http file upload -->

クライアントは次のヘッダを設定しても構いません。
<!--
The following headers may be set by the client:
-->

 * `X-LXD-uid`: 0
 * `X-LXD-gid`: 0
 * `X-LXD-mode`: 0700
 * `X-LXD-type`: `directory`, `file`, `symlink` のいずれか <!-- one of `directory`, `file` or `symlink` -->
 * `X-LXD-write`: overwrite (か append。 append は `file_append` API 拡張によって導入されます) <!-- overwrite (or append, introduced with API extension `file_append`) -->

これはコマンドラインあるいはウェブブラウザからでさえ簡単に使えるように
設計されています。
<!--
This is designed to be easily usable from the command line or even a web
browser.
-->

#### DELETE (`?path=/path/inside/the/instance`)
 * 説明: インスタンス内のファイルを削除します <!-- Description: delete a file in the instance -->
 * 導入: `file_delete` API 拡張によって <!-- Introduced: with API extension `file_delete` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

### `/1.0/instances/<name>/snapshots`
#### GET
 * 説明: スナップショットの一覧 <!-- Description: List of snapshots -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: このインスタンスのスナップショットの URL の一覧 <!-- Return: list of URLs for snapshots for this instance -->

戻り値
<!--
Return value:
-->

```json
[
    "/1.0/instances/blah/snapshots/snap0"
]
```

#### POST
 * 説明: 新しいスナップショットを作成します <!-- Description: create a new snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```js
{
    "name": "my-snapshot",          // スナップショットの名前
    "stateful": true                // 状態も含めるかどうか
}
```

<!--
```js
{
    "name": "my-snapshot",          // Name of the snapshot
    "stateful": true                // Whether to include state too
}
```
-->

### `/1.0/instances/<name>/snapshots/<name>`
#### GET
 * 説明: スナップショットの情報 <!-- Description: Snapshot information -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: スナップショットを表す dict <!-- Return: dict representing the snapshot -->

戻り値
<!--
Return:
-->

```json
{
    "architecture": "x86_64",
    "config": {
        "security.nesting": "true",
        "volatile.base_image": "a49d26ce5808075f5175bf31f5cb90561f5023dcd408da8ac5e834096d46b2d8",
        "volatile.eth0.hwaddr": "00:16:3e:ec:65:a8",
        "volatile.last_state.idmap": "[{\"Isuid\":true,\"Isgid\":false,\"Hostid\":100000,\"Nsid\":0,\"Maprange\":65536},{\"Isuid\":false,\"Isgid\":true,\"Hostid\":100000,\"Nsid\":0,\"Maprange\":65536}]",
    },
    "created_at": "2016-03-08T23:55:08Z",
    "devices": {
        "eth0": {
            "name": "eth0",
            "nictype": "bridged",
            "parent": "lxdbr0",
            "type": "nic"
        },
        "root": {
            "path": "/",
            "type": "disk"
        },
    },
    "ephemeral": false,
    "expanded_config": {
        "security.nesting": "true",
        "volatile.base_image": "a49d26ce5808075f5175bf31f5cb90561f5023dcd408da8ac5e834096d46b2d8",
        "volatile.eth0.hwaddr": "00:16:3e:ec:65:a8",
        "volatile.last_state.idmap": "[{\"Isuid\":true,\"Isgid\":false,\"Hostid\":100000,\"Nsid\":0,\"Maprange\":65536},{\"Isuid\":false,\"Isgid\":true,\"Hostid\":100000,\"Nsid\":0,\"Maprange\":65536}]",
    },
    "expanded_devices": {
        "eth0": {
            "name": "eth0",
            "nictype": "bridged",
            "parent": "lxdbr0",
            "type": "nic"
        },
        "root": {
            "path": "/",
            "type": "disk"
        },
    },
    "name": "blah",
    "profiles": [
        "default"
    ],
    "size": 738476032,
    "stateful": false
}
```

#### POST
 * 説明: スナップショットをリネーム／マイグレートします <!-- Description: used to rename/migrate the snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (スナップショットをリネーム)
<!--
Input (rename the snapshot):
-->

```json
{
    "name": "new-name"
}
```

入力 (マイグレーション元をセットアップ)
<!--
Input (setup the migration source):
-->

```json
{
    "name": "new-name",
    "migration": true,
    "live": "true"
}
```

戻り値 (migration=true の場合)
<!--
Return (with migration=true):
-->

```js
{
    "control": "secret1",       // マイグレーション制御ソケット
    "fs": "secret3"             // ファイルシステム転送ソケット
}
```

<!--
```js
{
    "control": "secret1",       // Migration control socket
    "fs": "secret3"             // Filesystem transfer socket
}
```
-->

既に存在する名前にリネームしようとすると 409 (Conflict) という HTTP ステータスコードが返ります。
<!--
Renaming to an existing name must return the 409 (Conflict) HTTP code.
-->

`default` プロファイルをリネームしようとすると 403 (Forbidden) という HTTP ステータスコードが返ります。
<!--
Attempting to rename the `default` profile will return the 403 (Forbidden) HTTP code.
-->

#### DELETE
 * 説明: スナップショットを削除します <!-- Description: remove the snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

#### PUT
 * 説明: スナップショットを更新します <!-- Description: update the snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

Input:

```json
{
    "expires_at": "2019-01-16T12:34:56+02:00"
}
```

この操作に対する HTTP ステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/instances/<name>/state`
#### GET
 * 説明: 現在の状態 <!-- Description: current state -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 現在の状態を表す dict <!-- Return: dict representing current state -->

出力
<!--
Output:
-->

```json
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "metadata": {
        "status": "Running",
        "status_code": 103,
        "cpu": {
            "usage": 4986019722
        },
        "disk": {
            "root": {
                "usage": 422330368
            }
        },
        "memory": {
            "usage": 51126272,
            "usage_peak": 70246400,
            "swap_usage": 0,
            "swap_usage_peak": 0
        },
        "network": {
            "eth0": {
                "addresses": [
                    {
                        "family": "inet",
                        "address": "10.0.3.27",
                        "netmask": "24",
                        "scope": "global"
                    },
                    {
                        "family": "inet6",
                        "address": "fe80::216:3eff:feec:65a8",
                        "netmask": "64",
                        "scope": "link"
                    }
                ],
                "counters": {
                    "bytes_received": 33942,
                    "bytes_sent": 30810,
                    "packets_received": 402,
                    "packets_sent": 178
                },
                "hwaddr": "00:16:3e:ec:65:a8",
                "host_name": "vethBWTSU5",
                "mtu": 1500,
                "state": "up",
                "type": "broadcast"
            },
            "lo": {
                "addresses": [
                    {
                        "family": "inet",
                        "address": "127.0.0.1",
                        "netmask": "8",
                        "scope": "local"
                    },
                    {
                        "family": "inet6",
                        "address": "::1",
                        "netmask": "128",
                        "scope": "local"
                    }
                ],
                "counters": {
                    "bytes_received": 86816,
                    "bytes_sent": 86816,
                    "packets_received": 1226,
                    "packets_sent": 1226
                },
                "hwaddr": "",
                "host_name": "",
                "mtu": 65536,
                "state": "up",
                "type": "loopback"
            },
            "lxdbr0": {
                "addresses": [
                    {
                        "family": "inet",
                        "address": "10.0.3.1",
                        "netmask": "24",
                        "scope": "global"
                    },
                    {
                        "family": "inet6",
                        "address": "fe80::68d4:87ff:fe40:7769",
                        "netmask": "64",
                        "scope": "link"
                    }
                ],
                "counters": {
                    "bytes_received": 0,
                    "bytes_sent": 570,
                    "packets_received": 0,
                    "packets_sent": 7
                },
                "hwaddr": "6a:d4:87:40:77:69",
                "host_name": "",
                "mtu": 1500,
                "state": "up",
                "type": "broadcast"
           },
           "zt0": {
                "addresses": [
                    {
                        "family": "inet",
                        "address": "29.17.181.59",
                        "netmask": "7",
                        "scope": "global"
                    },
                    {
                        "family": "inet6",
                        "address": "fd80:56c2:e21c:0:199:9379:e711:b3e1",
                        "netmask": "88",
                        "scope": "global"
                    },
                    {
                        "family": "inet6",
                        "address": "fe80::79:e7ff:fe0d:5123",
                        "netmask": "64",
                        "scope": "link"
                    }
                ],
                "counters": {
                    "bytes_received": 0,
                    "bytes_sent": 806,
                    "packets_received": 0,
                    "packets_sent": 9
                },
                "hwaddr": "02:79:e7:0d:51:23",
                "host_name": "",
                "mtu": 2800,
                "state": "up",
                "type": "broadcast"
            }
        },
        "pid": 13663,
        "processes": 32
    }
}
```

#### PUT
 * 説明: インスタンスの状態を変更する <!-- Description: change the instance state -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```js
{
    "action": "stop",       // 状態を変更するアクション (stop, start, restart, freeze, unfreeze のいずれか)
    "timeout": 30,          // 状態の変更が失敗したと判定するまでのタイムアウト
    "force": true,          // 状態の変更を強制する (現状では stop と restart でのみ有効で、インスタンスを強制停止することを意味します)
    "stateful": true        // 停止または開始する前の状態を保管または復元するかどうか (stop と start でのみ有効、デフォルトは false)
}
```

<!--
```js
{
    "action": "stop",       // State change action (stop, start, restart, freeze or unfreeze)
    "timeout": 30,          // A timeout after which the state change is considered as failed
    "force": true,          // Force the state change (currently only valid for stop and restart where it means killing the instance)
    "stateful": true        // Whether to store or restore runtime state before stopping or startiong (only valid for stop and start, defaults to false)
}
```
-->

### `/1.0/instances/<name>/logs`
#### GET
 * 説明: このインスタンスで利用可能なログファイルの一覧を返します。
   作成の失敗についてのログを取得できるようにするため、この操作は
   削除が完了した (あるいは一度も作られなかった) インスタンスに対しても
   動作します。 <!-- Description: Returns a list of the log files available for this instance.
  Note that this works on instances that have been deleted (or were never
  created) to enable people to get logs for failed creations. -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: 利用可能なログファイルの一覧 <!-- Return: a list of the available log files -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/instances/blah/logs/forkstart.log",
    "/1.0/instances/blah/logs/lxc.conf",
    "/1.0/instances/blah/logs/lxc.log"
]
```

### `/1.0/instances/<name>/logs/<logfile>`
#### GET
 * 説明: 特定のログファイルの中身を返します <!-- Description: returns the contents of a particular log file. -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 該当なし <!-- Operation: N/A -->
 * 戻り値: ログファイルの中身 <!-- Return: the contents of the log file -->

#### DELETE
 * 説明: 特定のログファイルを削除します <!-- Description: delete a particular log file. -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: 空のレスポンスまたは標準のエラー <!-- Return: empty response or standard error -->

### `/1.0/instances/<name>/metadata`
#### GET
 * 説明: インスタンスのメタデータ <!-- Description: Instance metadata -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: インスタンスのメタデータを表す dict <!-- Return: dict representing instance metadata -->

戻り値
<!--
Return:
-->

```json
{
    "architecture": "x86_64",
    "creation_date": 1477146654,
    "expiry_date": 0,
    "properties": {
        "architecture": "x86_64",
        "description": "BusyBox x86_64",
        "name": "busybox-x86_64",
        "os": "BusyBox"
    },
    "templates": {
        "/template": {
            "when": [
                ""
            ],
            "create_only": false,
            "template": "template.tpl",
            "properties": {}
        }
    }
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: インスタンスのメタデータを置き換える <!-- Description: Replaces instance metadata -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "architecture": "x86_64",
    "creation_date": 1477146654,
    "expiry_date": 0,
    "properties": {
        "architecture": "x86_64",
        "description": "BusyBox x86_64",
        "name": "busybox-x86_64",
        "os": "BusyBox"
    },
    "templates": {
        "/template": {
            "when": [
                ""
            ],
            "create_only": false,
            "template": "template.tpl",
            "properties": {}
        }
    }
}
```

### `/1.0/instances/<name>/metadata/templates`
#### GET
 * 説明: インスタンステンプレートの一覧 <!-- Description: List instance templates -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: インスタンステンプレート名の一覧 <!-- Return: a list with instance template names -->

戻り値
<!--
Return:
-->

```json
[
    "template.tpl",
    "hosts.tpl"
]
```

#### GET (`?path=<template>`)
 * 説明: インスタンステンプレートの中身 <!-- Description: Content of an instance template -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: テンプレートの中身 <!-- Return: the content of the template -->

#### POST (`?path=<template>`)
 * 説明: インスタンステンプレートを追加します <!-- Description: Add an instance template -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->
 * 標準的な HTTP のファイルアップロード <!-- Standard http file upload -->

#### PUT (`?path=<template>`)
 * 説明: テンプレートの中身を置き換えます <!-- Description: Replace content of a template -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->
 * 標準的な HTTP のファイルアップロード <!-- Standard http file upload -->

#### DELETE (`?path=<template>`)
 * 説明: コンテナーテンプレートを削除します <!-- Description: Delete a container template -->
 * 導入: `container_edit_metadata` API 拡張によって <!-- Introduced: with API extension `container_edit_metadata` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: Sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

### `/1.0/instances/<name>/backups`
#### GET
 * 説明: インスタンスのバックアップの一覧 <!-- Description: List of backups for the instance -->
 * 導入: `container_backup` API 拡張によって <!-- Introduced: with API extension `container_backup` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: インスタンスのバックアップの一覧 <!-- Return: a list of backups for the instance -->

戻り値
<!--
Return value:
-->

```json
[
    "/1.0/instances/c1/backups/c1/backup0",
    "/1.0/instances/c1/backups/c1/backup1",
]
```

#### POST
 * 説明: 新しいバックアップを作成します <!-- Description: Create a new backup -->
 * 導入: `container_backup` API 拡張によって <!-- Introduced: with API extension `container_backup` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```js
{
    "name": "backupName",      // バックアップのユニークな識別子
    "expiry": 3600,            // いつ自動的にバックアップを削除するか
    "instance_only": true,     // true の場合、スナップショットは含まれません
    "optimized_storage": true  // true の場合 btrfs send または zfs send がインスタンスとスナップショットに対して使用されます
}
```

<!--
```js
{
    "name": "backupName",      // unique identifier for the backup
    "expiry": 3600,            // when to delete the backup automatically
    "instance_only": true,     // if True, snapshots aren't included
    "optimized_storage": true  // if True, btrfs send or zfs send is used for instance and snapshots
}
```
-->

### `/1.0/instances/<name>/backups/<name>`
#### GET
 * 説明: バックアップの情報 <!-- Description: Backup information -->
 * 導入: `container_backup` API 拡張によって <!-- Introduced: with API extension `container_backup` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: バックアップの dict <!-- Returns: dict of the backup -->

出力
<!--
Output:
-->

```json
{
    "name": "backupName",
    "creation_date": "2018-04-23T12:16:09+02:00",
    "expiry_date": "2018-04-23T12:16:09+02:00",
    "instance_only": false,
    "optimized_storage": false
}
```

#### DELETE
 * 説明: バックアップを削除します <!-- Description: remove the backup -->
 * 導入: `container_backup` API 拡張によって <!-- Introduced: with API extension `container_backup` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

#### POST
 * 説明: バックアップをリネームします <!-- Description: used to rename the backup -->
 * 導入: `container_backup` API 拡張によって <!-- Introduced: with API extension `container_backup` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "new-name"
}
```

### `/1.0/instances/<name>/backups/<name>/export`
#### GET
 * 説明: バックアップの tarball を取得します <!-- Description: fetch the backup tarball -->
 * 導入: `container_backup` API 拡張によって <!-- Introduced: with API extension `container_backup` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: バックアップの tarball を含む dict <!-- Return: dict containing the backup tarball -->

出力
<!--
Output:
-->

```json
{
    "data": "<byte-stream>"
}
```

### `/1.0/events`
この URL は真の REST API エンドポイントではなく、代わりに GET クエリを
実行すると接続をウェブソケットにアップグレードし、そのウェブソケット上で
通知が送信されます。
<!--
This URL isn't a real REST API endpoint, instead doing a GET query on it
will upgrade the connection to a websocket on which notifications will
be sent.
-->

#### GET (`?type=operation,logging`)
 * 説明: ウェブソケットへのアップグレード <!-- Description: websocket upgrade -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: なし (イベントのフローが終わることはありません) <!-- Return: none (never ending flow of events) -->

サポートされる引数は次のとおりです。
<!--
Supported arguments are:
-->

 * type: 購読する通知のカンマ区切りリスト (デフォルトは all) <!-- type: comma separated list of notifications to subscribe to (defaults to all) -->

通知の種別は次のとおりです。
<!--
The notification types are:
-->

 * operation (作成、更新、終了という全てのバックグラウンド操作についての通知) <!-- operation (notification about creation, updates and termination of all background operations) -->
 * logging (サーバからの全てのログエントリー) <!-- logging (every log entry from the server) -->
 * lifecycle (インスタンスのライフサイクルイベント) <!-- lifecycle (instance lifecycle events) -->

このエンドポイントの出力が完了することはありません。それぞれの通知は個別の JSON dict として送られます。
<!--
This never returns. Each notification is sent as a separate JSON dict:
-->

```js
{
    "timestamp": "2015-06-09T19:07:24.379615253-06:00",                // 現在のタイムスタンプ
    "type": "operation",                                               // 通知の種別
    "metadata": {}                                                     // リソースまたはタイプに特有な追加のメタデータ
}
```

```json
{
    "timestamp": "2016-02-17T11:44:28.572721913-05:00",
    "type": "logging",
    "metadata": {
        "context": {
            "ip": "@",
            "method": "GET",
            "url": "/1.0/instances/xen/snapshots",
        },
        "level": "info",
        "message": "handling"
    }
}
```

<!--
```js
{
    "timestamp": "2015-06-09T19:07:24.379615253-06:00",                // Current timestamp
    "type": "operation",                                               // Notification type
    "metadata": {}                                                     // Extra resource or type specific metadata
}
```

```json
{
    "timestamp": "2016-02-17T11:44:28.572721913-05:00",
    "type": "logging",
    "metadata": {
        "context": {
            "ip": "@",
            "method": "GET",
            "url": "/1.0/containers/xen/snapshots",
        },
        "level": "info",
        "message": "handling"
    }
}
```
-->

### `/1.0/images`
#### GET
 * 説明: (public または private の) イメージの一覧 <!-- Description: list of images (public or private) -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: このサーバが提供しているイメージの URL の一覧 <!-- Return: list of URLs for images this server publishes -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/images/54c8caac1f61901ed86c68f24af5f5d3672bdc62c71d04f06df3a59e95684473",
    "/1.0/images/97d97a3d1d053840ca19c86cdd0596cf1be060c5157d31407f2a4f9f350c78cc",
    "/1.0/images/a49d26ce5808075f5175bf31f5cb90561f5023dcd408da8ac5e834096d46b2d8",
    "/1.0/images/c9b6e738fae75286d52f497415463a8ecc61bbcb046536f220d797b0e500a41f"
]
```

#### POST
 * 説明: 新しいイメージを作成し提供する <!-- Description: create and publish a new image -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (次のいずれか 1 つ)
<!--
Input (one of):
-->

 * 標準の HTTP ファイルアップロード <!-- Standard http file upload -->
 * 作成元のイメージの dict (リモートのイメージを転送する場合) <!-- Source image dictionary (transfers a remote image) -->
 * 作成元のインスタンスの dict (ローカルインスタンスからイメージを作成する場合) <!-- Source instance dictionary (makes an image out of a local instance) -->
 * リモートのイメージの URL の dict (リモートのイメージをダウンロードする場合) <!-- Remote image URL dictionary (downloads a remote image) -->

HTTP ファイルアップロードの場合、次のヘッダがクライアントにより設定可能です。
<!--
In the http file upload case, The following headers may be set by the client:
-->

 * `X-LXD-fingerprint`: SHA-256 (設定された場合はアップロードされたファイルのフィンガープリントが一致する必要があります) <!-- (if set, uploaded file must match) -->
 * `X-LXD-filename`: FILENAME (エクスポートの際に使用されます) <!-- (used for export) -->
 * `X-LXD-public`: true/false (デフォルトは false) <!-- (defaults to false) -->
 * `X-LXD-properties`: 重複したキーを除いたキーと値のペアを URL エンコードしたもの (任意で指定可能なプロパティ) <!-- URL-encoded key value pairs without duplicate keys (optional properties) -->

作成元としてイメージを使う場合、次の dict を使用する必要があります。
<!--
In the source image case, the following dict must be used:
-->

```js
{
    "filename": filename,                   // エクスポートの際に使用されます (任意で指定可能)
    "public": true,                         // 信頼されないユーザーがイメージをダウンロードしてよいか (デフォルトは false)
    "auto_update": true,                    // イメージを自動更新するかどうか (任意で指定可能、デフォルトは false)
    "properties": {                         // イメージのプロパティ (任意で指定可能、作成元のプロパティに追加して適用されます)
        "os": "Ubuntu"
    },
    "aliases": [                            // 初期のエイリアスを設定します ("image_create_aliases" API 拡張)
        {"name": "my-alias",
         "description": "A description"}
    ],
    "source": {
        "type": "image",
        "mode": "pull",                     // 現在は pull のみがサポートされています
        "server": "https://10.0.2.3:8443",  // リモートサーバ (pull モードのときのみ)
        "protocol": "lxd",                  // プロトコル (lxd または simplestreams、デフォルトは lxd)
        "secret": "my-secret-string",       // シークレット (pull モードのときのみ、 private なイメージのみ)
        "certificate": "PEM certificate",   // 任意で指定可能な PEM 証明書。指定されない場合はシステム CA が使用されます
        "fingerprint": "SHA256",            // イメージのフィンガープリント (エイリアスを指定しない場合は必須です)
        "alias": "ubuntu/devel",            // エイリアスの名前 (フィンガープリントを指定しない場合は必須です)
    }
}
```

<!--
```js
{
    "filename": filename,                   // Used for export (optional)
    "public": true,                         // Whether the image can be downloaded by untrusted users (defaults to false)
    "auto_update": true,                    // Whether the image should be auto-updated (optional; defaults to false)
    "properties": {                         // Image properties (optional, applied on top of source properties)
        "os": "Ubuntu"
    },
    "aliases": [                            // Set initial aliases ("image_create_aliases" API extension)
        {"name": "my-alias",
         "description": "A description"}
    ],
    "source": {
        "type": "image",
        "mode": "pull",                     // Only pull is supported for now
        "server": "https://10.0.2.3:8443",  // Remote server (pull mode only)
        "protocol": "lxd",                  // Protocol (one of lxd or simplestreams, defaults to lxd)
        "secret": "my-secret-string",       // Secret (pull mode only, private images only)
        "certificate": "PEM certificate",   // Optional PEM certificate. If not mentioned, system CA is used.
        "fingerprint": "SHA256",            // Fingerprint of the image (must be set if alias isn't)
        "alias": "ubuntu/devel",            // Name of the alias (must be set if fingerprint isn't)
    }
}
```
-->

作成元にインスタンスを使う場合、次の dict を使用する必要があります。
<!--
In the source instance case, the following dict must be used:
-->

```js
{
    "compression_algorithm": "xz",  // イメージの圧縮アルゴリズムをオーバーライドします (任意で指定可能)
    "filename": filename,           // エクスポートの際に使用されます (任意で指定可能)
    "public":   true,               // 信頼されないユーザーがイメージをダウンロードしてよいか (デフォルトは false)
    "properties": {                 // イメージのプロパティ (任意で指定可能)
        "os": "Ubuntu"
    },
    "aliases": [                    // 初期のエイリアスを設定します ("image_create_aliases" API 拡張)
        {"name": "my-alias",
         "description": "A description""}
    ],
    "source": {
        "type": "instance",         // "instance" か "snapshot" のいずれか
        "name": "abc"
    }
}
```

<!--
```js
{
    "compression_algorithm": "xz",  // Override the compression algorithm for the image (optional)
    "filename": filename,           // Used for export (optional)
    "public":   true,               // Whether the image can be downloaded by untrusted users (defaults to false)
    "properties": {                 // Image properties (optional)
        "os": "Ubuntu"
    },
    "aliases": [                    // Set initial aliases ("image_create_aliases" API extension)
        {"name": "my-alias",
         "description": "A description"}
    ],
    "source": {
        "type": "instance",         // One of "instance" or "snapshot"
        "name": "abc"
    }
}
```
-->

リモートイメージの URL の場合は、次の dict を使用する必要があります。
<!--
In the remote image URL case, the following dict must be used:
-->

```js
{
    "filename": filename,                           // エクスポートの際に使用されます (任意で指定可能)
    "public":   true,                               // 信頼されないユーザーがイメージをダウンロードしてよいか (デフォルトは false)
    "properties": {                                 // イメージのプロパティ (任意で指定可能)
        "os": "Ubuntu"
    },
    "aliases": [                                    // 初期のエイリアスを設定します ("image_create_aliases" API 拡張)
        {"name": "my-alias",
         "description": "A description"}
    ],
    "source": {
        "type": "url",
        "url": "https://www.some-server.com/image"  // イメージの URL
    }
}
```

<!--
```js
{
    "filename": filename,                           // Used for export (optional)
    "public":   true,                               // Whether the image can be downloaded by untrusted users  (defaults to false)
    "properties": {                                 // Image properties (optional)
        "os": "Ubuntu"
    },
    "aliases": [                                    // Set initial aliases ("image_create_aliases" API extension)
        {"name": "my-alias",
         "description": "A description"}
    ],
    "source": {
        "type": "url",
        "url": "https://www.some-server.com/image"  // URL for the image
    }
}
```
-->

LXD が入力を受け取った後、バックグラウンド操作が開始され、イメージを
ストアに追加し、場合によってバックエンドファイルシステムに特有な
なんらかの最適化を行います。
<!--
After the input is received by LXD, a background operation is started
which will add the image to the store and possibly do some backend
filesystem-specific optimizations.
-->

### `/1.0/images/<fingerprint>`
#### GET (`?secret=SECRET` を任意に指定可能) <!-- GET (optional `?secret=SECRET`) -->
 * 説明: イメージの説明とメタデータ <!-- Description: Image description and metadata -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: イメージのプロパティを表す dict <!-- Return: dict representing an image properties -->

出力
<!--
Output:
-->

```json
{
    "aliases": [
        {
            "name": "bionic",
            "description": "",
        }
    ],
    "architecture": "x86_64",
    "auto_update": true,
    "cached": false,
    "fingerprint": "54c8caac1f61901ed86c68f24af5f5d3672bdc62c71d04f06df3a59e95684473",
    "filename": "ubuntu-bionic-18.04-amd64-server-20180201.tar.xz",
    "properties": {
        "architecture": "x86_64",
        "description": "Ubuntu 18.04 LTS server (20180601)",
        "os": "ubuntu",
        "release": "bionic"
    },
    "update_source": {
        "server": "https://10.1.2.4:8443",
        "protocol": "lxd",
        "certificate": "PEM certificate",
        "alias": "ubuntu/bionic/amd64"
    },
    "public": false,
    "size": 123792592,
    "created_at": "2016-02-01T21:07:41Z",
    "expires_at": "1970-01-01T00:00:00Z",
    "last_used_at": "1970-01-01T00:00:00Z",
    "uploaded_at": "2016-02-16T00:44:47Z"
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: イメージのプロパティを置き換えたり、情報や公開状態を変更します <!-- Description: Replaces the image properties, update information and visibility -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "auto_update": true,
    "properties": {
        "architecture": "x86_64",
        "description": "Ubuntu 18.04 LTS server (20180601)",
        "os": "ubuntu",
        "release": "bionic"
    },
    "public": true,
}
```

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: イメージのプロパティを変更したり、情報や公開状態を変更します <!-- Description: Updates the image properties, update information and visibility -->
 * 導入: `patch` API 拡張によって <!-- Introduced: with API extension `patch` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "properties": {
        "os": "ubuntu",
        "release": "bionic"
    },
    "public": true,
}
```

#### DELETE
 * 説明: イメージを削除します <!-- Description: Remove an image -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

この操作に対する HTTP ステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/images/<fingerprint>/export`
#### GET (`?secret=SECRET` を任意で指定可能) <!-- GET (optional `?secret=SECRET`) -->
 * 説明: イメージの tarball をダウンロードします <!-- Description: Download the image tarball -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 生のファイルまたは標準のエラー <!-- Return: Raw file or standard error -->

信頼されていない LXD が別の LXD に保管されている private なイメージから
新しいインスタンスを起動する場合は secret の文字列が必要です。
<!--
The secret string is required when an untrusted LXD is spawning a new
instance from a private image stored on a different LXD.
-->

2 つの LXD の間で信頼関係を要求する代わりに、クライアントは
`/1.0/images/<fingerprint>/export` に `POST` のリクエストを
送ってシークレットトークンを取得し、それをエクスポート先の LXD に
渡すことが出来ます。エクスポート先の LXD はシークレットトークンを
渡してイメージを guest として取得します。
<!--
Rather than require a trust relationship between the two LXDs, the
client will `POST` to `/1.0/images/<fingerprint>/export` to get a secret
token which it'll then pass to the target LXD. That target LXD will then
GET the image as a guest, passing the secret token.
-->

#### POST
 * 説明: イメージ tarball をアップロードします <!-- Description: Upload the image tarball -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

これはイメージを source から target にコピーします。
<!--
This copies an image from the source to the target.
-->

入力。
<!--
Input:
-->

```json
{
    "target": "target URL",
    "secret": "secret",
    "certificate": "target certificate",
    "aliases": ["alias"]
}
```


### `/1.0/images/<fingerprint>/refresh`
#### POST
 * 説明: イメージをオリジンからリフレッシュします <!-- Description: Refresh an image from its origin -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

これは指定されたイメージをオリジンからリフレッシュする操作を作成します。
<!--
This creates an operation to refresh the specified image from its origin.
-->

### `/1.0/images/<fingerprint>/secret`
#### POST
 * 説明: ランダムなトークンを生成し、ゲストが使用することを LXD に伝えます <!-- Description: Generate a random token and tell LXD to expect it be used by a guest -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```json
{
}
```

戻り値
<!--
Return:
-->

```json
{
    "secret": "52e9ec5885562aa24d05d7b4846ebb8b5f1f7bf5cd6e285639b569d9eaf54c9b"
}
```

メタデータ内の "secret" に生成されたシークレットの文字列が設定された
標準のバックグランド操作です。
<!--
Standard backround operation with "secret" set to the generated secret
string in metadata.
-->

シークレットはそれを使うイメージ URL にアクセスされてから 5 秒後に自動的に
無効化されます。これによりイメージの情報の取得とその後の /export の呼び出し
の両方に同じシークレットが使えます。
<!--
The secret is automatically invalidated 5s after an image URL using it
has been accessed. This allows to both retried the image information and
then hit /export with the same secret.
-->

### `/1.0/images/aliases`
#### GET
 * 説明: エイリアスの一覧 (イメージの公開状態に応じて public または private なエイリアスが含まれます) <!-- Description: list of aliases (public or private based on image visibility) -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: このサーバが知っているエイリアスの URL の一覧 <!-- Return: list of URLs for aliases this server knows about -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/images/aliases/sl6",
    "/1.0/images/aliases/bionic",
    "/1.0/images/aliases/xenial"
]
```

#### POST
 * 説明: 新しいエイリアスを作成します <!-- Description: create a new alias -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "description": "The alias description",
    "target": "SHA-256",
    "name": "alias-name"
}
```

### `/1.0/images/aliases/<name>`
#### GET
 * 説明: エイリアスの説明とターゲット <!-- Description: Alias description and target -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: エイリアスの説明とターゲットを表す dict <!-- Return: dict representing an alias description and target -->

出力
<!--
Output:
-->

```json
{
    "name": "test",
    "description": "my description",
    "target": "c9b6e738fae75286d52f497415463a8ecc61bbcb046536f220d797b0e500a41f"
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: エイリアスのターゲットまたは説明を置き換えます <!-- Description: Replaces the alias target or description -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "description": "New description",
    "target": "54c8caac1f61901ed86c68f24af5f5d3672bdc62c71d04f06df3a59e95684473"
}
```

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: エイリアスのターゲットまたは説明を更新します <!-- Description: Updates the alias target or description -->
 * 導入: `patch` API 拡張 によって <!-- Introduced: with API extension `patch` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "description": "New description"
}
```

#### POST
 * 説明: エイリアスをリネームします <!-- Description: rename an alias -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "new-name"
}
```

既に存在する名前にリネームしようとすると 409 (Conflict) の HTTP ステータスコードを返します。
<!--
Renaming to an existing name must return the 409 (Conflict) HTTP code.
-->

#### DELETE
 * 説明: エイリアスを削除します <!-- Description: Remove an alias -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

### `/1.0/networks`
#### GET
 * 説明: ネットワークの一覧 <!-- Description: list of networks -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ホストに現在定義されているネットワークの URL の一覧 <!-- Return: list of URLs for networks that are current defined on the host -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/networks/eth0",
    "/1.0/networks/lxdbr0"
]
```

#### POST
 * 説明: 新しいネットワークを定義します <!-- Description: define a new network -->
 * 導入: `network` API 拡張によって <!-- Introduced: with API extension `network` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "my-network",
    "description": "My network",
    "config": {
        "ipv4.address": "none",
        "ipv6.address": "2001:470:b368:4242::1/64",
        "ipv6.nat": "true"
    }
}
```

### `/1.0/networks/<name>`
#### GET
 * 説明: ネットワークについての情報 <!-- Description: information about a network -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻りt: ネットワークを表す dict <!-- Return: dict representing a network -->

戻り値
<!--
Return:
-->

```json
{
    "config": {},
    "name": "lxdbr0",
    "managed": false,
    "type": "bridge",
    "used_by": [
        "/1.0/instances/blah"
    ]
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: ネットワークの情報を置き換えます <!-- Description: replace the network information -->
 * 導入: `network` API 拡張によって <!-- Introduced: with API extension `network` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "bridge.driver": "openvswitch",
        "ipv4.address": "10.0.3.1/24",
        "ipv6.address": "fd1:6997:4939:495d::1/64"
    }
}
```

初期の作成や GET での情報取得結果と同じ dict です。ただし、
config のキーだけが使用され、それ以外の全てのキーは無視されます。
<!--
Same dict as used for initial creation and coming from GET. Only the
config is used, everything else is ignored.
-->

#### PATCH (ETag supported) <!-- PATCH (ETag supported) -->
 * 説明: ネットワークの情報を更新します。 <!-- Description: update the network information -->
 * 導入: `network` API 拡張によって <!-- Introduced: with API extension `network` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "dns.mode": "dynamic"
    }
}
```

#### POST
 * 説明: ネットワークをリネームします <!-- Description: rename a network -->
 * 導入: `network` API 拡張によって <!-- Introduced: with API extension `network` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (ネットワークをリネームします)
<!--
Input (rename a network):
-->

```json
{
    "name": "new-name"
}
```

HTTP ステータスコードは 204 (No content) で Location ヘッダは
リネーム後のリソースの URL を指します。
<!--
HTTP return value must be 204 (No content) and Location must point to
the renamed resource.
-->

既に存在する名前にリネームしようとすると 409 (Conflict) の HTTP ステータスコードを返します。
<!--
Renaming to an existing name must return the 409 (Conflict) HTTP code.
-->

#### DELETE
 * 説明: ネットワークを削除します <!-- Description: remove a network -->
 * 導入: `network` API 拡張によって <!-- Introduced: with API extension `network` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

この操作に対する HTTP ステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/networks/<name>/state`
#### GET
 * 説明: ネットワークの状態 <!-- Description: network state -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ネットワークの状態を表す dict <!-- Return: dict representing a network's state -->

戻り値
<!--
Return:
-->

```json
{
    "addresses": [
        {
            "family": "inet",
            "address": "10.87.252.1",
            "netmask": "24",
            "scope": "global"
        },
        {
            "family": "inet6",
            "address": "fd42:6e0e:6542:a212::1",
            "netmask": "64",
            "scope": "global"
        },
        {
            "family": "inet6",
            "address": "fe80::3419:9ff:fe9b:f9aa",
            "netmask": "64",
            "scope": "link"
        }
    ],
    "counters": {
        "bytes_received": 0,
        "bytes_sent": 17724,
        "packets_received": 0,
        "packets_sent": 95
    },
    "hwaddr": "36:19:09:9b:f9:aa",
    "mtu": 1500,
    "state": "up",
    "type": "broadcast"
}
```

### `/1.0/operations`
#### GET
 * 説明: 操作の一覧 <!-- Description: list of operations -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 状態ごとの現在実行中またはキューに追加された操作の URL の一覧を表す dict <!-- Return: dict representing a list of URLs for operations that are currently going on/queued according to their status -->

戻り値
<!--
Return:
-->

```json
{
    "success": [
        "/1.0/operations/c0fc0d0d-a997-462b-842b-f8bd0df82507"
    ],
    "running": [
        "/1.0/operations/092a8755-fd90-4ce4-bf91-9f87d03fd5bc"
    ]
}
```

### `/1.0/operations/<uuid>`
#### GET
 * 説明: バックグラウンド操作 <!-- Description: background operation -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: バックグラウンド操作を表す dict <!-- Return: dict representing a background operation -->

戻り値
<!--
Return:
-->

```js
{
    "id": "b8d84888-1dc2-44fd-b386-7f679e171ba5",
    "class": "token",                                                                       // "task" (バックグラウンドのタスク), "websocket" (ウェブソケットと認証情報の組) あるいは "token" (一時的な認証情報)
    "created_at": "2016-02-17T16:59:27.237628195-05:00",                                    // 作成時のタイムスタンプ
    "updated_at": "2016-02-17T16:59:27.237628195-05:00",                                    // 最終更新時のタイムスタンプ
    "status": "Running",
    "status_code": 103,
    "resources": {                                                                          // 影響を受けるリソースの一覧
        "images": [
            "/1.0/images/54c8caac1f61901ed86c68f24af5f5d3672bdc62c71d04f06df3a59e95684473"
        ]
    },
    "metadata": {                                                                           // 操作についての追加情報 (action, target, ...)
        "secret": "c9209bee6df99315be1660dd215acde4aec89b8e5336039712fc11008d918b0d"
    },
    "may_cancel": true,                                                                     // (DELETE で) 操作をキャンセルできるかどうか
    "err": ""
}
```

<!--
```js
{
    "id": "b8d84888-1dc2-44fd-b386-7f679e171ba5",
    "class": "token",                                                                       // One of "task" (background task), "websocket" (set of websockets and crendentials) or "token" (temporary credentials)
    "created_at": "2016-02-17T16:59:27.237628195-05:00",                                    // Creation timestamp
    "updated_at": "2016-02-17T16:59:27.237628195-05:00",                                    // Last update timestamp
    "status": "Running",
    "status_code": 103,
    "resources": {                                                                          // List of affected resources
        "images": [
            "/1.0/images/54c8caac1f61901ed86c68f24af5f5d3672bdc62c71d04f06df3a59e95684473"
        ]
    },
    "metadata": {                                                                           // Extra information about the operation (action, target, ...)
        "secret": "c9209bee6df99315be1660dd215acde4aec89b8e5336039712fc11008d918b0d"
    },
    "may_cancel": true,                                                                     // Whether it's possible to cancel the operation (DELETE)
    "err": ""
}
```
-->

#### DELETE
 * 説明: 操作をキャンセルします。この API を呼び出すとエントリーを実際に削除するのではなく状態を "cancelling" に変更します <!-- Description: cancel an operation. Calling this will change the state to "cancelling" rather than actually removing the entry. -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

この操作に対する HTTP ステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/operations/<uuid>/wait`
#### GET (`?timeout=30` を任意で指定可能) <!-- GET (optional `?timeout=30`) -->
 * 説明: 操作が完了するのを待ちます <!-- Description: Wait for an operation to finish -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 操作が最終の状態に達した後の操作の dict <!-- Return: dict of the operation after it's reached its final state -->

入力 (最終の状態まで無限に待つ場合): 引数なし
<!--
Input (wait indefinitely for a final state): no argument
-->

入力 (同様に最終の状態まで待つが 30 秒後にタイムアウトする場合): ?timeout=30
<!--
Input (similar but times out after 30s): ?timeout=30
-->

### `/1.0/operations/<uuid>/websocket`
#### GET (`?secret=SECRET`)
 * 説明: この接続はウェブソケットの接続にアップグレードされ、操作の種別毎に
   定義されたプロトコルを話します。例えば exec 操作の場合、ウェブソケットは
   標準入力／標準出力／標準エラー出力のための双方向のパイプになり、インスタンス
   内のプロセスの入出力をやりとりします。 migration の場合はマイグレーション
   情報を通信するプライマリのインターフェースになります。ここで使用する
   シークレットは操作を作るときに指定していたのと同じものを指定します。
   正しいシークレットを指定すればゲストもこのエンドポイントに接続できます。
   <!-- Description: This connection is upgraded into a websocket connection
   speaking the protocol defined by the operation type. For example, in the
   case of an exec operation, the websocket is the bidirectional pipe for
   stdin/stdout/stderr to flow to and from the process inside the instance.
   In the case of migration, it will be the primary interface over which the
   migration information is communicated. The secret here is the one that was
   provided when the operation was created. Guests are allowed to connect
   provided they have the right secret. -->
 * 認証: guest または trusted <!-- Authentication: guest or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ウェブソケットのストリームまたは標準のエラー <!-- Return: websocket stream or standard error -->

### `/1.0/profiles`
#### GET
 * 説明: 設定プロファイルの一覧 <!-- Description: List of configuration profiles -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 定義されているプロファイルの URL の一覧 <!-- Return: list of URLs to defined profiles -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/profiles/default"
]
```

#### POST
 * 説明: 新しいプロファイルを定義します <!-- Description: define a new profile -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "my-profilename",
    "description": "Some description string",
    "config": {
        "limits.memory": "2GB"
    },
    "devices": {
        "kvm": {
            "type": "unix-char",
            "path": "/dev/kvm"
        }
    }
}
```

### `/1.0/profiles/<name>`
#### GET
 * 説明: プロファイルの設定 <!-- Description: profile configuration -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: プロファイルの設定を表す dict <!-- Return: dict representing the profile content -->

出力
<!--
Output:
-->

```json
{
    "name": "test",
    "description": "Some description string",
    "config": {
        "limits.memory": "2GB"
    },
    "devices": {
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        }
    },
    "used_by": [
        "/1.0/instances/blah"
    ]
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: プロファイルの情報を置き換えます <!-- Description: replace the profile information -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "limits.memory": "4GB"
    },
    "description": "Some description string",
    "devices": {
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        }
    }
}
```

初期の作成や GET での情報取得結果と同じ dict です。
name プロパティは変更できません (変更するには POST を参照してください)。
<!--
Same dict as used for initial creation and coming from GET. The name
property can't be changed (see POST for that).
-->

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: プロファイルの情報を変更します <!-- Description: update the profile information -->
 * 導入: `patch` API 拡張によって <!-- Introduced: with API extension `patch` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "limits.memory": "4GB"
    },
    "description": "Some description string",
    "devices": {
        "kvm": {
            "path": "/dev/kvm",
            "type": "unix-char"
        }
    }
}
```

#### POST
 * 説明: プロファイルをリネームします <!-- Description: rename a profile -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (プロファイルをリネームします)
<!--
Input (rename a profile):
-->

```json
{
    "name": "new-name"
}
```

HTTP ステータスコードは 204 (No content) で Location ヘッダは
リネーム後のリソースの URL を指します。
<!--
HTTP return value must be 204 (No content) and Location must point to
the renamed resource.
-->

既に存在する名前にリネームしようとすると 409 (Conflict) の HTTP ステータスコードを返します。
<!--
Renaming to an existing name must return the 409 (Conflict) HTTP code.
-->

#### DELETE
 * 説明: プロファイルを削除します <!-- Description: remove a profile -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

この操作に対する HTTP ステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

`default` プロファイルを削除しようとすると 403 (Forbidden) という HTTP ステータスコードが返ります。
<!--
Attempting to delete the `default` profile will return the 403 (Forbidden) HTTP code.
-->

### `/1.0/projects`
#### GET
 * 説明: プロジェクトの一覧 <!-- Description: List of projects -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 定義済みのプロジェクトの URL の一覧 <!-- Return: list of URLs to defined projects -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/projects/default"
]
```

#### POST
 * 説明: 新しいプロジェクトを定義します <!-- Description: define a new project -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "test",
    "config": {
        "features.images": "true",
        "features.profiles": "true",
    },
    "description": "Some description string"
}
```

### `/1.0/projects/<name>`
#### GET
 * 説明: プロジェクトの設定 <!-- Description: project configuration -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: プロジェクトの内容を表す dict <!-- Return: dict representing the project content -->

出力
<!--
Output:
-->

```json
{
    "name": "test",
    "config": {
        "features.images": "true",
        "features.profiles": "true",
    },
    "description": "Some description string",
    "used_by": [
        "/1.0/instances/blah"
    ]
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: プロジェクトの情報を置き換えます <!-- Description: replace the project information -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "features.images": "true",
        "features.profiles": "true",
    },
    "description": "Some description string"
}
```

初期作成への入力と GET の戻り値は同じ形式の dict が使用されます。
name プロパティは変更できません (その用途には POST を参照してください)。
<!--
Same dict as used for initial creation and coming from GET. The name
property can't be changed (see POST for that).
-->

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: プロジェクトの情報を更新します <!-- Description: replace the project information -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "features.images": "true",
    },
    "description": "Some description string"
}
```

#### POST
 * 説明: プロジェクトをリネームします <!-- Description: rename a project -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (プロジェクトをリネームします)
<!--
Input (rename a project):
-->

```json
{
    "name": "new-name"
}
```

HTTP ステータスコードは 204 (No content) で Location ヘッダは
リネーム後のリソースの URL を指します。
<!--
HTTP return value must be 204 (No content) and Location must point to
the renamed resource.
-->

既に存在する名前にリネームしようとすると 409 (Conflict) の HTTP ステータスコードを返します。
<!--
Renaming to an existing name must return the 409 (Conflict) HTTP code.
-->

`default` プロジェクトをリネームしようとすると 403 (Forbidden) という HTTP ステータスコードが返ります。
<!--
Attempting to rename the `default` project will return the 403 (Forbidden) HTTP code.
-->

#### DELETE
 * 説明: プロジェクトを削除します <!-- Description: remove a project -->
 * 導入: `projects` API 拡張によって <!-- Introduced: with API extension `projects` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

この操作に対する HTTP ステータスコードは 202 (Accepted) です。
<!--
HTTP code for this should be 202 (Accepted).
-->

`default` プロジェクトを削除しようとすると 403 (Forbidden) という HTTP ステータスコードが返ります。
<!--
Attempting to delete the `default` project will return the 403 (Forbidden) HTTP code.
-->

### `/1.0/storage-pools`
#### GET
 * 説明: ストレージプールの一覧 <!-- Description: list of storage pools -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ホストに現在定義されているストレージプールの一覧 <!-- Return: list of storage pools that are currently defined on the host -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/storage-pools/default",
    "/1.0/storage-pools/pool1",
    "/1.0/storage-pools/pool2",
    "/1.0/storage-pools/pool3",
    "/1.0/storage-pools/pool4"
]
```

#### POST
 * 説明: 新しいストレージプールを作成します <!-- Description: create a new storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "size": "10GB"
    },
    "driver": "zfs",
    "name": "pool1"
}
```

### `/1.0/storage-pools/<name>`
#### GET
 * 説明: ストレージプールの情報 <!-- Description: information about a storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ストレージプールを表す dict <!-- Return: dict representing a storage pool -->

戻り値
<!--
Return:
-->

```json
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "operation": "",
    "error_code": 0,
    "error": "",
    "metadata": {
        "name": "default",
        "driver": "zfs",
        "used_by": [
            "/1.0/instances/alp1",
            "/1.0/instances/alp10",
            "/1.0/instances/alp11",
            "/1.0/instances/alp12",
            "/1.0/instances/alp13",
            "/1.0/instances/alp14",
            "/1.0/instances/alp15",
            "/1.0/instances/alp16",
            "/1.0/instances/alp17",
            "/1.0/instances/alp18",
            "/1.0/instances/alp19",
            "/1.0/instances/alp2",
            "/1.0/instances/alp20",
            "/1.0/instances/alp3",
            "/1.0/instances/alp4",
            "/1.0/instances/alp5",
            "/1.0/instances/alp6",
            "/1.0/instances/alp7",
            "/1.0/instances/alp8",
            "/1.0/instances/alp9",
            "/1.0/images/62e850a334bb9d99cac00b2e618e0291e5e7bb7db56c4246ecaf8e46fa0631a6"
        ],
        "config": {
            "size": "61203283968",
            "source": "/home/chb/mnt/l2/disks/default.img",
            "volume.size": "0",
            "zfs.pool_name": "default"
        }
    }
}
```

#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: ストレージプールの情報を置き換えます <!-- Description: replace the storage pool information -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
 Input:
-->

```json
{
    "config": {
        "size": "15032385536",
        "source": "pool1",
        "volume.block.filesystem": "xfs",
        "volume.block.mount_options": "discard",
        "lvm.thinpool_name": "LXDThinPool",
        "lvm.vg_name": "pool1",
        "volume.size": "10737418240"
    }
}
```

#### PATCH
 * 説明: ストレージプールの設定を変更します <!-- Description: update the storage pool configuration -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {
        "volume.block.filesystem": "xfs",
    }
}
```

#### DELETE
 * 説明: ストレージプールを削除します <!-- Description: delete a storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

### `/1.0/storage-pools/<name>/resources`
#### GET
 * 説明: ストレージプールで利用可能なリソースに関する情報 <!-- Description: information about the resources available to the storage pool -->
 * 導入: `resources` API 拡張によって <!-- Introduced: with API extension `resources` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ストレージプールのリソースを表す dict <!-- Return: dict representing the storage pool resources -->

戻り値
<!--
Return:
-->

```json
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "operation": "",
    "error_code": 0,
    "error": "",
    "metadata": {
        "space": {
            "used": 207111192576,
            "total": 306027577344
        },
        "inodes": {
            "used": 3275333,
            "total": 18989056
        }
    }
}
```


### `/1.0/storage-pools/<name>/volumes`
#### GET
 * 説明: ストレージボリュームの一覧 <!-- Description: list of storage volumes -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 指定されたストレージプール上に現在存在するストレージボリュームの一覧 <!-- Return: list of storage volumes that currently exist on a given storage pool -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/storage-pools/default/volumes/container/alp1",
    "/1.0/storage-pools/default/volumes/container/alp10",
    "/1.0/storage-pools/default/volumes/container/alp11",
    "/1.0/storage-pools/default/volumes/container/alp12",
    "/1.0/storage-pools/default/volumes/container/alp13",
    "/1.0/storage-pools/default/volumes/container/alp14",
    "/1.0/storage-pools/default/volumes/container/alp15",
    "/1.0/storage-pools/default/volumes/container/alp16",
    "/1.0/storage-pools/default/volumes/container/alp17",
    "/1.0/storage-pools/default/volumes/container/alp18",
    "/1.0/storage-pools/default/volumes/container/alp19",
    "/1.0/storage-pools/default/volumes/container/alp2",
    "/1.0/storage-pools/default/volumes/container/alp20",
    "/1.0/storage-pools/default/volumes/container/alp3",
    "/1.0/storage-pools/default/volumes/container/alp4",
    "/1.0/storage-pools/default/volumes/container/alp5",
    "/1.0/storage-pools/default/volumes/container/alp6",
    "/1.0/storage-pools/default/volumes/container/alp7",
    "/1.0/storage-pools/default/volumes/container/alp8",
    "/1.0/storage-pools/default/volumes/container/alp9",
    "/1.0/storage-pools/default/volumes/image/62e850a334bb9d99cac00b2e618e0291e5e7bb7db56c4246ecaf8e46fa0631a6"
]
```

#### POST
 * 説明: 指定されたストレージプール上に新しいストレージボリュームを作成します <!-- Description: create a new storage volume on a given storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期または (既存のボリュームをコピーする場合は) 非同期 <!-- Operation: sync or async (when copying an existing volume) -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {},
    "name": "vol1",
    "type": "custom"
}
```

入力 (ボリュームをコピーする場合)
<!--
Input (when copying a volume):
-->

```json
{
    "config": {},
    "name": "vol1",
    "type": "custom",
    "source": {
        "pool": "pool2",
        "name": "vol2",
        "type": "copy"
    }
}
```

入力 (ボリュームをマイグレーションする場合)
<!--
Input (when migrating a volume):
-->

```js
{
    "config": {},
    "name": "vol1",
    "type": "custom"
    "source": {
        "pool": "pool2",
        "name": "vol2",
        "type": "migration",
        "mode": "pull"        // "pull" (デフォルト), "push", "relay" のいずれか
    }
}
```

<!--
```js
{
    "config": {},
    "name": "vol1",
    "type": "custom"
    "source": {
        "pool": "pool2",
        "name": "vol2",
        "type": "migration",
        "mode": "pull"        // One of "pull" (default), "push", "relay"
    }
}
```
-->

### `/1.0/storage-pools/<pool>/volumes/<type>`
#### POST
 * 説明: 指定のストレージプール上に特定のタイプのストレージボリュームを作成します <!-- Description: create a new storage volume of a particular type on a given storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期または (既存のボリュームをコピーする場合は) 非同期 <!-- Operation: sync or async (when copying an existing volume) -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "config": {},
    "name": "vol1",
}
```

入力 (ボリュームをコピーする場合)
<!--
Input (when copying a volume):
-->

```json
{
    "config": {},
    "name": "vol1",
    "source": {
        "pool": "pool2",
        "name": "vol2",
        "type": "copy"
    }
}
```

入力 (ボリュームをマイグレーションする場合)
<!--
Input (when migrating a volume):
-->

```js
{
    "config": {},
    "name": "vol1",
    "source": {
        "pool": "pool2",
        "name": "vol2",
        "type": "migration"
        "mode": "pull",      // "pull" (デフォルト), "push", "relay" のいずれか
    }
}
```

<!--
```js
{
    "config": {},
    "name": "vol1",
    "source": {
        "pool": "pool2",
        "name": "vol2",
        "type": "migration"
        "mode": "pull",      // One of "pull" (default), "push", "relay"
    }
}
```
-->

### `/1.0/storage-pools/<pool>/volumes/<type>/<name>`
#### POST
 * 説明: 指定のストレージプール上のストレージボリュームをリネームします <!-- Description: rename a storage volume on a given storage pool -->
 * 導入: `storage_api_volume_rename` API 拡張によって <!-- Introduced: with API extension `storage_api_volume_rename` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期または (別のプールに移動する場合は) 非同期 <!-- Operation: sync or async (when moving to a different pool) -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "vol1",
    "pool": "pool3"
}
```

入力 (lxd インスタンス間でマイグレーションする場合)
<!--
Input (migration across lxd instances):
-->

```json
{
    "name": "vol1",
    "pool": "pool3",
    "migration": true
}
```

誰か (つまり他の lxd インスタンス) が全てのウェブソケットに接続してソースと
交渉を始めるまでは、マイグレーションは実際には開始されません。
<!--
The migration does not actually start until someone (i.e. another lxd instance)
connects to all the websockets and begins negotiation with the source.
-->

出力のメタデータのセクション (マイグレーションの場合)
<!--
Output in metadata section (for migration):
-->

```js
{
    "control": "secret1",       // マイグレーション制御ソケット
    "fs": "secret2"             // ファイルシステム転送ソケット
}
```

<!--
```js
{
    "control": "secret1",       // Migration control socket
    "fs": "secret2"             // Filesystem transfer socket
}
```
-->

これらは作成の呼び出し時に渡すべきシークレットです。
<!--
These are the secrets that should be passed to the create call.
-->

#### GET
 * 説明: ストレージプール上の指定のタイプのストレージボリュームの情報 <!-- Description: information about a storage volume of a given type on a storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ストレージボリュームを表す dict <!-- Return: dict representing a storage volume -->

戻り値
<!--
Return:
-->

```json
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error": "",
    "metadata": {
        "type": "custom",
        "used_by": [],
        "name": "vol1",
        "config": {
            "block.filesystem": "ext4",
            "block.mount_options": "discard",
            "size": "10737418240"
        }
    }
}
```


#### PUT (ETag サポートあり) <!-- PUT (ETag supported) -->
 * 説明: ストレージボリュームの情報を置き換えるかスナップショットから復元します <!-- Description: replace the storage volume information or restore from snapshot -->
 * 導入: `storage`, `storage_api_volume_snapshots` API 拡張によって <!-- Introduced: with API extension `storage`, `storage_api_volume_snapshots` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
 Input:
-->

(訳注: ストレージボリュームの情報を置き換える場合の入力)

```json
{
    "config": {
        "size": "15032385536",
        "source": "pool1",
        "used_by": "",
        "volume.block.filesystem": "xfs",
        "volume.block.mount_options": "discard",
        "lvm.thinpool_name": "LXDThinPool",
        "lvm.vg_name": "pool1",
        "volume.size": "10737418240"
    }
}
```

(訳注: スナップショットから復元する場合の入力)

```json
{
    "restore": "snapshot-name"
}
```

#### PATCH (ETag サポートあり) <!-- PATCH (ETag supported) -->
 * 説明: ストレージボリュームの情報を変更します <!-- Description: update the storage volume information -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
 Input:
-->

```json
{
    "config": {
        "volume.block.mount_options": "",
    }
}
```

#### DELETE
 * 説明: 指定したストレージプール上の指定したタイプのストレージボリュームを削除します <!-- Description: delete a storage volume of a given type on a given storage pool -->
 * 導入: `storage` API 拡張によって <!-- Introduced: with API extension `storage` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```

### `/1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots`
#### GET
 * 説明: ボリュームスナップショットの一覧 <!-- Description: List of volume snapshots -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: このボリュームのスナップショットの URL の一覧 <!-- Return: list of URLs for snapshots for this volume -->

戻り値
<!--
Return value:
-->

```json
[
    "/1.0/storage-pools/default/volumes/custom/foo/snapshots/snap0"
]
```

#### POST
 * 説明: 新規のボリュームスナップショットを作成する <!-- Description: create a new volume snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```js
{
    "name": "my-snapshot",          // スナップショットの名前
}
```

<!--
```js
{
    "name": "my-snapshot",          // Name of the snapshot
}
```
-->

### `/1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/name`
#### GET
 * 説明: スナップショットの情報 <!-- Description: Snapshot information -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: スナップショットを表す dict <!-- Return: dict representing the snapshot -->

戻り値
<!--
Return:
-->

```json
{
    "config": {},
    "description": "",
    "name": "snap0"
}
```

#### PUT
 * 説明: ボリュームスナップショットの情報 <!-- Description: Volume snapshot information -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: ボリュームスナップショットを表す dict <!-- Return: dict representing the volume snapshot -->

入力
<!--
Input:
-->

```json
{
    "description": "new-description"
}
```

#### POST
 * 説明: ボリュームスナップショットをリネームするのに使用されます <!-- Description: used to rename the volume snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力
<!--
Input:
-->

```json
{
    "name": "new-name"
}
```

#### DELETE
 * 説明: ボリュームスナップショットを削除します <!-- Description: remove the volume snapshot -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

レスポンスの HTTP ステータスコードは 202 (Accepted)。
<!--
HTTP code for this should be 202 (Accepted).
-->

### `/1.0/resources`
#### GET
 * 説明: LXD サーバで利用可能なリソースに関する情報 <!-- Description: information about the resources available to the LXD server -->
 * 導入: `resources` API 拡張によって <!-- Introduced: with API extension `resources` -->
 * 認証: guest, untrusted または trusted <!-- Authentication: guest, untrusted or trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: システムリソースを表す dict <!-- Return: dict representing the system resources -->

戻り値
<!--
Return:
-->

```json
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "operation": "",
    "error_code": 0,
    "error": "",
    "metadata": {
        "cpu": {
            "sockets": [
               {
                   "cores": 2,
                   "frequency": 2691,
                   "frequency_turbo": 3400,
                   "name": "GenuineIntel",
                   "vendor": "Intel(R) Core(TM) i5-3340M CPU @ 2.70GHz",
                   "threads": 4
               }
            ],
            "total": 4
        },
        "memory": {
            "used": 4454240256,
            "total": 8271765504
        }
    }
}
```

### `/1.0/cluster`
#### GET
 * 説明: クラスタの情報 (ネットワークやストレージプールなど) <!-- Description: information about a cluster (such as networks and storage pools) -->
 * 導入: `clustering` API 拡張によって <!-- Introduced: with API extension `clustering` -->
 * 認証: trusted または untrusted <!-- Authentication: trusted or untrusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: クラスタを表す dict <!-- Return: dict representing a cluster -->

戻り値
<!--
Return:
-->

```json
{
    "server_name": "node1",
    "enabled": true,
    "member_config": [
        {
            "entity": "storage-pool",
            "name": "local",
            "key": "source",
            "description": "\"source\" property for storage pool \"local\"",
        },
        {
            "entity": "network",
            "name": "lxdbr0",
            "key": "bridge.external_interfaces",
            "description": "\"bridge.external_interfaces\" property for network \"lxdbr0\"",
        },
    ],
}
```

#### PUT
 * 説明: このノード上のクラスタをブートストラップしたりクラスタに参加したり、クラスタリングを無効にします <!-- Description: bootstrap or join a cluster, or disable clustering on this node -->
 * 導入: `clustering` API 拡張によって <!-- Introduced: with API extension `clustering` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期または非同期 <!-- Operation: sync or async -->
 * 戻り値: 入力によってさまざまなペイロード <!-- Return: various payloads depending on the input -->

入力 (新しいクラスタをブートストラップする場合)
<!--
Input (bootstrap a new cluster):
-->

```json
{
    "server_name": "lxd1",
    "enabled": true,
}
```

バックグラウンド操作か標準のエラーを返します。
<!--
Return background operation or standard error.
-->

入力 (既存のクラスタに参加を依頼する場合)
<!--
Input (request to join an existing cluster):
-->

```json
{
    "server_name": "node2",
    "server_address": "10.1.1.102:8443",
    "enabled": true,
    "cluster_address": "10.1.1.101:8443",
    "cluster_certificate": "-----BEGIN CERTIFICATE-----MIFf\n-----END CERTIFICATE-----",
    "cluster_password": "sekret",
    "member_config": [
        {
            "entity": "storage-pool",
            "name": "local",
            "key": "source",
            "value": "/dev/sdb"
        },
        {
            "entity": "network",
            "name": "lxdbr0",
            "key": "bridge.external_interfaces",
            "value": "vlan0"
        }
    ]
}
```

入力 (このノード上のクラスタリングを無効にする場合)
<!--
Input (disable clustering on the node):
-->

```json
{
    "enabled": false
}
```

### `/1.0/cluster/members`
#### GET
 * 説明: クラスタ内の LXD メンバの一覧 <!-- Description: list of LXD members in the cluster -->
 * 導入: `clustering` API 拡張によって <!-- Introduced: with API extension `clustering` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: クラスタメンバの一覧 <!-- Return: list of cluster members -->

戻り値
<!--
Return:
-->

```json
[
    "/1.0/cluster/members/lxd1",
    "/1.0/cluster/members/lxd2"
]
```

### `/1.0/cluster/members/<name>`
#### GET
 * 説明: メンバの情報と状態を取得します <!-- Description: retrieve the member's information and status -->
 * 導入: `clustering` API 拡張によって <!-- Introduced: with API extension `clustering` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: メンバを表す dict <!-- Return: dict representing the member -->

戻り値
<!--
Return:
-->

```json
{
    "server_name": "lxd1",
    "url": "https://10.1.1.101:8443",
    "database": true,
    "status": "Online",
    "message":"fully operational"
}
```

#### POST
 * 説明: クラスタメンバをリネームします <!-- Description: rename a cluster member -->
 * 導入: `clustering` API 拡張によって <!-- Introduced: with API extension `clustering` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 同期 <!-- Operation: sync -->
 * 戻り値: 標準の戻り値または標準のエラー <!-- Return: standard return value or standard error -->

入力
<!--
Input:
-->

```json
{
    "server_name": "node1"
}
```

#### DELETE (`?force=1` を任意で指定可能) <!-- DELETE (optional `?force=1`) -->
 * 説明: クラスタのメンバを削除します <!-- Description: remove a member of the cluster -->
 * 導入: `clustering` API 拡張によって <!-- Introduced: with API extension `clustering` -->
 * 認証: trusted <!-- Authentication: trusted -->
 * 操作: 非同期 <!-- Operation: async -->
 * 戻り値: バックグラウンド操作または標準のエラー <!-- Return: background operation or standard error -->

入力 (現在は何もなし)
<!--
Input (none at present):
-->

```json
{
}
```
