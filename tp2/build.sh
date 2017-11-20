#!/bin/bash
pandoc tp2.md -o tp2.pdf -V geometry:margin=1.2in
zip tp2.zip img/* *js *css *html *pdf
