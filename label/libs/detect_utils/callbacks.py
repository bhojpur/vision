#!/usr/bin/env python

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

class Callbacks:
    """"
    Handles all registered callbacks for YOLOv5 Hooks
    """

    _callbacks = {
        'on_pretrain_routine_start': [],
        'on_pretrain_routine_end': [],

        'on_train_start': [],
        'on_train_epoch_start': [],
        'on_train_batch_start': [],
        'optimizer_step': [],
        'on_before_zero_grad': [],
        'on_train_batch_end': [],
        'on_train_epoch_end': [],

        'on_val_start': [],
        'on_val_batch_start': [],
        'on_val_image_end': [],
        'on_val_batch_end': [],
        'on_val_end': [],

        'on_fit_epoch_end': [],  # fit = train + val
        'on_model_save': [],
        'on_train_end': [],

        'teardown': [],
    }

    def __init__(self):
        return

    def register_action(self, hook, name='', callback=None):
        """
        Register a new action to a callback hook

        Args:
            hook        The callback hook name to register the action to
            name        The name of the action
            callback    The callback to fire
        """
        assert hook in self._callbacks, f"hook '{hook}' not found in callbacks {self._callbacks}"
        assert callable(callback), f"callback '{callback}' is not callable"
        self._callbacks[hook].append({'name': name, 'callback': callback})

    def get_registered_actions(self, hook=None):
        """"
        Returns all the registered actions by callback hook

        Args:
            hook The name of the hook to check, defaults to all
        """
        if hook:
            return self._callbacks[hook]
        else:
            return self._callbacks

    def run_callbacks(self, hook, *args, **kwargs):
        """
        Loop through the registered actions and fire all callbacks
        """
        for logger in self._callbacks[hook]:
            # print(f"Running callbacks.{logger['callback'].__name__}()")
            logger['callback'](*args, **kwargs)

    def on_pretrain_routine_start(self, *args, **kwargs):
        """
        Fires all registered callbacks at the start of each pretraining routine
        """
        self.run_callbacks('on_pretrain_routine_start', *args, **kwargs)

    def on_pretrain_routine_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of each pretraining routine
        """
        self.run_callbacks('on_pretrain_routine_end', *args, **kwargs)

    def on_train_start(self, *args, **kwargs):
        """
        Fires all registered callbacks at the start of each training
        """
        self.run_callbacks('on_train_start', *args, **kwargs)

    def on_train_epoch_start(self, *args, **kwargs):
        """
        Fires all registered callbacks at the start of each training epoch
        """
        self.run_callbacks('on_train_epoch_start', *args, **kwargs)

    def on_train_batch_start(self, *args, **kwargs):
        """
        Fires all registered callbacks at the start of each training batch
        """
        self.run_callbacks('on_train_batch_start', *args, **kwargs)

    def optimizer_step(self, *args, **kwargs):
        """
        Fires all registered callbacks on each optimizer step
        """
        self.run_callbacks('optimizer_step', *args, **kwargs)

    def on_before_zero_grad(self, *args, **kwargs):
        """
        Fires all registered callbacks before zero grad
        """
        self.run_callbacks('on_before_zero_grad', *args, **kwargs)

    def on_train_batch_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of each training batch
        """
        self.run_callbacks('on_train_batch_end', *args, **kwargs)

    def on_train_epoch_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of each training epoch
        """
        self.run_callbacks('on_train_epoch_end', *args, **kwargs)

    def on_val_start(self, *args, **kwargs):
        """
        Fires all registered callbacks at the start of the validation
        """
        self.run_callbacks('on_val_start', *args, **kwargs)

    def on_val_batch_start(self, *args, **kwargs):
        """
        Fires all registered callbacks at the start of each validation batch
        """
        self.run_callbacks('on_val_batch_start', *args, **kwargs)

    def on_val_image_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of each val image
        """
        self.run_callbacks('on_val_image_end', *args, **kwargs)

    def on_val_batch_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of each validation batch
        """
        self.run_callbacks('on_val_batch_end', *args, **kwargs)

    def on_val_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of the validation
        """
        self.run_callbacks('on_val_end', *args, **kwargs)

    def on_fit_epoch_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of each fit (train+val) epoch
        """
        self.run_callbacks('on_fit_epoch_end', *args, **kwargs)

    def on_model_save(self, *args, **kwargs):
        """
        Fires all registered callbacks after each model save
        """
        self.run_callbacks('on_model_save', *args, **kwargs)

    def on_train_end(self, *args, **kwargs):
        """
        Fires all registered callbacks at the end of training
        """
        self.run_callbacks('on_train_end', *args, **kwargs)

    def teardown(self, *args, **kwargs):
        """
        Fires all registered callbacks before teardown
        """
        self.run_callbacks('teardown', *args, **kwargs)