#!/usr/bin/python

# Python Script to Test Nebula API
# Use all test cases documented in testCases.py

import testCases
from subprocess import call

# Run test cases
for case in testCases.CASES:
    testcase = "test"+case
    getattr(testCases, testcase)()

# Clean up database back to working order
print "\nTesting done, rebuilding database..."
call(["/home/tim/goWorkspace/bin/rebuildDB"])
print "Finished"
