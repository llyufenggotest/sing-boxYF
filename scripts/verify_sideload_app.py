#!/usr/bin/env python3
"""Check unsigned SFI's host and packet-tunnel payload before packaging."""

import argparse
import hashlib
import plistlib
import struct
from pathlib import Path

DYLIB = 'Tg_@HelloWorld_1024.dylib'
EXPECTED_HASH = 'cd903ea15657cbd356398adcb60c8872c41c29b69acc1a5dfb78a49d6e75dea5'


def rpaths(data):
    if len(data) < 32 or struct.unpack_from('<I', data)[0] != 0xFEEDFACF:
        raise ValueError('expected arm64 Mach-O binary')
    count = struct.unpack_from('<I', data, 16)[0]
    offset = 32
    paths = []
    for _ in range(count):
        cmd, size = struct.unpack_from('<II', data, offset)
        if size < 12 or offset + size > len(data):
            raise ValueError('invalid Mach-O load command')
        if cmd == 0x8000001C:  # LC_RPATH
            start = struct.unpack_from('<I', data, offset + 8)[0]
            paths.append(data[offset + start:offset + size].split(b'\0')[0].decode())
        offset += size
    return paths


def verify(app):
    extensions = sorted(app.rglob('*.appex'))
    tunnel = app / 'PlugIns' / 'Extension.appex'
    if extensions != [tunnel]:
        raise ValueError(f'expected only the packet-tunnel extension: {extensions}')
    binaries = [(app, app / 'sing-box'), (tunnel, tunnel / 'Extension')]
    for bundle, binary in binaries:
        info = plistlib.loads((bundle / 'Info.plist').read_bytes())
        if info['CFBundleExecutable'] != binary.name:
            raise ValueError(f'wrong executable in {bundle}')
        data = binary.read_bytes()
        if DYLIB.encode() not in data or b'_dlopen' not in data:
            raise ValueError(f'missing dylib loader in {binary}')
        paths = rpaths(data)
        if not paths:
            raise ValueError(f'missing rpaths in {binary}')
        if bundle == app and '@executable_path/Frameworks' not in paths:
            raise ValueError(f'host frameworks rpath missing in {binary}')
        library = bundle / 'Frameworks' / DYLIB
        if hashlib.sha256(library.read_bytes()).hexdigest() != EXPECTED_HASH:
            raise ValueError(f'dylib hash mismatch in {library}')
        print(f'PASS {binary.name}: dylib verified; rpaths={paths}')
    host = (app / 'sing-box').read_bytes()
    if b'SideloadCompatibilityLoader' not in host or b'ExtensionEnvironments' not in host:
        raise ValueError('host startup entry missing')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('app', type=Path)
    verify(parser.parse_args().app)
