import Foundation
import Library

class PacketTunnelProvider: ExtensionProvider {
    override func startTunnel(options: [String: NSObject]?, completionHandler: @escaping (Error?) -> Void) {
        SideloadCompatibilityLoader.load()
        super.startTunnel(options: options, completionHandler: completionHandler)
    }
}
