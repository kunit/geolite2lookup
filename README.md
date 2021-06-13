# geolite2lookup

[GeoLite2 Free Geolocation Data](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) を利用して、国もしくは都市の判定を行う

ライブラリとしては、 [IncSW/geoip2](https://github.com/IncSW/geoip2) を利用している

## 使い方

```
look up country/city using IP Address or hostname

Usage:
  geolite2lookup <ipaddress|hostname> [flags]

Flags:
  -d, --dir string    MMDB directory (default "/usr/share/GeoIP2")
  -f, --file string   MMDB filename (default "GeoLite2-[type].mmdb")
  -h, --help          help for geolite2lookup
  -i, --info          show additional information (only type "City")
  -t, --type string   MMDB Edition (default "Country")
  -v, --version       version for geolite2lookup
```

## 環境変数

| 環境変数名 | 説明 | デフォルト値 |
| --- | --- | --- |
| `GEOLITE2LOOKUP_MMDB_TYPE` | 参照する MMDB のタイプ | `Country` |
| `GEOLITE2LOOKUP_MMDB_DIR` | MMDB ファイルが配置してあるディレクトリ | `/usr/share/GeoIP2` |

## インストール方法

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export GEOLITE2LOOKUP_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/kunit/geolite2lookup/releases/download/v$GEOLITE2LOOKUP_VERSION/geolite2lookup_$GEOLITE2LOOKUP_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export GEOLITE2LOOKUP_VERSION=X.X.X
$ yum install https://github.com/kunit/geolite2lookup/releases/download/v$GEOLITE2LOOKUP_VERSION/geolite2lookup_$GEOLITE2LOOKUP_VERSION-1_amd64.rpm
```

**apk:**

Use [apk-add-from-url](https://github.com/k1LoW/apk-add-from-url)

``` console
$ export GEOLITE2LOOKUP_VERSION=X.X.X
$ curl -L https://git.io/apk-add-from-url | sh -s -- https://github.com/kunit/geolite2lookup/releases/download/v$GEOLITE2LOOKUP_VERSION/geolite2lookup_$GEOLITE2LOOKUP_VERSION-1_amd64.apk
```

**homebrew tap:**

```console
$ brew install kunit/tap/geolite2lookup
```

**manually:**

Download binary from [releases page](https://github.com/kunit/geolite2lookup/releases)

**go get:**

```console
$ go get github.com/kunit/geolite2lookup
```

## License

[MIT License](LICENSE).
