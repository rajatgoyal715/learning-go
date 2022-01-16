# FMT Cheat Sheet

## General

- %v (value in default format)
- %T (type)
- %% (literal %)

## Boolean

- %t (true or false)

## Integer

- %b (base 2)
- %o (base 8)
- %d (base 10)
- %x (base 16)

## Floating Points

- %e (scientific notation)
- %f / %F (decimal no exponent)
- %g (for large exponents)

## Strings

- %s (default)
- %q (double quoted string)

## Width & Precision

- %f (default width, default precision)
- %9f (width 9, default precision)
- %.2f (default width, precision 2) // This rounds off the digits after specified decimals
- %9.2f (width 9, precision 2)
- %9.f (width 9, precision 0)

## Padding

- %09d (pads digit to length 9 with preceding 0's)
- %-4d (pads with spaces (width 4, left justified))
