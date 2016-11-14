[![Build Status](https://secure.travis-ci.org/classmarkets/timex.png?branch=master)](http://travis-ci.org/classmarkets/timex)
[![GoDoc](https://godoc.org/github.com/classmarkets/timex?status.svg)](https://godoc.org/github.com/classmarkets/timex)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/classmarkets/timex/blob/master/LICENSE)

# timex: Additional time functions
 
This package provides functions for adding days, weeks or months to 
a `time.Time` while handling daylight saving time (DST) correctly.

The main use case for this is calculating periodically recurrent dates that
happen at greater intervals but always at the same clock time. 
