#import <Foundation/Foundation.h>
#import <objc/message.h>
#import <objc/runtime.h>

static NSURL *(*originalContainerURL)(id, SEL, NSString *);

static NSURL *fallbackContainerURL(NSString *identifier) {
    Class proxyClass = NSClassFromString(@"LSBundleProxy");
    SEL currentSelector = NSSelectorFromString(@"bundleProxyForCurrentProcess");
    if (proxyClass && [proxyClass respondsToSelector:currentSelector]) {
        id proxy = ((id (*)(id, SEL))objc_msgSend)(proxyClass, currentSelector);
        SEL groupsSelector = NSSelectorFromString(@"groupContainerURLs");
        if (proxy && [proxy respondsToSelector:groupsSelector]) {
            NSDictionary *groups = ((id (*)(id, SEL))objc_msgSend)(proxy, groupsSelector);
            NSURL *url = [groups objectForKey:identifier];
            if (url) {
                return url;
            }
            if (groups.count == 1) {
                return groups.allValues.firstObject;
            }
        }
    }

    NSArray<NSURL *> *support = [[NSFileManager defaultManager]
        URLsForDirectory:NSApplicationSupportDirectory
        inDomains:NSUserDomainMask];
    NSURL *base = support.firstObject;
    if (!base) {
        return nil;
    }
    NSURL *url = [base URLByAppendingPathComponent:@"sing-boxYF-sideload" isDirectory:YES];
    [[NSFileManager defaultManager] createDirectoryAtURL:url
                              withIntermediateDirectories:YES
                                               attributes:nil
                                                    error:nil];
    return url;
}

static NSURL *sideloadContainerURL(id self, SEL selector, NSString *identifier) {
    NSURL *url = originalContainerURL ? originalContainerURL(self, selector, identifier) : nil;
    return url ?: fallbackContainerURL(identifier);
}

__attribute__((constructor)) static void installSideloadCompatibility(void) {
    Class fileManager = [NSFileManager class];
    SEL selector = @selector(containerURLForSecurityApplicationGroupIdentifier:);
    Method method = class_getInstanceMethod(fileManager, selector);
    if (!method) {
        return;
    }
    originalContainerURL = (void *)method_getImplementation(method);
    method_setImplementation(method, (IMP)sideloadContainerURL);
}
