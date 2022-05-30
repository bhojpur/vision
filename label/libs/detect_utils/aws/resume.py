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

# Resume all interrupted trainings in yolov5/ dir including DDP trainings
# Usage: $ python utils/aws/resume.py

import os
import sys
from pathlib import Path

import torch
import yaml

sys.path.append('./')  # to run '$ python *.py' files in subdirectories

port = 0  # --master_port
path = Path('').resolve()
for last in path.rglob('*/**/last.pt'):
    ckpt = torch.load(last)
    if ckpt['optimizer'] is None:
        continue

    # Load opt.yaml
    with open(last.parent.parent / 'opt.yaml') as f:
        opt = yaml.safe_load(f)

    # Get device count
    d = opt['device'].split(',')  # devices
    nd = len(d)  # number of devices
    ddp = nd > 1 or (nd == 0 and torch.cuda.device_count() > 1)  # distributed data parallel

    if ddp:  # multi-GPU
        port += 1
        cmd = f'python -m torch.distributed.launch --nproc_per_node {nd} --master_port {port} train.py --resume {last}'
    else:  # single-GPU
        cmd = f'python train.py --resume {last}'

    cmd += ' > /dev/null 2>&1 &'  # redirect output to dev/null and run in daemon thread
    print(cmd)
    os.system(cmd)