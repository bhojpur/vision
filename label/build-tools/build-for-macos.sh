#!/bin/sh

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

brew install python@2
pip install --upgrade virtualenv

# clone Bhojpur Vision source
rm -rf /tmp/bhojpurSetup
mkdir /tmp/bhojpurSetup
cd /tmp/bhojpurSetup
curl https://codeload.github.com/bhojpur/vision/zip/master --output bhojpur.zip
unzip bhojpur.zip
rm bhojpur.zip

# setup python3 space
virtualenv --system-site-packages  -p python3 /tmp/bhojpurSetup/visionlab-py3
source /tmp/bhojpurSetup/visionlab-py3/bin/activate
cd visionlab-master

# build Bhojpur Vision app
pip install py2app
pip install PyQt5 lxml
make qt5py3
rm -rf build dist
python setup.py py2app -A
mv "/tmp/bhojpurSetup/visionlab-master/dist/visionlab.app" /Applications
# deactivate python3
deactivate
cd ../
rm -rf /tmp/bhojpurSetup
echo 'DONE'