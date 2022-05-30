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

# Packaging and Release
docker run --workdir=$(pwd)/ --volume="/home/$USER:/home/$USER" bhojpur/py2qt4 /bin/sh -c 'make qt4py2; make test;sudo python setup.py sdist;sudo python setup.py install'

while true; do
    read -p "Do you wish to deploy this to PyPI(twine upload dist/* or pip install dist/*)?" yn
    case $yn in
        [Yy]* ) docker run -it --rm --workdir=$(pwd)/ --volume="/home/$USER:/home/$USER" bhojpur/py2qt4; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
done
# python setup.py register
# python setup.py sdist upload
# Net pypi: twine upload dist/*

# Test before upladoing: pip install dist/visionlab.tar.gz