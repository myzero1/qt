package designer

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads
#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_DESIGNER_LIB -DQT_CORE_LIB -DQT_GUI_LIB -DQT_WIDGETS_LIB -DQT_UIPLUGIN_LIB -DQT_DESIGNERCOMPONENTS_LIB
#cgo CXXFLAGS: -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include -IC:/Qt/Qt5.7.0/5.7/mingw53_32/mkspecs/win32-g++
#cgo CXXFLAGS: -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/QtDesigner -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/QtCore -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/QtGui -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/QtWidgets -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/QtUiPlugin -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/QtDesignerComponents

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition
#cgo LDFLAGS: -LC:/Qt/Qt5.7.0/5.7/mingw53_32/lib -lQt5Designer -lQt5Core -lQt5Gui -lQt5Widgets -lQt5DesignerComponents -lmingw32 -lqtmain -lshell32
*/
import "C"
