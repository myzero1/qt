package location

/*
#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_POSITIONING_LIB -DQT_QML_LIB -DQT_QUICK_LIB -DQT_LOCATION_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/gcc_64/include -I/usr/local/Qt5.7.0/5.7/gcc_64/mkspecs/linux-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtCore -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtGui -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtNetwork -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtPositioning -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtQml -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtQuick -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtLocation

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/gcc_64 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/gcc_64/lib
#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/gcc_64/lib -L/usr/lib64 -lQt5Core -lQt5Gui -lQt5Network -lQt5Positioning -lQt5Qml -lQt5Quick -lQt5Location -lpthread
*/
import "C"
