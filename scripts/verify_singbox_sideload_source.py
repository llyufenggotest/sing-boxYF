#!/usr/bin/env python3
from pathlib import Path

source = Path("ios/SideloadSupport/SingBoxYFSideload.m").read_text()
required = (
    "containerURLForSecurityApplicationGroupIdentifier:",
    "bundleProxyForCurrentProcess",
    "groupContainerURLs",
    "method_setImplementation",
    "NSApplicationSupportDirectory",
    "sing-boxYF-sideload",
)
for marker in required:
    if marker not in source:
        raise SystemExit(f"missing sideload compatibility marker: {marker}")
if "NSUserDefaults" in source or "SecItem" in source:
    raise SystemExit("sing-boxYF dylib must not hook unrelated defaults or keychain APIs")
print("SING_BOX_YF_SIDELOAD_SOURCE_CONTRACT_PASS")
