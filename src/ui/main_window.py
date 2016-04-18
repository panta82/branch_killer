import sys
from PyQt5.QtWidgets import QMainWindow, QStatusBar


# noinspection PyArgumentList
class MainWindow(QMainWindow):
	def __init__(self):
		super().__init__()
		self.configure()
		pass

	def configure(self):
		self.setGeometry(300, 300, 300, 220)
		self.setWindowTitle('Hello!')
		sb = self.statusBar() # type: QStatusBar
		sb.showMessage("Ready")
