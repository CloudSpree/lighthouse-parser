# Lighouse results Prometheus exporter

Simple tool for parsing Lighthouse results (JSON) producing Prometheus metrics
for audit score and numericValue and category score.

## example

```bash
go run cmd/main.go -file ./test.json -commit 940819b | curl --data-binary @- http://localhost:9091/metrics/job/some_job
```

The best way to use this tool is to push metrics for the give URL to the same category e.g. `job/some_job`. Using multiple instances might introduce the duplicates hence avoid stuff like `job/some_job/instance/commit/${COMMIT_SHA}`.

# prefix

If needed, you can supply `-prefix` flag e.g. `-prefix=lighthouse`. 
Resulting metrics will be like `lighthouse_audit_score` then.

## output

```
audit_score{id="first-contentful-paint",host="xyz.co.uk",commit="940819b"} 0.11
audit_score{id="total-blocking-time",host="xyz.co.uk",commit="940819b"} 0.15
audit_score{id="third-party-summary",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="hreflang",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="is-on-https",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="image-size-responsive",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="aria-valid-attr-value",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="aria-valid-attr",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="bootup-time",host="xyz.co.uk",commit="940819b"} 0.49
audit_score{id="aria-hidden-focus",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="largest-contentful-paint",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="maskable-icon",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="tabindex",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="uses-responsive-images",host="xyz.co.uk",commit="940819b"} 0.75
audit_score{id="tap-targets",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="unsized-images",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="uses-optimized-images",host="xyz.co.uk",commit="940819b"} 0.75
audit_score{id="lcp-lazy-loaded",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="uses-http2",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="bf-cache",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="html-has-lang",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="unused-javascript",host="xyz.co.uk",commit="940819b"} 0.18
audit_score{id="robots-txt",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="unminified-css",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="modern-image-formats",host="xyz.co.uk",commit="940819b"} 0.58
audit_score{id="font-size",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="plugins",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="content-width",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="uses-text-compression",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="errors-in-console",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="legacy-javascript",host="xyz.co.uk",commit="940819b"} 0.88
audit_score{id="notification-on-start",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="valid-source-maps",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="paste-preventing-inputs",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="meta-description",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="font-display",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="meta-viewport",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="geolocation-on-start",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="link-name",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="uses-long-cache-ttl",host="xyz.co.uk",commit="940819b"} 0.02
audit_score{id="efficient-animated-content",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="dom-size",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="crawlable-anchors",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="render-blocking-resources",host="xyz.co.uk",commit="940819b"} 0.43
audit_score{id="first-meaningful-paint",host="xyz.co.uk",commit="940819b"} 0.26
audit_score{id="cumulative-layout-shift",host="xyz.co.uk",commit="940819b"} 0.06
audit_score{id="server-response-time",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="viewport",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="button-name",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="document-title",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="link-text",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="max-potential-fid",host="xyz.co.uk",commit="940819b"} 0.01
audit_score{id="duplicate-id-active",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="html-lang-valid",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="themed-omnibox",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="mainthread-work-breakdown",host="xyz.co.uk",commit="940819b"} 0.33
audit_score{id="third-party-facades",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="prioritize-lcp-image",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="total-byte-weight",host="xyz.co.uk",commit="940819b"} 0.85
audit_score{id="unminified-javascript",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="inspector-issues",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="image-alt",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="installable-manifest",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="aria-allowed-attr",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="color-contrast",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="offscreen-images",host="xyz.co.uk",commit="940819b"} 0.94
audit_score{id="uses-rel-preconnect",host="xyz.co.uk",commit="940819b"} 0.74
audit_score{id="redirects",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="splash-screen",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="deprecations",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="service-worker",host="xyz.co.uk",commit="940819b"} 0.00
audit_score{id="interactive",host="xyz.co.uk",commit="940819b"} 0.18
audit_score{id="no-unload-listeners",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="unused-css-rules",host="xyz.co.uk",commit="940819b"} 0.75
audit_score{id="speed-index",host="xyz.co.uk",commit="940819b"} 0.26
audit_score{id="aria-hidden-body",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="no-document-write",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="heading-order",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="doctype",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="is-crawlable",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="uses-passive-event-listeners",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="image-aspect-ratio",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="duplicated-javascript",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="charset",host="xyz.co.uk",commit="940819b"} 1.00
audit_score{id="http-status-code",host="xyz.co.uk",commit="940819b"} 1.00
# TYPE audit_numeric_value gauge
audit_numeric_value{id="first-contentful-paint",host="xyz.co.uk",commit="940819b"} 4853.87
audit_numeric_value{id="total-blocking-time",host="xyz.co.uk",commit="940819b"} 1453.74
audit_numeric_value{id="bootup-time",host="xyz.co.uk",commit="940819b"} 3529.84
audit_numeric_value{id="largest-contentful-paint",host="xyz.co.uk",commit="940819b"} 9496.79
audit_numeric_value{id="uses-responsive-images",host="xyz.co.uk",commit="940819b"} 300.00
audit_numeric_value{id="uses-optimized-images",host="xyz.co.uk",commit="940819b"} 300.00
audit_numeric_value{id="uses-http2",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="unused-javascript",host="xyz.co.uk",commit="940819b"} 3480.00
audit_numeric_value{id="unminified-css",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="modern-image-formats",host="xyz.co.uk",commit="940819b"} 600.00
audit_numeric_value{id="uses-text-compression",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="legacy-javascript",host="xyz.co.uk",commit="940819b"} 150.00
audit_numeric_value{id="uses-long-cache-ttl",host="xyz.co.uk",commit="940819b"} 1318503.12
audit_numeric_value{id="efficient-animated-content",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="dom-size",host="xyz.co.uk",commit="940819b"} 451.00
audit_numeric_value{id="render-blocking-resources",host="xyz.co.uk",commit="940819b"} 1355.00
audit_numeric_value{id="first-meaningful-paint",host="xyz.co.uk",commit="940819b"} 5194.23
audit_numeric_value{id="cumulative-layout-shift",host="xyz.co.uk",commit="940819b"} 0.74
audit_numeric_value{id="server-response-time",host="xyz.co.uk",commit="940819b"} 98.55
audit_numeric_value{id="max-potential-fid",host="xyz.co.uk",commit="940819b"} 760.00
audit_numeric_value{id="metrics",host="xyz.co.uk",commit="940819b"} 11508.00
audit_numeric_value{id="mainthread-work-breakdown",host="xyz.co.uk",commit="940819b"} 5004.48
audit_numeric_value{id="prioritize-lcp-image",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="total-byte-weight",host="xyz.co.uk",commit="940819b"} 2919076.00
audit_numeric_value{id="unminified-javascript",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="network-rtt",host="xyz.co.uk",commit="940819b"} 7.57
audit_numeric_value{id="installable-manifest",host="xyz.co.uk",commit="940819b"} 1.00
audit_numeric_value{id="offscreen-images",host="xyz.co.uk",commit="940819b"} 70.00
audit_numeric_value{id="uses-rel-preconnect",host="xyz.co.uk",commit="940819b"} 314.18
audit_numeric_value{id="redirects",host="xyz.co.uk",commit="940819b"} 0.00
audit_numeric_value{id="interactive",host="xyz.co.uk",commit="940819b"} 11507.56
audit_numeric_value{id="unused-css-rules",host="xyz.co.uk",commit="940819b"} 300.00
audit_numeric_value{id="speed-index",host="xyz.co.uk",commit="940819b"} 7553.94
audit_numeric_value{id="network-server-latency",host="xyz.co.uk",commit="940819b"} 28.65
audit_numeric_value{id="duplicated-javascript",host="xyz.co.uk",commit="940819b"} 0.00
# TYPE category_score gauge
category_score{id="seo",host="xyz.co.uk",commit="940819b"} 1.00
category_score{id="pwa",host="xyz.co.uk",commit="940819b"} 0.33
category_score{id="performance",host="xyz.co.uk",commit="940819b"} 0.10
category_score{id="accessibility",host="xyz.co.uk",commit="940819b"} 0.83
category_score{id="best-practices",host="xyz.co.uk",commit="940819b"} 1.00
```
