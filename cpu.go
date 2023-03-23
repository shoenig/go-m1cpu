//go:build darwin && arm64

package m1cpu

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// #include <CoreFoundation/CoreFoundation.h>
// #include <IOKit/IOKitLib.h>
//
// #define HzToGHz(hz) ((hz) / 1000000000.0)
//
// UInt64 global_pCoreHz;
// UInt64 global_eCoreHz;
//
// UInt64 getFrequency(CFTypeRef typeRef) {
// CFDataRef cfData = typeRef;
//
// CFIndex size = CFDataGetLength(cfData);
// UInt8 buf[size];
// CFDataGetBytes(cfData, CFRangeMake(0, size), buf);
//
// UInt8 b1 = buf[size-7];
// UInt8 b2 = buf[size-6];
// UInt8 b3 = buf[size-5];
// UInt8 b4 = buf[size-4];
//
// UInt64 pCoreHz = 0x00000000FFFFFFFF & ((b1<<24) | (b2 << 16) | (b3 << 8) | (b4));
// return pCoreHz;
// }
//
// void initialize() {
//   CFMutableDictionaryRef matching = IOServiceMatching("AppleARMIODevice");
//   io_iterator_t  iter;
//   IOServiceGetMatchingServices(kIOMainPortDefault, matching, &iter);
//
//   io_object_t obj;
//   while ((obj = IOIteratorNext(iter))) {
//     char class[512];
//     IOObjectGetClass(obj, class);
//     char name[512];
//     IORegistryEntryGetName(obj, name);
//
//     if (strcmp(name, "pmgr") == 0) {
//       CFTypeRef pCoreRef = IORegistryEntryCreateCFProperty(obj, CFSTR("voltage-states5-sram"), kCFAllocatorDefault, 0);
//       CFTypeRef eCoreRef = IORegistryEntryCreateCFProperty(obj, CFSTR("voltage-states1-sram"), kCFAllocatorDefault, 0);
//
//       long long pCoreHz = getFrequency(pCoreRef);
//       long long eCoreHz = getFrequency(eCoreRef);
//
//       global_pCoreHz = pCoreHz;
//       global_eCoreHz = eCoreHz;
//       return;
//     }
//   }
// }
//
// UInt64 eCoreHz() {
//   return global_eCoreHz;
// }
//
// UInt64 pCoreHz() {
//   return global_pCoreHz;
// }
//
// Float64 eCoreGHz() {
//   return HzToGHz(global_eCoreHz);
// }
//
// Float64 pCoreGHz() {
//   return HzToGHz(global_pCoreHz);
// }
import "C"

func init() {
	C.initialize()
}

// IsAppleSilicon returns true on this platform.
func IsAppleSilicon() bool {
	return true
}

// PCoreHZ returns the max frequency in Hertz of the P-Core of an Apple Silicon CPU.
func PCoreHz() uint64 {
	return uint64(C.pCoreHz())
}

// ECoreHZ returns the max frequency in Hertz of the E-Core of an Apple Silicon CPU.
func ECoreHz() uint64 {
	return uint64(C.eCoreHz())
}

// PCoreGHz returns the max frequency in Gigahertz of the P-Core of an Apple Silicon CPU.
func PCoreGHz() float64 {
	return float64(C.pCoreGHz())
}

// ECoreGHz returns the max frequency in Gigahertz of the E-Core of an Apple Silicon CPU.
func ECoreGHz() float64 {
	return float64(C.eCoreGHz())
}
