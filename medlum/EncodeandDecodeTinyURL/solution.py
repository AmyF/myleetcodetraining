class Codec:

    def __init__(self):
        self.map = {}
        self.i = 0
    
    def encode(self, longUrl):
        """Encodes a URL to a shortened URL.
        
        :type longUrl: str
        :rtype: str
        """
        self.map[self.i] = longUrl
        self.i += 1
        return "http://tinyurl.com/" + str(self.i-1)

    def decode(self, shortUrl):
        """Decodes a shortened URL to its original URL.
        
        :type shortUrl: str
        :rtype: str
        """
        tmp = int(shortUrl[len("http://tinyurl.com/"):])
        return self.map[tmp]

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))
