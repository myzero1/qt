package macextras

/*
#cgo CFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS9.3.sdk -miphoneos-version-min=7.0 -arch armv7 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC
#cgo CXXFLAGS: -pipe -stdlib=libc++ -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -fvisibility-inlines-hidden -Wno-non-virtual-dtor -Wno-overloaded-virtual -Wno-exit-time-destructors -O2 -std=gnu++11 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS9.3.sdk -miphoneos-version-min=7.0 -arch armv7 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC
#cgo CXXFLAGS: -DDARWIN_NO_CARBON -DQT_NO_PRINTER -DQT_NO_PRINTDIALOG -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_MACEXTRAS_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/ios/mkspecs/macx-ios-clang/ios -I/usr/local/Qt5.7.0/5.7/ios/include -I/usr/local/Qt5.7.0/5.7/ios/mkspecs/macx-ios-clang
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/ios/include/QtCore -I/usr/local/Qt5.7.0/5.7/ios/include/QtGui -I/usr/local/Qt5.7.0/5.7/ios/include/QtMacExtras

#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS9.3.sdk -miphoneos-version-min=7.0 -arch armv7
#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/ios/plugins/platforms -lqios -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L/usr/local/Qt5.7.0/5.7/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype -framework Security -framework SystemConfiguration -framework CoreBluetooth -L/usr/local/Qt5.7.0/5.7/ios/plugins/imageformats -lqdds -lqicns -lqico -lqtga -lqtiff -lqwbmp -lqwebp -lqtharfbuzzng -lz -lqtpcre -lm -lQt5Core -lQt5Gui -lQt5MacExtras -lQt5PlatformSupport
*/
import "C"
