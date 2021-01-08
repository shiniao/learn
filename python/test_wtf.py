

def test_string_interning():
    a = 'wtf!'
    b = 'wft!'
    result = a is b
    assert result != True

