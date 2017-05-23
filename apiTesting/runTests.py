#!/usr/bin/python

# Main Python Script to Test Nebula API
# Use all test cases documented in testCases.py
#
# Code written by:
# Tim Monfette (tjm354)

import testCases
from subprocess import call

# Run test cases
for case in testCases.CASES:
    testcase = "test"+case
    getattr(testCases, testcase)()

# Print out final results
print
testCases.finalResults()

# Clean up database back to working order
print "\nTesting done, rebuilding database..."
call(["/opt/api/bin/rebuildDB"])
print "Finished"
