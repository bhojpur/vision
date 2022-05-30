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

import os, sys
import glob
from PIL import Image

def txtLabel_to_xmlLabel(classes_file,source_pth,save_xml_pth):
    if not os.path.exists(save_xml_pth):
        os.makedirs(save_xml_pth)
    classes = open(classes_file).read().splitlines()
    print(classes)
    for file in os.listdir(source_pth):
        if not file.endswith(('.bmp', '.dib', '.png', '.jpg', '.jpeg', '.pbm', '.pgm', '.ppm', '.tif', '.tiff', '.JPG')):
            continue
        img_file = Image.open(os.path.join(source_pth,file))
        txt_file = open(os.path.join(source_pth,file.replace('.JPG','.txt').replace('.jpg','.txt').replace('.png','.txt').replace('.PNG','.txt'))).read().splitlines()
        print(txt_file)
        xml_file = open(os.path.join(save_xml_pth,file.replace('.JPG','.xml').replace('.jpg','.xml').replace('.png','.xml').replace('.PNG','.xml')), 'w')
        width, height = img_file.size
        xml_file.write('<annotation>\n')
        xml_file.write('\t<folder>simple</folder>\n')
        xml_file.write('\t<filename>' + str(file) + '</filename>\n')
        xml_file.write('\t<size>\n')
        xml_file.write('\t\t<width>' + str(width) + ' </width>\n')
        xml_file.write('\t\t<height>' + str(height) + '</height>\n')
        xml_file.write('\t\t<depth>' + str(3) + '</depth>\n')
        xml_file.write('\t</size>\n')
        for line in txt_file:
            print(line)
            line_split = line.split(' ')
            x_center = float(line_split[1])
            y_center = float(line_split[2])
            w = float(line_split[3])
            h = float(line_split[4])
            xmax = int((2*x_center*width + w*width)/2)
            xmin = int((2*x_center*width - w*width)/2)
            ymax = int((2*y_center*height + h*height)/2)
            ymin = int((2*y_center*height - h*height)/2)

            xml_file.write('\t<object>\n')
            xml_file.write('\t\t<name>'+ str(classes[int(line_split[0])]) +'</name>\n')
            xml_file.write('\t\t<pose>Unspecified</pose>\n')
            xml_file.write('\t\t<truncated>0</truncated>\n')
            xml_file.write('\t\t<difficult>0</difficult>\n')
            xml_file.write('\t\t<bndbox>\n')
            xml_file.write('\t\t\t<xmin>' + str(xmin) + '</xmin>\n')
            xml_file.write('\t\t\t<ymin>' + str(ymin) + '</ymin>\n')
            xml_file.write('\t\t\t<xmax>' + str(xmax) + '</xmax>\n')
            xml_file.write('\t\t\t<ymax>' + str(ymax) + '</ymax>\n')
            xml_file.write('\t\t</bndbox>\n')
            xml_file.write('\t</object>\n')
        xml_file.write('</annotation>')

#
# if __name__ == '__main__':
#     classes_file = r"/home/saiki/Documents/test-help/classes.txt"
#     txtLabel_to_xmlLabel(r'/home/saiki/Documents/test-help/',r'/home/saiki/Documents/test-help/xml/')