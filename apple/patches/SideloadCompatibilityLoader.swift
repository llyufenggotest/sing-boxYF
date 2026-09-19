import Darwin
import Foundation
import Libbox
import Library

/// Loads the approved sideload dylib from the process-local Frameworks folder.
/// The loader is compiled into both SFI and the packet-tunnel extension.
enum SideloadCompatibilityLoader {
    private static let name = "SingBoxYFSideload.dylib"
    private static var handle: UnsafeMutableRawPointer?

    static func load() {
        guard handle == nil else { return }
        let path = Bundle.main.privateFrameworksPath.map { URL(fileURLWithPath: $0).appendingPathComponent(name).path }
            ?? "@rpath/\(name)"
        handle = dlopen(path, RTLD_NOW | RTLD_LOCAL)
        if handle == nil {
            NSLog("sideload compatibility dylib unavailable")
        }
    }
}
