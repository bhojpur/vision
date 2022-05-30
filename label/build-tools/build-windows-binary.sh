#!/bin/bash

# Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

### Window requires pyinstall v2.1
wine msiexec -i python-2.7.8.msi
wine pywin32-218.win32-py2.7.exe
wine PyQt4-4.11.4-gpl-Py2.7-Qt4.8.7-x32.exe
wine lxml-3.7.3.win32-py2.7.exe

THIS_SCRIPT_PATH=`readlink -f $0`
THIS_SCRIPT_DIR=`dirname ${THIS_SCRIPT_PATH}`
cd pyinstaller
git checkout v2.1
cd ${THIS_SCRIPT_DIR}
echo ${THIS_SCRIPT_DIR}

#. venv_wine/bin/activate
rm -r build
rm -r dist
rm visionlab.spec

wine c:/Python27/python.exe pyinstaller/pyinstaller.py --hidden-import=xml \
            --hidden-import=xml.etree \
            --hidden-import=xml.etree.ElementTree \
            --hidden-import=lxml.etree \
             -D -F -n visionlab -c "../visionlab.py" -p ../libs -p ../

FOLDER=$(git describe --abbrev=0 --tags)
FOLDER="windows_"$FOLDER
rm -rf "$FOLDER"
mkdir "$FOLDER"
cp dist/visionlab.exe $FOLDER
cp -rf ../data $FOLDER/data
zip "$FOLDER.zip" -r $FOLDER