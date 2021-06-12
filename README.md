# geolite2lookup

[GeoLite2 Free Geolocation Data](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) を利用して、国の判定を行う

ライブラリとしては、 [IncSW/geoip2](https://github.com/IncSW/geoip2) を利用している

## 使い方

```
Usage:
  geolite2lookup <ipaddress> [flags]

Flags:
  -d, --dir string    MMDB direcotry (default "/usr/share/GeoIP2")
  -f, --file string   MMDB filename (default "GeoLite2-[type].mmdb")
  -h, --help          help for geolite2lookup
  -t, --type string   MMDB Edition (default "Country")
  -v, --version       version for geolite2lookup
```

## License

[MIT License](LICENSE).
