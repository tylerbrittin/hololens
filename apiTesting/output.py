# Provide color to an output string
# Green for success (true), red for failure (false)
def colorize(string, status):
    attr = []
    if status:
        attr.append('32')
    else:
        attr.append('31')

    return '\x1b[%sm%s\x1b[0m' % (';'.join(attr), string)
