# Colors output for Regression Tests
#
# Code written by:
# Tim Monfette (tjm354)

# Global values for number of passed and failed tests
numPassed = 0
numFailed = 0
total = 0

# Provide color to an output string
# Green for success (true), red for failure (false)
def colorize(string, status):
    attr = []
    if status:
        attr.append('32')
        global numPassed
        numPassed = numPassed + 1
    else:
        attr.append('31')
        global numFailed
        numFailed = numFailed + 1

    global total
    total = total + 1
    return '\x1b[%sm%s\x1b[0m' % (';'.join(attr), string)

# Getter for globals
def getResults():
  return [total, numPassed, numFailed]
