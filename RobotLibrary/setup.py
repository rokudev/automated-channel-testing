import setuptools

setuptools.setup(
    name="RokuRobotLibrary",
    version="2.2.0",
    packages=['Library'],
    install_requires=[
          'robotframework==3.1.2',
          'requests==2.22.0'
      ]
)