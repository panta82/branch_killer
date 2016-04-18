import sys
from PyQt5.QtWidgets import QApplication, QWidget

from ui.main_window import MainWindow

if __name__ == "__main__":
	application = QApplication(sys.argv)

	win = MainWindow()
	win.show()

	sys.exit(application.exec_())
