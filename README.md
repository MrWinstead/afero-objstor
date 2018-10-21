# Afero Object Storage Driver

[![CircleCI](https://circleci.com/gh/MrWinstead/afero-objstor/tree/master.svg?style=svg)](https://circleci.com/gh/MrWinstead/afero-objstor/tree/master)
[![GoDoc](https://godoc.org/github.com/MrWinstead/afero-objstor?status.svg)](https://godoc.org/github.com/MrWinstead/afero-objstor)

## Overview

A filesystem backend for the
[afero filesystem abstraction](https://github.com/spf13/afero) which is designed
to work with [go-cloud](https://github.com/google/go-cloud) supported object
storage systems. Additionally, it provides an extended interface beyond that of
afero which offers deadline sensitivity provided by golang contexts.

## Notes on Distributed System Use

This library is likely to be used in a cloud-native context with multiple
workers all using the same backend object storage bucket. This library does not
make promises regarding file locking, consistency, or multiple-writer
situations. This would best be solved with locks currently managed outside of
this library.

## Features
### Currently Supported

### Currently Unsupported
