# Lighouse results Prometheus exporter

Simple tool for parsing Lighthouse results (JSON) producing Prometheus metrics
for `score` and `numericValue`.

## example

```bash
go run cmd/main.go -file ./test.json | curl --data-binary @- http://localhost:9091/metrics/job/some_job
```

## output

```
# TYPE audit_score gauge
audit_score{id="largest-contentful-paint",host="test.xyz"} 0.00
audit_score{id="first-meaningful-paint",host="test.xyz"} 0.26
audit_score{id="plugins",host="test.xyz"} 1.00
audit_score{id="bf-cache",host="test.xyz"} 0.00
audit_score{id="third-party-summary",host="test.xyz"} 1.00
audit_score{id="total-byte-weight",host="test.xyz"} 0.85
audit_score{id="first-contentful-paint",host="test.xyz"} 0.11
audit_score{id="errors-in-console",host="test.xyz"} 1.00
audit_score{id="html-has-lang",host="test.xyz"} 1.00
audit_score{id="legacy-javascript",host="test.xyz"} 0.88
audit_score{id="cumulative-layout-shift",host="test.xyz"} 0.06
audit_score{id="uses-rel-preconnect",host="test.xyz"} 0.74
audit_score{id="maskable-icon",host="test.xyz"} 0.00
audit_score{id="mainthread-work-breakdown",host="test.xyz"} 0.33
audit_score{id="paste-preventing-inputs",host="test.xyz"} 1.00
audit_score{id="bootup-time",host="test.xyz"} 0.49
audit_score{id="unused-javascript",host="test.xyz"} 0.18
audit_score{id="notification-on-start",host="test.xyz"} 1.00
audit_score{id="aria-valid-attr",host="test.xyz"} 1.00
audit_score{id="unminified-css",host="test.xyz"} 1.00
audit_score{id="aria-allowed-attr",host="test.xyz"} 1.00
audit_score{id="is-on-https",host="test.xyz"} 1.00
audit_score{id="color-contrast",host="test.xyz"} 0.00
audit_score{id="lcp-lazy-loaded",host="test.xyz"} 1.00
audit_score{id="themed-omnibox",host="test.xyz"} 0.00
audit_score{id="unminified-javascript",host="test.xyz"} 1.00
audit_score{id="modern-image-formats",host="test.xyz"} 0.58
audit_score{id="efficient-animated-content",host="test.xyz"} 1.00
audit_score{id="dom-size",host="test.xyz"} 1.00
audit_score{id="no-document-write",host="test.xyz"} 1.00
audit_score{id="max-potential-fid",host="test.xyz"} 0.01
audit_score{id="no-unload-listeners",host="test.xyz"} 1.00
audit_score{id="html-lang-valid",host="test.xyz"} 1.00
audit_score{id="uses-responsive-images",host="test.xyz"} 0.75
audit_score{id="is-crawlable",host="test.xyz"} 1.00
audit_score{id="uses-http2",host="test.xyz"} 1.00
audit_score{id="link-text",host="test.xyz"} 1.00
audit_score{id="interactive",host="test.xyz"} 0.18
audit_score{id="installable-manifest",host="test.xyz"} 0.00
audit_score{id="meta-viewport",host="test.xyz"} 0.00
audit_score{id="uses-long-cache-ttl",host="test.xyz"} 0.02
audit_score{id="inspector-issues",host="test.xyz"} 1.00
audit_score{id="deprecations",host="test.xyz"} 1.00
audit_score{id="speed-index",host="test.xyz"} 0.26
audit_score{id="prioritize-lcp-image",host="test.xyz"} 1.00
audit_score{id="uses-text-compression",host="test.xyz"} 1.00
audit_score{id="image-size-responsive",host="test.xyz"} 1.00
audit_score{id="tap-targets",host="test.xyz"} 1.00
audit_score{id="render-blocking-resources",host="test.xyz"} 0.43
audit_score{id="duplicated-javascript",host="test.xyz"} 1.00
audit_score{id="robots-txt",host="test.xyz"} 1.00
audit_score{id="button-name",host="test.xyz"} 1.00
audit_score{id="document-title",host="test.xyz"} 1.00
audit_score{id="total-blocking-time",host="test.xyz"} 0.15
audit_score{id="content-width",host="test.xyz"} 1.00
audit_score{id="aria-hidden-body",host="test.xyz"} 1.00
audit_score{id="aria-hidden-focus",host="test.xyz"} 1.00
audit_score{id="tabindex",host="test.xyz"} 1.00
audit_score{id="uses-optimized-images",host="test.xyz"} 0.75
audit_score{id="hreflang",host="test.xyz"} 1.00
audit_score{id="font-size",host="test.xyz"} 1.00
audit_score{id="crawlable-anchors",host="test.xyz"} 1.00
audit_score{id="duplicate-id-active",host="test.xyz"} 1.00
audit_score{id="heading-order",host="test.xyz"} 1.00
audit_score{id="link-name",host="test.xyz"} 0.00
audit_score{id="unused-css-rules",host="test.xyz"} 0.75
audit_score{id="doctype",host="test.xyz"} 1.00
audit_score{id="service-worker",host="test.xyz"} 0.00
audit_score{id="geolocation-on-start",host="test.xyz"} 1.00
audit_score{id="third-party-facades",host="test.xyz"} 0.00
audit_score{id="aria-valid-attr-value",host="test.xyz"} 1.00
audit_score{id="http-status-code",host="test.xyz"} 1.00
audit_score{id="splash-screen",host="test.xyz"} 0.00
audit_score{id="unsized-images",host="test.xyz"} 1.00
audit_score{id="offscreen-images",host="test.xyz"} 0.94
audit_score{id="font-display",host="test.xyz"} 1.00
audit_score{id="charset",host="test.xyz"} 1.00
audit_score{id="meta-description",host="test.xyz"} 1.00
audit_score{id="server-response-time",host="test.xyz"} 1.00
audit_score{id="valid-source-maps",host="test.xyz"} 0.00
audit_score{id="image-alt",host="test.xyz"} 1.00
audit_score{id="redirects",host="test.xyz"} 1.00
audit_score{id="image-aspect-ratio",host="test.xyz"} 1.00
audit_score{id="uses-passive-event-listeners",host="test.xyz"} 1.00
audit_score{id="viewport",host="test.xyz"} 1.00
# TYPE audit_numeric_value gauge
audit_numeric_value{id="largest-contentful-paint",host="test.xyz"} 9496.79
audit_numeric_value{id="first-meaningful-paint",host="test.xyz"} 5194.23
audit_numeric_value{id="total-byte-weight",host="test.xyz"} 2919076.00
audit_numeric_value{id="first-contentful-paint",host="test.xyz"} 4853.87
audit_numeric_value{id="network-server-latency",host="test.xyz"} 28.65
audit_numeric_value{id="metrics",host="test.xyz"} 11508.00
audit_numeric_value{id="legacy-javascript",host="test.xyz"} 150.00
audit_numeric_value{id="cumulative-layout-shift",host="test.xyz"} 0.74
audit_numeric_value{id="uses-rel-preconnect",host="test.xyz"} 314.18
audit_numeric_value{id="mainthread-work-breakdown",host="test.xyz"} 5004.48
audit_numeric_value{id="bootup-time",host="test.xyz"} 3529.84
audit_numeric_value{id="unused-javascript",host="test.xyz"} 3480.00
audit_numeric_value{id="unminified-css",host="test.xyz"} 0.00
audit_numeric_value{id="unminified-javascript",host="test.xyz"} 0.00
audit_numeric_value{id="modern-image-formats",host="test.xyz"} 600.00
audit_numeric_value{id="efficient-animated-content",host="test.xyz"} 0.00
audit_numeric_value{id="dom-size",host="test.xyz"} 451.00
audit_numeric_value{id="max-potential-fid",host="test.xyz"} 760.00
audit_numeric_value{id="uses-responsive-images",host="test.xyz"} 300.00
audit_numeric_value{id="uses-http2",host="test.xyz"} 0.00
audit_numeric_value{id="interactive",host="test.xyz"} 11507.56
audit_numeric_value{id="installable-manifest",host="test.xyz"} 1.00
audit_numeric_value{id="uses-long-cache-ttl",host="test.xyz"} 1318503.12
audit_numeric_value{id="speed-index",host="test.xyz"} 7553.94
audit_numeric_value{id="prioritize-lcp-image",host="test.xyz"} 0.00
audit_numeric_value{id="uses-text-compression",host="test.xyz"} 0.00
audit_numeric_value{id="render-blocking-resources",host="test.xyz"} 1355.00
audit_numeric_value{id="duplicated-javascript",host="test.xyz"} 0.00
audit_numeric_value{id="total-blocking-time",host="test.xyz"} 1453.74
audit_numeric_value{id="uses-optimized-images",host="test.xyz"} 300.00
audit_numeric_value{id="unused-css-rules",host="test.xyz"} 300.00
audit_numeric_value{id="offscreen-images",host="test.xyz"} 70.00
audit_numeric_value{id="server-response-time",host="test.xyz"} 98.55
audit_numeric_value{id="network-rtt",host="test.xyz"} 7.57
audit_numeric_value{id="redirects",host="test.xyz"} 0.00
```
