# sing-boxYF custom protocol compatibility

This document is the implementation contract for the FlClashYF markers and wire
protocols. A marker is not considered supported until the corresponding wire
handshake, configuration validation, and interoperability test are present.

## Compatibility matrix

| Marker / feature | FlClashYF implementation | sing-boxYF integration point | Status |
| --- | --- | --- | --- |
| `#x365` | VLESS custom request/response header in `core/mihomo/transport/vless/conn.go` | VLESS client handshake decorator | pending |
| `#juzi` | VLESS request HMAC extension in `core/mihomo/transport/vless/conn.go` | VLESS client handshake decorator | pending |
| `#pure` | VLESS TLS 1.3-only TCP header and response prefix in `core/mihomo/transport/vless/pure.go` | VLESS client handshake decorator; reject UDP | pending |
| `blackstone` | Custom XHTTP payload decryption/validation in `adapter/outbound/xhttp.go` | Separate custom outbound; never hide it inside generic XHTTP | pending |
| standard XHTTP | `transport/xhttp/{config,client,conn,reuse,upload_queue,xpadding}.go` | VLESS transport variant | pending |
| `#vt` / ViewTurbo | local `sing-shadowsocks2` method wrapper | Shadowsocks method registration or dedicated fork dependency | pending |
| `#fastup` | Trojan password derivation plus forced H2 mux | Trojan option normalization and mux policy | pending |
| OPPA | `transport/oppa` and `adapter/outbound/oppa.go` | New outbound, option, registry, packet connection | pending |

## Required boundaries

- `#x365`, `#juzi`, and `#pure` are mutually exclusive VLESS marker suffixes.
  The parser must reject malformed combinations instead of silently falling
  back to standard VLESS.
- Standard XHTTP and the custom Blackstone dynamic protocol are different
  protocols. The latter must not be implemented by adding API fetching or
  private payload decryption to the generic VLESS transport.
- `#fastup` is a Trojan credential derivation marker, not a new wire protocol.
  The marker must never be sent to the server as part of the password.
- ViewTurbo must be registered only for its exact method/password contract;
  ordinary Shadowsocks must continue to use the upstream method implementation.
- No live credentials, private keys, API endpoints, or embedded production
  secrets may be copied into sing-boxYF. Runtime interoperability fixtures use
  injected CI secrets or synthetic values.

## Acceptance gates

For each entry above, CI must run:

1. option parsing and schema validation;
2. malformed-marker and unsupported-network tests;
3. deterministic wire/golden tests with synthetic identities;
4. local mock-server interoperability tests;
5. Linux Go tests before the Apple jobs;
6. iOS arm64 libbox compilation through the existing Apple workflow.

A successful Apple build proves compilation and linkage only. It does not prove
server interoperability; that requires the mock/live protocol tests separately.
