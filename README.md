# About emoji

Package emoji provides standardized ways for translating unicode code points
for Emoji to/from their [emoji cheat sheet](http://www.webpagefx.com/tools/emoji-cheat-sheet/)
encoding. This is useful when working with third party APIs such as Slack,
GitHub, etc.

This was written because the existing Go emoji packages only provided cheat
sheet names to unicode conversion, and not the opposite. Also, I was not able
to find any emoticon packages for Go.

# Gemoji Data

The data for this package is generated from GitHub's [gemoji](https://github.com/github/gemoji)
project.

# TODO

* Convert `UnicodeVersion` and `IOSVersion` fields of `Emoji` type to something more easily comparable (ie, int)
