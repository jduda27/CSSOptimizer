# CSS Optimizer
The goal of this project is to create a convenient tool that reads in a .css file and returns a optimized copy of the file. This is useful for disgustingly large css files that are in dire need of a rewrite but need to be organized first in order to better understand where to even begin.

It runs in a few different states, the first state is the analysis state, this analysis will rank the specificity of a specific rule set. We map the line number of the beginning of each rule set and calculate this value following a simple ranking.

To score each rule set we have 3 Integers (0 0 0) which is mapped to a line number.

ID tags # have the heaviest weight (1 0 0)
Classes and other non-generic selectors score (0 1 0)
Generics selectors score (0 0 1)

We also look for duplicate selectors. To do this we will make a simple hash of each selector and if it matches we record the line number of these rule sets to later be refactored.

After analysis we will make a pass to reorganize the document. We begin by creating a sorted list with the lowest specificity score being ranked highest and the most specific being ranked lowest. If two specificities are equal we organize the entries alphabetically.

After this ranking is determined we then parse the document once more this time jumping to the starting line number and copying each rule in the correct order to the exported document.

CSS Duplicates are also handled in this stage, if a duplicate exists we combine all styles from both selectors checking for duplicate styles and only inputing one if there are duplicates.

We can later include CSS minification.
